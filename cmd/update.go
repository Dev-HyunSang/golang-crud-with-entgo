package cmd

import (
	"context"
	"log"
	"time"

	"github.com/dev-hyunsang/golang-crud-with-entgo/database"
	"github.com/gofiber/fiber/v2"
)

func UpdateToDoHandler(c *fiber.Ctx) error {
	req := new(ToDos)
	if err := c.BodyParser(req); err != nil {
		log.Println("Failed to BodyParser")
		log.Println(err)
	}

	client, err := database.ConnectionSQLite()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{})
	}

	upodateToDo, err := client.ToDo.Update().
		SetTodo(req.ToDo).
		SetEditedAt(time.Now()).
		Save(context.Background())
	if err != nil {
		log.Println("Failed to Update Todo via SQLite")
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"success": false,
			"message": "",
			"time":    time.Now(),
		})
	}

	log.Println(upodateToDo)

	return c.Status(200).JSON(fiber.Map{
		"status":  "sucess",
		"sucess":  true,
		"message": "정상적으로 할 일을 업로드 하였습니다.",
		"time":    time.Now(),
	})
}
