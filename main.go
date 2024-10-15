package main

import (
	"fmt"
	"os/exec"

	"github.com/google/uuid"
)

// cli function to generate random uuid v4
func main() {
	id := uuid.New()
	fmt.Println(id.String())

	// Copy to clipboard
	err := toClipboard([]byte(id.String()), "linux")
	if err != nil {
		fmt.Println("Error copying to clipboard: ", err)
	} else {
		fmt.Println("Copied to clipboard")
	}
}

func toClipboard(output []byte, arch string) error {
	var copyCmd *exec.Cmd

	// Mac "OS"
	if arch == "darwin" {
		copyCmd = exec.Command("pbcopy")
	}
	// Linux
	if arch == "linux" {
		copyCmd = exec.Command("xclip", "-selection", "c")
	}

	in, err := copyCmd.StdinPipe()

	if err != nil {
		return err
	}

	if err := copyCmd.Start(); err != nil {
		return err
	}

	if _, err := in.Write([]byte(output)); err != nil {
		return err
	}

	if err := in.Close(); err != nil {
		return err
	}

	return copyCmd.Wait()
}
