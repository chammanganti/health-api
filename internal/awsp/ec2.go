package awsp

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	model "github.com/chammanganti/health-api/internal/models"
)

// EC2 interface
type EC2Interface interface {
	GetInstanceStatus() ([]model.EC2Instance, error)
}

// EC2
type EC2 struct {
	Config aws.Config
}

// New EC2
func NewEC2() EC2Interface {
	config, err := LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	return &EC2{
		Config: config,
	}
}

// Gets the EC2 instances' status
func (e *EC2) GetInstanceStatus() ([]model.EC2Instance, error) {
	client := ec2.NewFromConfig(e.Config)
	params := &ec2.DescribeInstancesInput{
		Filters: []types.Filter{
			{
				Name: aws.String("instance-state-name"),
				Values: []string{
					*aws.String("running"),
					*aws.String("pending"),
				},
			},
		},
	}

	res, err := client.DescribeInstances(context.TODO(), params)
	if err != nil {
		return []model.EC2Instance{}, err
	}

	instances := []model.EC2Instance{}
	for _, reservation := range res.Reservations {
		for _, instance := range reservation.Instances {
			instances = append(instances, model.EC2Instance{
				InstanceID:   *instance.InstanceId,
				InstanceType: instance.InstanceType,
				State:        instance.State.Name,
			})
		}
	}

	return instances, nil
}
