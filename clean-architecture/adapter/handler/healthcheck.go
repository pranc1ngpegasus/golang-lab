package handler

import (
	"net/http"
)

type Healthcheck http.Handler

func NewHealthcheck() Healthcheck {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	return mux
}
