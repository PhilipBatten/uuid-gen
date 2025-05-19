package main

import (
	"fmt"
	"uuid-gen/src/commands"
	"uuid-gen/src/utils"
)

// cli function to generate random uuid v4
func main() {
	id := commands.Run()

	// Copy to clipboard
	err := utils.ToClipboard([]byte(id), "linux")
	if err != nil {
		fmt.Println("Error copying to clipboard: ", err)
	} else {
		fmt.Println("Copied to clipboard")
	}
}
