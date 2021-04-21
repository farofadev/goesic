package repositories

import (
	"log"

	"github.com/farofadev/goesic/lib"
)

type Pedido struct {
	Id string `json:"id"`
	PessoaId string `json:"pessoa_id"`
	Situacao string `json:"situacao"`
	CriadoEm string `json:"criado_em"`
	DataPrazo string `json:"data_prazo"`
}

type PedidoRepository struct {}


func (repository *PedidoRepository) FetchAll() *[]Pedido {
	pedidos := []Pedido{}

	database := lib.DBConnect()

	rows, err := database.Query("SELECT id, pessoa_id, situacao, criado_em FROM pedidos;")

	if err != nil  {
		log.Fatal(err)
	}

	for rows.Next() {
		// Scan one customer record
		pedido := Pedido{}
		err := rows.Scan(&pedido.Id, &pedido.PessoaId, &pedido.Situacao, &pedido.CriadoEm)

		if err != nil {
			log.Println(err)
		}

		pedidos = append(pedidos, pedido)
	}

	defer rows.Close()
	defer database.Close()

	return &pedidos
}

func (repository *PedidoRepository) FindById(id string) *Pedido {
	pedido := Pedido{}

	database := lib.DBConnect()

	rows, err := database.Query("SELECT id, pessoa_id, situacao, criado_em FROM pedidos where id = ? LIMIT 1;", id)

	if err != nil  {
		log.Fatal(err)
	}

	for rows.Next() {
		err := rows.Scan(&pedido.Id, &pedido.PessoaId, &pedido.Situacao, &pedido.CriadoEm)

		if err != nil {
			log.Println(err)
		}
	}

	defer rows.Close()
	defer database.Close()

	return &pedido
}
