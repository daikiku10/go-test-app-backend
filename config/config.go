package config

import "github.com/caarlos0/env/v6"

type Config struct {
	Env           string `env:"GO_ENV" envDefault:""`
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
