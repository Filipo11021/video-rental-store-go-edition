package film

import (
	"app/modules/film/internal/film_domain"
	"fmt"

	"go.uber.org/fx"
	"gorm.io/gorm"
)

func Module() fx.Option {
	return fx.Module("film",
		fx.Provide(
			fx.Private,
			film_domain.NewGormFilmRepo,
		),
		fx.Provide(
			film_domain.NewApi,
		),
		fx.Invoke(func(db *gorm.DB) {
			err := db.AutoMigrate(&film_domain.Film{})

			if err != nil {
				panic(fmt.Sprintf("failed to auto-migrate database: %v", err))
			}
		}),
	)
}
