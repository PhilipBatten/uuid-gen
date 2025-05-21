package commands

import (
	"regexp"
	"testing"
)

func TestNewUUID4Command(t *testing.T) {
	cmd := NewUUID4Command()
	if cmd == nil {
		t.Error("NewUUID4Command returned nil")
	}
	if cmd.generator == nil {
		t.Error("UUID4Command generator is nil")
	}
}

func TestUUID4Command_Execute(t *testing.T) {
	cmd := NewUUID4Command()

	// Test that Execute returns a valid UUID v4
	uuid := cmd.Execute()

	// UUID v4 format: 8-4-4-4-12 hexadecimal digits
	uuidPattern := `^[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$`
	matched, err := regexp.MatchString(uuidPattern, uuid)
	if err != nil {
		t.Fatalf("Error matching UUID pattern: %v", err)
	}
	if !matched {
		t.Errorf("Generated UUID %s does not match UUID v4 format", uuid)
	}

	// Test that multiple executions generate different UUIDs
	uuid2 := cmd.Execute()
	if uuid == uuid2 {
		t.Error("Execute generated the same UUID twice")
	}
}
