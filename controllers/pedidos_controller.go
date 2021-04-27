package controllers

import (
	"log"
	"net/http"

	"github.com/farofadev/goesic/form_requests"
	"github.com/farofadev/goesic/models"
	"github.com/farofadev/goesic/repositories"
	"github.com/farofadev/goesic/responses"
	"github.com/gofiber/fiber/v2"
)

var pedidoRepository = repositories.NewPedidoCachedRepository()

func PedidosIndex(ctx *fiber.Ctx) error {
	page := ctx.Query("page")
	pedidos, err := pedidoRepository.FetchAll(page)

	if err != nil {
		log.Println(err)
		responses.SendResponseInternalServerError(ctx)
		return err
	}

	payload := responses.NewResponseDataPayload()
	payload.Data = pedidos

	return payload.Send(ctx)
}

func PedidosStore(ctx *fiber.Ctx) error {
	pedido := models.NewPedido()
	formRequest := form_requests.NewPedidoFormRequest()

	if err := ctx.BodyParser(formRequest); err != nil {
		log.Println(err)
		responses.SendResponseInternalServerError(ctx)
		return err
	}

	pedido.PessoaId = formRequest.PessoaId

	if _, err := pedidoRepository.Store(pedido); err != nil {
		log.Println(err)
		responses.SendResponseInternalServerError(ctx)
		return err
	}

	payload := responses.NewResponseDataPayload()
	payload.Data = pedido
	payload.StatusCode = http.StatusCreated

	return payload.Send(ctx)
}

func PedidosShow(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	pedido, err := pedidoRepository.FindById(id)

	if err != nil {
		log.Println(err)
		responses.SendResponseInternalServerError(ctx)
		return err
	}

	if pedido.Id == "" {
		responses.SendResponseNotFound(ctx)
		return err
	}

	payload := responses.NewResponseDataPayload()
	payload.Data = pedido

	return payload.Send(ctx)
}
