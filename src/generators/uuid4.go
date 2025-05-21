package generators

import (
	"github.com/google/uuid"
)

type UUIDGenerator interface {
	Generate() string
}

type UUID4Generator struct{}

func NewUUID4Generator() *UUID4Generator {
	return &UUID4Generator{}
}

func (g *UUID4Generator) Generate() string {
	id := uuid.New()
	return id.String()
}
