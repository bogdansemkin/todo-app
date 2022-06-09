package config

type Config struct{
	BindAddr string `toml:"bind_adddr"`
}

func NewConfig() *Config{
	return &Config{
		BindAddr: ":8081",
	}
}
