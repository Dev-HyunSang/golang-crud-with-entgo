package cmd

import (
	"context"
	"log"
	"time"

	"github.com/dev-hyunsang/golang-crud-with-entgo/database"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func CreateTodoHandler(c *fiber.Ctx) error {
	req := new(ToDos)
	if err := c.BodyParser(req); err != nil {
		log.Println("Failed to BodyParser")
		log.Println(err)
	}

	client, err := database.ConnectionSQLite()
	if err != nil {
		log.Println("Failed to Connection DataBase via SQLite")
		log.Println(err)
	}

	todoUUID, err := uuid.NewUUID()
	if err != nil {
		log.Println("Failed to Create UUID")
		log.Println(err)
	}

	createToDo, err := client.ToDo.
		Create().
		SetTodoUUID(todoUUID).
		SetTodo(req.ToDo).
		SetDone(false).
		SetCreatedAt(time.Now()).
		SetDeletedAt(time.Now()).
		SetEditedAt(time.Now()).
		Save(context.Background())
	if err != nil {
		log.Println("Failed to Create ToDo via SQLite")
		log.Println(err)

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"success": false,
			"message": "알 수 없는 오류가 발생되었습니다. 잠시후 시도해 주세요.",
			"time":    time.Now(),
		})
	}

	log.Println(createToDo.Done)

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"success": true,
		"message": "성공적으로 새로운 항목을 추가하였습니다.",
		"time":    time.Now(),
	})
}
