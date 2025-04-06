package film_domain

import (
	"app/modules/film/film_contracts"
)

type api struct {
	filmRepo FilmRepo
}

func NewApi(filmRepo FilmRepo) film_contracts.Api {
	return &api{filmRepo: filmRepo}
}

func (a *api) CreateFilm(filmDTO film_contracts.FilmDTO) error {
	newFilm := &Film{
		Title: filmDTO.Title,
		Type:  FilmType(filmDTO.Type),
	}

	return a.filmRepo.Create(newFilm)
}

func (a *api) GetAllFilms() ([]film_contracts.FilmDTO, error) {
	films, err := a.filmRepo.FindAll()
	if err != nil {
		return nil, err
	}

	filmsDTO := make([]film_contracts.FilmDTO, len(films))
	for i, film := range films {
		filmsDTO[i] = film.Dto()
	}

	return filmsDTO, nil
}

func (a *api) GetFilmById(id int) (*film_contracts.FilmDTO, error) {
	film, err := a.filmRepo.FindById(id)
	if err != nil {
		return nil, err
	}
	if film == nil {
		return nil, nil
	}
	dto := film.Dto()
	return &dto, nil
}
