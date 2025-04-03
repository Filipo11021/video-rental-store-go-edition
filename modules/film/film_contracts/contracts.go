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
