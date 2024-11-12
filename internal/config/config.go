package config

import "github.com/caarlos0/env/v11"

type Config struct {
	GeminiConfig GeminiConfig
}

type GeminiConfig struct {
	API_KEY string `env:"GEMINI_API_KEY"`
}

func New() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
