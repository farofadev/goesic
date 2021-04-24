package responses

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
)

const (
	ContentTypeApplicationJson = "application/json"
	ContentTypeApplicationXml  = "application/xml"
	DefaultContentType         = ContentTypeApplicationJson
)

var ResponseMarshalers = map[string]interface{}{
	ContentTypeApplicationJson: json.Marshal,
	ContentTypeApplicationXml:  xml.Marshal,
}

type ResponseDataPayload struct {
	Data        interface{} `json:"data"`
	Message     string      `json:"message"`
	StatusCode  int         `json:"status_code"`
	Error       bool        `json:"error"`
	ContentType string      `json:"-" xml:"-"`
}

func (payload *ResponseDataPayload) Send(res http.ResponseWriter) {
	if payload.StatusCode == 0 {
		payload.StatusCode = http.StatusOK
	}

	if payload.Message == "" {
		payload.Message = http.StatusText(payload.StatusCode)
	}

	if payload.ContentType == "" {
		payload.ContentType = DefaultContentType
	}

	if payload.StatusCode >= 400 && payload.StatusCode < 600 {
		payload.Error = true
	}

	bytes := payload.Marshal()

	res.Header().Set("Content-Type", payload.ContentType)
	res.WriteHeader(payload.StatusCode)
	res.Write(bytes)
}

func (payload *ResponseDataPayload) SetStatusCode(code int) *ResponseDataPayload {
	payload.StatusCode = code

	return payload
}

func (payload *ResponseDataPayload) Marshal() []byte {
	contentType := payload.ContentType

	if contentType == "" {
		contentType = DefaultContentType
	}

	marshal, ok := ResponseMarshalers[contentType]

	if !ok {
		return []byte("")
	}

	data, _ := marshal.(func(interface{}) ([]byte, error))(payload)

	return data
}

func NewResponseDataPayload() *ResponseDataPayload {
	return &ResponseDataPayload{}
}

func SendResponseInternalServerError(res http.ResponseWriter) {
	payload := ResponseDataPayload{
		StatusCode: http.StatusInternalServerError,
	}

	payload.Send(res)
}

func SendResponseNotFound(res http.ResponseWriter) {
	payload := ResponseDataPayload{
		StatusCode: http.StatusNotFound,
	}

	payload.Send(res)
}
