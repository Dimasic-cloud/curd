package main

import (
	"curd/handlers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// создаём экземпляр приложения
	app := fiber.New()

	// хандлеры
	// обрабатываем post запрос и создаём пользователя с записями
	app.Post("/create", handlers.CreatePost)
	// обрабатываем get запрос, чтобы пользователь мог получать свои посты
	app.Get("/posts/:username", handlers.GetPosts)

	// запускаем сервер на локальном порту :3000
	app.Listen(":3000")
}
