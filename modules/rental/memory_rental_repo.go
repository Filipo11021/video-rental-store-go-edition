package rental

import (
	"sync"
)

type memoryRentalRepo struct {
	mutex   sync.RWMutex
	rentals map[int]*rental
	nextID  int
}

func newMemoryRentalRepo() rentalRepo {
	return &memoryRentalRepo{
		rentals: make(map[int]*rental),
		nextID:  1,
	}
}

func (r *memoryRentalRepo) create(rental *rental) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	rental.ID = r.nextID
	r.rentals[rental.ID] = rental
	r.nextID++
	return nil
}

func (r *memoryRentalRepo) findAll() ([]*rental, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	rentals := make([]*rental, 0, len(r.rentals))
	for _, rental := range r.rentals {
		rentals = append(rentals, rental)
	}
	return rentals, nil
}

func (r *memoryRentalRepo) findById(id int) (*rental, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	rental, exists := r.rentals[id]
	if !exists {
		return nil, nil
	}
	return rental, nil
}

func (r *memoryRentalRepo) findByFilmId(filmId int) ([]*rental, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	var rentals []*rental
	for _, rental := range r.rentals {
		if rental.FilmID == filmId {
			rentals = append(rentals, rental)
		}
	}
	return rentals, nil
}

func (r *memoryRentalRepo) update(rental *rental) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.rentals[rental.ID]; !exists {
		return nil
	}
	r.rentals[rental.ID] = rental
	return nil
}
