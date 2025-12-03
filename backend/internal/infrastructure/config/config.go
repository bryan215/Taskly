package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Host      string `yaml:"host"`
	Port      string `yaml:"port"`
	User      string `yaml:"user"`
	Password  string `yaml:"password"`
	DBName    string `yaml:"dbname"`
	JWTSecret string `yaml:"jwt_secret"`
}

func LoadConfig(path string) (*Config, error) {
	config := &Config{}
	if host := os.Getenv("DB_HOST"); host != "" {
		config.Host = host
		config.Port = getEnvOrDefault("DB_PORT", "5432")
		config.User = getEnvOrDefault("DB_USER", "postgres")
		config.Password = getEnvOrDefault("DB_PASSWORD", "postgres")
		config.DBName = getEnvOrDefault("DB_NAME", "tareas_db")
		config.JWTSecret = getEnvOrDefault("JWT_SECRET", "secret")
		return config, nil
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	if err := yaml.Unmarshal(data, config); err != nil {
		return nil, fmt.Errorf("error parsing config file: %w", err)
	}

	config.JWTSecret = getEnvOrDefault("JWT_SECRET", config.JWTSecret)

	return config, nil
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
