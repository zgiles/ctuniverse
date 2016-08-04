package main

import (
	"log"
	"strconv"
	"time"
	"golang.org/x/net/websocket"
	"net/http"
	"github.com/tylerb/graceful" // "gopkg.in/tylerb/graceful.v1"
	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
	redis_universedb "github.com/zgiles/ctuniverse/db/redis/universedb"
	universestore "github.com/zgiles/ctuniverse/stores/universestore"
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

	// Local Variables
	var l_universestoredb universestore.StoreDBI
	var l_universestore universestore.StoreI


	switch config.Serverconfig.Maindb {
	case "redis":
		if config.Redisconfig.Enabled == false {
			log.Fatal("redis selected, but not enabled")
		}

		log.Println("redis: opening redis connection")
		redispool, rediserr := redisStart(config.Redisconfig) // this is a redispool *redis.Pool
		if rediserr != nil {
			log.Fatal(rediserr)
		}
		defer log.Println("redis: no close needed...")
		// defer db.Close()
		log.Println("redis: open")

		log.Println("redis: opening UniverseStoreDB")
		l_universestoredb = redis_universedb.New(redispool)
		log.Println("redis: opening UnivserStore")
		l_universestore = universestore.New(l_universestoredb)

	default:
		log.Fatal("no valid db selected as primary")

	}

	// app context
	appC := appContext{ universestore: l_universestore }
	log.Println("app ready")		
	log.Println(appC)

	commonHandlers := alice.New(context.ClearHandler, loggingHandler, recoverHandler)

	router := httprouter.New()
	// router.GET("/", wrapHandler(commonHandlers.ThenFunc(indexHandler)))
	router.GET("/", wrapHandler(http.FileServer(http.Dir("."))))
	router.GET("/ws", wrapHandler(websocket.Handler(wshandler)))
	router.NotFound = commonHandlers.ThenFunc(errorHandler)

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
