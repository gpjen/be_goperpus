package users

type Service interface {
	FindAll() ([]Users, error)
	FindById(ID uint64) (Users, error)
	Register(newUser UserRequest) (Users, error)
	// Login(email string, pass string) (Users, error)
	// Update(user Users) (Users, error)
	// SofDelete(user UserRequest) (Users, error)
}

type service struct {
	repository Repository
}

func NewUserService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Users, error) {
	return s.repository.FindAll()
}

func (s *service) FindById(ID uint64) (Users, error) {
	return s.repository.FindById(ID)
}

func (s *service) Register(newUser UserRequest) (Users, error) {
	userData := Users{
		Name:       newUser.Name,
		Email:      newUser.Email,
		ImgProfile: newUser.ImgProfile,
		Password:   newUser.Password,
	}

	newData, err := s.repository.Create(userData)
	return newData, err
}
