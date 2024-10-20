package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name      string    `gorm:"type:varchar(200)"`
	Email     string    `gorm:"type:varchar(200);unique"`
	Password  string    `gorm:"type:varchar(200)"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"` // Untuk soft delete
}
