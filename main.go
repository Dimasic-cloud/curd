package main

import (
	"github.com/gofiber/fiber/v2"
)

// структура для json запроса
type Post struct {
	Body     string `json:"body"`
	UserName string `json:"userName"`
}

// имитация базы данных
var dataBase = make(map[string]map[string]string)

func main() {
	// создаём экземпляр приложения
	app := fiber.New()

	// обработчик post запросса
	// который создаёт запись
	app.Post("/create", createPost)

	// активируем сервер на порту :3000
	app.Listen(":3000")
}

// функция обработчик
func createPost(c *fiber.Ctx) error {
	var post Post

	// проверка json на соответствие полей со структурой
	// при плохом завершении возвращаем 400 и ошибку
	// при положительном заносим эти данные в БД
	if err := c.BodyParser(&post); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error() + "\n",
		})
	}
	dataBase[post.UserName] = map[string]string{
		"body": post.Body,
	}

	//возвращаем положительный ответ
	return c.JSON(fiber.Map{
		"status": fiber.StatusCreated,
	})
}
