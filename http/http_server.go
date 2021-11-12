package main

import (
	"fmt"
	"net/http"
)

func main() {
	server := http.Server{
		Addr: "127.0.0.01:8080",
		Handler: nil,
	}
	fmt.Println(server)
	server.ListenAndServe()
}
