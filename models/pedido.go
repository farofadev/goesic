package models

import (
	"database/sql"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/farofadev/goesic/utils"
	"github.com/google/uuid"
)

type Pedido struct {
	Id        string `json:"id"`
	Protocolo string `json:"protocolo"`
	PessoaId  string `json:"pessoa_id"`
	Situacao  string `json:"situacao"`
	CriadoEm  string `json:"criado_em"`
	DataPrazo string `json:"data_prazo"`
}

const (
	PedidoSituacaoAberto     = "aberto"
	PedidoSituacaoRespondido = "respondido"
	PedidoSituacaoNegado     = "negado"
)

func (pedido *Pedido) SqlColumns() []string {
	return []string{"id", "protocolo", "pessoa_id", "situacao", "criado_em", "data_prazo"}
}

func (pedido *Pedido) RowValues() []interface{} {
	return []interface{}{pedido.Id, pedido.Protocolo, pedido.PessoaId, pedido.Situacao, pedido.CriadoEm, pedido.DataPrazo}
}

func (pedido *Pedido) ScanFromSqlRows(rows *sql.Rows) error {
	return rows.Scan(&pedido.Id, &pedido.Protocolo, &pedido.PessoaId, &pedido.Situacao, &pedido.CriadoEm, &pedido.DataPrazo)
}

func (pedido *Pedido) SqlReplacementsString() string {
	columns := pedido.SqlColumns()

	replacements := make([]string, len(columns))

	for i := range replacements {
		replacements[i] = "?"
	}

	return strings.Join(replacements, ",")
}

func (pedido *Pedido) SqlColumnsString() string {
	return strings.Join(pedido.SqlColumns(), ", ")
}

func (pedido *Pedido) MakeId() {
	pedido.Id = uuid.NewString()
}

func (pedido *Pedido) MakeProtocol() {
	superrand := rand.Int63()
	pedido.Protocolo = time.Now().Format("060102") + fmt.Sprint(superrand)[:16]
}

func (pedido *Pedido) MakeCriadoEm() {
	pedido.CriadoEm = utils.FormatDateTimeString(time.Now())
}

func (pedido *Pedido) MakeDataPrazoSituacaoAberto() {
	datetime := utils.ParseDateTimeStringToTime(pedido.CriadoEm)

	pedido.DataPrazo = utils.FormatDateTimeString(datetime.Add(15 * 24 * time.Hour))
}
