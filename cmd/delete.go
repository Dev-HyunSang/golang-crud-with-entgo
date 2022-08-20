package cmd

import (
	"context"
	"log"
	"time"

	"github.com/dev-hyunsang/golang-crud-with-entgo/database"
	"github.com/dev-hyunsang/golang-crud-with-entgo/ent/todo"
	"github.com/gofiber/fiber/v2"
)

func DeleteToDoHandler(c *fiber.Ctx) error {
	req := new(ToDos)
	if err := c.BodyParser(req); err != nil {
		log.Println("Failed to BodayParser")
		log.Println(err)
	}

	client, err := database.ConnectionSQLite()
	if err != nil {
		log.Println("Failed to Connection SQLite")
		log.Println(err)
	}

	todoDelete := client.ToDo.
		Query().
		Where(todo.TodoUUID(req.ToDoUUID)).
		OnlyX(context.Background())

	client.ToDo.
		DeleteOne(todoDelete).
		ExecX(context.Background())

	return c.Status(200).JSON(fiber.Map{
		"status":  "error",
		"success": true,
		"message": "성공적으로 할 일을 삭제하였습니다.",
		"time":    time.Now(),
	})
}
