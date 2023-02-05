package routes

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/sk25469/zocket-crud-api/pkg/controller"
)

func RegisterRoutes() {
	app := fiber.New()

	app.Get("/book", controller.GetAllBook)
	app.Post("/book", controller.AddNewBook)
	app.Put("/book/:id", controller.UpdateBookById)
	app.Delete("/book/:id", controller.DeleteBookById)

	log.Fatal(app.Listen(":3000"))
}
