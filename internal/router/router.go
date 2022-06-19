package router

import (
	handler "github.com/chammanganti/health-api/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetUpRoutes(f *fiber.App) {
	awsHandler := handler.NewAWSHandler()

	// Health
	f.Get("/health-check", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).SendString("Not dead!")
	})

	// AWS
	aws := f.Group("/api/aws")
	aws.Get("/ec2", awsHandler.GetEC2Status)
	aws.Get("/elb/:name", awsHandler.GetELBStatus)
}
