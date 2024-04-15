package main

import (
	"fmt"

	"c0deg3isha.io/todolist/database"
	"c0deg3isha.io/todolist/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func helloWorld(c *fiber.Ctx) error {
	return c.SendString("Hello world")
}

func initDatabase() {
	var err error
	dsn := "host=loalhost user=postgres password=admin123 dbname=todoDB port 54321"
	database.DBConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}
	fmt.Println("Database connected!")
	database.DBConn.AutoMigrate(&models.Todo{})
	fmt.Println("Migrated DB")
}

func setupRoutes(app *fiber.App) {
	app.Get("/todos", models.GetTodos)
	app.Post("/todos", models.CreateTodo)
	app.Get("/todos/:id", models.GetTodo)
	app.Put("/todos/:id", models.UpdateTodo)
	app.Delete("/todos/:id", models.DeleteTodo)
}

func main() {
	app := fiber.New()
	app.Use(cors.New())
	initDatabase()

	app.Get("/", helloWorld)
	setupRoutes(app)
	app.Listen(":8000")
}
