package main

import (
	"log"

	"github.com/alesanfra/gilgamesh/broker"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	log.Println("ciao")
	broker := broker.New()
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to Gilgamesh")
	})

	api := app.Group("/api/v1", logger.New())

	api.Post("/topic/:topic", func(c *fiber.Ctx) error {
		topic := c.Params("topic")
		if broker.TopicExists(topic) {
			return c.SendStatus(fiber.ErrConflict.Code)
		}

		broker.CreateTopic(topic)

		return c.JSON(&fiber.Map{
			"success": true,
			"topic":   topic,
		})
	})

	api.Post("/topic/:topic/message", func(c *fiber.Ctx) error {
		topic := c.Params("topic")
		if !broker.TopicExists(topic) {
			return c.SendStatus(fiber.ErrNotFound.Code)
		}

		broker.SendMessage(topic, "aaaaaaa")

		return c.JSON(&fiber.Map{
			"success": true,
			"topic":   topic,
		})
	})

	log.Fatal(app.Listen(":3000"))
}
