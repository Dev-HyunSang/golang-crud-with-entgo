package cmd

import (
	"context"
	"log"
	"time"

	"github.com/dev-hyunsang/golang-crud-with-entgo/database"
	"github.com/gofiber/fiber/v2"
)

func ReadToDoHandler(c *fiber.Ctx) error {
	client, err := database.ConnectionSQLite()
	if err != nil {
		log.Println("Failed to Connection SQLite")
		log.Println(err)
	}

	todos, err := client.ToDo.Query().All(context.Background())
	if err != nil {
		log.Println("Failed get todos via SQLite")
		log.Println(err)
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  "error",
		"success": true,
		"message": "성공적으로 할 일을 가지고 왔습니다.",
		"datas":   todos,
		"time":    time.Now(),
	})
}
