package file

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

const (
	LogDateFormat  = "2006-01-02"
	DefaultLogScan = 10    //定时检查文件是否需要切割间隔时间
	DefaultLogSeq  = 10000 //队列长度
)

type Logger struct {
	mu         *sync.RWMutex
	filePrefix string
	fileSuffix string
	date       *time.Time
	logFile    *os.File
	logChan    chan string
}

//访问日志初始化
func LogInit(prefix, suffix string) *Logger {
	lg := &Logger{
		mu:         new(sync.RWMutex),
		filePrefix: prefix,
		fileSuffix: suffix,
		logChan:    make(chan string, DefaultLogSeq),
	}

	t, _ := time.Parse(LogDateFormat, time.Now().Format(LogDateFormat))

	lg.date = &t
	lg.mu.Lock()
	defer lg.mu.Unlock()

	if !lg.isMustSplit() {
		filePath := lg.filePath(lg.date.Format(LogDateFormat))
		lg.logFile, _ = os.OpenFile(filePath, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644)
	} else {
		lg.split()
	}

	go lg.logWriter()
	go lg.fileMonitor()

	return lg
}

//追加日志
func (lg *Logger) Append(s string) {
	nowFormat := time.Now().Format("2006-01-02 15:04:05")
	lg.logChan <- fmt.Sprintf("[%v]\t%s", nowFormat, s)
}

//异步写日志
func (lg *Logger) logWriter() {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("logWritter() catch panic: %v", err)
		}
	}()

	for {
		select {
		case str := <-lg.logChan:
			lg.write(str)
		}
	}
}

//将字符串写入到文件
func (lg *Logger) write(str string) {
	lg.mu.RLock()
	defer lg.mu.RUnlock()
	buf := []byte(str)
	if len(str) == 0 || str[len(str)-1] != '\n' {
		buf = append(buf, '\n')
	}
	_, _ = lg.logFile.Write(buf)
}

//判断是否需要切割
func (lg *Logger) isMustSplit() bool {
	t, _ := time.Parse(LogDateFormat, time.Now().Format(LogDateFormat))
	if t.After(*lg.date) {
		return true
	}
	return false
}

//切割日志
func (lg *Logger) split() {
	logFile := lg.filePath(time.Now().Format(LogDateFormat))
	if !lg.isExist(logFile) {
		if lg.logFile != nil {
			_ = lg.logFile.Close()
		}

		t, _ := time.Parse(LogDateFormat, time.Now().Format(LogDateFormat))
		lg.date = &t
		lg.logFile, _ = os.Create(logFile)
	}
}

func (lg *Logger) filePath(date string) string {
	filePath := lg.filePrefix + date + lg.fileSuffix
	return filePath
}

//文件监控
func (lg *Logger) fileMonitor() {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("fileMonitor() catch panic: %v", err)
		}
	}()

	timer := time.NewTicker(time.Duration(DefaultLogScan) * time.Second)
	for {
		select {
		case <-timer.C:
			lg.fileCheck()
		}
	}
}

func (lg *Logger) fileCheck() {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("fileCheck() catch panic: %v", err)
		}
	}()

	if lg.isMustSplit() {
		lg.mu.Lock()
		defer lg.mu.Unlock()
		lg.split()
	}
}

// Determine a file or a path exists in the os
func (lg *Logger) isExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}
