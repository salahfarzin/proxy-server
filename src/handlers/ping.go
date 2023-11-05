package handlers

import "net/http"

const (
	pong = "pong"
)

func Ping(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(pong))
}
