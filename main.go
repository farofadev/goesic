package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/farofadev/goesic/controllers"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
)

func Index(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
    fmt.Fprint(res, "Welcome!\n")
}

func main() {
	godotenv.Load()

	log.Println("Iniciando...")

    router := httprouter.New()

	pedidosController := &controllers.PedidosController{}

    router.GET("/", Index)
    router.GET("/pedidos", pedidosController.Index)
	router.POST("/pedidos", pedidosController.Criar)
	router.GET("/pedidos/:id", pedidosController.Exibir)

	appHost := os.Getenv("APP_HOST")

	if appHost == "" {
		appHost = "0.0.0.0:8080"
	}

    log.Fatal(http.ListenAndServe(appHost, router))
}


// func Hello(res http.ResponseWriter, req *http.Request, param httprouter.Params) {
//     fmt.Fprintf(res, "hello, %s!\n", param.ByName("name"))
// }
