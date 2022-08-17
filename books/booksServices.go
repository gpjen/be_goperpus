package books

type Services interface {
	FindAll() ([]Books, error)
	FindById(Id int) (Books, error)
	Create(bookRequest BookRequest) (Books, error)
	Update(bookRequest BookRequest, Id int) (Books, error)
	Delete(Id int) (Books, error)
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

	bookData := Books{
		Title:    bookRequest.Title,
		Author:   bookRequest.Author,
		Desc:     bookRequest.Desc,
		Image:    bookRequest.Image,
		Price:    bookRequest.Price,
		Discound: bookRequest.Discound,
		Rating:   bookRequest.Rating,
	}

	data, err := s.repository.Create(bookData)
	return data, err
}

func (s *services) Update(bookRequest BookRequest, Id int) (Books, error) {

	getData, err := s.repository.FindById(Id)

	if err != nil {
		return getData, err
	}

	getData.Title = bookRequest.Title
	getData.Author = bookRequest.Author
	getData.Desc = bookRequest.Desc
	getData.Image = bookRequest.Image
	getData.Price = bookRequest.Price
	getData.Discound = bookRequest.Discound

	bookUpdate, err := s.repository.Update(getData)

	return bookUpdate, err
}

func (s *services) Delete(Id int) (Books, error) {
	getData, err := s.repository.FindById(Id)
	if err != nil {
		return getData, err
	}

	data, err := s.repository.Delete(getData)
	return data, err
}
