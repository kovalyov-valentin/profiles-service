package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
	"time"
)

const (
	envPrefix   = ""
	envFilePath = "configs/.env"
)

type Config struct {
	DB
	HTTPServer
	Api
}

type Api struct {
	AgeUrl         string `envconfig:"AGE_URL"`
	GenderUrl      string `envconfig:"GENDER_URL"`
	NationalityUrl string `envconfig:"NATIONALITY_URL"`
}

type DB struct {
	Username           string `envconfig:"DB_USERNAME" default:"mobile"`
	Host               string `envconfig:"DB_HOST" default:"localhost"`
	Port               string `envconfig:"DB_PORT" default:"5040"`
	DBName             string `envconfig:"DB_NAME" default:"profilesdb"`
	Password           string `envconfig:"DB_PASSWORD" default:"password"`
	SSLMode            string `envconfig:"DB_SSLMODE" default:"disable"`
	DatabaseURL        string `envconfig:"DATABASE_URL" required:"true"`
	MaxOpenConnections int    `envconfig:"DATABASE_MAX_OPEN_CONNECTIONS" default:"10"`
}

type HTTPServer struct {
	Address     string        `envconfig:"HTTP_SERVER_ADDRESS" default:"localhost:8080"`
	Timeout     time.Duration `envconfig:"HTTP_SERVER_TIMEOUT" default:"4s"`
	IdleTimeout time.Duration `envconfig:"HTTP_SERVER_CTX_TIMEOUT" default:"60s"`
	CtxTimeout  time.Duration `envconfig:"HTTP_SERVER_IDLE_TIMEOUT" default:"60s"`
}

func MustLoad() *Config {
	if err := godotenv.Load(envFilePath); err != nil {
		logrus.Fatalf("Error loading .env file: %s", err)
	}
	var cfg Config
	err := envconfig.Process(envPrefix, &cfg)
	if err != nil {
		logrus.Fatalf("Error filling in the structure: %s", err)
		return nil
	}

	logrus.Printf("DatabaseURL: %s\n", cfg.DatabaseURL)
	return &cfg
}
