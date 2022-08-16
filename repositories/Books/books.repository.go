package books

import "gorm.io/gorm"

type Repository interface {
	Create(book Books) error
	FindAll() ([]Books, error)
	FindById(Id int) (Books, error)
	// Update()
	// Delete()
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Books, error) {
	var books []Books
	err := r.db.Find(&books).Error
	return books, err
}

func (r *repository) FindById(Id int) (Books, error) {
	var book Books
	err := r.db.Find(&book, Id).Error
	return book, err
}

func (r *repository) Create(book Books) error {
	err := r.db.Create(&book).Error
	return err
}
