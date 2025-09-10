package friendships

import (
	"github.com/google/uuid"
	"github.com/jafferhussain11/celeb-social/models/users"
	"gorm.io/gorm"
)

type Friendship struct {
	gorm.Model
	UserID   uuid.UUID `gorm:"uniqueIndex:idx_user_friend;type:uuid;not null" json:"user_id"`
	FriendID uuid.UUID `gorm:"uniqueIndex:idx_user_friend;type:uuid;not null" json:"friend_id"`

	User   users.User `gorm:"foreignKey:UserID";references:"ID" json:"-"`
	Friend users.User `gorm:"foreignKey:FriendID";references:"ID" json:"-"`
}
