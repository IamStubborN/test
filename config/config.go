package config

import (
	"log"
	"time"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type Config struct {
	Logger Logger
	API    API
	Cache  Cache
	DB     DB
}

type API struct {
	Port            int           `env:"API_PORT" envDefault:"5555"`
	WriteTimeout    time.Duration `env:"API_WRITE_TIMEOUT" envDefault:"10s"`
	ReadTimeout     time.Duration `env:"API_READ_TIMEOUT" envDefault:"5s"`
	GracefulTimeout time.Duration `env:"API_GRACEFUL_TIMEOUT" envDefault:"10s"`
}

type Logger struct {
	Level        string   `env:"LOGGER_LEVEL" envDefault:"info"`
	EncodingType string   `env:"LOGGER_ENCODING_TYPE" envDefault:"console"`
	OutputPaths  []string `env:"LOGGER_OUTPUTS" envSeparator:"," envDefault:"stdout"`
}

type Cache struct {
	BackupFreq time.Duration `env:"CACHE_BACKUP_FREQ" envDefault:"10s"`
}

type DB struct {
	Host       string `env:"DB_HOST" envDefault:"localhost"`
	Port       int    `env:"DB_PORT" envDefault:"5432"`
	Name       string `env:"DB_NAME" envDefault:"test"`
	User       string `env:"DB_USER" envDefault:"postgres"`
	Password   string `env:"DB_PASSWORD" envDefault:"test"`
	RetryCount int    `env:"DB_RETRY_COUNT" envDefault:"60"`
}

func LoadConfig() *Config {
	var cfg Config

	err := godotenv.Load()
	if err != nil {
		log.Println("no .env file")
	}

	if err := env.Parse(&cfg); err != nil {
		log.Fatalln(err)
	}

	return &cfg
}
