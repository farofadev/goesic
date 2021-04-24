package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/farofadev/goesic/controllers"
	"github.com/farofadev/goesic/models"
	"github.com/farofadev/goesic/repositories"
	"github.com/farofadev/goesic/utils"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
)

func Index(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
    fmt.Fprint(res, "Welcome!\n")
}

func generatePedido() {

   pedidoRepository := repositories.PedidoRepository{}

   pedido := models.Pedido {
	    PessoaId: uuid.NewString(),
        Situacao: "aberto",
        CriadoEm: utils.FormatDateTimeString(time.Now()),
        DataPrazo: "2021-04-30", 
   }
   pedidoRepository.Store(&pedido)
}

func generatePedidoLoop() {
	count := 100

	for i := 0; i < count; i++ {
		time.Sleep(5 * time.Minute)

		generatePedido()
	}
	
}


func main() {
	
	godotenv.Load()

	go generatePedidoLoop()

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

