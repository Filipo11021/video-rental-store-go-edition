package rental

import (
	"fmt"

	"go.uber.org/fx"
	"gorm.io/gorm"
)

func Module() fx.Option {
	return fx.Module("rental",
		fx.Provide(newGormRentalRepo),
		fx.Provide(newPriceCalculator),
		fx.Provide(newApi),
		fx.Invoke(func(db *gorm.DB) {
			err := db.AutoMigrate(&rental{})
			if err != nil {
				panic(fmt.Sprintf("failed to auto-migrate database: %v", err))
			}
		}),
	)
}
