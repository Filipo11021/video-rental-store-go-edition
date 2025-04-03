package rental

import (
	"sync"
)

type memoryRentalRepo struct {
	mutex   sync.RWMutex
	rentals map[int]*Rental
	nextID  int
}

func newMemoryRentalRepo() rentalRepo {
	return &memoryRentalRepo{
		rentals: make(map[int]*Rental),
		nextID:  1,
	}
}

func (r *memoryRentalRepo) create(rental *Rental) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	rental.ID = r.nextID
	r.rentals[rental.ID] = rental
	r.nextID++
	return nil
}

func (r *memoryRentalRepo) findAll() ([]*Rental, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	rentals := make([]*Rental, 0, len(r.rentals))
	for _, rental := range r.rentals {
		rentals = append(rentals, rental)
	}
	return rentals, nil
}

func (r *memoryRentalRepo) findById(id int) (*Rental, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	rental, exists := r.rentals[id]
	if !exists {
		return nil, nil
	}
	return rental, nil
}

func (r *memoryRentalRepo) findByFilmId(filmId int) ([]*Rental, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	var rentals []*Rental
	for _, rental := range r.rentals {
		if rental.FilmID == filmId {
			rentals = append(rentals, rental)
		}
	}
	return rentals, nil
}

func (r *memoryRentalRepo) update(rental *Rental) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.rentals[rental.ID]; !exists {
		return nil
	}
	r.rentals[rental.ID] = rental
	return nil
}
