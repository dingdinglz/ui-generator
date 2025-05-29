package main

import (
	"ui/route"

	"github.com/gofiber/fiber/v2"
)

func BindRoutes(server *fiber.App) {
	apiGroup := server.Group("/api")
	apiGroup.Post("/generate", route.GenerateRoute)
	apiGroup.Get("/file", route.FileRoute)
	apiGroup.Post("/save", route.FileUpdateRoute)
	apiGroup.Get("/load", route.FileListRoute)
	apiGroup.Get("/change", route.ChangeFileRoute)
	apiGroup.Get("/continue", route.ContinueRoute)
	apiGroup.Get("/add", route.AddRoute)
	apiGroup.Get("/download", route.DownLoadAllRoute)
}
