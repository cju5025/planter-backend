package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	_ "github.com/joho/godotenv/autoload"
	"github.com/qinains/fastergoding"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())
	fastergoding.Run()

	apiAccessToken := os.Getenv("API_ACCESS_TOKEN")

	app.Get("/plants", func(c *fiber.Ctx) error {
		resp, err := http.Get("https://trefle.io/api/v1/plants?token=" + apiAccessToken)
		if err != nil {
			log.Fatalln(err)
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}

		return c.Send(body)
	})

	app.Listen(":3000")
}
