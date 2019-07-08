package controllers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type VideoController struct {
}

func (controller *VideoController) Streaming(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}

func (controller *VideoController) Upload(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}
