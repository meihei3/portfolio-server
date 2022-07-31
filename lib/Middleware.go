package lib

import (
	"log"
	"net/http"
	"time"
)

func WithLoggerMiddleware(f func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("time:%s remote_ip:%s host:%s method:%s uri:%s", time.Now().Format(time.RFC3339), r.Header.Get("X-Real-IP"), r.Host, r.Method, r.URL)
		f(w, r)
	}
}
