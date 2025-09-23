package repository

type MemoryHelloRepository struct {
	message string
}

func NewMemoryHelloRepository() *MemoryHelloRepository {
	return &MemoryHelloRepository{message: "Hello, World!"}
}

func (r *MemoryHelloRepository) GetGreeting() (string, error) {
	return r.message, nil
}
