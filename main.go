package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"

	"github.com/hoalt14/go-api/handler"
)

var (
	host = ""
	port = 8080
)

func main() {
	r := chi.NewRouter()
	r.Get("/hello", handler.GetHello())
	r.Get("/", handler.GetRoot())

	addr := fmt.Sprintf("%s:%d", host, port)
	svr := http.Server{
		Addr:         addr,
		Handler:      r,
		TLSConfig:    nil,
		WriteTimeout: time.Second,
		ReadTimeout:  time.Second,
		IdleTimeout:  time.Second,
	}

	if err := svr.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
