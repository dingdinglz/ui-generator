package route

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"ui/config"
	"ui/global"
	"ui/tool"

	"github.com/dingdinglz/openai"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/valyala/fasthttp"
)

type GenerateMessage struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

type TaskList struct {
	Data []struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	} `json:"data"`
}

func MakeGenerateMessage(t string, message string) string {
	data, _ := json.Marshal(GenerateMessage{
		Type:    t,
		Message: message,
	})
	return string(data)
}

func SendSSEMessage(w *bufio.Writer, message string) {
	fmt.Fprintf(w, "data: %s\n\n", message)
	w.Flush()
}

func GenerateRoute(ctx *fiber.Ctx) error {
	ctx.Set("Content-Type", "text/event-stream")
	ctx.Set("Cache-Control", "no-cache")
	ctx.Set("Connection", "keep-alive")
	ctx.Set("Transfer-Encoding", "chunked")
	idea := ctx.Query("idea")
	mode := ctx.Query("mode")
	ctx.Status(fiber.StatusOK).Context().SetBodyStreamWriter(fasthttp.StreamWriter(func(w *bufio.Writer) {
		if idea == "" {
			SendSSEMessage(w, MakeGenerateMessage("error", "idea不能为空"))
			return
		}
		sessionID := uuid.New().String()
		tool.DirCreate(filepath.Join(global.RootPath, "data", sessionID))
		SendSSEMessage(w, MakeGenerateMessage("start", sessionID))
		messages := []openai.Message{
			{
				Role:    "system",
				Content: global.CorePrompt,
			},
			{
				Role:    "user",
				Content: idea,
			},
		}
		if mode == "web" {
			messages[0].Content = global.CorePromptWeb
		}
		taskListData := ""
		e := global.OpenaiClient.ChatStream(config.ConfigVar.Model.Model, messages, func(s string) {
			taskListData += s
			SendSSEMessage(w, MakeGenerateMessage("update", taskListData))
		})
		if e != nil {
			SendSSEMessage(w, MakeGenerateMessage("error", e.Error()))
			return
		}

		messages = append(messages, openai.Message{
			Role:    "assistant",
			Content: taskListData,
		})

		taskListData = tool.StringBetweenContain(taskListData, "{", "}")
		taskList := TaskList{}
		e = json.Unmarshal([]byte(taskListData), &taskList)
		if e != nil {
			SendSSEMessage(w, MakeGenerateMessage("error", "task list生成出错:"+taskListData))
			return
		}
		SendSSEMessage(w, MakeGenerateMessage("task", taskListData))
		for _, item := range taskList.Data {
			messages = append(messages, openai.Message{
				Role:    "user",
				Content: item.Name,
			})
			SendSSEMessage(w, MakeGenerateMessage("task_start", item.Name))
			fileContent := ""
			e := global.OpenaiClient.ChatStream(config.ConfigVar.Model.Model, messages, func(s string) {
				fileContent += s
				SendSSEMessage(w, MakeGenerateMessage("update", fileContent))
			})
			if e != nil {
				SendSSEMessage(w, MakeGenerateMessage("error", e.Error()))
				return
			}
			fileContent = tool.StringBetween(fileContent, "```html", "```")
			messages = append(messages, openai.Message{
				Role:    "assistant",
				Content: fileContent,
			})
			os.WriteFile(filepath.Join(global.RootPath, "data", sessionID, item.Name), []byte(fileContent), os.ModePerm)
			SendSSEMessage(w, MakeGenerateMessage("task_end", item.Name))
		}
		SendSSEMessage(w, MakeGenerateMessage("update", "生成完成！如遇任何问题请联系dinglz"))
		messageSave, _ := json.Marshal(messages)
		os.WriteFile(filepath.Join(global.RootPath, "data", sessionID, "messages.json"), messageSave, os.ModePerm)
		SendSSEMessage(w, MakeGenerateMessage("end", "生成完成"))
	}))
	return nil
}

func ChangeFileRoute(ctx *fiber.Ctx) error {
	ctx.Set("Content-Type", "text/event-stream")
	ctx.Set("Cache-Control", "no-cache")
	ctx.Set("Connection", "keep-alive")
	ctx.Set("Transfer-Encoding", "chunked")
	id := ctx.Query("id")
	name := ctx.Query("name")
	prompt := ctx.Query("prompt")
	ctx.Status(fiber.StatusOK).Context().SetBodyStreamWriter(fasthttp.StreamWriter(func(w *bufio.Writer) {
		if id == "" || prompt == "" || name == "" {
			SendSSEMessage(w, MakeGenerateMessage("error", "不能为空！"))
			return
		}
		workPath := filepath.Join(global.RootPath, "data", id)
		if !tool.FileExist(filepath.Join(workPath, "messages.json")) {
			SendSSEMessage(w, MakeGenerateMessage("error", "历史记录缺失！"))
			return
		}
		historyContent, _ := os.ReadFile(filepath.Join(workPath, "messages.json"))
		historyMessages := []openai.Message{}
		json.Unmarshal(historyContent, &historyMessages)
		historyMessages = append(historyMessages, openai.Message{
			Role:    "user",
			Content: "下面你需要重新生成" + name + "，要求是" + prompt + "\n注意，你需要和之前的生成规则保持一致，即仅返回html的内容即可",
		})
		fileContent := ""
		e := global.OpenaiClient.ChatStream(config.ConfigVar.Model.Model, historyMessages, func(s string) {
			fileContent += s
			SendSSEMessage(w, MakeGenerateMessage("update", fileContent))
		})
		if e != nil {
			SendSSEMessage(w, MakeGenerateMessage("error", e.Error()))
			return
		}
		fileContent = tool.StringBetween(fileContent, "```html", "```")
		os.WriteFile(filepath.Join(workPath, name), []byte(fileContent), os.ModePerm)
		updateHistoryMessage(id, name, fileContent)
		SendSSEMessage(w, MakeGenerateMessage("update", "生成完成！如遇任何问题请联系dinglz"))
		SendSSEMessage(w, MakeGenerateMessage("end", "生成完成"))
	}))
	return nil
}
