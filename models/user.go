package models

import (
	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email" gorm:"unique"`
	Password string    `json:"password"`
	Posts    []Post    `gorm:"foreignKey:AuthorID" json:"posts"`
}
