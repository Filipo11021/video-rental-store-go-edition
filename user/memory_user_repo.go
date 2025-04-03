package user

import (
	"sync"
)

type memoryUserRepo struct {
	users  map[string]*User
	mu     sync.RWMutex
	nextID int
}

func newMemoryUserRepo() userRepo {
	return &memoryUserRepo{
		users:  make(map[string]*User),
		nextID: 1,
	}
}

func (r *memoryUserRepo) findById(id string) (*User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	user, exists := r.users[id]
	if !exists {
		return nil, nil
	}
	return user, nil
}
