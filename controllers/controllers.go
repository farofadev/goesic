package controllers

import (
	"encoding/json"
	"net/http"
)

type ResponseDataPayload struct {
	Data       interface{} `json:"data"`
	Message    string      `json:"message"`
	StatusCode int         `json:"status_code"`
	Error      bool        `json:"error"`
}

func (payload *ResponseDataPayload) Send(res http.ResponseWriter) {

	if payload.StatusCode == 0 {
		payload.StatusCode = http.StatusOK
	}

	if payload.Message == "" {
		payload.Message = http.StatusText(payload.StatusCode)
	}

	if payload.StatusCode >= 400 && payload.StatusCode < 600 {
		payload.Error = true
	}

	jsonBytes, _ := json.Marshal(payload)

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(payload.StatusCode)
	res.Write(jsonBytes)
}

func SendResponseInternalServerError(res http.ResponseWriter) {
	payload := ResponseDataPayload{
		StatusCode: http.StatusInternalServerError,
	}

	payload.Send(res)
}

func SendResponseNotFound(res http.ResponseWriter) {
	payload := ResponseDataPayload{
		StatusCode: http.StatusNotFound,
	}

	payload.Send(res)
}
