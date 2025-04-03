package rental

import (
	"time"

	"app/film"
)

type Facade interface {
	CreateRental(dto CreateRentalDTO) error
	GetAllRentals() ([]RentalDTO, error)
	GetRentalById(id int) (*RentalDTO, error)
	ReturnRental(id int) error
	GetRentalsByFilmId(filmId int) ([]RentalDTO, error)
	CalculatePrice(filmId int, days int) (float64, error)
	CalculateLateCharge(filmId int, extraDays int) (float64, error)
}

type facade struct {
	rentalRepo      rentalRepo
	priceCalculator PriceCalculator
	filmFacade      film.Facade
}

func newFacade(rentalRepo rentalRepo, priceCalculator PriceCalculator, filmFacade film.Facade) Facade {

	return &facade{
		rentalRepo:      rentalRepo,
		priceCalculator: priceCalculator,
		filmFacade:      filmFacade,

	}
}

func (f *facade) CreateRental(dto CreateRentalDTO) error {
	startDate := time.Now()
	endDate := startDate.AddDate(0, 0, dto.Days)

	rental := &Rental{
		FilmID:    dto.FilmID,
		StartDate: startDate,
		EndDate:   endDate,
	}

	return f.rentalRepo.create(rental)
}

func (f *facade) GetAllRentals() ([]RentalDTO, error) {
	rentals, err := f.rentalRepo.findAll()
	if err != nil {
		return nil, err
	}

	rentalsDTO := make([]RentalDTO, len(rentals))
	for i, rental := range rentals {
		rentalsDTO[i] = rental.dto()
	}

	return rentalsDTO, nil
}

func (f *facade) GetRentalById(id int) (*RentalDTO, error) {
	rental, err := f.rentalRepo.findById(id)
	if err != nil {
		return nil, err
	}
	if rental == nil {
		return nil, nil
	}
	dto := rental.dto()
	return &dto, nil
}

func (f *facade) ReturnRental(id int) error {
	rental, err := f.rentalRepo.findById(id)
	if err != nil {
		return err
	}
	if rental == nil {
		return nil
	}

	rental.Returned = true
	return f.rentalRepo.update(rental)
}

func (f *facade) GetRentalsByFilmId(filmId int) ([]RentalDTO, error) {
	rentals, err := f.rentalRepo.findByFilmId(filmId)
	if err != nil {
		return nil, err
	}

	rentalsDTO := make([]RentalDTO, len(rentals))
	for i, rental := range rentals {
		rentalsDTO[i] = rental.dto()
	}

	return rentalsDTO, nil
}

func (f *facade) CalculatePrice(filmId int, days int) (float64, error) {
	film, err := f.filmFacade.GetFilmById(filmId)
	if err != nil {
		return -1, err
	}
	return f.priceCalculator.calculatePrice(film.Type, days), nil
}

func (f *facade) CalculateLateCharge(filmId int, extraDays int) (float64, error) {	
	film, err := f.filmFacade.GetFilmById(filmId)
	if err != nil {
		return -1, err
	}
	return f.priceCalculator.calculateLateCharge(film.Type, extraDays), nil
}
