package controllers

import (
	"log"
	"net/http"
)

type PedidosController struct {
    
}

func (c PedidosController) Index(response http.ResponseWriter, resquest *http.Request) {
    response.Header().Set("Content-Type", "application/json")
	
	response.WriteHeader(http.StatusOK)
    response.Write([]byte(`{"message": "Index"}`))
}

func (c PedidosController) Criar(response http.ResponseWriter, request *http.Request) {
    response.Header().Set("Content-Type", "application/json")
	
	response.WriteHeader(http.StatusOK)
    response.Write([]byte(`{"message": "Criar"}`))
}

func  (c PedidosController) Exibir(response http.ResponseWriter, request *http.Request) {
    log.Printf("ID: %s", request.URL.Query().Get("id"))
    response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
    response.Write([]byte(`{"message": "Exibir"}`))
}