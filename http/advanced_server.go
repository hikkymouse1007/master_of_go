package main

import "net/http"

func main() {
	// /usr/local/Cellar/go/1.14/libexec/src/net/http/server.go
	server := http.Server{
		Addr: "127.0.0.1:8080",
		Handler: nil,
	}
	server.ListenAndServe()
}
