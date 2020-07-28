package myservice

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

type HelloServer struct{}

func (s * HelloServer) Run() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/headers", headers)
	http.ListenAndServe(":8080", nil)
}
