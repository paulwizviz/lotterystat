// Package config implements application configuration files
package config

import (
	"errors"
	"fmt"
	"os"
	"path"
	"runtime"
	"time"

	"gopkg.in/yaml.v2"
)

const (
	SQLiteFile      = "sqlite.db"
	SettingFileName = "settings.yaml"
)

var (
	ErrConfig = errors.New("config err")
)

type DBConfig struct {
	SQLiteDB        string
	ConnMaxIdleTime time.Duration
	ConnMaxLifeTime time.Duration
	MaxIdleConn     int
	MaxOpenConn     int
}

type Detail struct {
	DBConfig DBConfig `yaml:"dbconfig"`
	Path     string   `yaml:"path"`
}

// Initialize app config
func Initialize() (*Detail, error) {
	p, err := location()
	if err != nil {
		return nil, fmt.Errorf("%w-%s", ErrConfig, err.Error())
	}
	settingFile := path.Join(p, SettingFileName)

	_, err = os.Stat(settingFile)
	if errors.Is(err, os.ErrExist) {
		d, err := unmarshalSettingFile(settingFile)
		if err != nil {
			return nil, fmt.Errorf("%w-%s", ErrConfig, err.Error())
		}
		return d, nil
	}

	_, err = os.Stat(p)
	if errors.Is(err, os.ErrNotExist) {
		err = os.MkdirAll(p, 0775)
		if err != nil {
			return nil, fmt.Errorf("%w-%s", ErrConfig, err.Error())
		}
	}

	settingDetail := Detail{
		DBConfig: DBConfig{
			SQLiteDB: path.Join(p, "data.db"),
		},
		Path: p,
	}

	b, err := yaml.Marshal(settingDetail)
	if err != nil {
		return nil, fmt.Errorf("%w-%s", ErrConfig, err.Error())
	}
	f, err := os.Create(settingFile)
	if err != nil {
		return nil, fmt.Errorf("%w-%s", ErrConfig, err.Error())
	}
	defer f.Close()
	_, err = f.Write(b)
	if err != nil {
		return nil, fmt.Errorf("%w-%s", ErrConfig, err.Error())
	}
	err = os.Chmod(settingFile, 0666)
	if err != nil {
		return nil, fmt.Errorf("%w-%s", ErrConfig, err.Error())
	}
	return &settingDetail, nil
}

func unmarshalSettingFile(f string) (*Detail, error) {
	content, err := os.ReadFile(f)
	if err != nil {
		return nil, fmt.Errorf("%w-%s", ErrConfig, err.Error())
	}
	var setting Detail
	err = yaml.Unmarshal(content, &setting)
	if err != nil {
		return nil, fmt.Errorf("%w-%s", ErrConfig, err.Error())
	}
	return &setting, nil
}

// createIfNotExistDir a director name "$HOME/.ebz" or
// %APPDATA%/ebz if it does not exists
func createIfNotExistDir(p string) error {

	_, err := os.Stat(p)
	if err == nil {
		return nil
	}

	err = os.MkdirAll(p, 0775)
	if err != nil {
		return fmt.Errorf("%w-%s", ErrConfig, err.Error())
	}

	return nil
}

// createIfNotExistFile a file named "settings.yml" in the path
// $HOME/.ebz or %APPDATA%/ebz if it does not exists
func createIfNotExistFile(cPath string) error {
	settingFile := path.Join(cPath, SettingFileName)
	_, err := os.Stat(settingFile)
	if err == nil {
		return nil
	}
	sqliteDB := path.Join(cPath, SQLiteFile)
	d := DBConfig{
		SQLiteDB: sqliteDB,
	}
	f, err := os.Create(settingFile)
	if err != nil {
		return fmt.Errorf("%w-%s", ErrConfig, err.Error())
	}
	defer f.Close()
	s := Detail{
		DBConfig: d,
		Path:     cPath,
	}
	b, err := yaml.Marshal(s)
	if err != nil {
		return fmt.Errorf("%w-%s", ErrConfig, err.Error())
	}
	_, err = f.Write(b)
	if err != nil {
		return fmt.Errorf("%w-%s", ErrConfig, err.Error())
	}
	err = os.Chmod(settingFile, 0666)
	if err != nil {
		return fmt.Errorf("%w-%s", ErrConfig, err.Error())
	}
	return nil
}

// location returns $HOME/.bz or %APPDATA%/ebz
func location() (string, error) {
	var dir string
	switch runtime.GOOS {
	case "windows":
		dir = os.Getenv("AppData")
		if dir == "" {
			return "", fmt.Errorf("%w-APPDATA not set", ErrConfig)
		}
		dir = path.Join(dir, "ebz")
	default:
		dir = os.Getenv("HOME")
		if dir == "" {
			return "", fmt.Errorf("%w-$HOME not set", ErrConfig)
		}
		dir = path.Join(dir, ".ebz")
	}
	return dir, nil
}
