package main

import (
	"fmt"

	"github.com/AtheerAPeter/go-fiber-tutorial/book"
	"github.com/AtheerAPeter/go-fiber-tutorial/database"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func getFunc(c *fiber.Ctx) {
	c.Send("hello")
}

func setUpRoutes(app *fiber.App) {
	app.Get("/api/v1/book", book.GetBooks)
	app.Get("/api/v1/book/:id", book.GetBook)
	app.Post("/api/v1/book", book.NewBook)
	app.Delete("/api/v1/book/:id", book.DeleteBook)
}

func initDatabase() {
	var err error
	database.DBconn, err = gorm.Open("sqlite3", "books.db")
	if err != nil {
		panic("Failed to connect ")
	}
	fmt.Println(("success"))

	database.DBconn.AutoMigrate(&book.Book{})
	fmt.Println("Database Migrated")
}
func main() {
	app := fiber.New()
	initDatabase()
	defer database.DBconn.Close()
	setUpRoutes(app)

	app.Listen(3000)
}
