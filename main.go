package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Print("helloworld: received a request")
	target := os.Getenv("TARGET")
	if target == "" {
		target = "World"
	}
	fmt.Fprintf(w, "Hello %s!\n", target)
}

func versionHandler(w http.ResponseWriter, r *http.Request) {
	version := os.Getenv("BUILD_VERSION")
	if len(version) == 0 {
		version = "dev"
	}
	fmt.Fprintf(w, "BUILD_VERSION: %s", version)
}

func main() {
	log.Print("helloworld: starting server...")

	http.HandleFunc("/", handler)
	http.HandleFunc("/version", versionHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("helloworld: listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
