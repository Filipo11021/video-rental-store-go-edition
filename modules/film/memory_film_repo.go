package film

import "sync"

type memoryFilmRepo struct {
	mutex  sync.RWMutex
	films  map[int]*film
	nextID int
}

func newMemoryFilmRepo() filmRepo {
	return &memoryFilmRepo{
		films:  make(map[int]*film),
		nextID: 1,
	}
}

func (r *memoryFilmRepo) create(film *film) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	film.ID = r.nextID
	r.films[film.ID] = film
	r.nextID++

	return nil
}

func (r *memoryFilmRepo) findAll() ([]*film, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	films := make([]*film, 0, len(r.films))
	for _, film := range r.films {
		films = append(films, film)
	}

	return films, nil
}

func (r *memoryFilmRepo) findById(id int) (*film, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	film, exists := r.films[id]
	if !exists {
		return nil, nil
	}
	return film, nil
}
