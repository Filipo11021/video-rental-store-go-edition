package rental

import (
	"time"
)

type rental struct {
	ID        int       `gorm:"primaryKey"`
	FilmID    int       `gorm:"not null"`
	StartDate time.Time `gorm:"not null"`
	EndDate   time.Time `gorm:"not null"`
	Returned  bool      `gorm:"default:false"`
}

func (rental) TableName() string {
	return "rentals"
}

func (r *rental) dto() RentalDTO {
	return RentalDTO{
		ID:        r.ID,
		FilmID:    r.FilmID,
		StartDate: r.StartDate,
		EndDate:   r.EndDate,
		Returned:  r.Returned,
	}
}
