package user

import "github.com/google/uuid"

type Service interface {
	Createuser(u *user) error
	Getuser(id uuid.UUID) (*user, error)
	GetAllusers() ([]user, error)
	Updateuser(id uuid.UUID, u *user) error
	Deleteuser(id uuid.UUID) error
	BulkCreateusers(u []user) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) Createuser(u *user) error {
	return s.repo.Create(u)
}

func (s *service) Getuser(id uuid.UUID) (*user, error) {
	return s.repo.Get(id)
}

func (s *service) GetAllusers() ([]user, error) {
	return s.repo.GetAll()
}

func (s *service) Updateuser(id uuid.UUID, u *user) error {
	u.ID = id
	return s.repo.Update(u)
}

func (s *service) Deleteuser(id uuid.UUID) error {
	return s.repo.Delete(id)
}

func (s *service) BulkCreateusers(u []user) error {
	return s.repo.BulkCreate(u)
}
