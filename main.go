package main

import (
	crand "crypto/rand"
	"fmt"
	"log"
	"math"
	"math/big"
	"math/rand"
	"net/http"
	"text/template"
	"time"

	"github.com/meihei3/portfolio-server/lib"
)

const (
	port = 8080
)

var (
	frontendContentsPath = "./local" // 本番環境では`-ldflags`フラグで上書きする
)

type Option struct {
	CSPNonce string
}

func main() {
	seed, _ := crand.Int(crand.Reader, big.NewInt(math.MaxInt64))
	rand.Seed(seed.Int64()) // math/rand は seed 値が固定なので、生成する必要がある。

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

	t, err := template.ParseFiles(frontendContentsPath + "/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, Option{CSPNonce: lib.GenerateRandomStr(32)})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
