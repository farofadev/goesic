package repositories

import (
	"fmt"
	"log"

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

	// 	ModelImpl
	// 	PedidoId       string `json:"pedido_id"`
	// 	PessoaId       string `json:"pessoa_id"`
	// 	UsuarioId      string `json:"usuario_id"`
	// 	Mensagem       string `json:"mensagem"`
	// 	SituacaoPedido string `json:"situacao_pedido"`
	// 	Tipo           string `json:"tipo"`
	// }

	mensagem.SetId()
	mensagem.SetCriadoEm()

	log.Println("collumns string: ", mensagem.SqlColumnsString())

	log.Println("collumns string: ", mensagem.SqlReplacementsString())

	log.Println("values: ", mensagem.RowValues())

	rawSql := fmt.Sprintf("INSERT INTO mensagens (%s) VALUES (%s);", mensagem.SqlColumnsString(), mensagem.SqlReplacementsString())
	_, err = db.Exec(rawSql, mensagem.RowValues()...)

	return mensagem, err

}
