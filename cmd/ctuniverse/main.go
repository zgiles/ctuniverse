// Code by Zachary Giles
// This code is under the MIT License, a copy of which is found in the LICENSE file.

package main

import (
	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
	"github.com/tylerb/graceful" // "gopkg.in/tylerb/graceful.v1"
	"github.com/zgiles/ctuniverse/logging"
	"log"
	"net/http"
	"strconv"
	"time"
)

func main() {
	// Options Parse

	// Config Stage
	config, configerr := loadConfig("config.toml")
	if configerr != nil {
		log.Fatal(configerr)
	}

	log.Println("App Setting up...")
	hub := newHub()
	go hub.run()

	// Handlers
	commonHandlers := alice.New(context.ClearHandler, logging.LoggingHandler, logging.RecoverHandler)

	router := httprouter.New()
	router.GET("/", wrapHandler(http.FileServer(http.Dir("static/"))))
	router.GET("/ws", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		context.Set(r, "params", ps)
		wshandler(hub, w, r)
	})
	router.NotFound = commonHandlers.ThenFunc(logging.ErrorHandler)

	log.Println("App running...")
	// Server
	httpsrv := &graceful.Server{
		Timeout: time.Duration(config.Serverconfig.Closetimeout) * time.Second,
		Server: &http.Server{
			Addr:    config.Serverconfig.Ip + ":" + strconv.FormatInt(config.Serverconfig.Port, 10),
			Handler: router,
		},
	}

	httperr := httpsrv.ListenAndServe()
	if httperr != nil {
		log.Fatal(httperr)
	}

	log.Println("main: end of main")

}
