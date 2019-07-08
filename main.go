package main

import (
	"fmt"
	"go-streaming-server/conf"
	"go-streaming-server/routers"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	app := httprouter.New()
	router := routers.NewRouter(app)

	err := router.ConfigureRouter()

	if err != nil {
		log.Fatal(err)
	}

	httpServer := http.Server{
		Addr:    fmt.Sprintf("%s:%d", conf.HOST, conf.PORT),
		Handler: router,
	}

	log.Printf("Server started on %s:%d", conf.HOST, conf.PORT)
	err = httpServer.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}
