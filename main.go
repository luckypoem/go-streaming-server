package main

import (
	"fmt"
	"go-streaming-server/conf"
	_ "go-streaming-server/models"
	"go-streaming-server/routers"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	app := httprouter.New()
	config, err := conf.LoadConfigFromFile("./config.toml")

	if err != nil {
		log.Fatal(err)
	}

	router := routers.NewRouter(app, config.MaxConnection)

	err = router.ConfigureRouter()

	if err != nil {
		log.Fatal(err)
	}

	httpServer := http.Server{
		Addr:    fmt.Sprintf("%s:%d", config.Host, config.Port),
		Handler: router,
	}

	log.Printf("Server started on %s:%d", config.Host, config.Port)
	err = httpServer.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}
