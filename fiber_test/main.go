package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/newchakrit/fiber_test/bookList"
)

func main() {
	//http.HandleFunc("/hello", helloHandler)

	//fmt.Printf("Starting sever at port 8080\n")
	//if err := http.ListenAndServe(":8080", nil); err != nil {
	//	log.Fatal(err)
	//}

	app := fiber.New()

	bookList.Books()

	app.Get("/books", bookList.GetBooks)
	app.Get("/books/:id", bookList.GetBook)
	app.Post("/books", bookList.CreateBook)
	app.Put("/books/:id", bookList.UpdateBook)
	app.Delete("/books/:id", bookList.DeleteBook)

	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello World!")
	})

	app.Listen(":8080")
}

// func helloHandler(w http.ResponseWriter, r *http.Request) {
//	if r.URL.Path != "/hello" {
//		http.Error(w, "404 not found.", http.StatusNotFound)
//		return
//	}
//
//	if r.Method != "GET" {
//		http.Error(w, "Method is not supported.", http.StatusNotFound)
//		return
//	}
//	fmt.Fprintf(w, "Hello World!")
//}
