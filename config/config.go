package config

import (
	"errors"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
)

type Config struct {
	Tg `json:"tg" yaml:"tg"`
	Db `json:"db" yaml:"db"`
}
type Tg struct {
	ApiToken string ` json:"apitoken "yaml:"apitoken" env-default:"apitoken"`
}
type Db struct {
	Username string `yaml:"username" env-default:"postgres"`
	Password string `yaml:"password" env-default:"postgres"`
	Host     string `yaml:"host" env-default:"localhost"`
	Port     string `yaml:"port" env-default:"5432"`
	Name     string `yaml:"name" env-default:"postgres"`
}

func Load() (*Config, error) {
	configPath := "/Users/vladislavtrofimov/GolandProjects/TgDbMai/config/config.yaml"
	if configPath == "" {
		return nil, errors.New("CONFIG_PATH is not set")
	}
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return nil, errors.New("config file does not exist")

	}
	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		return nil, errors.New("cannot read config file")
	}
	return &cfg, nil
}
