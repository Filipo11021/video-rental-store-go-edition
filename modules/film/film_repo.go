package film

import (
	"gorm.io/gorm"
)

// "if errors.Is(err, ErrFilmAlreadyExist)"
// var ErrFilmAlreadyExist = errors.New("film already exists")

type filmRepo interface {
	create(film *film) error
	findAll() ([]*film, error)
	findById(id int) (*film, error)
}

type gormFilmRepo struct {
	db *gorm.DB
}

func newGormFilmRepo(db *gorm.DB) filmRepo {
	return &gormFilmRepo{db: db}
}

func (r *gormFilmRepo) create(film *film) error {
	return r.db.Create(film).Error
}

func (r *gormFilmRepo) findAll() ([]*film, error) {
	var films []*film
	if err := r.db.Find(&films).Error; err != nil {
		return nil, err
	}
	return films, nil
}

func (r *gormFilmRepo) findById(id int) (*film, error) {
	var film film
	if err := r.db.First(&film, id).Error; err != nil {
		return nil, err
	}
	return &film, nil
}
