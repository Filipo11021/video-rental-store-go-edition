package rental

import (
	"app/modules/film/film_contracts"
)

const (
	NewReleaseBasePrice = 40.0
	RegularBasePrice    = 30.0
	OldBasePrice        = 30.0
)

type PriceCalculator interface {
	calculatePrice(filmType film_contracts.FilmTypeDto, days int) float64
	calculateLateCharge(filmType film_contracts.FilmTypeDto, extraDays int) float64
}

type priceCalculator struct{}

func newPriceCalculator() PriceCalculator {
	return &priceCalculator{}
}

func (pc *priceCalculator) calculatePrice(filmType film_contracts.FilmTypeDto, days int) float64 {
	switch filmType {
	case film_contracts.NewRelease:
		return NewReleaseBasePrice * float64(days)
	case film_contracts.Regular:
		if days <= 3 {
			return RegularBasePrice
		}
		return RegularBasePrice + RegularBasePrice*float64(days-3)
	case film_contracts.Old:
		if days <= 5 {
			return OldBasePrice
		}
		return OldBasePrice + OldBasePrice*float64(days-5)
	default:
		return 0
	}
}

func (pc *priceCalculator) calculateLateCharge(filmType film_contracts.FilmTypeDto, extraDays int) float64 {
	return pc.calculatePrice(filmType, extraDays)
}
