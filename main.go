package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/farofadev/goesic/controllers"
	"github.com/farofadev/goesic/lib"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
)

func Index(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
    fmt.Fprint(res, "Welcome!\n")
}

func tryDBConnect() {
	log.Println("Testando Conexão com o Banco de Dados...")

	database := lib.DBConnect()
	err := database.Close()

	if err != nil {
		log.Fatalln("Erro ao conectar com o Banco de Dados",err)
	}

	log.Println("Banco de Dados: OK")
}

func main() {
	godotenv.Load()

	log.Println("Iniciando...")

	tryDBConnect()

    router := httprouter.New()

	pedidosController := &controllers.PedidosController{}

    router.GET("/", Index)
    router.GET("/pedidos", pedidosController.Index)
	router.POST("/pedidos", pedidosController.Store)
	router.GET("/pedidos/:id", pedidosController.Show)

	log.Fatal(http.ListenAndServe(":8080", router))
}

