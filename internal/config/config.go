// Package config implements application configuration files
package config

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path"
	"runtime"

	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

const (
	SQLiteFile      = "sqlite.db"
	SettingFileName = "settings.yaml"
)

var (
	ErrConfig = errors.New("config err")
)

const (
	configType = "yaml"
	configName = "setting"
	dbName     = "data.db"
)

func Path() string {
	return viper.GetString("path")
}

func Name() string {
	return viper.GetString("name")
}

func DBConn() string {
	return viper.GetString("db_conn")
}

type Detail struct {
	Path   string `yaml:"path"`
	Name   string `yaml:"name"`
	DBConn string `yaml:"db_conn"`
}

func Initilalize() error {
	p, err := location()
	if err != nil {
		return fmt.Errorf("%w-%s", ErrConfig, err.Error())
	}
	d := Detail{
		Path:   p,
		Name:   configName,
		DBConn: fmt.Sprintf("file:%s%s?cache=shared&mode=memory", p, dbName),
	}
	cFile := path.Join(p, fmt.Sprintf("%s.%s", d.Name, configType))
	_, err = os.Stat(cFile)
	if errors.Is(err, os.ErrExist) {
		err := initViper(p)
		if err != nil {
			return err
		}
		return nil
	}
	_, err = os.Stat(p)
	if errors.Is(err, os.ErrNotExist) {
		err := os.MkdirAll(p, 0777)
		if err != nil {
			return fmt.Errorf("%w-%s", ErrConfig, err.Error())
		}
	}
	b, err := yaml.Marshal(d)
	if err != nil {
		return fmt.Errorf("%w-%s", ErrConfig, err.Error())
	}
	f, err := os.Create(cFile)
	if err != nil {
		return fmt.Errorf("%w-%s", ErrConfig, err.Error())
	}
	defer f.Close()
	_, err = f.Write(b)
	if err != nil {
		return fmt.Errorf("%w-%s", ErrConfig, err.Error())
	}
	err = initViper(p)
	if err != nil {
		return err
	}
	return nil
}

func initViper(p string) error {
	viper.AddConfigPath(p)
	viper.SetConfigType(configType)
	viper.SetConfigName(configName)
	err := viper.ReadInConfig()
	if err != nil {
		return fmt.Errorf("%w-%s", ErrConfig, err.Error())
	}
	log.Printf("%s used", viper.ConfigFileUsed())
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
