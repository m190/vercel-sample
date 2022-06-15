package api

import "net/http"

func HandlerFirst(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	handle(w, "first")
}

func handle(w http.ResponseWriter, str string) {
	w.Write([]byte(str))
}
