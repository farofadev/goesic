package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/farofadev/goesic/repositories"
	"github.com/julienschmidt/httprouter"
)

type PedidosController struct {
    
}

func (ctrl *PedidosController) Index(res http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
    res.Header().Set("Content-Type", "application/json")
	
	res.WriteHeader(http.StatusOK)


    repository := &repositories.PedidoRepository{}
   
    pedidos := repository.FetchAll()

    var data struct {
        Data *[]repositories.Pedido `json:"data"`
        Error bool `json:"error"`
    }

    data.Data = pedidos
    data.Error = false

    re, _ := json.Marshal(data)

    res.Write(re)
}

func (ctrl *PedidosController) Criar(res http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
    res.Header().Set("Content-Type", "application/json")
	
	res.WriteHeader(http.StatusOK)
    fmt.Fprintf(res, `{
        "pedido" : {"Id": "wehre"}
    }
    `)
}

func  (ctrl *PedidosController) Exibir(res http.ResponseWriter, _ *http.Request, params httprouter.Params) {
    res.Header().Set("Content-Type", "application/json")
	
	res.WriteHeader(http.StatusOK)

    id := params.ByName("id")

    repository := &repositories.PedidoRepository{}

    pedido := repository.FindById(id)

    var data struct {
        Data *repositories.Pedido `json:"data"`
        Error bool `json:"error"`
    }

    data.Data = pedido
    data.Error = false

    re, _ := json.Marshal(data)

    res.Write(re)
}