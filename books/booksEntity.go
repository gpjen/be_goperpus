package books

import "time"

// entity books
type Books struct {
	ID        uint   `gorm:"primary_key;autoIncrement;not null"`
	Title     string `gorm:"size:255"`
	Author    string `gorm:"size:100"`
	Desc      string `gorm:"type:text"`
	Image     string `gorm:"size:255"`
	Price     uint
	Discound  uint `gorm:"size:3"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
