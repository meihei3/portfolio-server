package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	port = 8080
)

var (
	frontendContentsPath = "./local" // 本番環境では`-ldflags`フラグで上書きする
)

func main() {
	http.HandleFunc("/", indexHandler)

	log.Println("start server")
	log.Printf("http://localhost:%d", port)
	log.Println("------------------------------------------------------")

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Fatalf("failed server: %v", err)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("time:%s remote_ip:%s host:%s method:%s uri:%s", time.Now().Format(time.RFC3339), r.RemoteAddr, r.Host, r.Method, r.URL)
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	f, err := os.Open(frontendContentsPath + "/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	d, err := f.Stat()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.ServeContent(w, r, frontendContentsPath+"/index.html", d.ModTime(), f)
}
