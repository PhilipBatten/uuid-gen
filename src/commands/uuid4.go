package commands

import (
	"uuid-gen/src/generators"
)

type Command interface {
	Execute() string
}

type UUID4Command struct {
	generator generators.UUIDGenerator
}

func NewUUID4Command() *UUID4Command {
	return &UUID4Command{
		generator: generators.NewUUID4Generator(),
	}
}

func (c *UUID4Command) Execute() string {
	return c.generator.Generate()
}
