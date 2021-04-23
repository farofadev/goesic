package repositories

import (
	"fmt"

	"github.com/farofadev/goesic/lib"
	"github.com/farofadev/goesic/models"
)

type PedidoRepository struct {}

func (*PedidoRepository) FetchAll() (*[]models.Pedido, error) {
	pedidos := []models.Pedido{}

	database := lib.GetDefaultDBConnection()

	rawSql := fmt.Sprintf("SELECT %s FROM pedidos;", (&models.Pedido{}).SqlColumnsString())
	rows, err := database.Query(rawSql)

	if err != nil  {
		return &pedidos,err
	}

	defer rows.Close()

	for rows.Next() {
		// Scan one customer record
		pedido := models.Pedido{}
		err := pedido.ScanFromSqlRows(rows)

		if err != nil {
			return &pedidos, err
		}

		pedidos = append(pedidos, pedido)
	}

	return &pedidos, nil
}

func (repository *PedidoRepository) FindById(id string) (*models.Pedido, error) {
	return repository.FindBy("id", id)
}

func (*PedidoRepository) FindBy(field string, value interface{}) (*models.Pedido, error) {
	pedido := models.Pedido{}

	database := lib.GetDefaultDBConnection()

	rawSql := fmt.Sprintf("SELECT %s FROM pedidos where %s = ? LIMIT 1;", pedido.SqlColumnsString(), field)
	rows, err := database.Query(rawSql, value)

	if err != nil  {
		return &pedido, err
	}

	defer rows.Close()

	for rows.Next() {
		// Scan one customer record
		err := pedido.ScanFromSqlRows(rows)

		if err != nil {
			return &pedido, err
		}
	}

	return &pedido, nil
}

func (*PedidoRepository) Store(pedido *models.Pedido) (*models.Pedido, error) {
	database := lib.GetDefaultDBConnection()

	if pedido.Id == "" {
		pedido.MakeId()
		pedido.MakeProtocol()
	}

	if pedido.Situacao == "" {
		pedido.Situacao = "aberto"
	}

	rawSql := fmt.Sprintf("INSERT INTO pedidos (%s) VALUES (%s);", pedido.SqlColumnsString(), pedido.SqlReplacementsString())
	rows, err := database.Query(rawSql, pedido.RowValues()...)

	if err != nil {
		return pedido, err
	}

	defer rows.Close()

	return pedido, nil
}