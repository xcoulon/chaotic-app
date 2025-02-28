// ⚡️ Fiber is an Express inspired web framework written in Go with ☕️
// 🤖 Github Repository: https://github.com/gofiber/fiber
// 📌 API Documentation: https://docs.gofiber.io

package main

import (
	"log"
	"math/rand"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Fiber instance
	app := fiber.New()

	// Routes
	app.Get("/hello", hello)
	app.Get("/_health/ready", ready)
	app.Get("/_health/alive", alive)

	log.Print("starting chaotic-app")
	// Start server
	log.Fatal(app.Listen(":3000"))
}

// `/hello` handler
func hello(c *fiber.Ctx) error {
	return c.SendString("Hello, World 👋!")
}

// `/_health/ready` handler
func ready(c *fiber.Ctx) error {
	return c.SendStatus(http.StatusOK)
}

// `/_health/alive` handler
func alive(c *fiber.Ctx) error {
	// randomly alive
	if rand.Intn(100) > 50 {
		return c.SendStatus(http.StatusInternalServerError)
	}
	return c.SendStatus(http.StatusOK)
}
