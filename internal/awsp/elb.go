package awsp

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	elb "github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
	model "github.com/chammanganti/health-api/internal/models"
)

// ELB interface
type ELBInterface interface {
	GetELBStatus(name string) ([]model.ELB, error)
}

// ELB
type ELB struct {
	Config aws.Config
}

// New ELB
func NewELB() ELBInterface {
	config, err := LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	return &ELB{
		Config: config,
	}
}

// Get ELB status
func (e *ELB) GetELBStatus(name string) ([]model.ELB, error) {
	client := elb.NewFromConfig(e.Config)
	params := &elb.DescribeLoadBalancersInput{
		Names: []string{name},
	}

	res, err := client.DescribeLoadBalancers(context.TODO(), params)
	if err != nil {
		return []model.ELB{}, err
	}

	elbs := []model.ELB{}
	for _, elb := range res.LoadBalancers {
		elbs = append(elbs, model.ELB{
			Name:  *elb.LoadBalancerName,
			State: elb.State.Code,
		})
	}

	return elbs, nil
}
