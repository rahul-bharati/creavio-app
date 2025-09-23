package repository

type UserRepository interface {
	GetGreeting() (string, error)
}
