package config

import (
	"errors"
	"log"
	"os"
	"reflect"

	"github.com/joho/godotenv"
)

type Config struct {
	APIBaseURI         string `env:"API_BASE_URI"`
	WebBaseURI         string `env:"WEB_BASE_URI"`
	PostgresURI        string `env:"POSTGRES_URI"`
	JWTSecret          string `env:"JWT_SECRET"`
	APIPort            string `env:"API_PORT"`
	DBName             string `env:"DB_NAME"`
	GoogleClientID     string `env:"GOOGLE_CLIENT_ID"`
	GoogleClientSecret string `env:"GOOGLE_CLIENT_SECRET"`
	OpenAIApiKey       string `env:"OPENAI_API_KEY"`
	TriggerSecretKey   string `env:"TRIGGER_SECRET_KEY"`
}

func Load() Config {
	if err := godotenv.Load(); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			log.Println(".env not found, using environment variables as default")
		} else {
			log.Fatal("error loading .env", err)
		}
	}

	var cfg Config
	loadFromEnv(&cfg)
	return cfg
}

func loadFromEnv(cfg *Config) {
	v := reflect.ValueOf(cfg).Elem()
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := t.Field(i)

		if envTag := fieldType.Tag.Get("env"); envTag != "" {
			if envValue := os.Getenv(envTag); envValue != "" {
				field.SetString(envValue)
			}
		}
	}
}
