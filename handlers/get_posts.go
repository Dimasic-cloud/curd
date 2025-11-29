package handlers

import (
	"curd/database"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func GetPosts(c *fiber.Ctx) error {
	userName := c.Params("userName")
	mapPosts := database.DataBase[userName]

	fmt.Println(mapPosts)
	return c.SendString("ok")
}
