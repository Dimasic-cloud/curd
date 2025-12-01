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

// проверяем, что метод DeletePost, правильно обрабатывает коректные данные
func TestDeletePostValidData(t *testing.T) {
	app := fiber.New()
	app.Delete("/delete/:userName", DeletePost)

	// заносим пробное значение в БД
	database.DataBase["dima"] = "test"

	// создаём и отправляем запрос, проверяя на ошибки
	req := httptest.NewRequest("DELETE", "/delete/dima", nil)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)

	// парсим, обрабатываем и сравниваем данные из ответа
	body, err := io.ReadAll(resp.Body)
	var res map[string]string
	json.Unmarshal(body, &res)
	assert.NoError(t, err)
	assert.Equal(t, "successful deletion", res["status"])
}

// функция для проверки отрицательного случая
func TestDeletePostInvalidData(t *testing.T) {
	app := fiber.New()
	app.Delete("/delete/:username", DeletePost)

	// имитация БД
	database.DataBase["dima"] = "test"

	// создаём и отпровляем некоректный запрос
	req := httptest.NewRequest("DELETE", "/delete/nina", nil)
	resp, err := app.Test(req)
	assert.NoError(t, err)

	// полученные данные обрабатываем и сравниваем
	body, err := io.ReadAll(resp.Body)
	var res map[string]string
	json.Unmarshal(body, &res)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	assert.Equal(t, "does not exist", res["error"])
}
