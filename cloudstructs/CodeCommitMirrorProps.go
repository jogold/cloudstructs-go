package cloudstructs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// Properties for a CodeCommitMirror.
type CodeCommitMirrorProps struct {
	// The ECS cluster where to run the mirroring operation.
	Cluster awsecs.ICluster `field:"required" json:"cluster" yaml:"cluster"`
	// The source repository.
	Repository CodeCommitMirrorSourceRepository `field:"required" json:"repository" yaml:"repository"`
	// The schedule for the mirroring operation.
	// Default: - everyday at midnight.
	//
	Schedule awsevents.Schedule `field:"optional" json:"schedule" yaml:"schedule"`
	// Where to run the mirroring Fargate tasks.
	// Default: - public subnets.
	//
	SubnetSelection *awsec2.SubnetSelection `field:"optional" json:"subnetSelection" yaml:"subnetSelection"`
}

