package handlers

import (
	"curd/database"

	"github.com/gofiber/fiber/v2"
)

// функция обнавления тела поста
func UpdatePost(c *fiber.Ctx) error {
	var post database.Post

	// проверяем правильность JSON, полученного из запроса
	// если он не верен, то возвращаем ошибку 400
	if err := c.BodyParser(&post); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid JSON",
		})
	}

	// проверяем правильность данных, записанных в структуру
	// если они не верны, то возвращаем ошибку 400
	if err := Validate.Struct(&post); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid JSON",
		})
	}

	// проверяем существует ли такой user и есть ли у него пост
	// если такого пользователя несуществует, то возвращаем ошибку 400
	// в обратном случае обновляем данные
	if body := database.DataBase[post.UserName]; body == "" {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "user not found",
		})
	} else {
		database.DataBase[post.UserName] = post.Body
	}

	// возвращаем положительный ответ
	return c.SendStatus(fiber.StatusOK)
}
