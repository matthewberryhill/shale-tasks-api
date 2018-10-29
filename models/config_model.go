package models

type Config struct {
	Name        string `json:"name"`
	Version     string `json:"version"`
	Environment string `json:"environment"`
	Error       string `json:"error"`
}

func GetConfig() *Config {
	c := new(Config)
	c.Name = "shale-tasks-api"
	c.Version = "0.0.1"
	c.Environment = "prod"
	c.Error = ""

	return c
}
