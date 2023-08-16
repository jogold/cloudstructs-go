package cloudstructs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for a SlackEvents.
type SlackEventsProps struct {
	// The signing secret of the Slack app.
	SigningSecret awscdk.SecretValue `field:"required" json:"signingSecret" yaml:"signingSecret"`
	// A name for the API Gateway resource.
	// Default: SlackEventsApi.
	//
	ApiName *string `field:"optional" json:"apiName" yaml:"apiName"`
	// Whether to use a custom event bus.
	// Default: false.
	//
	CustomEventBus *bool `field:"optional" json:"customEventBus" yaml:"customEventBus"`
}

