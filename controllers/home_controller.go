package controllers

import (
	"net/http"

	"github.com/farofadev/goesic/responses"
)

func HomeIndex(res http.ResponseWriter, req *http.Request) {
	payload := responses.NewResponseDataPayload()

	payload.Data = "Bem-vindo!"

	payload.Send(res)
}
