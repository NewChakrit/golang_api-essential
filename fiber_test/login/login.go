package login

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

var memberUser = User{ // defualt
	Email:    "sudlhor@gmail.com",
	Password: "password",
}

func Login(c *fiber.Ctx) error {
	user := new(User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	//if user.Email == memberUser.Email && user.Password == memberUser.Password {
	//	return c.JSON(fiber.Map{
	//		"message": "Login success",
	//	})
	//}

	// user, pass ไม่ตรง = Unauthorized
	if user.Email != memberUser.Email || user.Password != memberUser.Password {
		return fiber.ErrUnauthorized
	}

	// Create Token
	token := jwt.New(jwt.SigningMethodHS256) // สร้าง pattern ของ token

	// Set claims
	// claims คือการเก็บข้อมูลของ token และทำการ encrypt ออกมาเป็นข้อมูล set นึง
	// เราจะสามารถดึงข้อมูลของ set นี้ได้ก็ต่อเมื่อมี key ตรงกัน
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = user.Email
	claims["role"] = "admin"
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	// Generate encoded token and send it as response
	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{
		"message": "Login success",
		"token":   t,
	})
}
