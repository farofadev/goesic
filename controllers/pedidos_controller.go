package controllers

import (
	"log"
	"net/http"

	"github.com/farofadev/goesic/form_requests"
	"github.com/farofadev/goesic/models"
	"github.com/farofadev/goesic/repositories"
	"github.com/julienschmidt/httprouter"
)

type PedidosController struct{}

func (*PedidosController) Index(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	repository := &repositories.PedidoRepository{}

	pedidos, err := repository.FetchAll(req.URL.Query().Get("page"))

	if err != nil {
		log.Println(err)

		payload := &ResponseDataPayload{}
		payload.StatusCode = http.StatusInternalServerError

		payload.Send(res)

		return
	}

	payload := &ResponseDataPayload{}
	payload.Data = pedidos

	payload.Send(res)

}

func (*PedidosController) Store(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	repository := &repositories.PedidoRepository{}

	pedido := models.Pedido{}
	formRequest := form_requests.PedidoFormRequest{}

	if _, err := form_requests.DecodeRequestBody(&formRequest, req); err != nil {
		log.Println(err)
		SendResponseInternalServerError(res)
		return
	}

	pedido.PessoaId = formRequest.PessoaId

	if _, err := repository.Store(&pedido); err != nil {
		log.Println(err)
		SendResponseInternalServerError(res)
		return
	}

	payload := &ResponseDataPayload{
		Data:       pedido,
		StatusCode: http.StatusCreated,
	}

	payload.Send(res)
}

func (*PedidosController) Show(res http.ResponseWriter, _ *http.Request, params httprouter.Params) {
	id := params.ByName("id")

	repository := &repositories.PedidoRepository{}

	pedido, err := repository.FindById(id)

	if err != nil {
		log.Println(err)
		SendResponseInternalServerError(res)
		return
	}

	if pedido.Id == "" {
		SendResponseNotFound(res)
		return
	}

	payload := &ResponseDataPayload{
		Data: pedido,
	}

	payload.Send(res)
}
