package api

import (
	"os"
	"testing"
)

func TestReadFile(t *testing.T) {
	// Create a temporary test file
	content := "line1\nline2\nline3\n"
	tmpFile, err := os.CreateTemp("", "testfile")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	// Write content to the temp file
	if _, err := tmpFile.WriteString(content); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	tmpFile.Close()

	// Test ReadFile with the temp file
	result, err := ReadFile(tmpFile.Name())
	if err != nil {
		t.Fatalf("ReadFile failed: %v", err)
	}

	// Verify the content
	expected := content
	if result != expected {
		t.Errorf("ReadFile returned %q, expected %q", result, expected)
	}
}

func TestReadFile_EmptyFile(t *testing.T) {
	// Create an empty temporary test file
	tmpFile, err := os.CreateTemp("", "testfile")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())
	tmpFile.Close()

	// Test ReadFile with the empty file
	result, err := ReadFile(tmpFile.Name())
	if err != nil {
		t.Fatalf("ReadFile failed: %v", err)
	}

	// Verify the content
	expected := ""
	if result != expected {
		t.Errorf("ReadFile returned %q, expected %q", result, expected)
	}
}

func TestReadFile_NonexistentFile(t *testing.T) {
	// Test ReadFile with a nonexistent file
	_, err := ReadFile("nonexistentfile.txt")
	if err == nil {
		t.Error("ReadFile should return an error for a nonexistent file")
	}
}
