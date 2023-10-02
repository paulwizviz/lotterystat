// Package config implements application configuration files
package config

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path"
	"runtime"
	"time"

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
	Type           = "yaml"
	Name           = "setting"
	MaxIdleTimeKey = "max_idle_time"
	MaxLifeTimeKey = "max_life_time"
	MaxIdleConnKey = "max_idle_conn"
	MaxOpenConnKey = "max_open_comn"
	DBConnKey      = "db_conn"
	DBConnVal      = "sqlite.db"
)

type Detail struct {
	Path        string        `yaml:"path"`
	Name        string        `yaml:"name"`
	DBConn      string        `yaml:"db_conn"`
	MaxIdleTime time.Duration `yaml:"max_idle_time"`
	MaxLifeTime time.Duration `yaml:"max_life_time"`
	MaxIdleConn int           `yaml:"max_idle_conn"`
	MaxOpenConn int           `yaml:"max_open_comn"`
}

func (d Detail) DBConnVal() string {
	return path.Join(d.Path, DBConnVal)
}

func (d Detail) ConfigFile() string {
	return path.Join(d.Path, fmt.Sprintf("%s.%s", d.Name, Type))
}

func Initilalize() error {
	p, err := Location()
	if err != nil {
		return fmt.Errorf("%w-%s", ErrConfig, err.Error())
	}
	d := Detail{
		Path:   p,
		Name:   Name,
		DBConn: DBConnVal,
	}
	cFile := d.ConfigFile()
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
	f, err := os.Create(d.ConfigFile())
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
	viper.SetConfigType(Type)
	viper.SetConfigName(Name)

	err := viper.ReadInConfig()
	if err != nil {
		return fmt.Errorf("%w-%s", ErrConfig, err.Error())
	}
	log.Println("Using config file:", viper.ConfigFileUsed())
	return nil
}

// Location returns $HOME/.bz or %APPDATA%/ebz
func Location() (string, error) {
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
