// Package ebzconfig implements application configuration files
package ebzconfig

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"path"

	"github.com/paulwizviz/lotterystat/internal/sqlops"
	"github.com/paulwizviz/lotterystat/internal/tball"
	"github.com/spf13/viper"
)

var (
	ErrConfig = errors.New("config err")
)

var locationFunc = location

const (
	dbName     = "lottery.db"
	configFile = "ebz.yaml"
)

// Configuration represents the application configuration
type Configuration struct {
	TballCache       string `mapstructure:"tball_cache"`
	EuromillionCache string `mapstructure:"euromillion_cache"`
	SflCache         string `mapstructure:"sfl_cache"`
	LottoCache       string `mapstructure:"lotto_cache"`
	DatabasePath     string `mapstructure:"database_path"`
}

// AppConfig is the global configuration instance
var AppConfig Configuration

// Initialize sets up the application configuration
func Initialize() error {
	appHome, err := locationFunc()
	if err != nil {
		return fmt.Errorf("%w: %w", ErrConfig, err)
	}

	if _, err := os.Stat(appHome); errors.Is(err, os.ErrNotExist) {
		if err := os.MkdirAll(appHome, 0777); err != nil {
			return fmt.Errorf("%w: %v", ErrConfig, err)
		}
	}

	viper.SetConfigName("ebz")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(appHome)

	// Set default values for cache paths
	viper.SetDefault("tball_cache", path.Join(appHome, "cache", "tball"))
	viper.SetDefault("euromillion_cache", path.Join(appHome, "cache", "euro"))
	viper.SetDefault("sfl_cache", path.Join(appHome, "cache", "sfl"))
	viper.SetDefault("lotto_cache", path.Join(appHome, "cache", "lotto"))
	viper.SetDefault("database_path", path.Join(appHome, dbName))

	if err := viper.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			if err := viper.SafeWriteConfig(); err != nil {
				return fmt.Errorf("%w: error writing default config: %v", ErrConfig, err)
			}
		} else {
			return fmt.Errorf("%w: error reading config file: %v", ErrConfig, err)
		}
	}

	if err := viper.Unmarshal(&AppConfig); err != nil {
		return fmt.Errorf("%w: unable to decode into struct: %v", ErrConfig, err)
	}

	// Ensure cache directories exist
	caches := []string{AppConfig.TballCache, AppConfig.EuromillionCache, AppConfig.SflCache, AppConfig.LottoCache}
	for _, c := range caches {
		if _, err := os.Stat(c); errors.Is(err, os.ErrNotExist) {
			if err := os.MkdirAll(c, 0700); err != nil {
				return fmt.Errorf("%w: error creating cache dir %s: %v", ErrConfig, c, err)
			}
		}
	}

	log.Printf("Thunderball cache is: %v", AppConfig.TballCache)
	log.Printf("Euro Million cache is: %v", AppConfig.EuromillionCache)
	log.Printf("Set For Life cache is: %v", AppConfig.SflCache)
	log.Printf("Lotto cache is: %v", AppConfig.LottoCache)

	if err := createDBTables(context.Background(), AppConfig.DatabasePath); err != nil {
		return fmt.Errorf("%w: %v", ErrConfig, err)
	}

	return nil
}

func createDBTables(ctx context.Context, dbFile string) error {
	db, err := sqlops.NewSQLiteFile(dbFile)
	if err != nil {
		return err
	}
	defer db.Close()
	if err := sqlops.CreateTables(ctx, db, tball.CreateTableFn); err != nil {
		return err
	}

	return nil
}

// location returns $HOME/.ebz
func location() (string, error) {
	dir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("%w: unable to get home dir", ErrConfig)
	}
	return path.Join(dir, ".ebz"), nil
}
