package routers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Router struct {
	context *httprouter.Router
}

func NewRouter(ctx *httprouter.Router) *Router {
	return &Router{
		context: ctx,
	}
}

func (router *Router) ConfigureRouter() error {
	return nil
}

func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	router.context.ServeHTTP(w, r)
}
