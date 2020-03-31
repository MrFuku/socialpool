package main

import (
	"log"
	"net/http"
)

const addr string = ":8084"

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("public"))))
	log.Println("Webサイトのアドレス:", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Println("サーバーの起動に失敗しました：", err)
	}
}
