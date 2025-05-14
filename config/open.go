package config

import (
	"encoding/json"
	"fmt"
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
		fmt.Println("请修改config.json文件后重新启动！")
		os.Exit(0)
	}
	data, _ := os.ReadFile(filepath.Join(global.RootPath, "config.json"))
	e := json.Unmarshal(data, &ConfigVar)
	return e
}
