package commands

import (
	"fmt"

	"github.com/google/uuid"
)

func Run() string {
	id := uuid.New()
	fmt.Println(id.String())
	return id.String()
}
