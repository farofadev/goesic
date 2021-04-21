package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/farofadev/goesic/controllers"
	"github.com/julienschmidt/httprouter"
)

func Index(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
    fmt.Fprint(res, "Welcome!\n")
}

func Hello(res http.ResponseWriter, req *http.Request, param httprouter.Params) {
    fmt.Fprintf(res, "hello, %s!\n", param.ByName("name"))
}

func main() {
    router := httprouter.New()

	pedidosController := &controllers.PedidosController{}

	http.HandleFunc("/api/pedidos", pedidosController.Index)
	// http.HandleFunc("/api/pedidos/criar", pedidosController.Criar)
	// http.HandleFunc("/api/pedidos/exibir", pedidosController.Exibir)

    router.GET("/", Index)
    router.GET("/hello/:name", Hello)

    log.Fatal(http.ListenAndServe(":8080", router))
}


	


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

