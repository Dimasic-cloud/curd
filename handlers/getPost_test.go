package handlers

import (
	"curd/database"
	"encoding/json"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

// проверяем, что метод GetPost, правильно обрабатывает коректные данные
func TestGetPostValidData(t *testing.T) {
	app := fiber.New()
	app.Get("/post/:username", GetPost)

	// заносим пробное значение в БД
	database.DataBase["dima"] = "test"

	// создаём и отправляем запрос, проверяя на ошибки
	req := httptest.NewRequest("GET", "/post/dima", nil)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)

	// парсим, обрабатываем и сравниваем данные из ответа
	body, err := io.ReadAll(resp.Body)
	var res map[string]string
	json.Unmarshal(body, &res)
	assert.NoError(t, err)
	assert.Equal(t, "test", res["body"])
}

// функция для проверки отрицательного случая
func TestGetPostInvalidData(t *testing.T) {
	app := fiber.New()
	app.Get("/post/:username", GetPost)

	// имитация БД
	database.DataBase["dima"] = "test"

	// создаём и отпровляем запрос
	req := httptest.NewRequest("GET", "/post/nina", nil)
	resp, err := app.Test(req)
	assert.NoError(t, err)

	// полученные данные обрабатываем и сравниваем
	body, err := io.ReadAll(resp.Body)
	var res map[string]string
	json.Unmarshal(body, &res)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusNotFound, resp.StatusCode)
	assert.Equal(t, "user not found", res["error"])
}
