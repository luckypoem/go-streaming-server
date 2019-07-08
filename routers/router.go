package routers

import (
	"go-streaming-server/limiters"
	"go-streaming-server/response"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Router struct {
	context *httprouter.Router
	limiter *limiters.ConnectionLimiter
}

func NewRouter(ctx *httprouter.Router, maxconnection int) *Router {
	return &Router{
		context: ctx,
		limiter: limiters.NewConnectionLimiter(maxconnection),
	}
}

func (router *Router) ConfigureRouter() error {
	return nil
}

func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if !router.limiter.GetConnection() {
		response.SendResponse(w, &response.ErrorResponse{
			Code:    http.StatusTooManyRequests,
			Message: "Too many requests.",
		})
	}

	router.context.ServeHTTP(w, r)

	defer router.limiter.FreeConnection()
}
