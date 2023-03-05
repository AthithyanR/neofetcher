package main

import (
	"log"
	"os/exec"
	"time"

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

	app.Get("/return-slowly", func(c *fiber.Ctx) error {
		ticker := time.NewTicker(1 * time.Second)
		for i := 0; i < 5; i++ {
			c.Write([]byte(time.Now().String()))
			<-ticker.C
		}
		return nil
	})

	log.Fatal(app.Listen(":6969"))
}
