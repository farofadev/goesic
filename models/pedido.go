package models

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/google/uuid"
)
type Pedido struct {
	Id string `json:"id"`
	Protocolo string `json:"protocolo"`
	PessoaId string `json:"pessoa_id"`
	Situacao string `json:"situacao"`
	CriadoEm string `json:"criado_em"`
	DataPrazo string `json:"data_prazo"`
}

func (pedido *Pedido) MakeId() {
	pedido.Id = uuid.NewString()
}

func (pedido *Pedido) MakeProtocol() {
	protocol := time.Now().Format("060102")

	superrand := rand.Int63()
	protocol = protocol + fmt.Sprint(superrand)[:16]

	pedido.Protocolo = protocol
}
