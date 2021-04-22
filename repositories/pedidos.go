package repositories

import (
	"log"

	"github.com/farofadev/goesic/lib"
	"github.com/farofadev/goesic/models"
	"github.com/google/uuid"
)

type PedidoRepository struct {}

func (*PedidoRepository) FetchAll() *[]models.Pedido {
	pedidos := []models.Pedido{}

	database := lib.DBConnect()

	rows, err := database.Query("SELECT id, pessoa_id, situacao, criado_em FROM pedidos;")

	if err != nil  {
		log.Fatal(err)
	}

	for rows.Next() {
		// Scan one customer record
		pedido := models.Pedido{}
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

func (*PedidoRepository) FindById(id string) *models.Pedido {
	pedido := models.Pedido{}

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

func (*PedidoRepository) Store(pedido *models.Pedido) (*models.Pedido, error) {
	
	database := lib.DBConnect()

	if pedido.Id == "" {
		pedido.Id = uuid.NewString()
	}

	if pedido.Situacao == "" {
		pedido.Situacao = "aberto"
	}

	rows, err := database.Query("INSERT INTO pedidos (id, pessoa_id, situacao, criado_em, data_prazo) VALUES (?,?,?,?,?);", pedido.Id, pedido.PessoaId, pedido.Situacao, pedido.CriadoEm, pedido.DataPrazo)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	defer database.Close()

	return pedido, err
}