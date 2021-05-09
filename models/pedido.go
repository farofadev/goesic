package models

import (
	"database/sql"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/farofadev/goesic/utils"
)

type Pedido struct {
	ModelImpl
	Protocolo string `json:"protocolo"`
	PessoaId  string `json:"pessoa_id"`
	Situacao  string `json:"situacao"`
	DataPrazo string `json:"data_prazo"`
}

const (
	PedidoSituacaoAberto     = "aberto"
	PedidoSituacaoRespondido = "respondido"
	PedidoSituacaoNegado     = "negado"
)

const PedidoPrazoInicial = 15 * 24 * time.Hour

func NewPedido() *Pedido {
	return &Pedido{}
}

func (pedido *Pedido) SqlColumns() []string {
	return []string{"id", "protocolo", "pessoa_id", "situacao", "criado_em", "data_prazo"}
}

func (pedido *Pedido) RowValues() []interface{} {
	return []interface{}{&pedido.Id, &pedido.Protocolo, &pedido.PessoaId, &pedido.Situacao, &pedido.CriadoEm, &pedido.DataPrazo}
}

func (pedido *Pedido) ScanFromSqlRows(rows *sql.Rows) error {
	return rows.Scan(pedido.RowValues()...)

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

func (pedido *Pedido) MakeProtocol() {
	superrand := rand.Int63()
	pedido.Protocolo = time.Now().Format("060102") + fmt.Sprint(superrand)[:16]
}

func (pedido *Pedido) MakeDataPrazoSituacaoAberto() {
	datetime := utils.ParseDateTimeStringToTime(pedido.CriadoEm)

	pedido.DataPrazo = utils.FormatDateTimeString(datetime.Add(PedidoPrazoInicial))
}

func GetCacheKeyForPedidoId(id string) string {
	return "pedido:" + id
}

func GetCacheKeyForPedidosPaginator(params *utils.PaginatorParams) string {
	return fmt.Sprintf("pedidos:%s", params.GetCacheKey())
}

func (pedido *Pedido) GetCacheKeyForId() string {
	return GetCacheKeyForPedidoId(pedido.Id)
}
