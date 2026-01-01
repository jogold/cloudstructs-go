package cloudstructs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsscheduler"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
)

// Properties for a SslServerTest.
type SslServerTestProps struct {
	// The hostname to test.
	Host *string `field:"required" json:"host" yaml:"host"`
	// The email registered with SSL Labs API.
	// See: https://github.com/ssllabs/ssllabs-scan/blob/master/ssllabs-api-docs-v4.md#register-for-scan-api-initiation-and-result-fetching
	//
	RegistrationEmail *string `field:"required" json:"registrationEmail" yaml:"registrationEmail"`
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
	ScheduleExpression awsscheduler.ScheduleExpression `field:"optional" json:"scheduleExpression" yaml:"scheduleExpression"`
}

