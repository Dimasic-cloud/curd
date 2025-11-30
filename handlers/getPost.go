package handlers

import (
	"curd/database"

	"github.com/gofiber/fiber/v2"
)

// обработчик для поиска постов по username
func GetPost(c *fiber.Ctx) error {
	// получаем username из запроса
	userName := c.Params("userName")

	// проверяем на существование
	if body := database.DataBase[userName]; body == "" {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "user not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"body": database.DataBase[userName],
	})
}
