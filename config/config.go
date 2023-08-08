package config

import "github.com/caarlos0/env/v6"

type Config struct {
	Env           string `env:"GO_ENV" envDefault:"development"`
	Port          int    `env:"PORT" envDefault:"8080"`
	DBHost        string `env:"DB_HOST" envDefault:"db"`
	DBPort        int    `env:"DB_POST" envDefault:"3306"`
	DBUser        string `env:"DB_USER" envDefault:"go_test"`
	DBPassword    string `env:"DB_PASSWORD" envDefault:"password"`
	DBName        string `env:"DB_NAME" envDefault:"go_database"`
	FrontEndpoint string `env:"FRONT_ENDPOINT" envDefault:"http://localhost:3000"`
}

// 環境変数の構造体を返却
//
// @return *Config 環境変数の構造体
//
// @return error エラー
func New() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
