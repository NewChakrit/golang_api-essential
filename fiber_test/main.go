package main

// cml run :  nodemon --exec go run . --signal SIGTERM

import (
	"github.com/gofiber/swagger"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/jwt/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
	"github.com/newchakrit/fiber_test/bookList"
	_ "github.com/newchakrit/fiber_test/docs"
	"github.com/newchakrit/fiber_test/env"
	"github.com/newchakrit/fiber_test/login"
	"github.com/newchakrit/fiber_test/middleware"
	"github.com/newchakrit/fiber_test/views"
)

// @title Book API
// @description This is a sample server for a book API.
// @version 1.0
// @termOfService http:swagger.io/terms/
// @contect.name API Support
// @contect.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/license/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("load .env error")
	}

	//http.HandleFunc("/hello", helloHandler)

	//fmt.Printf("Starting sever at port 8080\n")
	//if err := http.ListenAndServe(":8080", nil); err != nil {
	//	log.Fatal(err)
	//}

	// -- View Template --
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// -- Swager --

	app.Get("/swagger/*", swagger.HandlerDefault) // default

	// -- ENV --
	app.Get("/config", env.GetEnv)

	// -- Book List CRUD --
	//app := fiber.New()

	bookList.Books()

	// -- Login --
	app.Post("login", login.Login)

	// -- Middleware --
	//middleware คือ ทางผ่านของการยิง request เข้ามา
	//app.Use(middleware.CheckMiddleware) // ขั้นการยิง api ด้วย middleware

	// -- JWT Middleware --
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	}))

	app.Use(middleware.CheckMiddleware)

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
