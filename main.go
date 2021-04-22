package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

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

	db := lib.DBConnectAndSetDefault();

    router := httprouter.New()

	pedidosController := &controllers.PedidosController{}

    router.GET("/", Index)
    router.GET("/pedidos", pedidosController.Index)
	router.POST("/pedidos", pedidosController.Store)
	router.GET("/pedidos/:id", pedidosController.Show)

	defer lib.GetDefaultDBConnection().Close()

	go func() {
		for {
			log.Println("Verificando banco de dados")

			_, err := db.Query("SHOW TABLES;")

			if err != nil {
				log.Println("Conexão falhou, tentando reconectar...")

				lib.DBConnectAndSetDefault();
			}

			time.Sleep(30 * time.Second)
		}
	}()

	log.Fatal(http.ListenAndServe(":8080", router))
}

