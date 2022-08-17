package books

import (
	"errors"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]Books, error)
	FindById(Id int) (Books, error)
	Create(book Books) (Books, error)
	Update(book Books) (Books, error)
	Delete(book Books) (Books, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

// get books
func (r *repository) FindAll() ([]Books, error) {
	var books []Books
	err := r.db.Find(&books).Error
	return books, err
}

// get book by Id
func (r *repository) FindById(Id int) (Books, error) {
	var book Books
	err := r.db.Find(&book, Id).Error
	if book.ID < 1 {
		err := errors.New("no data on this id")
		return Books{}, err
	}
	return book, err
}

// create
func (r *repository) Create(book Books) (Books, error) {
	err := r.db.Create(&book).Error
	return book, err
}

// update book
func (r *repository) Update(book Books) (Books, error) {
	err := r.db.Save(&book).Error
	return book, err
}

// delete book
func (r *repository) Delete(book Books) (Books, error) {
	err := r.db.Delete(&book).Error
	return book, err
}
