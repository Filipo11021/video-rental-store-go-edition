package rental

import (
	"gorm.io/gorm"
)

type rentalRepo interface {
	create(rental *rental) error
	findAll() ([]*rental, error)
	findById(id int) (*rental, error)
	findByFilmId(filmId int) ([]*rental, error)
	update(rental *rental) error
}

type gormRentalRepo struct {
	db *gorm.DB
}

func newGormRentalRepo(db *gorm.DB) rentalRepo {
	return &gormRentalRepo{db: db}
}

func (r *gormRentalRepo) create(rental *rental) error {
	return r.db.Create(rental).Error
}

func (r *gormRentalRepo) findAll() ([]*rental, error) {
	var rentals []*rental
	if err := r.db.Find(&rentals).Error; err != nil {
		return nil, err
	}
	return rentals, nil
}

func (r *gormRentalRepo) findById(id int) (*rental, error) {
	var rental rental
	if err := r.db.First(&rental, id).Error; err != nil {
		return nil, err
	}
	return &rental, nil
}

func (r *gormRentalRepo) findByFilmId(filmId int) ([]*rental, error) {
	var rentals []*rental
	if err := r.db.Where("film_id = ?", filmId).Find(&rentals).Error; err != nil {
		return nil, err
	}
	return rentals, nil
}

func (r *gormRentalRepo) update(rental *rental) error {
	return r.db.Save(rental).Error
}
