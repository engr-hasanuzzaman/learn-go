package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func header(w http.ResponseWriter, req *http.Request) {
	for k, v := range req.Header {
		fmt.Fprintf(w, "%s=%s", k, v)
	}
}

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/header", header)

	http.ListenAndServe(":8000", nil)
}
