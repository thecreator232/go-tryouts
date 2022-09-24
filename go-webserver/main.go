package main

import (
	"fmt"
	"log"
	"net/http"
)

func handleHelloWorld(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		fmt.Fprintf(w, "kat lo")
	}

	if r.Method != "GET" {
		fmt.Fprintf(w, "galat method")
	}
}

func main() {
	http.HandleFunc("/hello", handleHelloWorld)
	fmt.Print("starting server on 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
