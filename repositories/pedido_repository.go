package repositories

import (
	"fmt"

	"github.com/farofadev/goesic/database"
	"github.com/farofadev/goesic/models"
	"github.com/farofadev/goesic/utils"
)

type PedidoRepository struct{}

func NewPedidoRepository() *PedidoRepository {
	return &PedidoRepository{}
}

func (*PedidoRepository) FetchAll(a ...interface{}) (*[]models.Pedido, error) {
	pedidos := []models.Pedido{}

	paginatorParams, err := utils.GetPaginatorParams(1, 25, a)

	if err != nil {
		return &pedidos, err
	}

	db, e1 := database.DBConnectDefault()

	if e1 != nil {
		return &pedidos, e1
	}

	defer db.Close()

	rawSql := fmt.Sprintf("SELECT %s FROM pedidos LIMIT ? OFFSET ?;", (&models.Pedido{}).SqlColumnsString())
	rows, err := db.Query(rawSql, paginatorParams.PageSize, paginatorParams.GetOffset())

	if err != nil {
		return &pedidos, err
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
	pedido := models.NewPedido()

	db, e1 := database.DBConnectDefault()

	if e1 != nil {
		return pedido, e1
	}

	defer db.Close()

	rawSql := fmt.Sprintf("SELECT %s FROM pedidos where %s = ? LIMIT 1;", pedido.SqlColumnsString(), field)
	rows, err := db.Query(rawSql, value)

	if err != nil {
		return pedido, err
	}

	defer rows.Close()

	for rows.Next() {
		err := pedido.ScanFromSqlRows(rows)

		if err != nil {
			return pedido, err
		}
	}

	return pedido, nil
}

func (*PedidoRepository) Store(pedido *models.Pedido) (*models.Pedido, error) {
	db, e1 := database.DBConnectDefault()

	if e1 != nil {
		return pedido, e1
	}

	defer db.Close()

	if pedido.Id == "" || pedido.Protocolo == "" {
		pedido.Situacao = models.PedidoSituacaoAberto
		pedido.SetId()
		pedido.MakeProtocol()
		pedido.SetCriadoEm()
		pedido.MakeDataPrazoSituacaoAberto()
	}

	if pedido.Situacao == "" {
		pedido.Situacao = "aberto"
	}

	rawSql := fmt.Sprintf("INSERT INTO pedidos (%s) VALUES (%s);", pedido.SqlColumnsString(), pedido.SqlReplacementsString())
	_, err := db.Exec(rawSql, pedido.RowValues()...)

	return pedido, err
}
