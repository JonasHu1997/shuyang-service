package curl

import (
	"bytes"
	"compress/gzip"
	"crypto/tls"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net"
	"net/http"
	"net/http/cookiejar"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"
)

type UCurl struct {
}

// HTTPSettings is the http.Client setting
type HTTPSettings struct {
	ShowDebug        bool
	UserAgent        string
	ConnectTimeout   time.Duration
	ReadWriteTimeout time.Duration
	TLSClientConfig  *tls.Config
	Proxy            func(*http.Request) (*url.URL, error)
	Transport        http.RoundTripper
	EnableCookie     bool
	Gzip             bool
	DumpBody         bool
}

// HTTPRequest provides more useful methods for Requesting one url than http.Request.
type HTTPRequest struct {
	url     string
	req     *http.Request
	params  map[string][]string
	files   map[string]string
	setting HTTPSettings
	resp    *http.Response
	body    []byte
	header  map[string][]string
	dump    []byte
}

var defaultCookieJar http.CookieJar
var settingMutex sync.Mutex
var defaultSetting = HTTPSettings{
	UserAgent:        "Go-http-client/1.1",
	ConnectTimeout:   5,
	ReadWriteTimeout: 10,
	Gzip:             true,
	DumpBody:         true,
}

// String returns the body string in response.
// it calls Response inner.
func (b *HTTPRequest) String() (string, error) {
	data, err := b.Bytes()
	if err != nil {
		return "", err
	}

	return string(data), nil
}

// Body adds Request raw body.
// it supports string and []byte.
func (b *HTTPRequest) Body(data interface{}) *HTTPRequest {
	switch t := data.(type) {
	case string:
		bf := bytes.NewBufferString(t)
		b.req.Body = ioutil.NopCloser(bf)
		b.req.ContentLength = int64(len(t))
	case []byte:
		bf := bytes.NewBuffer(t)
		b.req.Body = ioutil.NopCloser(bf)
		b.req.ContentLength = int64(len(t))
	}
	return b
}

func (b *HTTPRequest) getResponse() (*http.Response, error) {
	if b.resp.StatusCode != 0 {
		return b.resp, nil
	}
	resp, err := b.DoRequest()
	if err != nil {
		return nil, err
	}
	b.resp = resp
	return resp, nil
}

// createDefaultCookie creates a global cookiejar to store cookies.
func createDefaultCookie() {
	settingMutex.Lock()
	defer settingMutex.Unlock()
	defaultCookieJar, _ = cookiejar.New(nil)
}

// Header add header item string in Request.
func (b *HTTPRequest) Header(key, value string) *HTTPRequest {
	if strings.ToLower(key) == "host" {
		b.req.Host = value
	} else {
		b.req.Header.Set(key, value)
	}
	return b
}

// TimeoutDialer returns functions of connection dialer with timeout settings for http.Transport Dial field.
func TimeoutDialer(cTimeout time.Duration, rwTimeout time.Duration) func(net, addr string) (c net.Conn, err error) {
	return func(netW, addr string) (net.Conn, error) {
		conn, err := net.DialTimeout(netW, addr, cTimeout)
		if err != nil {
			return nil, err
		}
		err = conn.SetDeadline(time.Now().Add(rwTimeout))
		return conn, err
	}
}

// SetTimeout sets connect time out and read-write time out for HTTPRequest.
func (b *HTTPRequest) SetTimeout(connectTimeout, readWriteTimeout time.Duration) *HTTPRequest {
	b.setting.ConnectTimeout = connectTimeout
	b.setting.ReadWriteTimeout = readWriteTimeout
	return b
}

// DoRequest will do the client.Do
func (b *HTTPRequest) DoRequest() (*http.Response, error) {
	var paramBody string
	if len(b.params) > 0 {
		var buf bytes.Buffer
		for k, v := range b.params {
			for _, vv := range v {
				buf.WriteString(url.QueryEscape(k))
				buf.WriteByte('=')
				buf.WriteString(url.QueryEscape(vv))
				buf.WriteByte('&')
			}
		}
		paramBody = buf.String()
		paramBody = paramBody[0 : len(paramBody)-1]
	}

	b.buildURL(paramBody)
	uri, err := url.Parse(b.url)
	if err != nil {
		return nil, err
	}

	b.req.URL = uri

	trans := b.setting.Transport

	if trans == nil {
		// create default transport
		trans = &http.Transport{
			TLSClientConfig: b.setting.TLSClientConfig,
			Proxy:           b.setting.Proxy,
			Dial:            TimeoutDialer(b.setting.ConnectTimeout, b.setting.ReadWriteTimeout),
		}
	} else {
		// if b.transport is *http.Transport then set the settings.
		if t, ok := trans.(*http.Transport); ok {
			if t.TLSClientConfig == nil {
				t.TLSClientConfig = b.setting.TLSClientConfig
			}
			if t.Proxy == nil {
				t.Proxy = b.setting.Proxy
			}
		}
	}

	var jar http.CookieJar
	if b.setting.EnableCookie {
		if defaultCookieJar == nil {
			createDefaultCookie()
		}
		jar = defaultCookieJar
	}

	client := &http.Client{
		Transport: trans,
		Jar:       jar,
	}

	if b.setting.UserAgent != "" && b.req.Header.Get("User-Agent") == "" {
		b.req.Header.Set("User-Agent", b.setting.UserAgent)
	}

	if b.setting.ShowDebug {
		dump, err := httputil.DumpRequest(b.req, b.setting.DumpBody)
		if err != nil {
			log.Println(err.Error())
		}
		b.dump = dump
	}
	return client.Do(b.req)
}

func (b *HTTPRequest) buildURL(paramBody string) {
	// build GET url with query string
	if b.req.Method == "GET" && len(paramBody) > 0 {
		if strings.Index(b.url, "?") != -1 {
			b.url += "&" + paramBody
		} else {
			b.url = b.url + "?" + paramBody
		}
		return
	}

	// build POST/PUT/PATCH url and body
	if (b.req.Method == "POST" || b.req.Method == "PUT" || b.req.Method == "PATCH") && b.req.Body == nil {
		// with files
		if len(b.files) > 0 {
			pr, pw := io.Pipe()
			bodyWriter := multipart.NewWriter(pw)
			go func() {
				for name, filename := range b.files {
					fileWriter, err := bodyWriter.CreateFormFile(name, filename)
					if err != nil {
						log.Println("Httplib:", err)
					}
					fh, err := os.Open(filename)
					if err != nil {
						log.Println("Httplib:", err)
					}
					//iocopy
					_, err = io.Copy(fileWriter, fh)
					_ = fh.Close()
					if err != nil {
						log.Println("Httplib:", err)
					}
				}
				for k, v := range b.params {
					for _, vv := range v {
						_ = bodyWriter.WriteField(k, vv)
					}
				}
				_ = bodyWriter.Close()
				_ = pw.Close()
			}()
			b.Header("Content-Type", bodyWriter.FormDataContentType())
			b.req.Body = ioutil.NopCloser(pr)
			return
		}

		// with params
		if len(paramBody) > 0 {
			b.Header("Content-Type", "application/x-www-form-urlencoded")
			b.Body(paramBody)
		}
	}
}

// Bytes returns the body []byte in response.
// it calls Response inner.
func (b *HTTPRequest) Bytes() ([]byte, error) {
	b1, _, err := b.Response()
	return b1, err
}

func (b *HTTPRequest) Response() ([]byte, map[string][]string, error) {
	var h map[string][]string
	if b.body != nil {
		return b.body, h, nil
	}
	resp, err := b.getResponse()
	if err != nil {
		return nil, h, err
	}
	if resp.Body == nil {
		return nil, h, nil
	}
	h = resp.Header
	defer func() {
		_ = resp.Body.Close()
	}()
	if b.setting.Gzip && resp.Header.Get("Content-Encoding") == "gzip" {
		reader, err := gzip.NewReader(resp.Body)
		if err != nil {
			return nil, h, err
		}
		b.body, err = ioutil.ReadAll(reader)
	} else {
		b.body, err = ioutil.ReadAll(resp.Body)
	}
	return b.body, h, err
}

// params build query string as ?key1=value1&key2=value2...
func (b *HTTPRequest) Param(key, value string) *HTTPRequest {
	if param, ok := b.params[key]; ok {
		b.params[key] = append(param, value)
	} else {
		b.params[key] = []string{value}
	}
	return b
}

func (*UCurl) construct(rawUrl, method string) *HTTPRequest {
	var resp http.Response
	u, err := url.Parse(rawUrl)
	if err != nil {
		log.Println("Httplib:", err)
	}
	req := http.Request{
		URL:        u,
		Method:     method,
		Header:     make(http.Header),
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		ProtoMinor: 1,
	}
	return &HTTPRequest{
		url:     rawUrl,
		req:     &req,
		params:  map[string][]string{},
		files:   map[string]string{},
		setting: defaultSetting,
		resp:    &resp,
	}
}

func GET(rawUrl string) *HTTPRequest {
	var resp http.Response
	u, err := url.Parse(rawUrl)
	if err != nil {
		log.Println(fmt.Sprintf("Httplib:%s", err))
	}
	req := http.Request{
		URL:        u,
		Method:     "GET",
		Header:     make(http.Header),
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		ProtoMinor: 1,
	}
	handle := &HTTPRequest{
		url:     rawUrl,
		req:     &req,
		params:  map[string][]string{},
		files:   map[string]string{},
		setting: defaultSetting,
		resp:    &resp,
	}
	handle.SetTimeout(5*time.Second, 5*time.Second)
	return handle
}

func POST(rawUrl string) *HTTPRequest {
	var resp http.Response
	u, err := url.Parse(rawUrl)
	if err != nil {
		log.Println(fmt.Sprintf("Httplib:%s", err))
	}
	req := http.Request{
		URL:        u,
		Method:     "POST",
		Header:     make(http.Header),
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		ProtoMinor: 1,
	}
	handle := &HTTPRequest{
		url:     rawUrl,
		req:     &req,
		params:  map[string][]string{},
		files:   map[string]string{},
		setting: defaultSetting,
		resp:    &resp,
	}
	handle.SetTimeout(5*time.Second, 5*time.Second)
	return handle
}
