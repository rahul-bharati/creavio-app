package service

import "github.com/creavio/services/user/internal/repository"

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}
func (s *UserService) GetGreeting() (string, error) {
	return s.repo.GetGreeting()
}
