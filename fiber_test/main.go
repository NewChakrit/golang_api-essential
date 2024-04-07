package main

// cml run :  nodemon --exec go run . --signal SIGTERM

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/newchakrit/fiber_test/bookList"
	"github.com/newchakrit/fiber_test/views"
)

func main() {
	//http.HandleFunc("/hello", helloHandler)

	//fmt.Printf("Starting sever at port 8080\n")
	//if err := http.ListenAndServe(":8080", nil); err != nil {
	//	log.Fatal(err)
	//}

	// View Template
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	//app := fiber.New()

	bookList.Books()

	app.Get("/books", bookList.GetBooks)
	app.Get("/books/:id", bookList.GetBook)
	app.Post("/books", bookList.CreateBook)
	app.Put("/books/:id", bookList.UpdateBook)
	app.Delete("/books/:id", bookList.DeleteBook)

	app.Get("test-html", views.TestHTML)

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
