package memory

type MemoryNotificationRepository struct {
	message string
}

func NewMemoryNotificationRepository() *MemoryNotificationRepository {
	return &MemoryNotificationRepository{message: "Hello from Notification Service!"}
}

func (r *MemoryNotificationRepository) GetGreeting() (string, error) {
	return r.message, nil
}
