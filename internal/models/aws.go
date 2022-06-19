package model

import (
	ec2Types "github.com/aws/aws-sdk-go-v2/service/ec2/types"
	elbTypes "github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2/types"
)

// EC2 instance
type EC2Instance struct {
	InstanceID   string                     `json:"instance_id"`
	InstanceType ec2Types.InstanceType      `json:"instance_type"`
	State        ec2Types.InstanceStateName `json:"state"`
}

// EC2 instance response
type EC2InstanceResponse struct {
	Instances []EC2Instance `json:"instances"`
}

// ELB
type ELB struct {
	Name  string                         `json:"name"`
	State elbTypes.LoadBalancerStateEnum `json:"state"`
}
