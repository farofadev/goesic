package controllers

import (
	"github.com/farofadev/goesic/responses"
	"github.com/gofiber/fiber/v2"
)

func HomeIndex(ctx *fiber.Ctx) error {
	payload := responses.NewResponseDataPayload()

	payload.Data = map[string]string{"message": "Bem-vindo!"}

	return payload.Send(ctx)
}
