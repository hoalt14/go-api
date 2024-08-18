package handler

import (
	"encoding/json"
	"net/http"

	"rsc.io/quote/v3"
)

func GetRoot() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Test Something"))
	}
}

func GetHello() http.HandlerFunc {
	type reponse struct {
		Message string `json:"message"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		res := reponse{quote.HelloV3()}
		b, err := json.Marshal(res)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Write(b)
	}
}
