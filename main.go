package main

import (
	"fmt"
	"uuid-gen/src/commands"
	"uuid-gen/src/utils"
)

// cli function to generate random uuid v4
func main() {
	command := commands.NewUUID4Command()
	id := command.Execute()

	fmt.Println(id)

	// Copy to clipboard
	err := utils.ToClipboard([]byte(id))
	if err != nil {
		fmt.Println("Error copying to clipboard: ", err)
	} else {
		fmt.Println("Copied to clipboard")
	}
}
