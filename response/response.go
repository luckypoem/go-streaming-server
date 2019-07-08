package response

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func SendResponse(w http.ResponseWriter, v interface{}) {
	data, _ := json.Marshal(v)

	w.Write(data)
}
