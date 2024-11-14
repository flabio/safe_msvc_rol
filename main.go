package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/msvc_rol/infrastructure/routers"
)

func main() {
	app := fiber.New()
	// Custom CORS configuration

	// app.Use(cors.New(cors.Config{
	// 	AllowOrigins:     "http://34.207.102.227:3080", // Specify allowed origins
	// 	AllowMethods:     "GET,POST,PUT,DELETE",
	// 	AllowHeaders:     "Content-Type, Authorization",
	// 	ExposeHeaders:    "Content-Length",
	// 	AllowCredentials: true, // Allow credentials
	// }))
	routers.NewRouter(app)
	routers.NewInstanceModuleRouter(app)
	app.Listen(":3001")
}
