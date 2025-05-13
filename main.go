package main

import (
	"strconv"
	"ui/config"
	"ui/global"
	"ui/tool"

	"github.com/dingdinglz/openai"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	global.Init()
	e := config.OpenConfig()
	if e != nil {
		tool.ErrorLog(e)
		return
	}
	global.OpenaiClient = openai.NewClient(&openai.ClientConfig{
		BaseUrl: config.ConfigVar.Model.Base,
		ApiKey:  config.ConfigVar.Model.Key,
	})

	server := fiber.New()

	server.Use(logger.New(), recover.New())
	BindRoutes(server)

	server.Static("/", "./data")
	server.Static("/", "./web")

	e = server.Listen(config.ConfigVar.Host + ":" + strconv.Itoa(config.ConfigVar.Port))
	if e != nil {
		tool.ErrorLog(e)
	}
}
