package repositories

import (
	"fmt"
	"log"

	"github.com/farofadev/goesic/database"
	"github.com/farofadev/goesic/models"
	"github.com/farofadev/goesic/utils"
)

type PedidoRepository struct{}

func NewPedidoRepository() *PedidoRepository {
	return &PedidoRepository{}
}

func (*PedidoRepository) GetPaginatorConfig() *utils.PaginatorConfig {
	return utils.NewPaginatorConfig()
}

func (repo *PedidoRepository) FetchAll(a ...interface{}) (*[]models.Pedido, error) {
	pedidos := []models.Pedido{}

	paginatorParams, err := utils.GetPaginatorParams(repo.GetPaginatorConfig(), a...)

	if err != nil {
		return &pedidos, err
	}

	db, err := database.DBConnectDefault()

	if err != nil {
		return &pedidos, err
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

	db, err := database.DBConnectDefault()

	if err != nil {
		return pedido, err
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
	db, err := database.DBConnectDefault()

	if err != nil {
		return pedido, err
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
	_, err = db.Exec(rawSql, pedido.RowValues()...)

	return pedido, err
}

//Responder Salva a mensagem no banco de dados com seus respectivos dados e chama método que atualiza estado do pedido (aberto, respondido, negado)
func (repository *PedidoRepository) Responder(id string, tipo string) (interface{}, error) {

	log.Println("Estou no pedido repository", id)

	pedido, err := repository.FindById(id)

	if err != nil {
		return pedido, err
	}

	db, err := database.DBConnectDefault()

	if err != nil {
		return db, err
	}
	defer db.Close()

	updatePedido, err := db.Prepare("UPDATE pedidos SET situacao=? WHERE id=?")

	tx, err := db.Begin()

	res, err := tx.Stmt(updatePedido).Exec("respondido", id)

	log.Println("Aqui o pedido encontrado", pedido.Protocolo, err)

	return res, err

}
