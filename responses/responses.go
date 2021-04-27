package responses

import (
	"encoding/json"
	"encoding/xml"
	"net/http"

	"github.com/gofiber/fiber/v2"
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
	ContentType string      `json:"-" xml:"-"`
	Data        interface{} `json:"data"`
	Error       bool        `json:"error"`
	StatusCode  int         `json:"status_code"`
	StatusText  string      `json:"status_text"`
}

func (payload *ResponseDataPayload) Send(ctx *fiber.Ctx) error {
	if payload.StatusCode == 0 {
		payload.StatusCode = http.StatusOK
	}

	if payload.StatusText == "" {
		payload.StatusText = http.StatusText(payload.StatusCode)
	}

	if payload.ContentType == "" {
		payload.ContentType = DefaultContentType
	}

	if payload.StatusCode >= 400 && payload.StatusCode < 600 {
		payload.Error = true
	}

	bytes := payload.Marshal()

	ctx.Set("Content-Type", payload.ContentType)

	return ctx.Send(bytes)
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

func SendResponseInternalServerError(ctx *fiber.Ctx) error {
	payload := NewResponseDataPayload()
	payload.StatusCode = http.StatusInternalServerError

	return payload.Send(ctx)
}

func SendResponseNotFound(ctx *fiber.Ctx) error {
	payload := NewResponseDataPayload()
	payload.StatusCode = http.StatusNotFound

	return payload.Send(ctx)
}
