package config

type Config struct {
	Host  string      `json:"host"`
	Port  int         `json:"port"`
	Model ModelConfig `json:"model"`
}

type ModelConfig struct {
	Base  string `json:"base"`
	Key   string `json:"key"`
	Model string `json:"model"`
}

var ConfigVar Config
