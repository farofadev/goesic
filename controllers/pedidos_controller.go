package controllers

import (
	"log"
	"net/http"

	"github.com/farofadev/goesic/form_requests"
	"github.com/farofadev/goesic/models"
	"github.com/farofadev/goesic/repositories"
	"github.com/farofadev/goesic/responses"
	"github.com/gorilla/mux"
)

var pedidoRepository = repositories.NewPedidoCachedRepository()

func PedidosIndex(res http.ResponseWriter, req *http.Request) {
	pedidos, err := pedidoRepository.FetchAll(req.URL.Query().Get("page"))

	if err != nil {
		log.Println(err)
		responses.SendResponseInternalServerError(res)
		return
	}

	payload := responses.NewResponseDataPayload()
	payload.Data = pedidos

	payload.Send(res)
}

func PedidosStore(res http.ResponseWriter, req *http.Request) {
	pedido := models.NewPedido()
	formRequest := form_requests.NewPedidoFormRequest()

	if _, err := form_requests.DecodeRequestBody(formRequest, req); err != nil {
		log.Println(err)
		responses.SendResponseInternalServerError(res)
		return
	}

	pedido.PessoaId = formRequest.PessoaId

	if _, err := pedidoRepository.Store(pedido); err != nil {
		log.Println(err)
		responses.SendResponseInternalServerError(res)
		return
	}

	payload := responses.NewResponseDataPayload()
	payload.Data = pedido
	payload.StatusCode = http.StatusCreated

	payload.Send(res)
}

func PedidosShow(res http.ResponseWriter, req *http.Request) {

	params := mux.Vars(req)

	pedido, err := pedidoRepository.FindById(params["id"])

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
