package film_http

import (
	"app/modules/film"
	"app/modules/film/film_contracts"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

func handler(app *fiber.App, api film.Api) {
	app.Get("/films", func(c *fiber.Ctx) error {
		films, err := api.GetAllFilms()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		return c.JSON(films)
	})

	app.Get("/films/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		idInt, err := strconv.Atoi(id)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}
		film, err := api.GetFilmById(idInt)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		return c.JSON(film)
	})

	app.Post("/films", func(c *fiber.Ctx) error {
		var film film_contracts.FilmDTO
		if err := c.BodyParser(&film); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}
		err := api.CreateFilm(film)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		return c.JSON(film)
	})
}

func Module() fx.Option {
	return fx.Invoke(
		handler,
	)
}
