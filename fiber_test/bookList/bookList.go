package bookList

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var BookLists []Book

func Books() {
	BookLists = append(BookLists, Book{ID: 1, Title: "Money", Author: "Jom"})
	BookLists = append(BookLists, Book{ID: 2, Title: "Lift Style", Author: "Phone"})
	BookLists = append(BookLists, Book{ID: 3, Title: "Health", Author: "Yam"})
	BookLists = append(BookLists, Book{ID: 4, Title: "Game", Author: "Fast"})
	BookLists = append(BookLists, Book{ID: 5, Title: "Invest", Author: "New"})
}

func GetBooks(c *fiber.Ctx) error {
	Books()
	return c.JSON(BookLists)
}

func GetBook(c *fiber.Ctx) error {
	Books()
	bookId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	for _, book := range BookLists {
		if book.ID == bookId {
			return c.JSON(book)
		}
	}

	// SendStatus(เพิ่มคำว่า Send) เป็นการส่ง status อย่างเดียว ไม่มี message
	//return c.SendStatus(fiber.StatusNotFound)

	return c.Status(fiber.StatusNotFound).SendString("Data not found")
}
