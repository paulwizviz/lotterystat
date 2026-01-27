package config

import (
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLocation(t *testing.T) {
	home, err := os.UserHomeDir()
	if err != nil {
		t.Fatalf("Failed to get user home dir: %v", err)
	}
	expected := path.Join(home, ".ebz")
	actual, err := location()
	if err != nil {
		t.Fatalf("location failed: %v", err)
	}
	assert.Equal(t, expected, actual)
}
func TestInitialize(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "test-config")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Set environment variable to control config location
	oldLocationFunc := locationFunc
	locationFunc = func() (string, error) {
		return path.Join(tmpDir, ".ebz"), nil
	}
	t.Cleanup(func() {
		locationFunc = oldLocationFunc
	})

	err = Initilalize()
	if err != nil {
		t.Fatalf("Initilalize failed: %v", err)
	}

	// Verify config file and directory
	configDir := path.Join(tmpDir, ".ebz")

	dbFile := path.Join(configDir, "lottery.db")
	_, err = os.Stat(dbFile)
	if err != nil {
		t.Fatalf("db file not created: %v", err)
	}

}
