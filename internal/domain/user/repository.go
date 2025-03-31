package user

import "github.com/google/uuid"

type Repository interface {
	Create(u *user) error
	Get(id uuid.UUID) (*user, error)
	GetAll() ([]user, error)
	Update(u *user) error
	Delete(id uuid.UUID) error
	BulkCreate(u []user) error
}
