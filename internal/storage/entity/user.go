package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uint      `gorm:"primaryKey,column:id"`
	Name      string    `gorm:"column:name"`
	Age       float64   `gorm:"column:age"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

type Comment struct {
	gorm.Model
	ID        uint      `gorm:"primaryKey,column:id"`
	PostID    uint      `gorm:"column:post_id"`   // postID
	ParentID  uint      `gorm:"column:parent_id"` // parentID
	Comment   string    `gorm:"column:comment"`
	UserID    uint      `gorm:"column:user_id"`
	CreatedAt time.Time `gorm:"column:created_at"` // index
	UpdatedAt time.Time `gorm:"column:updated_at"`
}
