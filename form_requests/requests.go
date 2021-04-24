package form_requests

import (
	"encoding/json"
	"net/http"
)

type FormRequest interface{}

func DecodeRequestBody(form FormRequest, req *http.Request) (FormRequest, error) {
	err := json.NewDecoder(req.Body).Decode(form)

	return form, err
}
