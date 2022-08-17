package books

import (
	"time"
)

// entity books
type Books struct {
	ID        uint64    `gorm:"primary_key;auto_increment"`
	Title     string    `gorm:"type:varchar(100);not null"`
	Author    string    `gorm:"type:varchar(50);not null"`
	Desc      string    `gorm:"type:varchar(200)"`
	Image     string    `gorm:"type:varchar(200);default:noimagebooks.jpg"`
	Price     uint      `gorm:"default:0"`
	Discound  uint      `gorm:"size:3;default:0"`
	Rating    uint      `gorm:"size:1;default:0"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}
