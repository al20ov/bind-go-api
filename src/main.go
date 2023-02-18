package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

type Payload struct {
	Name string `json:"name"`
	Ttl  uint32 `json:"ttl"`
	Type string `json:"type"`
	Data string `json:"data"`
}

// Generate a tsig key like so: `tsig-keygen <name of the key> > <name of the key>.key`

func main() {
	config := LoadConfiguration()

	tsig := config.TsigKey

	app := fiber.New()

	app.Get("/zones/:zone", func(c *fiber.Ctx) error {
		c.Response().Header.Add("Content-Type", "application/json")
		zone := c.Params("zone")
		zoneContents, err := GetZone(zone, config.DNSServer, tsig)
		if err != nil {
			return c.SendStatus(404)
		}
		return c.SendString(string(zoneContents))
	})

	app.Get("/zones", func(c *fiber.Ctx) error {
		c.Response().Header.Add("Content-Type", "application/json")
		zoneType := c.Query("type")
		zones, err := GetZones(config.StatisticsURL, zoneType)
		if err != nil {
			return c.SendStatus(401)
		}
		return c.SendString(string(zones))
	})

	app.Post("/zones/:zone", func(c *fiber.Ctx) error {
		c.Response().Header.Add("Content-Type", "application/json")
		payload := Payload{}

		if err := c.BodyParser(&payload); err != nil {
			return err
		}

		zone := c.Params("zone")
		PostRecord(zone, config.DNSServer, tsig, payload.Name, payload.Ttl, payload.Type, payload.Data)
		return c.SendString("SUCCESS")
	})

	app.Delete("/zones/:zone", func(c *fiber.Ctx) error {
		c.Response().Header.Add("Content-Type", "application/json")
		payload := Payload{}

		if err := c.BodyParser(&payload); err != nil {
			return err
		}

		zone := c.Params("zone")
		DeleteRecord(zone, config.DNSServer, tsig, payload.Name, payload.Ttl, payload.Type, payload.Data)

		return c.SendString("SUCCESS")
	})

	log.Fatal(app.Listen(config.ListenAddress))
}
