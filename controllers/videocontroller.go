package controllers

import (
	"fmt"
	"go-streaming-server/conf"
	"go-streaming-server/response"
	"net/http"
	"os"
	"time"

	"github.com/julienschmidt/httprouter"
)

type VideoController struct {
}

func NewVideoController() *VideoController {
	return &VideoController{}
}

func (controller *VideoController) Streaming(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	videoid := p.ByName("vid")
	videopath := fmt.Sprintf("%s/%s", conf.VIDEO_DIR, videoid)

	_, err := os.Stat(videopath)

	if os.IsNotExist(err) {
		response.SendResponse(w, &response.ErrorResponse{
			Code:    http.StatusNotFound,
			Message: "404 video not found.",
		})
		return
	}

	video, err := os.Open(videopath)

	if err != nil {
		response.SendResponse(w, &response.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "File open Error",
		})
		return
	}

	w.Header().Set("Content-Type", "video/x-flv")
	http.ServeContent(w, r, "", time.Now(), video)

	defer video.Close()
}

func (controller *VideoController) Upload(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
}
