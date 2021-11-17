package main

import (
	"fmt"
	"net/http"
)

// Handler周りの参考記事
// https://teratail.com/questions/224653
// https://journal.lampetty.net/entry/understanding-http-handler-in-go


// 構造体とメゾッド: https://ema-hiro.hatenablog.com/entry/20170510/1494427060
//Handlerはinterface型なので、定義先はどんな型でもいいのだけれど、"ServeHTTPというメソッドを持っていなくてはならない。"
//つまりHelloHandlerの構造体であってもServeHTTPを定義することができる
// type Handler: https://pkg.go.dev/net/http#Handler

// 結局、ServeHTTPメゾッドを持つ構造体は、自動的にHandler型になる。

type HelloHandler struct {

}

func (h *HelloHandler) ServeHTTP (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}

type WorldHandler struct {

}

func (h *WorldHandler) ServeHTTP (w http.ResponseWriter, r *http.Request) {
	fmt.Println("request", r)
	fmt.Println("&request", &r)
	fmt.Println("response writer", w)
	fmt.Println("&response writer", &w)
	fmt.Fprintf(w, "world")
}


func main() {
	hello := HelloHandler{}
	world := WorldHandler{}

	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.Handle("/hello", &hello)
	http.Handle("/world", &world)

	server.ListenAndServe()

}


// まとめると、ServeHTTPメゾッドを持ったHandlerの構造体を作って、http.