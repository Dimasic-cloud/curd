package handlers

import (
	"curd/database"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var Validate = validator.New()

// функция обрабатывающая post запрос для создания поста
func CreatePost(c *fiber.Ctx) error {
	var post database.Post

	// проверка json запроса на соответствие с полями структуры
	// при плохом завершении возвращаем 400 и ошибку
	// при положительном заносим данные в БД
	if err := c.BodyParser(&post); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid JSON\n",
		})
	}

	// проверяем попали ли в структуру пустые строки
	// если да, возвращаем код 400 о невалидных данных
	if err := Validate.Struct(&post); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid JSON\n",
		})
	}

	// добавляем запись
	database.DataBase[post.UserName] = post.Body

	//возвращаем положительный ответ
	return c.Status(fiber.StatusCreated).JSON(post)
}
