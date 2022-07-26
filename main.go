package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = 8080

func main() {
	http.HandleFunc("/", indexHandler)

	log.Println("start server")
	log.Printf("http://localhost:%d", port)

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Fatalf("failed server: %v", err)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}
