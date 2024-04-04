package bookList

import "github.com/gofiber/fiber/v2"

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var BookLists []Book

func GetBook(c *fiber.Ctx) error {
	Books()
	return c.JSON(BookLists)
}

func Books() {
	BookLists = append(BookLists, Book{ID: 1, Title: "Money", Author: "Jom"})
	BookLists = append(BookLists, Book{ID: 2, Title: "Lift Style", Author: "Phone"})
	BookLists = append(BookLists, Book{ID: 3, Title: "Health", Author: "Yam"})
	BookLists = append(BookLists, Book{ID: 4, Title: "Game", Author: "Fast"})
	BookLists = append(BookLists, Book{ID: 5, Title: "Invest", Author: "New"})
}