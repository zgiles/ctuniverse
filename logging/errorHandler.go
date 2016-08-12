package main

import (
	"net/http"
)

func errorHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(500)
	w.Write([]byte("Internal Error"))
}
