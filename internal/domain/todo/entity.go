package todo

import (
	"test-api/internal/pkg/shared"
)

type todo struct {
	shared.Model
	name string `json:"name" gorm:"name"`status boolean `json:"status" gorm:"status"`user_id string `json:"user_id" gorm:"user_id"`
}
