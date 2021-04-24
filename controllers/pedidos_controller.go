package controllers

import (
	"log"
	"net/http"

	"github.com/farofadev/goesic/form_requests"
	"github.com/farofadev/goesic/models"
	"github.com/farofadev/goesic/repositories"
	"github.com/farofadev/goesic/responses"
	"github.com/julienschmidt/httprouter"
)

func PedidosIndex(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	repository := &repositories.PedidoRepository{}

	pedidos, err := repository.FetchAll(req.URL.Query().Get("page"))

	if err != nil {
		log.Println(err)

		payload := responses.NewResponseDataPayload()
		payload.StatusCode = http.StatusInternalServerError
		payload.Send(res)

		return
	}

	payload := responses.NewResponseDataPayload()
	payload.Data = pedidos

	payload.Send(res)
}

func PedidosStore(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	repository := &repositories.PedidoRepository{}

	pedido := models.Pedido{}
	formRequest := form_requests.PedidoFormRequest{}

	if _, err := form_requests.DecodeRequestBody(&formRequest, req); err != nil {
		log.Println(err)
		responses.SendResponseInternalServerError(res)
		return
	}

	pedido.PessoaId = formRequest.PessoaId

	if _, err := repository.Store(&pedido); err != nil {
		log.Println(err)
		responses.SendResponseInternalServerError(res)
		return
	}

	payload := responses.NewResponseDataPayload()
	payload.Data = pedido
	payload.StatusCode = http.StatusCreated

	payload.Send(res)
}

func PedidosShow(res http.ResponseWriter, _ *http.Request, params httprouter.Params) {
	id := params.ByName("id")

	repository := &repositories.PedidoRepository{}

	pedido, err := repository.FindById(id)

	if err != nil {
		log.Println(err)
		responses.SendResponseInternalServerError(res)
		return
	}

	if pedido.Id == "" {
		responses.SendResponseNotFound(res)
		return
	}

	payload := responses.NewResponseDataPayload()
	payload.Data = pedido
	payload.Send(res)
}
