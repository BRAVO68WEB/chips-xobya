package todo

import "github.com/google/uuid"

type Service interface {
	Createtodo(u *todo) error
	Gettodo(id uuid.UUID) (*todo, error)
	GetAlltodos() ([]todo, error)
	Updatetodo(id uuid.UUID, u *todo) error
	Deletetodo(id uuid.UUID) error
	BulkCreatetodos(u []todo) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) Createtodo(u *todo) error {
	return s.repo.Create(u)
}

func (s *service) Gettodo(id uuid.UUID) (*todo, error) {
	return s.repo.Get(id)
}

func (s *service) GetAlltodos() ([]todo, error) {
	return s.repo.GetAll()
}

func (s *service) Updatetodo(id uuid.UUID, u *todo) error {
	u.ID = id
	return s.repo.Update(u)
}

func (s *service) Deletetodo(id uuid.UUID) error {
	return s.repo.Delete(id)
}

func (s *service) BulkCreatetodos(u []todo) error {
	return s.repo.BulkCreate(u)
}
