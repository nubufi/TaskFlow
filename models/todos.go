package models

import "gorm.io/gorm"

type TodoItem struct {
	gorm.Model
	Title  string `json:"title"`
	UserID string `json:"user_id"`
}
