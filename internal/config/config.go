package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	DB         `yaml:"db"`
	HTTPServer `yaml:"http_server"`
}

type DB struct {
	Username string `yaml:"username" env-default:"mobile"`
	Host     string `yaml:"host" env-default:"localhost"`
	Port     string `yaml:"port" env-default:"5040"`
	DBName   string `yaml:"dbname" env-default:"profilesdb"`
	Password string `yaml:"password" env-default:"password"`
	SSLMode  string `yaml:"sslmode" env-default:"disable"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env-default:"localhost:8080"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
	CtxTimeout  time.Duration `yaml:"ctx_timeout" env-default:"60s"`
}

// export CONFIG_PATH="configs/config.yaml

func MustLoad() *Config {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("CONFIG_PATH is not set")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exists: %s", configPath)
	}

	var cfg Config
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	return &cfg
}
