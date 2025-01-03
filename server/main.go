package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Println("welcome to go lang")

	app := fiber.New()

	app.Get("/healthcheck", func(c *fiber.Ctx) error {
		return c.SendString("OK Api Response")
	})
	log.Fatal(app.Listen(":8000"))
}
