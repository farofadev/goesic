package models

import (
	"database/sql"
	"strings"
)

type Mensagem struct {
	ModelImpl
	PedidoId       string `json:"pedido_id"`
	PessoaId       string `json:"pessoa_id"`
	UsuarioId      string `json:"usuario_id"`
	Mensagem       string `json:"mensagem"`
	SituacaoPedido string `json:"situacao_pedido"`
	Tipo           string `json:"tipo"`
}

func NewMensagem() *Mensagem {
	return &Mensagem{}
}

func (mensagem *Mensagem) SqlColumns() []string {

	return []string{"id", "pedido_id", "pessoa_id", "usuario_id", "mensagem", "situacao_pedido", "tipo", "criado_em"}
}

func (mensagem *Mensagem) RowValues() []interface{} {

	return []interface{}{&mensagem.Id, &mensagem.PedidoId, &mensagem.PessoaId, &mensagem.UsuarioId, &mensagem.Mensagem, &mensagem.Tipo}
}

func (mensagem *Mensagem) ScanFromSqlRows(rows *sql.Rows) error {

	return rows.Scan(mensagem.RowValues()...)

}

func (mensagem *Mensagem) SqlReplacementsString() string {
	columns := mensagem.SqlColumns()

	replacements := make([]string, len(columns))

	for i := range replacements {
		replacements[i] = "?"
	}

	return strings.Join(replacements, "")

}

func (mensagem *Mensagem) SqlColumnsString() string {

	return strings.Join(mensagem.SqlColumns(), ",")
}
