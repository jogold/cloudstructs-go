package cloudstructs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsscheduler"
)

// Properties for a ToolkitCleaner.
type ToolkitCleanerProps struct {
	// The timeout for the clean function.
	// Default: Duration.minutes(30)
	//
	CleanAssetsTimeout awscdk.Duration `field:"optional" json:"cleanAssetsTimeout" yaml:"cleanAssetsTimeout"`
	// Only output number of assets and total size that would be deleted but without actually deleting assets.
	DryRun *bool `field:"optional" json:"dryRun" yaml:"dryRun"`
	// Retain unused assets that were created recently.
	// Default: - all unused assets are removed.
	//
	RetainAssetsNewerThan awscdk.Duration `field:"optional" json:"retainAssetsNewerThan" yaml:"retainAssetsNewerThan"`
	// Whether to clean on schedule.
	//
	// If you'd like to run the cleanup manually
	// via the console, set to `false`.
	// Default: true.
	//
	ScheduleEnabled *bool `field:"optional" json:"scheduleEnabled" yaml:"scheduleEnabled"`
	// The schedule for the cleaner.
	// Default: - every day.
	//
	ScheduleExpression awsscheduler.ScheduleExpression `field:"optional" json:"scheduleExpression" yaml:"scheduleExpression"`
}

