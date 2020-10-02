package endpoints

import (
	"fmt"

	"github.com/drew138/games/api/authentication"
	"github.com/drew138/games/api/authorization"
	"github.com/drew138/games/database"
	"github.com/drew138/games/database/models"
	"github.com/gofiber/fiber/v2"
)

// CreateUser add new user to database
func CreateUser(c *fiber.Ctx) error {

	if !HasJSONBody(c) {
		return fmt.Errorf("Body does not contain JSON format")
	}
	var User models.User
	if UnmarshalJSON(c, &User) {
		return fmt.Errorf("Invalid user properties")
	}
	validationError := authentication.ValidatePassword(User.Password)
	if validationError != nil {
		c.Status(400).Send([]byte(validationError.Error()))
		return validationError
	}
	User.Password = authentication.HashGenerator([]byte(User.Password))
	dbError := database.DBConn.Create(&User).Error
	if dbError != nil {
		c.Status(500).Send([]byte(dbError.Error()))
		return dbError
	}
	tokenPair, err := authorization.GenerateJWT(&User)
	if err != nil {
		c.Status(500).Send([]byte(err.Error()))
		return err
	}
	userMap := map[string]interface{}{
		"email":        User.Email,
		"name":         User.Name,
		"surname":      User.Surname,
		"isAdmin":      User.IsAdmin, //TODO remove this field
		"accessToken":  tokenPair["accessToken"],
		"refreshToken": tokenPair["refreshToken"],
	}
	c.Status(201)
	c.JSON(userMap) // convert to json and send response
	return nil
}
