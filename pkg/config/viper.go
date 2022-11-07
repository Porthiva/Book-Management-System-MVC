package config

import (
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type Env struct {
	DB_NAME     string `mapstructure:"DB_NAME"`
	DB_USERNAME string `mapstructure:"DB_USERNAME"`
	DB_PASSWORD string `mapstructure:"DB_PASSWORD"`
	DB_HOST     string `mapstructure:"DB_HOST"`
	DB_PORT     string `mapstructure:"DB_PORT"`
	PORT        string
}

var ENV Env
var WORKING_DIR string

func LoadEnv() {
	if err := tryLoadFromENVFile(); err != nil {
		tryLoadFromOSENV()
	}

	ENV.PORT = os.Getenv("PORT")
	if ENV.PORT == "" {
		ENV.PORT = "9010"
	}

	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	WORKING_DIR = pwd
}

func tryLoadFromENVFile() error {
	envFile := filepath.Join(".env")
	viper.SetConfigFile(envFile)

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	if err := viper.Unmarshal(&ENV); err != nil {
		return err
	}

	return nil
}

func tryLoadFromOSENV() {
	ENV.DB_NAME = os.Getenv("DB_NAME")
	ENV.DB_USERNAME = os.Getenv("DB_USERNAME")
	ENV.DB_PASSWORD = os.Getenv("DB_PASSWORD")
	ENV.DB_HOST = os.Getenv("DB_HOST")
	ENV.DB_PORT = os.Getenv("DB_PORT")
}
