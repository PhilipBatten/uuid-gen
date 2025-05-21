package generators

import (
	"testing"

	"github.com/google/uuid"
)

func TestUUID4Generator_Generate(t *testing.T) {
	generator := NewUUID4Generator()

	// Test that we can generate a UUID
	uuidStr := generator.Generate()
	if uuidStr == "" {
		t.Error("Generated UUID should not be empty")
	}

	// Test that the generated string is a valid UUID
	parsedUUID, err := uuid.Parse(uuidStr)
	if err != nil {
		t.Errorf("Generated string is not a valid UUID: %v", err)
	}

	// Test that it's a v4 UUID (random)
	if parsedUUID.Version() != 4 {
		t.Errorf("Expected UUID version 4, got version %d", parsedUUID.Version())
	}

	// Test that multiple generations produce different UUIDs
	uuidSet := make(map[string]bool)
	for i := 0; i < 1000; i++ {
		newUUID := generator.Generate()
		if uuidSet[newUUID] {
			t.Errorf("Generated duplicate UUID: %s", newUUID)
		}
		uuidSet[newUUID] = true
	}
}
