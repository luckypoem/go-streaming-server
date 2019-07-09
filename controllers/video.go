package controllers

import (
	"fmt"
	"go-streaming-server/conf"
	"go-streaming-server/functions"
	"go-streaming-server/models"
	"go-streaming-server/response"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
)

type VideoController struct {
	config *conf.Config
}

func NewVideoController() *VideoController {
	c, err := conf.LoadConfigFromFile("./config.toml")

	if err != nil {
		return nil
	}

	return &VideoController{
		config: c,
	}
}

func (controller *VideoController) Streaming(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	videoid := p.ByName("vid")
	token := p.ByName("token")
	videopath := fmt.Sprintf("%s/%s", controller.config.VideoDir, videoid)

	id, err := strconv.Atoi(videoid)

	if err != nil {
		response.SendResponse(w, http.StatusInternalServerError, &response.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Parse ID Error",
		})
		return
	}

	if !models.CheckVideo(id) {
		response.SendResponse(w, http.StatusNotFound, &response.ErrorResponse{
			Code:    http.StatusNotFound,
			Message: "404 video not found.",
		})
		return
	}

	if !models.CheckToken(id, token) {
		response.SendResponse(w, http.StatusUnauthorized, &response.ErrorResponse{
			Code:    http.StatusUnauthorized,
			Message: "Unautherized.",
		})
		return
	}

	_, err = os.Stat(videopath)

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
	token := functions.RandomString(6)
	videopath := fmt.Sprintf("%s/%s", controller.config.VideoDir, videoid)

	id, err := strconv.Atoi(videoid)

	if err != nil {
		response.SendResponse(w, http.StatusInternalServerError, &response.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "atoi error.",
		})
		return
	}

	if models.CheckVideo(id) {
		response.SendResponse(w, http.StatusBadRequest, &response.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "The file key is already in the folder.",
		})
		return
	}

	_, err = os.Stat(videopath)

	if !os.IsNotExist(err) {
		response.SendResponse(w, http.StatusBadRequest, &response.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "The file key is already in the folder.",
		})
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, controller.config.MaxUploadSize*1024*1024)

	if err := r.ParseMultipartForm(controller.config.MaxUploadSize * 1024 * 1024); err != nil {
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

	err = models.CreateVideo(&models.Video{
		FileId: id,
		Token:  token,
	})

	if err != nil {
		response.SendResponse(w, http.StatusInternalServerError, &response.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "db error.",
		})
		return
	}

	response.SendResponse(w, http.StatusOK, &response.TokenResponse{
		Code:    http.StatusOK,
		Message: "Successfully uploaded the video",
		Token:   token,
	})
}

func (controller *VideoController) Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	videoid := p.ByName("vid")
	videopath := fmt.Sprintf("%s/%s", controller.config.VideoDir, videoid)

	_, err := os.Stat(videopath)

	if os.IsNotExist(err) {
		response.SendResponse(w, http.StatusNotFound, &response.ErrorResponse{
			Code:    http.StatusNotFound,
			Message: "404 video not found.",
		})
		return
	}

	err = os.Remove(videopath)

	if err != nil {
		response.SendResponse(w, http.StatusInternalServerError, &response.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Remove file error.",
		})
		return
	}

	response.SendResponse(w, http.StatusOK, &response.Response{
		Code:    http.StatusOK,
		Message: "Video was successfully deleted.",
	})
}
