package books

import "gorm.io/gorm"

type Repository interface {
	Create(book Books) error
	FindAll() ([]Books, error)
	FindById(Id int) (Books, error)
	Update(book Books) error
	Delete(book Books) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

// create
func (r *repository) Create(book Books) error {
	return r.db.Create(&book).Error
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
	return book, err
}

// update book
func (r *repository) Update(book Books) error {
	return r.db.Save(&book).Error
}

// delete book
func (r *repository) Delete(book Books) error {
	return r.db.Delete(&book).Error
}
