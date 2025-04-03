package film

type FilmDTO struct {
	ID    int      `json:"id"`
	Title string   `json:"title"`
	Type  FilmType `json:"type"`
}

type Api interface {
	CreateFilm(filmDTO FilmDTO) error
	GetAllFilms() ([]FilmDTO, error)
	GetFilmById(id int) (*FilmDTO, error)
}

type api struct {
	filmRepo filmRepo
}

func newApi(filmRepo filmRepo) Api {
	return &api{filmRepo: filmRepo}
}

func (a *api) CreateFilm(filmDTO FilmDTO) error {
	newFilm := &film{
		Title: filmDTO.Title,
		Type:  filmDTO.Type,
	}

	return a.filmRepo.create(newFilm)
}

func (a *api) GetAllFilms() ([]FilmDTO, error) {
	films, err := a.filmRepo.findAll()
	if err != nil {
		return nil, err
	}

	filmsDTO := make([]FilmDTO, len(films))
	for i, film := range films {
		filmsDTO[i] = film.dto()
	}

	return filmsDTO, nil
}

func (a *api) GetFilmById(id int) (*FilmDTO, error) {
	film, err := a.filmRepo.findById(id)
	if err != nil {
		return nil, err
	}
	if film == nil {
		return nil, nil
	}
	dto := film.dto()
	return &dto, nil
}
