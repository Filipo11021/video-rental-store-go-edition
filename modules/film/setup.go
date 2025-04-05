package film

import (
	"app/modules/film/internal/film_internal"
	"fmt"

	"go.uber.org/fx"
	"gorm.io/gorm"
)

func Module() fx.Option {
	return fx.Module("film",
		fx.Provide(
			fx.Private,
			film_internal.NewGormFilmRepo,
		),
		fx.Provide(
			film_internal.NewApi,
		),
		fx.Invoke(func(db *gorm.DB) {
			err := db.AutoMigrate(&film_internal.Film{})

			if err != nil {
				panic(fmt.Sprintf("failed to auto-migrate database: %v", err))
			}
		}),
	)
}
