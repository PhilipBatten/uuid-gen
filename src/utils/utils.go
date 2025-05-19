package utils

import "os/exec"

func ToClipboard(output []byte, arch string) error {
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
