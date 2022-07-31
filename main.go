package main

import (
	crand "crypto/rand"
	"fmt"
	"log"
	"math"
	"math/big"
	"math/rand"
	"net/http"
	"strings"
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

type CSPHeader struct {
	DefaultSRC     []string
	StyleSRC       []string
	ScriptSRC      []string
	FontSRC        []string
	ImgSRC         []string
	BaseURI        []string
	FormAction     []string
	FrameAncestors []string
	ConnectSRC     []string
}

func (c *CSPHeader) String() string {
	return fmt.Sprintf("default-src %s; style-src %s; script-src %s; font-src %s; img-src %s; base-uri %s; form-action %s; frame-ancestors %s; connect-src %s;",
		strings.Join(c.DefaultSRC, " "), strings.Join(c.StyleSRC, " "), strings.Join(c.ScriptSRC, " "), strings.Join(c.FontSRC, " "), strings.Join(c.ImgSRC, " "), strings.Join(c.BaseURI, " "), strings.Join(c.FormAction, " "), strings.Join(c.FrameAncestors, " "), strings.Join(c.ConnectSRC, " "))
}

func makeCSPHeader() *CSPHeader {
	// 個人用なので、汎用的に使えなくてもいい。
	return &CSPHeader{
		DefaultSRC:     []string{"'none'"},
		StyleSRC:       []string{"'self'"},
		ScriptSRC:      []string{"'strict-dynamic'"},
		FontSRC:        []string{"'self'"},
		ImgSRC:         []string{"'self'"},
		BaseURI:        []string{"'none'"},
		FormAction:     []string{"'none'"},
		FrameAncestors: []string{"'none'"},
		ConnectSRC:     []string{"'self'", "*.google-analytics.com"},
	}
}

func main() {
	seed, _ := crand.Int(crand.Reader, big.NewInt(math.MaxInt64))
	rand.Seed(seed.Int64()) // math/rand は seed 値が固定なので、生成する必要がある。

	http.HandleFunc("/privacy-policy/", privacyPolicyHandler)
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
	log.Printf("method:indexHandler time:%s remote_ip:%s host:%s method:%s uri:%s", time.Now().Format(time.RFC3339), r.RemoteAddr, r.Host, r.Method, r.URL)
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	nonce := lib.GenerateRandomStr(32)
	c := makeCSPHeader()
	c.ScriptSRC = append(c.ScriptSRC, fmt.Sprintf("'nonce-%s'", nonce))
	w.Header().Add("Content-Security-Policy", fmt.Sprintf("%v", c))

	t, err := template.ParseFiles(frontendContentsPath + "/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, Option{CSPNonce: nonce})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func privacyPolicyHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("method:privacyPolicyHandler time:%s remote_ip:%s host:%s method:%s uri:%s", time.Now().Format(time.RFC3339), r.RemoteAddr, r.Host, r.Method, r.URL)

	nonce := lib.GenerateRandomStr(32)
	c := makeCSPHeader()
	c.ScriptSRC = append(c.ScriptSRC, fmt.Sprintf("'nonce-%s'", nonce))
	w.Header().Add("Content-Security-Policy", fmt.Sprintf("%v", c))

	t, err := template.ParseFiles(frontendContentsPath + "/privacy-policy/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, Option{CSPNonce: nonce})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
