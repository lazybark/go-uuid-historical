package uuidv7

import (
	"strings"
	"testing"
	"time"
)

func TestGenerateUUIDv7(t *testing.T) {
	// Test with current time.
	uuid, err := GenerateUUIDv7()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(uuid) != 36 {
		t.Fatalf("Expected UUID length of 36, got %d", len(uuid))
	}

	if uuid[14] != '7' {
		t.Fatalf("Expected version 7, got %c", uuid[14])
	}

	// Test with specific timestamp.
	layout := "2006-01-02 15:04:05.000 -0700"

	ts, err := time.Parse(layout, "2024-10-25 08:58:09.662 +0000")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	uuid, err = GenerateUUIDv7(ts)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(uuid) != 36 {
		t.Fatalf("Expected UUID length of 36, got %d", len(uuid))
	}

	if uuid[14] != '7' {
		t.Fatalf("Expected version 7, got %c", uuid[14])
	}

	// Check if the timestamp part is correctly encoded.
	expectedPrefix := "0192c2e5-bf7e-7000-"
	if !strings.HasPrefix(uuid, expectedPrefix) {
		t.Fatalf("Expected UUID to start with %s, got %s", expectedPrefix, uuid)
	}
}
