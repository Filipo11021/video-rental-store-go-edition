package rental

import (
	"time"
)

type Rental struct {
	ID        int       `gorm:"primaryKey"`
	FilmID    int       `gorm:"not null"`
	StartDate time.Time `gorm:"not null"`
	EndDate   time.Time `gorm:"not null"`
	Returned  bool      `gorm:"default:false"`
}

func (Rental) TableName() string {
	return "rentals"
}

func (r *Rental) dto() RentalDTO {
	return RentalDTO{
		ID:        r.ID,
		FilmID:    r.FilmID,
		StartDate: r.StartDate,
		EndDate:   r.EndDate,
		Returned:  r.Returned,
	}
}
