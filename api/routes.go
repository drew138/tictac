package api

import (
	"github.com/drew138/games/api/endpoints"
	"github.com/gofiber/fiber/v2"
)

// ResgisterEndPoints applies specified routes to fiber app
func ResgisterEndPoints(app *fiber.App) {
	// GET Endpoints
	// app.Get("/api/v1/solve", endpoints.SolveEquation)
	// app.Get("/")
	// POST Endpoints
	app.Post("/api/v1/register", endpoints.CreateUser)
	app.Post("/api/v1/login", endpoints.Login)
	app.Post("/api/v1/refresh", endpoints.RefreshJWT)
	//PUT Endpoints

	//PATCH Endpoints

}
