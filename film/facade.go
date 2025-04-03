package film

type FilmDTO struct {
	ID    int      `json:"id"`
	Title string   `json:"title"`
	Type  FilmType `json:"type"`
}

type Facade interface {
	CreateFilm(filmDTO FilmDTO) error
	GetAllFilms() ([]FilmDTO, error)
	GetFilmById(id int) (*FilmDTO, error)
}

type facade struct {
	filmRepo filmRepo
}

func newFacade(filmRepo filmRepo) Facade {
	return &facade{filmRepo: filmRepo}
}

func (f *facade) CreateFilm(filmDTO FilmDTO) error {
	newFilm := &film{
		Title: filmDTO.Title,
		Type:  filmDTO.Type,
	}

	return f.filmRepo.create(newFilm)
}

func (f *facade) GetAllFilms() ([]FilmDTO, error) {
	films, err := f.filmRepo.findAll()
	if err != nil {
		return nil, err
	}

	filmsDTO := make([]FilmDTO, len(films))
	for i, film := range films {
		filmsDTO[i] = film.dto()
	}

	return filmsDTO, nil
}

func (f *facade) GetFilmById(id int) (*FilmDTO, error) {
	film, err := f.filmRepo.findById(id)
	if err != nil {
		return nil, err
	}
	if film == nil {
		return nil, nil
	}
	dto := film.dto()
	return &dto, nil
}
