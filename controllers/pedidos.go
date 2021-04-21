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

func (_ *PedidosController) Index(res http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
    repository := &repositories.PedidoRepository{}
   
    pedidos := repository.FetchAll()

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

func (_ *PedidosController) Store(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
    repository := &repositories.PedidoRepository{}
           
    decoder := json.NewDecoder(req.Body)

	var pedido models.Pedido

	err := decoder.Decode(&pedido)
    
	if err != nil {
		panic(err)
	}

    repository.Store(&pedido)
    
    var payload struct {
        Data *models.Pedido `json:"data"`
        Error bool `json:"error"`
    }

    payload.Data = &pedido

    re, _ := json.Marshal(payload)

    res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
    res.Write(re)
}

func  (_ *PedidosController) Show(res http.ResponseWriter, _ *http.Request, params httprouter.Params) {
    id := params.ByName("id")

    repository := &repositories.PedidoRepository{}

    pedido := repository.FindById(id)

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