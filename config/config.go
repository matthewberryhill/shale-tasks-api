package config

type Config struct {
	Mongo []string `env:"MONGO"`
}
