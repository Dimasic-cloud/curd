package database

// структура для создания постов
type Post struct {
	Body     string `json:"body" validate:"required"`
	UserName string `json:"userName" validate:"required"`
}

// имитация базы данных
var DataBase = make(map[string]string)
