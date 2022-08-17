package users

type Services interface {
	FindAll() ([]Users, error)
	Login(email string, pass string) (Users, error)
	Register(newUser UserRequest) (Users, error)
}
