package handlers

import (
	"bytes"
	"curd/database"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

// проверяем, что метод UpdatePost, правильно обрабатывает коректные данные
func TestUpdatePostValidData(t *testing.T) {
	app := fiber.New()
	app.Put("/update", UpdatePost)

	// заносим пробное значение в БД
	database.DataBase["dima"] = "test"

	// подготавливаем данные к запросу
	postData := database.Post{
		UserName: "dima",
		Body:     "new test",
	}
	body, _ := json.Marshal(&postData)

	// создаём и отправляем запрос
	req := httptest.NewRequest("PUT", "/update", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)

	// проверяем коректность полученных данных
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	assert.Equal(t, "new test", database.DataBase["dima"])
}

// функция для проверки ошибок
func TestUpdatePostInvalidData(t *testing.T) {
	app := fiber.New()
	app.Put("/update", UpdatePost)

	// имитация БД
	database.DataBase["dima"] = "test"

	// подготавливаем данные
	postData := database.Post{
		UserName: "nina",
		Body:     "new test",
	}
	body, _ := json.Marshal(&postData)

	// создаём и отпровляем запрос
	req := httptest.NewRequest("PUT", "/update", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	assert.NoError(t, err)

	// сравниваем статус код
	assert.Equal(t, fiber.StatusNotFound, resp.StatusCode)

	// проверяем возвращённый json
	var bodyErr map[string]string
	err = json.NewDecoder(resp.Body).Decode(&bodyErr)
	assert.NoError(t, err)
	assert.Equal(t, "user not found", bodyErr["error"])
}
