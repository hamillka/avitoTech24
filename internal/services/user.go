package services

type IUserRepository interface {
}

type UserService struct {
	repo IUserRepository
}

func NewUserService(repository IUserRepository) *UserService {
	return &UserService{repo: repository}
}
