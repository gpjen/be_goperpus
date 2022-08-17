package users

import (
	"errors"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]Users, error)
	FindById(Id uint64) (Users, error)
	Create(user Users) (Users, error)
	Update(user Users) (Users, error)
	Delete(user Users) (Users, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Users, error) {
	var users []Users
	err := r.db.Find(&users).Error
	return users, err
}

func (r *repository) FindById(Id uint64) (Users, error) {
	var user Users
	err := r.db.Find(&user, Id).Error
	if user.ID < 1 {
		newErr := errors.New("no data on this id")
		return user, newErr
	}
	return user, err
}

func (r *repository) Create(user Users) (Users, error) {
	err := r.db.Create(&user).Error
	return user, err
}

func (r *repository) Update(user Users) (Users, error) {
	err := r.db.Save(&user).Error
	return user, err
}

func (r *repository) Delete(user Users) (Users, error) {
	err := r.db.Delete(&user).Error
	return user, err
}
