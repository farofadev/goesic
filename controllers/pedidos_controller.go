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

var pedidoCacheRepository = repositories.NewPedidoCachedRepository()
var pedidoRepository = repositories.NewPedidoRepository()
var mensagemRepository = repositories.NewMensagemRepository()

//PedidosIndex Essa função faz tal coisa...
func PedidosIndex(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	pedidos, err := pedidoCacheRepository.FetchAll(req.URL.Query().Get("page"))

	if err != nil {
		log.Println(err)
		responses.SendResponseInternalServerError(res)
		return
	}

	payload := responses.NewResponseDataPayload()
	payload.Data = pedidos

	payload.Send(res)
}

func PedidosStore(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	pedido := models.NewPedido()
	formRequest := form_requests.NewPedidoFormRequest()

	if _, err := form_requests.DecodeRequestBody(formRequest, req); err != nil {
		log.Println(err)
		responses.SendResponseInternalServerError(res)
		return
	}

	//pedido.PessoaId = formRequest.PessoaId

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

func PedidosShow(res http.ResponseWriter, _ *http.Request, params httprouter.Params) {
	id := params.ByName("id")

	pedido, err := pedidoCacheRepository.FindById(id)

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

func PedidosResponder(res http.ResponseWriter, req *http.Request, params httprouter.Params) {

	PedidoId := params.ByName("id")
	mensagem := models.NewMensagem()

	mensagem.PedidoId = PedidoId

	log.Println("MENSAGEM COM PEDIDO ID AQUI: ", mensagem.PedidoId)

	if _, err := form_requests.DecodeRequestBody(mensagem, req); err != nil {
		log.Println(err)
		responses.SendResponseInternalServerError(res)
		return
	}

	if _, err := pedidoRepository.Update(PedidoId, mensagem.Tipo); err != nil {
		log.Println(err)
		responses.SendResponseInternalServerError(res)
		return
	}

	if _, err := mensagemRepository.Store(mensagem); err != nil {
		log.Println(err)
		responses.SendResponseInternalServerError(res)
		return
	}

	payload := responses.NewResponseDataPayload()
	payload.Data = mensagem
	payload.Send(res)

}
