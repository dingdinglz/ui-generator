package config

import (
	"encoding/json"
	"os"
	"path/filepath"
	"ui/global"
	"ui/tool"
)

func OpenConfig() error {
	if !tool.FileExist(filepath.Join(global.RootPath, "config.json")) {
		data, _ := json.MarshalIndent(Config{
			Host:  "0.0.0.0",
			Port:  3000,
			Model: ModelConfig{},
		}, "", "    ")
		os.WriteFile(filepath.Join(global.RootPath, "config.json"), data, os.ModePerm)
	}
	data, _ := os.ReadFile(filepath.Join(global.RootPath, "config.json"))
	e := json.Unmarshal(data, &ConfigVar)
	return e
}
