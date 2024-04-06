package controllers

import (
	"net/http"
	"todo/webservice/internal/entities"
	"todo/webservice/internal/usecases"

	"github.com/gofiber/fiber/v2"
)

type TodoController struct {
	todoUsecase *usecases.TodoUsecase
}

func NewTodoController(todoUsecase *usecases.TodoUsecase) *TodoController {
	return &TodoController{
		todoUsecase: todoUsecase,
	}
}

func (t *TodoController) GetByID(c *fiber.Ctx) error {
	todo, e := t.todoUsecase.GetByID(c.Params("id"))
	if e != nil {
		return c.SendStatus(http.StatusNotFound)
	}
	return c.JSON(todo)
}

func (t *TodoController) Create(c *fiber.Ctx) error {
	draft := entities.DraftTodo{}
	e := c.BodyParser(&draft)
	if e != nil {
		return c.SendStatus(http.StatusBadRequest)
	}

	todo, e := t.todoUsecase.Create(&draft)
	if e != nil {
		return c.SendStatus(http.StatusNotFound)
	}

	return c.JSON(todo)
}
