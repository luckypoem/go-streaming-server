package main

import (
	"fmt"
	"go-streaming-server/conf"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()

	return router
}

func main() {
	router := RegisterHandlers()
	config, err := conf.LoadConfigFromFile("./config.toml")
	if err != nil {
		log.Fatal(err)
	}

	httpServer := http.Server{
		Addr:    fmt.Sprintf("%s:%s", config.Host, strconv.Itoa(config.Port)),
		Handler: router,
	}

	err = httpServer.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
