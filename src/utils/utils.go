package utils

import (
	"fmt"
	"os/exec"
	"runtime"
)

// DetectPlatform returns the current platform (darwin or linux)
func DetectPlatform() string {
	return runtime.GOOS
}

func ToClipboard(output []byte) error {
	var copyCmd *exec.Cmd
	arch := DetectPlatform()

	// Mac "OS"
	if arch == "darwin" {
		copyCmd = exec.Command("pbcopy")
	} else if arch == "linux" {
		// Linux
		copyCmd = exec.Command("xclip", "-selection", "c")
	} else {
		return fmt.Errorf("unsupported platform: %s. Only darwin and linux are supported", arch)
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
