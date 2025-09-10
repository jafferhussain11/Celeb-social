package posts

import (
	"github.com/google/uuid"
	"github.com/jafferhussain11/celeb-social/models/users"
)

type Post struct {
	ID        uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()" json:"id"`
	UserID    uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
	Content   string    `json:"content"`
	CreatedAt int64     `json:"created_at"`
	UpdatedAt int64     `json:"updated_at"`

	User users.User `gorm:"foreignKey:UserID;references:ID" json:"-"`
}
