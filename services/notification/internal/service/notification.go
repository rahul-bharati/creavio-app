package service

import "github.com/creavio/services/notification/internal/repository"

type NotificationService struct {
	repo repository.NotificationRepository
}

func NewNotificationService(repo repository.NotificationRepository) *NotificationService {
	return &NotificationService{repo}
}

func (s *NotificationService) GetGreeting() (string, error) {
	return s.repo.GetGreeting()
}
