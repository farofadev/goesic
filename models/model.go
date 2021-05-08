package models

import (
	"time"

	"github.com/farofadev/goesic/utils"
	"github.com/google/uuid"
)

type Model interface {
	SetId() string
	SetCriadoEm() string
}

type ModelImpl struct {
	Id       string `json:"id"`
	CriadoEm string `json:"criado_em"`
}

func (m *ModelImpl) SetId() {
	m.Id = uuid.NewString()

}

func (m *ModelImpl) SetCriadoEm() {
	m.CriadoEm = utils.FormatDateTimeString(time.Now())
}
