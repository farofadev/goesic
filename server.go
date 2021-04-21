package main

import (
	"log"
	"net/http"

	"github.com/farofadev/goesic/controllers"
)

// func get(http, caminho, metodo) {
// 	http.HandleFunc(caminho, func (response, request) {
// 		if (request.method != "GET") {
// 			response.WriteHeader(http.StatusNotFound)
// 			response.Write([]byte(`Page not found. (Method not allowed)`))
// 			return
// 		}

// 		metodo()
// 	})
// }

func main() {
	pedidosController := &controllers.PedidosController{}

	http.HandleFunc("/api/pedidos", pedidosController.Index)
	http.HandleFunc("/api/pedidos/criar", pedidosController.Criar)
	http.HandleFunc("/api/pedidos/exibir", pedidosController.Exibir)

    log.Fatal(http.ListenAndServe(":8888", nil))
}