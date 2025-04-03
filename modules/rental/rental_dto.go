package rental

import "time"

type RentalDTO struct {
	ID        int       `json:"id"`
	FilmID    int       `json:"film_id"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	Returned  bool      `json:"returned"`
}

type CreateRentalDTO struct {
	FilmID int `json:"film_id"`
	Days   int `json:"days"`
}
