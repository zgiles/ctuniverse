// Copyright 2016 Zachary Giles
// MIT License (Expat)
//
// Please see the LICENSE file

package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/tylerb/graceful"
	"log"
	"net/http"
	"strconv"
	"time"
)

//go:generate go-bindata -o=bindata.go -pkg=main static/...
//go:generate gofmt -w -s .

var appname = "ctuniverse"
var buildtime = "NoDateTimeProvided"
var githash = "NoGitHashProvided"

func panichandler(w http.ResponseWriter, r *http.Request, ps interface{}) {
	w.Write([]byte("<html>Not Found</html>"))
}

func main() {
	// Options Parse

	// Config Stage
	config, configerr := loadConfig("config.toml")
	if configerr != nil {
		log.Fatal(configerr)
	}

	log.Printf("AppName: %s", appname)
	log.Printf("GitHash: %s", githash)
	log.Printf("BuildTime: %s", buildtime)
	log.Println("App Setting up...")

	// app state
	hub := newHub()
	go hub.run()

	// Handlers
	router := httprouter.New()
	router.RedirectTrailingSlash = true
	router.RedirectFixedPath = true
	router.HandleOPTIONS = true
	router.PanicHandler = panichandler

	router.GET("/", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		log.Printf("[%s] %q %v\n", r.Method, r.URL.String(), time.Now())
		a, _ := Asset("static/index.html")
		w.Write(a)
	})

	router.GET("/ws", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		wshandler(hub, w, r)
	})

	log.Println("App running...")
	// Server
	httpsrv := &graceful.Server{
		Timeout: time.Duration(config.Serverconfig.Closetimeout) * time.Second,
		Server: &http.Server{
			Addr:    config.Serverconfig.IP + ":" + strconv.FormatInt(config.Serverconfig.Port, 10),
			Handler: router,
		},
	}

	httperr := httpsrv.ListenAndServe()
	if httperr != nil {
		log.Fatal(httperr)
	}

	log.Println("main: end of main")

}
