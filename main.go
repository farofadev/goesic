package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/farofadev/goesic/controllers"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	rand.Seed(time.Now().UnixNano())

	app := fiber.New()

	app.Get("/", controllers.HomeIndex)
	app.Get("/pedidos", controllers.PedidosIndex)
	app.Post("/pedidos", controllers.PedidosStore)
	app.Get("/pedidos/:id", controllers.PedidosShow)

	log.Fatal(app.Listen(":8080"))
}
