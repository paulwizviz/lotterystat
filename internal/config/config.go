package config

import (
	"errors"
	"fmt"
	"os"
	"path"
	"runtime"

	"gopkg.in/yaml.v2"
)

const (
	SQLitePathKey = "sqlite.path"
	SQLiteFileKey = "sqlite.file"
)

var (
	ErrConfig = errors.New("config err")
)

// CreateIfNotExist config files if not exists
func CreateIfNotExist() error {

	configPath, err := Path()
	if err != nil {
		return err
	}

	err = createIfNotExistSettingDir(configPath)
	if errors.Is(err, ErrConfig) {
		return err
	}

	err = createIfNotExistSetting(configPath)
	if errors.Is(err, ErrConfig) {
		return err
	}

	return nil
}

// Path returns $HOME/.bz or %APPDATA%/ebz
func Path() (string, error) {
	var dir string

	switch runtime.GOOS {
	case "windows":
		dir = os.Getenv("AppData")
		if dir == "" {
			return "", fmt.Errorf("%w: APPDATA not set", ErrConfig)
		}
		dir = path.Join(dir, "ebz")
	default:
		dir = os.Getenv("HOME")
		if dir == "" {
			return "", fmt.Errorf("%w: $HOME not set", ErrConfig)
		}
		dir = path.Join(dir, ".ebz")
	}

	return dir, nil
}

// SQLiteSetting for sqlite db
type SQLiteSetting struct {
	Path string `yaml:"path"`
	File string `yaml:"file"`
}

// Setting data type representing app configuration
type Setting struct {
	SQLiteSetting `yaml:"sqlite"`
}

// createIfNotExistSettingDir a director name "$HOME/.ebz" or
// %APPDATA%/ebz if it does not exists
func createIfNotExistSettingDir(sPath string) error {

	_, err := os.Stat(sPath)
	if err == nil {
		return nil
	}

	err = os.MkdirAll(sPath, 0775)
	if err != nil {
		return fmt.Errorf("%w: %s", ErrConfig, err.Error())
	}

	return nil
}

// createIfNotExistSetting a file named "settings.yml" in the path
// $HOME/.ebz or %APPDATA%/ebz if it does not exists
func createIfNotExistSetting(sPath string) error {

	settingFile := path.Join(sPath, "settings.yaml")
	_, err := os.Stat(settingFile)
	if err == nil {
		return nil
	}

	f, err := os.Create(settingFile)
	if err != nil {
		return fmt.Errorf("%w: %s", ErrConfig, err.Error())
	}
	defer f.Close()

	p, err := Path()
	if err != nil {
		return err
	}

	s := Setting{
		SQLiteSetting: SQLiteSetting{
			Path: p,
			File: "data.db",
		},
	}

	b, err := yaml.Marshal(s)
	if err != nil {
		return fmt.Errorf("%w: %s", ErrConfig, err.Error())
	}

	_, err = f.Write(b)
	if err != nil {
		return fmt.Errorf("%w: %s", ErrConfig, err.Error())
	}

	err = os.Chmod(settingFile, 0666)
	if err != nil {
		return fmt.Errorf("%w: %s", ErrConfig, err.Error())
	}
	return nil
}
