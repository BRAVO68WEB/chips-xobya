package user

import (
	"test-api/internal/pkg/shared"
)

type user struct {
	shared.Model
	name string `json:"name" gorm:"name"`email string `json:"email" gorm:"email"`password string `json:"password" gorm:"password"`
}
