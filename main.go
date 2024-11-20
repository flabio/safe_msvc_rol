package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/msvc_rol/infrastructure/routers"
)

func main() {
	app := fiber.New()
	routers.NewRolModuleRouter(app)
	routers.NewRouter(app)
	routers.NewInstanceModuleRouter(app)

	app.Listen(":3001")
}
