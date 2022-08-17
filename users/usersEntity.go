package users

import "time"

type Users struct {
	ID         uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Name       string    `gorm:"type:varchar(100);not null" json:"name"`
	Email      string    `gorm:"type:varchar(100);not null;uniqueIndex" json:"email"`
	ImgProfile string    `gorm:"type:varchar(255);default:noimageprofile.jpg" json:"img_profile"`
	Password   string    `gorm:"type:varchar(255);->;<-;not null" json:"-"`
	IsDelete   bool      `gorm:"default:false" json:"-"`
	CreatedAt  time.Time `gorm:""  json:"-"`
	UpdatedAt  time.Time `gorm:""  json:"-"`
	Token      string    `gorm:"-" json:"token,omitempty"`
}
