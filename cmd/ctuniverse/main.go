package main

import (
	"log"
	"strconv"
	"time"
	"net/http"
	"github.com/tylerb/graceful" // "gopkg.in/tylerb/graceful.v1"
	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

type appContext struct {
	universestore universestore.StoreI
}

func main() {
	// Options Parse

	// Config Stage
	config, configerr := loadConfig("config.toml")
	if configerr != nil {
		log.Fatal(configerr)
	}

	// app context
	appC := appContext{}

	// Handlers
	commonHandlers := alice.New(context.ClearHandler, loggingHandler, recoverHandler)

	router := httprouter.New()
	router.GET("/", wrapHandler(http.FileServer(http.Dir("static/"))))
	router.GET("/ws", wrapHandler(commonHandlers.ThenFunc(wshandler)))
	router.NotFound = commonHandlers.ThenFunc(errorHandler)

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
