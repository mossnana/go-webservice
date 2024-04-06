package server

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

type GoFiber struct {
	app *fiber.App
}

func NewGoFiber() *GoFiber {
	return &GoFiber{
		app: fiber.New(),
	}
}

func (fiber *GoFiber) Get(path string, handler fiber.Handler) fiber.Router {
	return fiber.app.Get(path, handler)
}

func (fiber *GoFiber) Post(path string, handler fiber.Handler) fiber.Router {
	return fiber.app.Post(path, handler)
}

func (fiber *GoFiber) Listen(url string) {
	log.Fatalln(fiber.app.Listen(url))
}
