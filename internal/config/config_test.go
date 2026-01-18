package config

import (
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLocation(t *testing.T) {
	testCases := []struct {
		name     string
		goos     string
		setup    func(t *testing.T)
		expected func(t *testing.T) string
	}{
		{
			name: "Windows",
			goos: "windows",
			setup: func(t *testing.T) {
				t.Setenv("APPDATA", "C:\\Users\\test\\AppData\\Roaming")
			},
			expected: func(t *testing.T) string {
				return path.Join("C:\\Users\\test\\AppData\\Roaming", "ebz")
			},
		},
		{
			name: "Linux",
			goos: "linux",
			setup: func(t *testing.T) {
				// No setup needed for linux
			},
			expected: func(t *testing.T) string {
				home, err := os.UserHomeDir()
				if err != nil {
					t.Fatalf("Failed to get user home dir: %v", err)
				}
				return path.Join(home, ".config", "ebz")
			},
		},
		{
			name: "Darwin",
			goos: "darwin",
			setup: func(t *testing.T) {
				// No setup needed for darwin
			},
			expected: func(t *testing.T) string {
				home, err := os.UserHomeDir()
				if err != nil {
					t.Fatalf("Failed to get user home dir: %v", err)
				}
				return path.Join(home, ".config", "ebz")
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.setup(t)
			actual, err := locationForOS(tc.goos)
			if err != nil {
				t.Fatalf("locationForOS failed: %v", err)
			}
			assert.Equal(t, tc.expected(t), actual)
		})
	}
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
		return path.Join(tmpDir, "ebz"), nil
	}
	t.Cleanup(func() {
		locationFunc = oldLocationFunc
	})

	err = Initilalize()
	if err != nil {
		t.Fatalf("Initilalize failed: %v", err)
	}

	// Verify config file and directory
	configDir := path.Join(tmpDir, "ebz")

	configFile := path.Join(configDir, "setting.yaml")
	_, err = os.Stat(configFile)
	if err != nil {
		t.Fatalf("Config file not created: %v", err)
	}

	// Verify config content
	assert.Equal(t, configDir, Path(), "Path should be correct")
	assert.Equal(t, "setting", Name(), "Name should be correct")
	assert.Equal(t, "file:"+path.Join(configDir, "data.db"), DBConn(), "DBConn should be correct")
}
