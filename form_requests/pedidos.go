package form_requests

type PedidoFormRequest struct {
	FormRequest
	PessoaId string `json:"pessoa_id"`
}
