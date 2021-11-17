package main

import (
	"fmt"
	"net/http"
)

type MyHandler struct {

}

// ＆と*の違い: https://qiita.com/tmzkysk/items/1b73eaf415fee91aaad3
// ポインタ渡し: https://qiita.com/kotaonaga/items/4a93ec40718c279154f5
// Struct, Method, Interface: https://qiita.com/S-Masakatsu/items/6fb8e765cd443e2edd7f
// 構造体とメゾッド: https://ema-hiro.hatenablog.com/entry/20170510/1494427060
// type Handler: https://pkg.go.dev/net/http#Handler

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func main() {
	handler := MyHandler{}
	fmt.Println("handler:", handler)
	fmt.Println("&handler:", &handler)
	server := http.Server{
		Addr: "127.0.0.1:8080",
		Handler: &handler,
	}
	server.ListenAndServe()
}
