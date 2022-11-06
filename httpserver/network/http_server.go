package network

import (
	"fmt"
	"httpserver/log"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	log.Logs.Info("收到http请求-/test")
	fmt.Fprintln(w, "hello world")
}

func InitHttpServer() {
	http.HandleFunc("/test", IndexHandler)
	var e = http.ListenAndServe(":8000", nil)
	if e != nil {
		log.Logs.Error(e.Error())
	}
}
