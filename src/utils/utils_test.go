package utils

import (
	"os/exec"
	"runtime"
	"testing"
)

// checkXclipAvailable checks if xclip is available on the system
func checkXclipAvailable() bool {
	cmd := exec.Command("which", "xclip")
	return cmd.Run() == nil
}

func TestDetectPlatform(t *testing.T) {
	platform := DetectPlatform()
	if platform != "darwin" && platform != "linux" {
		t.Errorf("DetectPlatform() returned unexpected platform: %s", platform)
	}
}

func TestToClipboard(t *testing.T) {
	// Skip test on Linux if xclip is not available
	if runtime.GOOS == "linux" && !checkXclipAvailable() {
		t.Skip("Skipping test: xclip is not available")
	}

	tests := []struct {
		name    string
		input   []byte
		wantErr bool
	}{
		{
			name:    "valid input",
			input:   []byte("test data"),
			wantErr: false,
		},
		{
			name:    "empty input",
			input:   []byte(""),
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ToClipboard(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToClipboard() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// TestToClipboardUnsupportedPlatform tests the behavior on unsupported platforms
// This test is skipped by default as it requires mocking the platform
func TestToClipboardUnsupportedPlatform(t *testing.T) {
	t.Skip("This test requires platform mocking to be implemented")

	// This test would verify that ToClipboard returns an error
	// when running on an unsupported platform
	// Implementation would require mocking runtime.GOOS
}
