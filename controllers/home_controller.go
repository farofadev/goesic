package controllers

import (
	"net/http"

	"github.com/farofadev/goesic/responses"
	"github.com/julienschmidt/httprouter"
)

func HomeIndex(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	payload := responses.NewResponseDataPayload()

	payload.Data = "Bem-vindo!"

	payload.Send(res)
}
