package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/farofadev/goesic/controllers"
	"github.com/farofadev/goesic/models"
	"github.com/farofadev/goesic/repositories"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

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

	// router := httprouter.New()

	// router.GET("/", controllers.HomeIndex)
	// router.GET("/pedidos", controllers.PedidosIndex)
	// router.POST("/pedidos", controllers.PedidosStore)
	// router.GET("/pedidos/:id", controllers.PedidosShow)

	r := mux.NewRouter()
	r.HandleFunc("/", controllers.HomeIndex).Methods("GET")
	r.HandleFunc("/pedidos", controllers.PedidosIndex).Methods("GET")
	r.HandleFunc("/pedidos", controllers.PedidosStore).Methods("POST")
	r.HandleFunc("/pedidos/{id}", controllers.PedidosShow).Methods("GET")

	go log.Println("Aplicação iniciada.")
	defer log.Println("Finalizada.")

	server := &http.Server{
		Addr:         ":8080",
		Handler:      r,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}
