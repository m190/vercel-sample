package api

import "net/http"

func HandlerSecond(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	handle(w, "second")
}
