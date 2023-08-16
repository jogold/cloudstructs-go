package cloudstructs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
)

// Properties for a SslServerTest.
type SslServerTestProps struct {
	// The hostname to test.
	Host *string `field:"required" json:"host" yaml:"host"`
	// The topic to which the results must be sent when the grade is below the minimum grade.
	// Default: - a new topic is created.
	//
	AlarmTopic awssns.ITopic `field:"optional" json:"alarmTopic" yaml:"alarmTopic"`
	// Minimum grade for the test. The grade is calculated using the worst grade of all endpoints.
	//
	// Used to send the results to an alarm SNS topic.
	// Default: SslServerTestGrade.A_PLUS
	//
	MinimumGrade SslServerTestGrade `field:"optional" json:"minimumGrade" yaml:"minimumGrade"`
	// The schedule for the test.
	// Default: - every day.
	//
	Schedule awsevents.Schedule `field:"optional" json:"schedule" yaml:"schedule"`
}

