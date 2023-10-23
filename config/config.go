package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
	"os"
)

type (
	// Config -.
	Config struct {
		App `yaml:"app"`
		Log `yaml:"logger"`
		PG  `yaml:"postgres"`
		Tg  `yaml:"tg"`
	}
	Tg struct {
		Apitoken string `env-required:"true" yaml:"apitoken"    env:"API_TOKEN"`
	}
	App struct {
		Name    string `env-required:"true" yaml:"name"    env:"APP_NAME"`
		Version string `env-required:"true" yaml:"version" env:"APP_VERSION"`
	}

	Log struct {
		Level string `env-required:"true" yaml:"log_level"   env:"LOG_LEVEL"`
	}

	PG struct {
		PoolMax  int    `env-required:"true" yaml:"pool_max" env:"PG_POOL_MAX"`
		Username string `env-required:"true" yaml:"username" env-default:"postgres"`
		Password string `env-required:"true" yaml:"password" env-default:"postgres"`
		Host     string `env-required:"true" yaml:"host" env-default:"localhost"`
		Port     string `env-required:"true" yaml:"port" env-default:"5432"`
		Name     string `env-required:"true" yaml:"name" env-default:"postgres"`
	}
)

func New() (*Config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}
	currPath, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	filePath := currPath + os.Getenv("path") + "/" + os.Getenv("env") + ".yml"
	fmt.Println(filePath)
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	d := yaml.NewDecoder(file)
	var cfg *Config

	if err := d.Decode(&cfg); err != nil {
		return nil, err
	}

	return cfg, nil

}
