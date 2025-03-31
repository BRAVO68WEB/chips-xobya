package router

import (
	"test-api/internal/domain/user"
	"test-api/internal/infrastructure/postgres"
	"test-api/internal/transport/handler"
)

func (r *Router) userRouter() {
	repo := postgres.NewuserRepo(r.App.DB)
	service := user.NewService(repo)
	h := handler.NewuserHandler(service)

	g := r.baseGroup.Group("/user")
	{
		g.Post("/", h.Createuser())
		g.Get("/", h.GetAllusers())
		g.Get("/:id", h.Getuser())
		g.Put("/:id", h.Updateuser())
		g.Delete("/:id", h.Deleteuser())
		g.Post("/bulk", h.BulkCreateusers())
	}
}
