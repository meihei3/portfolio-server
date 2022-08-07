package app

import (
	"fmt"
	"html/template"
	"net/http"
	"os"

	"github.com/meihei3/portfolio-server/lib"
)

type RouterInterface interface {
	Index(w http.ResponseWriter, req *http.Request, cnf *Config)
	PrivacyPolicy(w http.ResponseWriter, req *http.Request, cnf *Config)
	OpenStaticFile(w http.ResponseWriter, req *http.Request, cnf *Config, fn string)
	NotFound(w http.ResponseWriter, req *http.Request, cnf *Config)
}

type Router struct{}

func (r *Router) Index(w http.ResponseWriter, req *http.Request, cnf *Config) {
	nonce := lib.GenerateRandomStr(32)
	c := NewCSPHeader()
	c.ScriptSRC = append(c.ScriptSRC, fmt.Sprintf("'nonce-%s'", nonce))
	w.Header().Add("Content-Security-Policy", fmt.Sprintf("%v", c))

	t, err := template.ParseFiles(cnf.FrontendContentsPath + "/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, Transfer{CSPNonce: nonce})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (r *Router) PrivacyPolicy(w http.ResponseWriter, req *http.Request, cnf *Config) {
	nonce := lib.GenerateRandomStr(32)
	c := NewCSPHeader()
	c.ScriptSRC = append(c.ScriptSRC, fmt.Sprintf("'nonce-%s'", nonce))
	w.Header().Add("Content-Security-Policy", fmt.Sprintf("%v", c))

	t, err := template.ParseFiles(cnf.FrontendContentsPath + "/privacy-policy/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, Transfer{CSPNonce: nonce})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (r *Router) OpenStaticFile(w http.ResponseWriter, req *http.Request, cnf *Config, fn string) {
	// Routerで存在するファイルだけを呼び出す。
	f, err := os.Open(cnf.FrontendContentsPath + "/" + fn)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	d, err := f.Stat()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.ServeContent(w, req, cnf.FrontendContentsPath+"/"+fn, d.ModTime(), f)
}

func (r *Router) NotFound(w http.ResponseWriter, req *http.Request, cnf *Config) {
	w.WriteHeader(404)

	nonce := lib.GenerateRandomStr(32)
	c := NewCSPHeader()
	c.ScriptSRC = append(c.ScriptSRC, fmt.Sprintf("'nonce-%s'", nonce))
	w.Header().Add("Content-Security-Policy", fmt.Sprintf("%v", c))

	t, err := template.ParseFiles(cnf.FrontendContentsPath + "/404.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, Transfer{CSPNonce: nonce})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
