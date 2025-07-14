package models

import (
	"time"

	"github.com/google/uuid"
)

type Post struct {
	ID        uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content" gorm:"type:text"`
	AuthorID  uuid.UUID `json:"author_id" gorm:"type:uuid"`
	Author    User      `json:"author" gorm:"foreignKey:AuthorID"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}


