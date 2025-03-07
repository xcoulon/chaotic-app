// ⚡️ Fiber is an Express inspired web framework written in Go with ☕️
// 🤖 Github Repository: https://github.com/gofiber/fiber
// 📌 API Documentation: https://docs.gofiber.io

package main

import (
	"math/rand"
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Fiber instance
	app := fiber.New()

	// Routes
	app.Get("/hello", hello)
	app.Get("/_health/ready", ready)
	app.Get("/_health/alive", alive)

	log.Info("starting chaotic-app")
	// Start server
	log.Fatal(app.Listen(":3000"))
}

// `/hello` handler
func hello(c *fiber.Ctx) error {
	log.Infof("'hello' handler called by user-agent '%s'", getUserAgent(c))
	return c.SendString("Hello, World 👋!")
}

// `/_health/ready` handler
func ready(c *fiber.Ctx) error {
	log.Infof("'ready' handler called by user-agent '%s'", getUserAgent(c))
	return c.SendStatus(http.StatusOK)
}

// `/_health/alive` handler
func alive(c *fiber.Ctx) error {
	log.Infof("'alive' handler called by user-agent '%s'", getUserAgent(c))
	// randomly alive
	if rand.Intn(100) > 50 {
		return c.SendStatus(http.StatusInternalServerError)
	}
	return c.SendStatus(http.StatusOK)
}

func getUserAgent(c *fiber.Ctx) string {
	if ua, ok := c.GetReqHeaders()["User-Agent"]; ok && len(ua) > 0 {
		return ua[0]
	}
	return "unknown"
}
