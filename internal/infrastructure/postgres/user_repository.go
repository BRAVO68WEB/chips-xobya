package postgres

import (
	"github.com/google/uuid"
	"gorm.io/gorm"

	"test-api/internal/domain/user"
)

type userRepo struct {
	DB *gorm.DB
}

func NewuserRepo(db *gorm.DB) user.Repository {
	return &userRepo{DB: db}
}

func (r *userRepo) Create(u *user.user) error {
	if err := r.DB.Create(u).Error; err != nil {
		return err
	}

	return nil
}

func (r *userRepo) Get(id uuid.UUID) (*user.user, error) {
	u := new(user.user)

	if err := r.DB.First(u, id).Error; err != nil {
		return nil, err
	}

	return u, nil
}

func (r *userRepo) GetAll() ([]user.user, error) {
	var users []user.user

	if err := r.DB.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (r *userRepo) Update(u *user.user) error {
	if err := r.DB.Where("id = ?", u.ID.String()).Updates(&u).Error; err != nil {
		return err
	}

	return nil
}

func (r *userRepo) Delete(id uuid.UUID) error {
	if err := r.DB.Delete(&user.user{}, "id = ? ", id.String()).Error; err != nil {
		return err
	}

	return nil
}

func (r *userRepo) BulkCreate(u []user.user) error {
	if err := r.DB.Create(&u).Error; err != nil {
		return err
	}

	return nil
}

