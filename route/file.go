package route

import (
	"os"
	"path/filepath"
	"ui/global"

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
