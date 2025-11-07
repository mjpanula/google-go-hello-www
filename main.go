package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World! 🌍\n")
	log.Printf("Request received: %s %s from %s", r.Method, r.URL.Path, r.RemoteAddr)
}

func main() {
	http.HandleFunc("/", helloHandler)
	
	port := "8080"
	log.Printf("Starting server on port %s...", port)
	
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
