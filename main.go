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

func main() {
	godotenv.Load()

	log.Println("Iniciando...")

	lib.DBConnectAndSetDefault();
	
	// Fecha conexão com o baco de dados após a execução/falha da função main()
	defer lib.GetDefaultDBConnection().Close()
	
	// Execução assíncrona do loop de verificação da conexão com banco de dados
	go lib.LoopVerifyDefaultDBConnection()

    router := httprouter.New()

	pedidosController := &controllers.PedidosController{}

    router.GET("/", Index)
    router.GET("/pedidos", pedidosController.Index)
	router.POST("/pedidos", pedidosController.Store)
	router.GET("/pedidos/:id", pedidosController.Show)

	log.Fatal(http.ListenAndServe(":8080", router))
}

