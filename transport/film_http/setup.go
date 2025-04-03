package film_http

import (
	"app/film"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

func handler(app *fiber.App, facade film.Facade) {
	app.Get("/films", func(c *fiber.Ctx) error {
		films, err := facade.GetAllFilms()
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
		film, err := facade.GetFilmById(idInt)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		return c.JSON(film)
	})

	app.Post("/films", func(c *fiber.Ctx) error {
		var film film.FilmDTO
		if err := c.BodyParser(&film); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}
		err := facade.CreateFilm(film)
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
