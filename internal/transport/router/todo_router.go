package router

import (
	"test-api/internal/domain/todo"
	"test-api/internal/infrastructure/postgres"
	"test-api/internal/transport/handler"
)

func (r *Router) todoRouter() {
	repo := postgres.NewtodoRepo(r.App.DB)
	service := todo.NewService(repo)
	h := handler.NewtodoHandler(service)

	g := r.baseGroup.Group("/todo")
	{
		g.Post("/", h.Createtodo())
		g.Get("/", h.GetAlltodos())
		g.Get("/:id", h.Gettodo())
		g.Put("/:id", h.Updatetodo())
		g.Delete("/:id", h.Deletetodo())
		g.Post("/bulk", h.BulkCreatetodos())
	}
}
