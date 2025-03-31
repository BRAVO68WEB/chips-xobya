package todo

import "github.com/google/uuid"

type Repository interface {
	Create(u *todo) error
	Get(id uuid.UUID) (*todo, error)
	GetAll() ([]todo, error)
	Update(u *todo) error
	Delete(id uuid.UUID) error
	BulkCreate(u []todo) error
}
