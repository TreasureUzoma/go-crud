package posts

import (
	"go-crud/db"
	"go-crud/models"

	"github.com/gofiber/fiber/v2"
)

func CreatePost(c *fiber.Ctx) error {
	post := new(models.Post)
	if err := c.BodyParser(post); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Invalid request body",
			"details": err.Error(),
		})
	}

	if err := db.DB.Create(&post).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to create post",
			"details": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(post)
}

func GetPosts(c *fiber.Ctx) error {
	var posts []models.Post
	if err := db.DB.Find(&posts).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to fetch posts",
			"details": err.Error(),
		})
	}
	return c.JSON(posts)
}

func GetPostById(c *fiber.Ctx) error {
	id := c.Params("id")
	var post models.Post
	if err := db.DB.First(&post, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   "Post not found",
			"details": err.Error(),
		})
	}
	return c.JSON(post)
}

func UpdatePost(c *fiber.Ctx) error {
	id := c.Params("id")

	var existing models.Post
	if err := db.DB.First(&existing, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   "Post not found",
			"details": err.Error(),
		})
	}

	updated := new(models.Post)
	if err := c.BodyParser(updated); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Invalid request body",
			"details": err.Error(),
		})
	}

	if err := db.DB.Model(&existing).Updates(updated).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to update post",
			"details": err.Error(),
		})
	}

	return c.JSON(existing)
}

func DeletePost(c *fiber.Ctx) error {
	id := c.Params("id")
	var post models.Post

	if err := db.DB.First(&post, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   "Post not found",
			"details": err.Error(),
		})
	}

	if err := db.DB.Delete(&post).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to delete post",
			"details": err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
