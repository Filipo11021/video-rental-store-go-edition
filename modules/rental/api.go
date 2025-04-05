package rental

import (
	"time"

	"app/modules/film/film_contracts"
)

type Api interface {
	CreateRental(dto CreateRentalDTO) error
	GetAllRentals() ([]RentalDTO, error)
	GetRentalById(id int) (*RentalDTO, error)
	ReturnRental(id int) error
	GetRentalsByFilmId(filmId int) ([]RentalDTO, error)
	CalculatePrice(filmId int, days int) (float64, error)
	CalculateLateCharge(filmId int, extraDays int) (float64, error)
}

type api struct {
	rentalRepo      rentalRepo
	priceCalculator priceCalculator
	filmApi         film_contracts.Api
}

func newApi(rentalRepo rentalRepo, priceCalculator priceCalculator, filmApi film_contracts.Api) Api {

	return &api{
		rentalRepo:      rentalRepo,
		priceCalculator: priceCalculator,
		filmApi:         filmApi,
	}
}

func (a *api) CreateRental(dto CreateRentalDTO) error {
	startDate := time.Now()
	endDate := startDate.AddDate(0, 0, dto.Days)

	rental := &rental{
		FilmID:    dto.FilmID,
		StartDate: startDate,
		EndDate:   endDate,
	}

	return a.rentalRepo.create(rental)
}

func (a *api) GetAllRentals() ([]RentalDTO, error) {
	rentals, err := a.rentalRepo.findAll()
	if err != nil {
		return nil, err
	}

	rentalsDTO := make([]RentalDTO, len(rentals))
	for i, rental := range rentals {
		rentalsDTO[i] = rental.dto()
	}

	return rentalsDTO, nil
}

func (a *api) GetRentalById(id int) (*RentalDTO, error) {
	rental, err := a.rentalRepo.findById(id)
	if err != nil {
		return nil, err
	}
	if rental == nil {
		return nil, nil
	}
	dto := rental.dto()
	return &dto, nil
}

func (a *api) ReturnRental(id int) error {
	rental, err := a.rentalRepo.findById(id)
	if err != nil {
		return err
	}
	if rental == nil {
		return nil
	}

	rental.Returned = true
	return a.rentalRepo.update(rental)
}

func (a *api) GetRentalsByFilmId(filmId int) ([]RentalDTO, error) {
	rentals, err := a.rentalRepo.findByFilmId(filmId)
	if err != nil {
		return nil, err
	}

	rentalsDTO := make([]RentalDTO, len(rentals))
	for i, rental := range rentals {
		rentalsDTO[i] = rental.dto()
	}

	return rentalsDTO, nil
}

func (a *api) CalculatePrice(filmId int, days int) (float64, error) {
	film, err := a.filmApi.GetFilmById(filmId)
	if err != nil {
		return -1, err
	}
	return a.priceCalculator.calculatePrice(film.Type, days), nil
}

func (a *api) CalculateLateCharge(filmId int, extraDays int) (float64, error) {
	film, err := a.filmApi.GetFilmById(filmId)
	if err != nil {
		return -1, err
	}
	return a.priceCalculator.calculateLateCharge(film.Type, extraDays), nil
}
