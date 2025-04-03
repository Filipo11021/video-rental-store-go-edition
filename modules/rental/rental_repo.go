package rental

import (
	"gorm.io/gorm"
)

type rentalRepo interface {
	create(rental *Rental) error
	findAll() ([]*Rental, error)
	findById(id int) (*Rental, error)
	findByFilmId(filmId int) ([]*Rental, error)
	update(rental *Rental) error
}

type gormRentalRepo struct {
	db *gorm.DB
}

func newGormRentalRepo(db *gorm.DB) rentalRepo {
	return &gormRentalRepo{db: db}
}

func (r *gormRentalRepo) create(rental *Rental) error {
	return r.db.Create(rental).Error
}

func (r *gormRentalRepo) findAll() ([]*Rental, error) {
	var rentals []*Rental
	if err := r.db.Find(&rentals).Error; err != nil {
		return nil, err
	}
	return rentals, nil
}

func (r *gormRentalRepo) findById(id int) (*Rental, error) {
	var rental Rental
	if err := r.db.First(&rental, id).Error; err != nil {
		return nil, err
	}
	return &rental, nil
}

func (r *gormRentalRepo) findByFilmId(filmId int) ([]*Rental, error) {
	var rentals []*Rental
	if err := r.db.Where("film_id = ?", filmId).Find(&rentals).Error; err != nil {
		return nil, err
	}
	return rentals, nil
}

func (r *gormRentalRepo) update(rental *Rental) error {
	return r.db.Save(rental).Error
}
