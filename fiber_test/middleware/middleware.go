package middleware

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func CheckMiddleware(c *fiber.Ctx) error {
	//start := time.Now()
	//fmt.Printf("URL = %s, Method = %s, Time = %s \n", c.OriginalURL(), c.Method(), start)

	user := c.Locals("user").(*jwt.Token) // c.Local อ่านตัวแปรของ context หรือตัวแปรที่มีการรับฝากเอาไว้จาก context ของตัวอื่น
	claims := user.Claims.(jwt.MapClaims)

	fmt.Println(claims) // map[email:sudlhor@gmail.com exp:1.712569913e+09 role:admin]

	if claims["role"] != "admin" {
		return fiber.ErrUnauthorized
	}

	return c.Next() // middleware สามารถอนุญาตให้ผ่านไปต่อได้
}
