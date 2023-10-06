package config

import (
	"os"
	"path"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLocation(t *testing.T) {
	actual, err := location()
	if err != nil {
		t.Fatal("error not expected")
	}
	switch runtime.GOOS {
	case "windows":
		uprofile := os.Getenv("USERPROFILE")
		expected := path.Join(uprofile, "AppData", "Roaming", "ebz")
		assert.Equal(t, expected, actual, "ebzcli config path for windows")
	case "linux", "darwin":
		homePath := os.Getenv("HOME")
		expected := path.Join(homePath, ".ebz")
		assert.Equal(t, expected, actual, "ebzcli config path for macOS and Linux")
	default:
		t.Fatal("this should not happen")
	}
}
