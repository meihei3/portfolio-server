package app

import (
	"log"
	"net/http"
	"os"
	"time"
)

type MiddlewareInterface interface {
	With(func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request)
}

type LoggerMiddleware struct {
	logger *log.Logger
}

func (m *LoggerMiddleware) With(f func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		m.logger.Printf("method:%s uri:%s remote_ip:%s time:%s host:%s", r.Method, r.URL, r.Header.Get("X-Real-IP"), time.Now().Format(time.RFC3339), r.Host)
		f(w, r)
	}
}

func NewLoggerMiddleware() *LoggerMiddleware {
	return &LoggerMiddleware{
		// これ以外は使わないと思う
		logger: log.New(os.Stdout, "", log.LstdFlags),
	}
}
