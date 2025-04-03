package film

import (
	"fmt"

	"go.uber.org/fx"
	"gorm.io/gorm"
)

func Module() fx.Option {
	return fx.Options(
		fx.Provide(newGormFilmRepo),
		fx.Provide(newFacade),
		fx.Invoke(func(db *gorm.DB) {
			err := db.AutoMigrate(&film{})
			if err != nil {
				panic(fmt.Sprintf("failed to auto-migrate database: %v", err))
			}
		}),
	)
}
