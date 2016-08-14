// Copyright 2016 Zachary Giles
// MIT License (Expat)
//
// Please see the LICENSE file

package main

import (
	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func wrapHandler(h http.Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		context.Set(r, "params", ps)
		h.ServeHTTP(w, r)
	}
}
