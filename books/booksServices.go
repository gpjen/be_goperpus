package books

type Serfices interface {
	FindAll() ([]Books, error)
	FindById(Id int) (Books, error)
	Create(bookRequest BookRequest) (Books, error)
	// Update(bookRequest BookRequest, Id int) (Books, error)
	// Delete(bookRequest BookRequest) (Books, error)
}

type services struct {
	repository Repository
}

func NewServices(repository Repository) *services {
	return &services{repository}
}

func (s *services) FindAll() ([]Books, error) {
	return s.repository.FindAll()
}

func (s *services) FindById(Id int) (Books, error) {
	return s.repository.FindById(Id)
}

func (s *services) Create(bookRequest BookRequest) (Books, error) {

	price, _ := bookRequest.Price.Int64()
	discount, _ := bookRequest.Price.Int64()
	rating, _ := bookRequest.Rating.Int64()

	bookData := Books{
		Title:    bookRequest.Title,
		Author:   bookRequest.Author,
		Desc:     bookRequest.Desc,
		Image:    bookRequest.Image,
		Price:    uint(price),
		Discound: uint(discount),
		Rating:   uint(rating),
	}

	data, err := s.repository.Create(bookData)
	return data, err
}

// func (s *services) Update(bookRequest BookRequest, Id int) (Books, error) {

// 	price, _ := bookRequest.Price.Int64()
// 	discound, _ := bookRequest.Discound.Int64()

// 	dataUpdate := Books{
// 		Title:    bookRequest.Title,
// 		Author:   bookRequest.Author,
// 		Desc:     bookRequest.Desc,
// 		Image:    bookRequest.Image,
// 		Price:    uint(price),
// 		Discound: uint(discound),
// 	}

// 	return s.repository.Update(&dataUpdate, Id).Error()
// }
