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

	// создаём несколько некоректных json запросов
	invalidJson1 := `{"userName": "dima", "body": ""}`
	invalidJson2 := `{"userName": "", "body": "test"}`

	// создаём запроссы
	req1 := httptest.NewRequest("POST", "/create", bytes.NewReader([]byte(invalidJson1)))
	req2 := httptest.NewRequest("POST", "/create", bytes.NewReader([]byte(invalidJson2)))
	req1.Header.Set("Content-Type", "application/json")
	req2.Header.Set("Content-Type", "application/json")

	// отпавляем запросы на серверр и проверяем полученные ошибки
	resp1, err1 := app.Test(req1)
	resp2, err2 := app.Test(req2)
	assert.NoError(t, err1)
	assert.NoError(t, err2)
	assert.Equal(t, fiber.StatusBadRequest, resp1.StatusCode)
	assert.Equal(t, fiber.StatusBadRequest, resp2.StatusCode)
}
