package endpoints

import (
	"fmt"

	"github.com/drew138/tictac/api/authentication"
	"github.com/drew138/tictac/api/authorization"
	"github.com/drew138/tictac/database"
	"github.com/drew138/tictac/database/models"
	"github.com/gofiber/fiber/v2"
)

// Login - Grant access and permissions by providing jwt
func Login(c *fiber.Ctx) error {

	user := new(models.User) // request user
	if UnmarshalJSON(c, &user) {
		return fmt.Errorf("Invalid user properties")
	}
	var User models.User // user in database
	database.DBConn.Where("email = ?", user.Email).First(&User)
	err := authentication.AssertPassword(User.Password, []byte(user.Password))
	if err != nil {
		c.Status(401).JSON(err)
		return err
	}
	tokenPair, err := authorization.GenerateJWT(user)
	if err != nil {
		c.Status(401).JSON(err)
		return err
	}
	c.Status(201)
	c.JSON(tokenPair)
	return nil
}
