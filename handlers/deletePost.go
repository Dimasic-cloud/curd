package handlers

import (
	"curd/database"

	"github.com/gofiber/fiber/v2"
)

// обработчик удоления постов
func DeletePost(c *fiber.Ctx) error {
	userName := c.Params("userName")

	// проверяем на существование, если true, то удоляем и возвращаем статус ОК
	if _, exists := database.DataBase[userName]; exists {
		delete(database.DataBase, userName)

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status": "successful deletion",
		})
	}

	// в этом случае возвращаем ошибку
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"error": "does not exist",
	})
}
