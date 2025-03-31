package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"test-api/internal/domain/user"
)

type todoHandler struct {
	service todo.Service
}

func NewtodoHandler(service todo.Service) *todoHandler {
	return &todoHandler{service: service}
}

func (h *todoHandler) Createtodo() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		u := new(todo.todo)
		if err := ctx.BodyParser(u); err != nil {
			return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
				"error": "Invalid request payload : " + err.Error(),
			})
		}

		if err := h.service.Createtodo(u); err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to create todo: " + err.Error(),
			})
		}

		return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{})
	}
}

func (h *todoHandler) Gettodo() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id, err := uuid.Parse(ctx.Params("id"))
		if err != nil {
			return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
				"error": "Invalid id : " + err.Error(),
			})
		}

		u, err := h.service.Gettodo(id)
		if err != nil {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "todo not found: " + err.Error(),
			})
		}

		return ctx.Status(fiber.StatusOK).JSON(u)
	}
}

func (h *todoHandler) GetAlltodos() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		todos, err := h.service.GetAlltodos()
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to get todos: " + err.Error(),
			})
		}

		return ctx.Status(fiber.StatusOK).JSON(todos)
	}
}

func (h *todoHandler) Updatetodo() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id, err := uuid.Parse(ctx.Params("id"))
		if err != nil {
			return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
				"error": "Invalid id : " + err.Error(),
			})
		}

		u := new(todo.todo)
		if err := ctx.BodyParser(u); err != nil {
			return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
				"error": "Invalid request payload : " + err.Error(),
			})
		}

		if err := h.service.Updatetodo(id, u); err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to update todo: " + err.Error(),
			})
		}

		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{})
	}
}

func (h *todoHandler) Deletetodo() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id, err := uuid.Parse(ctx.Params("id"))
		if err != nil {
			return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
				"error": "Invalid id : " + err.Error(),
			})
		}

		if err := h.service.Deletetodo(id); err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to delete todo: " + err.Error(),
			})
		}

		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{})
	}
}

func (h *todoHandler) BulkCreatetodos() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		todos := make([]todo.todo, 0)
		if err := ctx.BodyParser(&todos); err != nil {
			return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
				"error": "Invalid request payload : " + err.Error(),
			})
		}

		if err := h.service.BulkCreatetodos(todos); err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to bulk create todos: " + err.Error(),
			})
		}

		return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{})
	}
}

