// blog.go - contains blog model

package model

import "gorm.io/gorm"

type Blog struct {
	gorm.Model
	Title   string `json:"title"`
	Content string `json:"content"`
	UserID  uint   `json:"user_id"`
}
