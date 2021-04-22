package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/farofadev/goesic/models"
	"github.com/farofadev/goesic/repositories"
	"github.com/julienschmidt/httprouter"
)

type PedidosController struct {
    
}

func (*PedidosController) Index(res http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
    repository := &repositories.PedidoRepository{}
   
    pedidos, err := repository.FetchAll()

    if err != nil {
        res.Header().Set("Content-Type", "application/json")
        res.WriteHeader(http.StatusInternalServerError)
        res.Write([]byte(`{"message": "Erro não esperado.", "error": true}`))
        return
    }

    var payload struct {
        Data *[]models.Pedido `json:"data"`
        Error bool `json:"error"`
    }

    payload.Data = pedidos
    payload.Error = false

    re, _ := json.Marshal(payload)

    res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
    res.Write(re)
}

func (*PedidosController) Store(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
    repository := &repositories.PedidoRepository{}
           
    decoder := json.NewDecoder(req.Body)

	var pedido models.Pedido

	decoder.Decode(&pedido)

    _, err := repository.Store(&pedido)
    
    if err != nil {
        res.Header().Set("Content-Type", "application/json")
        res.WriteHeader(http.StatusInternalServerError)
        res.Write([]byte(`{"message": "Erro não esperado.", "error": true}`))
        return
    }

    var payload struct {
        Data *models.Pedido `json:"data"`
        Error bool `json:"error"`
    }

    payload.Data = &pedido

    re, _ := json.Marshal(payload)

    res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated)
    res.Write(re)
}

func  (*PedidosController) Show(res http.ResponseWriter, _ *http.Request, params httprouter.Params) {
    id := params.ByName("id")

    repository := &repositories.PedidoRepository{}

    pedido, err := repository.FindById(id)

    if err != nil {
        res.Header().Set("Content-Type", "application/json")
        res.WriteHeader(http.StatusInternalServerError)
        res.Write([]byte(`{"message": "Erro não esperado.", "error": true}`))
        return
    }

    if pedido.Id == "" {
        res.Header().Set("Content-Type", "application/json")
        res.WriteHeader(http.StatusNotFound)
        res.Write([]byte(`{"message": "Pedido não encontrado", "error": true}`))

        return
    }

    var payload struct {
        Data *models.Pedido `json:"data"`
        Error bool `json:"error"`
    }

    payload.Data = pedido
    payload.Error = false

    re, _ := json.Marshal(payload)

    res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
    res.Write(re)
}