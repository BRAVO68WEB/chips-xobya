package postgres

import (
	"github.com/google/uuid"
	"gorm.io/gorm"

	"test-api/internal/domain/todo"
)

type todoRepo struct {
	DB *gorm.DB
}

func NewtodoRepo(db *gorm.DB) user.Repository {
	return &todoRepo{DB: db}
}

func (r *todoRepo) Create(u *todo.todo) error {
	if err := r.DB.Create(u).Error; err != nil {
		return err
	}

	return nil
}

func (r *todoRepo) Get(id uuid.UUID) (*todo.todo, error) {
	u := new(todo.todo)

	if err := r.DB.First(u, id).Error; err != nil {
		return nil, err
	}

	return u, nil
}

func (r *todoRepo) GetAll() ([]todo.todo, error) {
	var todos []todo.todo

	if err := r.DB.Find(&todos).Error; err != nil {
		return nil, err
	}

	return todos, nil
}

func (r *todoRepo) Update(u *todo.todo) error {
	if err := r.DB.Where("id = ?", u.ID.String()).Updates(&u).Error; err != nil {
		return err
	}

	return nil
}

func (r *todoRepo) Delete(id uuid.UUID) error {
	if err := r.DB.Delete(&todo.todo{}, "id = ? ", id.String()).Error; err != nil {
		return err
	}

	return nil
}

func (r *todoRepo) BulkCreate(u []todo.todo) error {
	if err := r.DB.Create(&u).Error; err != nil {
		return err
	}

	return nil
}

