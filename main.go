package main

import (
	"app/film"
	"app/rental"
	"app/transport/film_http"
	"app/user"
	"context"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.uber.org/fx"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type App struct {
	fx.In

	RentalFacade rental.Facade
	FilmFacade   film.Facade
	UserFacade   user.Facade
}

func provideDatabase() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("./test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return db
}

func provideFiberApp() *fiber.App {
	app := fiber.New(fiber.Config{
	})
	return app
}

func registerHooks(lifecycle fx.Lifecycle, app *fiber.App) {
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				if err := app.Listen(":3000"); err != nil {
					log.Printf("Error starting Fiber server: %v\n", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return app.Shutdown()
		},
	})
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	workOSAPIKey := os.Getenv("WORKOS_API_KEY")
	workOSClientID := os.Getenv("WORKOS_CLIENT_ID")

	if workOSAPIKey == "" || workOSClientID == "" {
		panic("WORKOS_API_KEY and WORKOS_CLIENT_ID must be set")
	}

	fx.New(
		fx.Provide(provideDatabase),
		fx.Provide(provideFiberApp),

		film.Module(),
		rental.Module(),
		user.Module(user.CreateConfigDTO{
			WorkOSAPIKey:   workOSAPIKey,
			WorkOSClientID: workOSClientID,
		}),

		fx.Invoke(registerHooks),
		film_http.Module(),
	).Run()
}
