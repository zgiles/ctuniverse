// Copyright 2016 Zachary Giles
// MIT License (Expat)
//
// Please see the LICENSE file

// Package logging contains handlers to print data to the Go Log
package logging

import (
	"log"
	"net/http"
	"time"
)

// TimeHandler is a wrapper that returns how long an HTTP Call takes.
func TimeHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		t1 := time.Now()
		next.ServeHTTP(w, r)
		t2 := time.Now()
		log.Printf("[%s] %q %v\n", r.Method, r.URL.String(), t2.Sub(t1))
	}
	return http.HandlerFunc(fn)
}

// ErrorHandler is the default handler in-case all other ones fail.
func ErrorHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(500)
	w.Write([]byte("Internal Error"))
}

// RecoverHandler handles situtations where inside functions exit abnormally.
func RecoverHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("panic: %+v", err)
				http.Error(w, http.StatusText(500), 500)
			}
		}()
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
