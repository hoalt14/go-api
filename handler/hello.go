package handler

import (
	"net/http"

	"rsc.io/quote/v3"
)

func GetHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(quote.HelloV3()))
}
