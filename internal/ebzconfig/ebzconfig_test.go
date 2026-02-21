package ebzconfig

import (
	"os"
	"path"
	"testing"

	"github.com/spf13/viper"
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
		viper.Reset()
	})

	err = Initialize()
	if err != nil {
		t.Fatalf("Initialize failed: %v", err)
	}

	// Verify config file and directory
	configDir := path.Join(tmpDir, ".ebz")
	configFile := path.Join(configDir, "ebz.yaml")

	_, err = os.Stat(configFile)
	if err != nil {
		t.Fatalf("config file not created: %v", err)
	}

	// Verify database file is created
	_, err = os.Stat(AppConfig.DatabasePath)
	if err != nil {
		t.Fatalf("database file not created: %v", err)
	}

	// Verify default values in struct
	assert.Contains(t, AppConfig.TballCache, path.Join(configDir, "cache", "tball"))
	assert.Contains(t, AppConfig.EuromillionCache, path.Join(configDir, "cache", "euro"))
	assert.Contains(t, AppConfig.SflCache, path.Join(configDir, "cache", "sfl"))
	assert.Contains(t, AppConfig.LottoCache, path.Join(configDir, "cache", "lotto"))
	assert.Equal(t, path.Join(configDir, "lottery.db"), AppConfig.DatabasePath)
}
