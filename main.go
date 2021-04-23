package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/farofadev/goesic/controllers"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
)

func Index(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
    fmt.Fprint(res, "Welcome!\n")
}

func main() {
	godotenv.Load()

	rand.Seed(time.Now().UnixNano())

	log.Println("Iniciando...")

    router := httprouter.New()
	
	pedidosController := &controllers.PedidosController{}

    router.GET("/", Index)
    router.GET("/pedidos", pedidosController.Index)
	router.POST("/pedidos", pedidosController.Store)
	router.GET("/pedidos/:id", pedidosController.Show)

	defer log.Println("Finalizada.")
	
	server := &http.Server{
		Addr: ":8080",
		Handler: router,
		ReadTimeout: 30*time.Second,
		WriteTimeout: 30*time.Second,
	}

	log.Fatal(server.ListenAndServe())
}

