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
	return c.JSON(BookLists)
}

func GetBook(c *fiber.Ctx) error {
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

func CreateBook(c *fiber.Ctx) error {
	book := new(Book)

	if err := c.BodyParser(book); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	BookLists = append(BookLists, *book)

	return c.JSON(BookLists)
}

func UpdateBook(c *fiber.Ctx) error {
	bookId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	bookUpdate := new(Book)
	if err := c.BodyParser(bookUpdate); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	for i, book := range BookLists {
		if book.ID == bookId {
			BookLists[i].Title = bookUpdate.Title
			BookLists[i].Author = bookUpdate.Author
			return c.JSON(BookLists[i])
		}
	}

	return c.Status(fiber.StatusNotFound).SendString("Data not found")
}

func DeleteBook(c *fiber.Ctx) error {
	bookId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	bookDelete := new(Book)
	if err := c.BodyParser(bookDelete); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	for i, book := range BookLists {
		if book.ID == bookId {
			BookLists = append(BookLists[:i], BookLists[i+1:]...) // ... คือกระจาย slice element แต่ละตัว แยกออกจากกัน แล้วเอามาต่อกันอีกที่ (เหมือนเอา element แต่ละตัวมาทำการ append แยกกันอีกที )
			//if i = 3
			// [1,2,3,4,5,6]
			// [1,2] + [4,5,6] = [1,2,4,5,6]
			return c.SendStatus(fiber.StatusNoContent)
		}
	}

	return c.Status(fiber.StatusNotFound).SendString("Data not found")
}
