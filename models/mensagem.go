package models

type mensagem struct {
	ModelImpl
	PedidoId       string `json:"pedido_id"`
	PessoaId       string `json:"pessoa_id"`
	UsuarioId      string `json:"usuario_id"`
	Mensagem       string `json:"mensagem"`
	SituacaoPedido string `json:"situacao_pedido"`
	Tipo           string `json:"tipo"`
}

func newMensagem() *mensagem {
	return &mensagem{}
}
