package film_domain

import (
	"gorm.io/gorm"
)

// "if errors.Is(err, ErrFilmAlreadyExist)"
// var ErrFilmAlreadyExist = errors.New("film already exists")

type FilmRepo interface {
	Create(film *Film) error
	FindAll() ([]*Film, error)
	FindById(id int) (*Film, error)
}

type gormFilmRepo struct {
	db *gorm.DB
}

func NewGormFilmRepo(db *gorm.DB) FilmRepo {
	return &gormFilmRepo{db: db}
}

func (r *gormFilmRepo) Create(film *Film) error {
	return r.db.Create(film).Error
}

func (r *gormFilmRepo) FindAll() ([]*Film, error) {
	var films []*Film
	if err := r.db.Find(&films).Error; err != nil {
		return nil, err
	}
	return films, nil
}

func (r *gormFilmRepo) FindById(id int) (*Film, error) {
	var film Film
	if err := r.db.First(&film, id).Error; err != nil {
		return nil, err
	}
	return &film, nil
}
