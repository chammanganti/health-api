package handler

import (
	"github.com/chammanganti/health-api/internal/awsp"
	model "github.com/chammanganti/health-api/internal/models"
	"github.com/gofiber/fiber/v2"
)

// AWS handler interface
type AWSHandlerInterface interface {
	GetEC2Status(c *fiber.Ctx) error
	GetELBStatus(c *fiber.Ctx) error
}

// AWS handler
type AWSHandler struct {
	EC2 awsp.EC2Interface
	ELB awsp.ELBInterface
}

// New AWS handler
func NewAWSHandler() AWSHandlerInterface {
	ec2 := awsp.NewEC2()
	elb := awsp.NewELB()

	return &AWSHandler{
		EC2: ec2,
		ELB: elb,
	}
}

// Gets the status of all the EC2 instances
func (a *AWSHandler) GetEC2Status(c *fiber.Ctx) error {
	instances, err := a.EC2.GetInstanceStatus()
	if err != nil {
		return RespondWithError(c, fiber.StatusBadRequest, err)
	}

	return c.Status(fiber.StatusOK).JSON(model.EC2InstanceResponse{
		Instances: instances,
	})
}

// Gets the status of an ELB
func (a *AWSHandler) GetELBStatus(c *fiber.Ctx) error {
	name := c.Params("name")

	elb, err := a.ELB.GetELBStatus(name)
	if err != nil {
		return RespondWithError(c, fiber.StatusBadRequest, err)
	}

	return c.Status(fiber.StatusOK).JSON(elb)
}
