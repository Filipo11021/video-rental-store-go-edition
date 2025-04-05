package film_contracts

type FilmDTO struct {
	ID    int      `json:"id"`
	Title string   `json:"title"`
	Type  FilmTypeDto `json:"type"`
}

type FilmTypeDto string

const (
	NewRelease FilmTypeDto = "NEW_RELEASE"
	Regular    FilmTypeDto = "REGULAR"
	Old        FilmTypeDto = "OLD"
)

type Api interface {
	CreateFilm(filmDTO FilmDTO) error
	GetAllFilms() ([]FilmDTO, error)
	GetFilmById(id int) (*FilmDTO, error)
}