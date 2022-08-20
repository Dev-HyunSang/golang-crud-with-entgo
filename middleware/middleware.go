package middleware

import (
	"github.com/dev-hyunsang/golang-crud-with-entgo/cmd"
	"github.com/gofiber/fiber/v2"
)

func Middleware(app *fiber.App) {
	app.Post("/create", cmd.CreateTodoHandler)
	app.Post("/update", cmd.UpdateToDoHandler)
	app.Post("/read", cmd.ReadToDoHandler)
	app.Delete("/delete", cmd.DeleteToDoHandler)
}
