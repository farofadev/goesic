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
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
)

func Index(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	fmt.Fprint(res, "Welcome!\n")
}

func generatePedido() {
	pedidoRepository := repositories.PedidoRepository{}

	pedido := models.Pedido{
		PessoaId: uuid.NewString(),
	}

	_, err := pedidoRepository.Store(&pedido)

	if err != nil {
		log.Println("Erro ao gerar pedido", err)
	} else {
		log.Println("Pedido gerado:", pedido.Id)
	}
}

func generatePedidoLoop() {
	count := 100

	for i := 0; i < count; i++ {
		time.Sleep(2 * time.Minute)

		generatePedido()
	}

	log.Println("Pedidos gerados com sucesso.")
}

func main() {

	godotenv.Load()

	go generatePedidoLoop()

	rand.Seed(time.Now().UnixNano())

	router := httprouter.New()

	router.GET("/", Index)
	router.GET("/pedidos", controllers.PedidosIndex)
	router.POST("/pedidos", controllers.PedidosStore)
	router.GET("/pedidos/:id", controllers.PedidosShow)

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
