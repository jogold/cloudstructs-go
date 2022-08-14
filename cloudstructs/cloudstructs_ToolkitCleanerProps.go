// High-level constructs for AWS CDK
package cloudstructs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// Properties for a ToolkitCleaner.
type ToolkitCleanerProps struct {
	// Only output number of assets and total size that would be deleted but without actually deleting assets.
	DryRun *bool `field:"optional" json:"dryRun" yaml:"dryRun"`
	// Retain unused assets that were created recently.
	RetainAssetsNewerThan awscdk.Duration `field:"optional" json:"retainAssetsNewerThan" yaml:"retainAssetsNewerThan"`
	// The schedule for the cleaner.
	Schedule awsevents.Schedule `field:"optional" json:"schedule" yaml:"schedule"`
	// Whether to clean on schedule.
	//
	// If you'd like to run the cleanup manually
	// via the console, set to `false`.
	ScheduleEnabled *bool `field:"optional" json:"scheduleEnabled" yaml:"scheduleEnabled"`
}

