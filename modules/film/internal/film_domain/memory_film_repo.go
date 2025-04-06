package film_domain

import "sync"

type memoryFilmRepo struct {
	mutex  sync.RWMutex
	films  map[int]*Film
	nextID int
}

func NewMemoryFilmRepo() FilmRepo {
	return &memoryFilmRepo{
		films:  make(map[int]*Film),
		nextID: 1,
	}
}

func (r *memoryFilmRepo) Create(film *Film) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	film.ID = r.nextID
	r.films[film.ID] = film
	r.nextID++

	return nil
}

func (r *memoryFilmRepo) FindAll() ([]*Film, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	films := make([]*Film, 0, len(r.films))
	for _, film := range r.films {
		films = append(films, film)
	}

	return films, nil
}

func (r *memoryFilmRepo) FindById(id int) (*Film, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	film, exists := r.films[id]
	if !exists {
		return nil, nil
	}
	return film, nil
}
