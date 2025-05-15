package main

import (
	"ui/route"

	"github.com/gofiber/fiber/v2"
)

func BindRoutes(server *fiber.App) {
	apiGroup := server.Group("/api")
	apiGroup.Get("/generate", route.GenerateRoute)
	apiGroup.Get("/file", route.FileRoute)
	apiGroup.Post("/save", route.FileUpdateRoute)
}
