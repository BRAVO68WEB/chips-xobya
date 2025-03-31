package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"test-api/internal/server"
)

type Router struct {
	App       *server.App
	baseGroup fiber.Router
}

func New(a *server.App) *Router {
	a.Use(cors.New())

	return &Router{
		App:       a,
		baseGroup: a.Group("/api/v1"),
	}
}

func (r *Router) RegisterRoutes() {
	r.userRouter()
}
