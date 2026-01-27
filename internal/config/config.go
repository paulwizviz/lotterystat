// Package config implements application configuration files
package config

import (
	"errors"
	"fmt"
	"os"
	"path"

	"github.com/paulwizviz/lotterystat/internal/sqlops"
)

var (
	ErrConfig = errors.New("config err")
)

var locationFunc = location

const (
	dbName = "lottery.db"
)

func Initilalize() error {
	p, err := locationFunc()
	if err != nil {
		return fmt.Errorf("%w:%w", ErrConfig, err)
	}
	_, err = os.Stat(p)
	if errors.Is(err, os.ErrNotExist) {
		err := os.MkdirAll(p, 0777)
		if err != nil {
			return fmt.Errorf("%w-%s", ErrConfig, err.Error())
		}
	}

	dbFile := path.Join(p, dbName)
	f, err := os.Create(dbFile)
	if err != nil {
		return fmt.Errorf("%w-%s", ErrConfig, err.Error())
	}
	f.Close()

	_, err = sqlops.NewSQLiteFile(dbFile)
	if err != nil {
		return err
	}
	return nil
}

// location returns $HOME/.ebz
func location() (string, error) {
	dir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("%w-unable to get home dir", ErrConfig)
	}
	return path.Join(dir, ".ebz"), nil
}
