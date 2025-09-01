package main

import (
	"go-crud/db"
	posts "go-crud/handlers"
	"go-crud/middleware"
	"go-crud/models"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Use(middleware.Logger)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello from go fiber")
	})

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON("Server is Healthy")
	})

	db.ConnectDB()
	db.DB.AutoMigrate(&models.Post{})

	api := app.Group("/api/v1")

	postsapi := api.Group("/posts")
	postsapi.Post("/", posts.CreatePost)
	postsapi.Get("/", posts.GetPosts)
	postsapi.Get("/:id", posts.GetPostById)
	postsapi.Patch("/:id", posts.UpdatePost)
	postsapi.Delete("/:id", posts.DeletePost)

	app.Listen(":3000")
}