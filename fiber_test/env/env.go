package env

import (
	"github.com/gofiber/fiber/v2"
	"os"
)

func GetEnv(c *fiber.Ctx) error {
	//if value, exists := os.LookupEnv("SECRET"); exists { // value คือการเก็บค่่าของตัวนั้นเอาไว้, exists คือค่าที่เช็คว่า value มีค่าหรือไม่
	//	return c.JSON(fiber.Map{
	//		"SECRET": value,
	//	})
	//}

	//return c.JSON(fiber.Map{
	//	"SECRET": "defaultsecret",
	//})

	secret := os.Getenv("SECRET")

	if secret == "" {
		secret = "defualtsecret"
		return c.JSON(fiber.Map{
			"SECRET": secret,
		})
	}

	return c.JSON(fiber.Map{
		"SECRET": os.Getenv("SECRET"),
	})
}
