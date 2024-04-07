package middleware

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"time"
)

func CheckMiddleware(c *fiber.Ctx) error {
	start := time.Now()
	fmt.Printf("URL = %s, Method = %s, Time = %s \n", c.OriginalURL(), c.Method(), start)

	return c.Next()
}
