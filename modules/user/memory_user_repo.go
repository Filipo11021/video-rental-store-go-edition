package user

import (
	"sync"
)

type memoryUserRepo struct {
	users  map[string]*user
	mu     sync.RWMutex
	nextID int
}

func newMemoryUserRepo() userRepo {
	return &memoryUserRepo{
		users:  make(map[string]*user),
		nextID: 1,
	}
}

func (r *memoryUserRepo) findById(id string) (*user, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	user, exists := r.users[id]
	if !exists {
		return nil, nil
	}
	return user, nil
}
