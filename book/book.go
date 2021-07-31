package book

import (
	"github.com/AtheerAPeter/go-fiber-tutorial/database"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

type Book struct {
	gorm.Model
	Title  string `json:"title"`
	Author string `json:"author"`
	Rating int    `json:"rating"`
}

func GetBooks(c *fiber.Ctx) {
	db := database.DBconn
	var books []Book
	db.Find(&books)
	c.JSON(books)
}
func GetBook(c *fiber.Ctx) {
	id := c.Params("id")
	var book Book
	database.DBconn.Find(&book, id)
	c.JSON(book)

}
func NewBook(c *fiber.Ctx) {

	book := new(Book)
	if err := c.BodyParser(book); err != nil {
		c.Status(500).Send("error")
		return
	}
	database.DBconn.Create(&book)
	c.JSON(book)

}
func DeleteBook(c *fiber.Ctx) {
	id := c.Params("id")

	var book Book
	database.DBconn.First(&book, id)
	if book.Title == "" {
		c.Status(500).Send("book not found")
		return
	}
	database.DBconn.Delete(&book)
	c.Send("Book deleted")
}
