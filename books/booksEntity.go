package books

import (
	"encoding/json"
	"time"
)

// entity books
type Books struct {
	ID        uint64      `gorm:"primary_key;auto_increment"`
	Title     string      `gorm:"size:255;not null"`
	Author    string      `gorm:"size:100;not null"`
	Desc      string      `gorm:"type:text"`
	Image     string      `gorm:"size:255"`
	Price     json.Number `gorm:"not null;default:0"`
	Discound  json.Number `gorm:"size:3;not null;default:0"`
	Rating    json.Number `gorm:"size:1;not null;default:0"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
