package controllers

import (
	"fmt"
	"go-streaming-server/conf"
	"go-streaming-server/response"
	"io/ioutil"
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
		response.SendResponse(w, http.StatusNotFound, &response.ErrorResponse{
			Code:    http.StatusNotFound,
			Message: "404 video not found.",
		})
		return
	}

	video, err := os.Open(videopath)

	if err != nil {
		response.SendResponse(w, http.StatusInternalServerError, &response.ErrorResponse{
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
	videoid := p.ByName("vid")
	videopath := fmt.Sprintf("%s/%s", conf.VIDEO_DIR, videoid)

	r.Body = http.MaxBytesReader(w, r.Body, conf.MAX_UPLOAD_SIZE*1024*1024)

	if err := r.ParseMultipartForm(conf.MAX_UPLOAD_SIZE * 1024 * 1024); err != nil {
		response.SendResponse(w, http.StatusInternalServerError, &response.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "parse form error.",
		})
		return
	}

	file, fileheader, err := r.FormFile("video")

	if err != nil {
		response.SendResponse(w, http.StatusInternalServerError, &response.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "parse form error.",
		})
		return
	}

	if fileheader.Header.Get("Content-Type") != "video/x-flv" {
		response.SendResponse(w, http.StatusBadRequest, &response.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "You can only upload `flv` files",
		})
		return
	}

	videodata, err := ioutil.ReadAll(file)

	err = ioutil.WriteFile(videopath, videodata, 06666)

	if err != nil {
		response.SendResponse(w, http.StatusInternalServerError, &response.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Save `flv` file error.",
		})
		return
	}

	response.SendResponse(w, http.StatusOK, &response.Response{
		Code:    http.StatusOK,
		Message: "Successfully uploaded the video",
	})
}
