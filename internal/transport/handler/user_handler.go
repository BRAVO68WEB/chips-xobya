package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"test-api/internal/domain/user"
)

type userHandler struct {
	service user.Service
}

func NewuserHandler(service user.Service) *userHandler {
	return &userHandler{service: service}
}

func (h *userHandler) Createuser() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		u := new(user.user)
		if err := ctx.BodyParser(u); err != nil {
			return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
				"error": "Invalid request payload : " + err.Error(),
			})
		}

		if err := h.service.Createuser(u); err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to create user: " + err.Error(),
			})
		}

		return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{})
	}
}

func (h *userHandler) Getuser() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id, err := uuid.Parse(ctx.Params("id"))
		if err != nil {
			return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
				"error": "Invalid id : " + err.Error(),
			})
		}

		u, err := h.service.Getuser(id)
		if err != nil {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "user not found: " + err.Error(),
			})
		}

		return ctx.Status(fiber.StatusOK).JSON(u)
	}
}

func (h *userHandler) GetAllusers() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		users, err := h.service.GetAllusers()
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to get users: " + err.Error(),
			})
		}

		return ctx.Status(fiber.StatusOK).JSON(users)
	}
}

func (h *userHandler) Updateuser() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id, err := uuid.Parse(ctx.Params("id"))
		if err != nil {
			return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
				"error": "Invalid id : " + err.Error(),
			})
		}

		u := new(user.user)
		if err := ctx.BodyParser(u); err != nil {
			return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
				"error": "Invalid request payload : " + err.Error(),
			})
		}

		if err := h.service.Updateuser(id, u); err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to update user: " + err.Error(),
			})
		}

		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{})
	}
}

func (h *userHandler) Deleteuser() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id, err := uuid.Parse(ctx.Params("id"))
		if err != nil {
			return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
				"error": "Invalid id : " + err.Error(),
			})
		}

		if err := h.service.Deleteuser(id); err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to delete user: " + err.Error(),
			})
		}

		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{})
	}
}

func (h *userHandler) BulkCreateusers() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		users := make([]user.user, 0)
		if err := ctx.BodyParser(&users); err != nil {
			return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
				"error": "Invalid request payload : " + err.Error(),
			})
		}

		if err := h.service.BulkCreateusers(users); err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to bulk create users: " + err.Error(),
			})
		}

		return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{})
	}
}

