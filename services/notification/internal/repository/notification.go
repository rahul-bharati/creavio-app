package repository

type NotificationRepository interface {
	GetGreeting() (string, error)
}
