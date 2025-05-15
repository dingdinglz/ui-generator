package route

import (
	"os"
	"path/filepath"
	"strings"
	"ui/global"
	"ui/tool"

	"github.com/gofiber/fiber/v2"
)

func FileRoute(c *fiber.Ctx) error {
	sessionID := c.Query("id")
	name := c.Query("name")
	if sessionID == "" || name == "" {
		return c.SendString("")
	}
	res, _ := os.ReadFile(filepath.Join(global.RootPath, "data", sessionID, name))
	return c.SendString(string(res))
}

func FileUpdateRoute(c *fiber.Ctx) error {
	id := c.FormValue("id")
	file := c.FormValue("file")
	content := c.FormValue("content")
	if strings.Count(id, "/") > 0 || strings.Count(file, "/") > 0 {
		return c.SendStatus(403)
	}
	if !tool.FileExist(filepath.Join(global.RootPath, "data", id, file)) {
		return c.SendStatus(403)
	}
	os.WriteFile(filepath.Join(global.RootPath, "data", id, file), []byte(content), os.ModePerm)
	return c.JSON(map[string]interface{}{
		"message": "ok",
	})
}
