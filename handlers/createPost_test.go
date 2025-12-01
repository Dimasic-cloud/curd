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

// проверка создания поста с коректным json
func TestCreatePostValidJSON(t *testing.T) {
	app := fiber.New()
	app.Post("/create", CreatePost)

	// создаём post запрос к серверу
	postData := database.Post{
		UserName: "lostcra",
		Body:     "test",
	}
	body, _ := json.Marshal(&postData)
	req := httptest.NewRequest("POST", "/create", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	// отправляем запрос к серверу и проверяем ошибку
	resp, err := app.Test(req)
	assert.NoError(t, err)
	defer resp.Body.Close()

	// сравниваем код статуса, возвращённого от сервера с правильным
	assert.Equal(t, fiber.StatusCreated, resp.StatusCode)

	// проверяем насколько коректно записываются данные в БД
	var post database.Post
	err = json.NewDecoder(resp.Body).Decode(&post)
	assert.NoError(t, err)
	assert.Equal(t, "lostcra", post.UserName)
	assert.Equal(t, "test", post.Body)
}

// проверка создания поста с некоректным json
func TestCreatePostInvalidJSON(t *testing.T) {
	app := fiber.New()
	app.Post("/create", CreatePost)

	// создаём некоректный json запрос
	invalidJson := `{"userName": "", "body": "test"}`

	// создаём запрос
	req := httptest.NewRequest("POST", "/create", bytes.NewReader([]byte(invalidJson)))
	req.Header.Set("Content-Type", "application/json")

	// отпавляем запрос на серверр
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
}
