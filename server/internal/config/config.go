package config

import "github.com/brumhard/alligotor"

type Config struct {
	Address string

	PrometheusDatasource PrometheusDatasourceConfig
}

type PrometheusDatasourceConfig struct {
	Address  string
	Username string
	Password string
}

func Get() (Config, error) {
	configSource := alligotor.New(
		alligotor.NewEnvSource("API"),
		alligotor.NewFlagsSource(),
	)

	cfg := Config{
		Address: ":8080",
		PrometheusDatasource: PrometheusDatasourceConfig{
			Address: "http://localhost:9090",
		},
	}

	err := configSource.Get(&cfg)

	return cfg, err
}
