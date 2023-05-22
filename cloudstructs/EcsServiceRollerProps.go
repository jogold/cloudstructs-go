package cloudstructs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
)

// Properties for a EcsServiceRoller.
type EcsServiceRollerProps struct {
	// The ECS cluster where the services run.
	Cluster awsecs.ICluster `field:"required" json:"cluster" yaml:"cluster"`
	// The ECS service for which tasks should be rolled.
	Service awsecs.IService `field:"required" json:"service" yaml:"service"`
	// The rule or schedule that should trigger a roll.
	Trigger RollTrigger `field:"optional" json:"trigger" yaml:"trigger"`
}

