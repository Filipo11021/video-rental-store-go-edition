package rental

import "go.uber.org/fx"

func Module() fx.Option {
	return fx.Module("rental",
		fx.Provide(newGormRentalRepo),
		fx.Provide(newPriceCalculator),
		fx.Provide(newApi),
	)
}
