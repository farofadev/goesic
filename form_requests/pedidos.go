package form_requests

type PedidoFormRequest struct {
	PessoaId string `json:"pessoa_id"`
}

func NewPedidoFormRequest() *PedidoFormRequest {
	return &PedidoFormRequest{}
}
