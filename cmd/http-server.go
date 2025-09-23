package main

import (
	"io"
	"net/http"
)

func EchoHttp(w http.ResponseWriter, req *http.Request) {
	io.Copy(w, req.Body)
}

func main() {
	http.Handle("/echo", http.HandlerFunc(EchoHttp))
	err := http.ListenAndServe(":12346", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
