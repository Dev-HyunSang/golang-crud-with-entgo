package main

import (
	"context"
	"log"

	"github.com/dev-hyunsang/golang-crud-with-entgo/database"
	"github.com/dev-hyunsang/golang-crud-with-entgo/middleware"
	"github.com/gofiber/fiber/v2"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	app := fiber.New()

	middleware.Middleware(app)
	client, err := database.ConnectionSQLite()
	if err != nil {
		log.Println("Failed to Connection SQLite")
		log.Println(err)
	}

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Println("Failed to Create Schema")
		log.Println(err)
	}

	if err := app.Listen(":3000"); err != nil {
		log.Println()
	}
}
