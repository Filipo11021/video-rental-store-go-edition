package film_internal

import "app/modules/film/film_contracts"

type FilmType string

const (
	NewRelease FilmType = "NEW_RELEASE"
	Regular    FilmType = "REGULAR"
	Old        FilmType = "OLD"
)

type Film struct {
	ID    int      `gorm:"primaryKey"`
	Title string   `gorm:"column:name"`
	Type  FilmType `gorm:"column:type"`
}

func (Film) TableName() string {
	return "films"
}

func (a *Film) Dto() film_contracts.FilmDTO {
	return film_contracts.FilmDTO{
		ID:    a.ID,
		Title: a.Title,
		Type:  film_contracts.FilmTypeDto(a.Type),
	}
}
