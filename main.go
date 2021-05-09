package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/farofadev/goesic/controllers"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
)

func main() {

	godotenv.Load()

	rand.Seed(time.Now().UnixNano())

	router := httprouter.New()

	router.GET("/", controllers.HomeIndex)
	router.GET("/pedidos", controllers.PedidosIndex)
	router.POST("/pedidos", controllers.PedidosStore)
	router.GET("/pedidos/:id", controllers.PedidosShow)
	router.POST("/pedidos/responder/:id", controllers.PedidosResponder)

	go log.Println("Aplicação iniciada.")
	defer log.Println("Finalizada.")

	server := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}
