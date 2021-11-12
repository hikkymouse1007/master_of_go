package main

import (
	"fmt"
	"net/http"
	"reflect"
)

func main() {
	server := http.Server{
		Addr: "127.0.0.01:8080",
		Handler: nil,
	}
	fmt.Println(reflect.TypeOf(server))
	server.ListenAndServe()
}
