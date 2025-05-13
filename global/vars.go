package global

import (
	"github.com/dingdinglz/openai"
)

var (
	Version       string = "v0.1 beta"
	RootPath      string = ""
	OpenaiClient  *openai.Client
	CorePrompt    string
	CorePromptWeb string
)
