package handlers

import (
	"curd/database"

	"github.com/gofiber/fiber/v2"
)

func UpdatePost(c *fiber.Ctx) error {
	var post database.Post

	if err := c.BodyParser(&post); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid JSON\n",
		})
	}

	if err := Validate.Struct(&post); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid JSON\n",
		})
	}

	if body := database.DataBase[post.UserName]; body == "" {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "user not found\n",
		})
	} else {
		database.DataBase[post.UserName] = post.Body
	}

	return c.SendStatus(fiber.StatusOK)
}
