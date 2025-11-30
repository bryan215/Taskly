package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
}

func LoadConfig(path string) (*Config, error) {
	config := &Config{}

	// Intentar cargar desde variables de entorno primero (para Docker)
	if host := os.Getenv("DB_HOST"); host != "" {
		config.Host = host
		config.Port = getEnvOrDefault("DB_PORT", "5432")
		config.User = getEnvOrDefault("DB_USER", "postgres")
		config.Password = getEnvOrDefault("DB_PASSWORD", "postgres")
		config.DBName = getEnvOrDefault("DB_NAME", "tareas_db")
		return config, nil
	}

	// Fallback a archivo YAML
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	if err := yaml.Unmarshal(data, config); err != nil {
		return nil, fmt.Errorf("error parsing config file: %w", err)
	}

	return config, nil
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
