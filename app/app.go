package app

import (
	"fmt"
	"log"
	"net/http"
)

type Config struct {
	ServerHost           string
	ServerPort           int
	FrontendContentsPath string
}

func WithFrontendContentsPathConfig(p string) *Config {
	return &Config{
		ServerHost:           "localhost",
		ServerPort:           8080,
		FrontendContentsPath: "./local",
	}
}

type Transfer struct {
	CSPNonce string
}

type App struct {
	Config     *Config
	Transfer   *any
	router     RouterInterface
	middleware []MiddlewareInterface
}

func NewAppWithConfig(c *Config) *App {
	return &App{
		Config:     c,
		Transfer:   nil,
		router:     &Router{},
		middleware: make([]MiddlewareInterface, 0),
	}
}

func (a *App) Run() error {
	http.HandleFunc("/", a.resolve())

	log.Println("start server")
	log.Printf("http://%s:%d", a.Config.ServerHost, a.Config.ServerPort)
	log.Println("------------------------------------------------------")

	return http.ListenAndServe(fmt.Sprintf("%s:%d", a.Config.ServerHost, a.Config.ServerPort), nil)
}

func (a *App) resolve() func(w http.ResponseWriter, r *http.Request) {
	resolve := func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			a.router.Index(w, r, a.Config)
		} else if r.URL.Path == "/privacy-policy/" {
			a.router.PrivacyPolicy(w, r, a.Config)
		} else if r.URL.Path == "/sitemap.xml" {
			a.router.OpenStaticFile(w, r, a.Config, "sitemap.xml")
		} else if r.URL.Path == "/robots.txt" {
			a.router.OpenStaticFile(w, r, a.Config, "robots.txt")
		} else {
			a.router.NotFound(w, r, a.Config)
		}
	}

	for _, m := range a.middleware {
		resolve = m.With(resolve)
	}

	return resolve
}

func (a *App) Use(m MiddlewareInterface) {
	a.middleware = append(a.middleware, m)
}
