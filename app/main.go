// main
package main

import (
	"log"
	"net/http"
)

const addr = ":8080"

func main() {
	mux := http.NewServeMux()
	handler := &MyHandler{}
	mux.Handle("/favicon.ico", http.NotFoundHandler())
	mux.Handle("/", handler)
	log.Printf("Now listening on %s...\n", addr)
	server := http.Server{Handler: mux, Addr: addr}
	log.Fatal(server.ListenAndServe())
}
