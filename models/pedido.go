package models

type Pedido struct {
	Id string `json:"id"`
	PessoaId string `json:"pessoa_id"`
	Situacao string `json:"situacao"`
	CriadoEm string `json:"criado_em"`
	DataPrazo string `json:"data_prazo"`
}
