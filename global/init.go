package global

import (
	"os"
	"path/filepath"
	"ui/tool"
)

func Init() {
	RootPath, _ = os.Getwd()

	tool.DirCreate(filepath.Join(RootPath, "data"))

	prompt, e := os.ReadFile(filepath.Join(RootPath, "prompt_web.txt"))
	if e != nil {
		tool.ErrorLog(e)
		os.Exit(0)
	}
	CorePromptWeb = string(prompt)

	prompt, e = os.ReadFile(filepath.Join(RootPath, "prompt.txt"))
	if e != nil {
		tool.ErrorLog(e)
		os.Exit(0)
	}
	CorePrompt = string(prompt)
}
