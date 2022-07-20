package main

import (
	"autocomplete/storage"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	conn, err := storage.NewSQLiteConn("./dictionary.db")
	if err != nil {
		log.Fatal(err)
	}
	
	storage := storage.NewStorage(conn)

	app := fiber.New()

	app.Get("/:word", func(c *fiber.Ctx) error {
		word := c.Params("word")
		words, err := storage.Dict.GetCompleteWord(&word)
		if err != nil {
			c.SendStatus(fiber.StatusInternalServerError)
			return err
		}

		return c.JSON(words)
	})

	app.Listen(":3000")
}
