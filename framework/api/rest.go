package api

import (
	"net/http"
)

// Get get url 调用 默认返回不支持
func (b *RestAPI) Get() {
	b.CustomResponse(http.StatusMethodNotAllowed, 405, nil, "unspported http method")
}

// Post get url 调用 默认返回不支持
func (b *RestAPI) Post() {
	b.CustomResponse(http.StatusMethodNotAllowed, 405, nil, "unspported http method")
}

// Put get url 调用 默认返回不支持
func (b *RestAPI) Put() {
	b.CustomResponse(http.StatusMethodNotAllowed, 405, nil, "unspported http method")
}

// Patch get url 调用 默认返回不支持
func (b *RestAPI) Patch() {
	b.CustomResponse(http.StatusMethodNotAllowed, 405, nil, "unspported http method")
}

// Delete get url 调用 默认返回不支持
func (b *RestAPI) Delete() {
	b.CustomResponse(http.StatusMethodNotAllowed, 405, nil, "unspported http method")
}
