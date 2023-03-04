package main

import (
	"log"
	"os/exec"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName: "neofetcher",
	})

	app.Get("/neofetch", func(c *fiber.Ctx) error {
		cmd := exec.Command("neofetch")
		output, err := cmd.Output()
		if err != nil {
			return c.SendString(err.Error())
		}
		stringOutput := string(output)
		return c.SendString(stringOutput)
	})

	log.Fatal(app.Listen(":6969"))
}
