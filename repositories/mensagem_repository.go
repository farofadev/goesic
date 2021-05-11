package repositories

import (
	"fmt"

	"github.com/farofadev/goesic/database"
	"github.com/farofadev/goesic/models"
)

type MensagemRepository struct{}

func NewMensagemRepository() *MensagemRepository {
	return &MensagemRepository{}
}

func (*MensagemRepository) Store(mensagem *models.Mensagem) (*models.Mensagem, error) {
	db, err := database.DBConnectDefault()

	if err != nil {
		return mensagem, err
	}
	defer db.Close()

	mensagem.SetId()
	mensagem.SetCriadoEm()

	switch mensagem.Tipo {
	case "resposta_pedido":
		mensagem.SituacaoPedido = models.PedidoSituacaoRespondido
	case "negado_pedido":
		mensagem.SituacaoPedido = models.PedidoSituacaoNegado
	default:
		mensagem.SituacaoPedido = models.PedidoSituacaoAberto

	}

	rawSql := fmt.Sprintf("INSERT INTO mensagens (%s) VALUES (%s);", mensagem.SqlColumnsString(), mensagem.SqlReplacementsString())
	_, err = db.Exec(rawSql, mensagem.RowValues()...)

	return mensagem, err

}
