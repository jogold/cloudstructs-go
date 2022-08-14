// High-level constructs for AWS CDK
package cloudstructs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for a SlackTextract.
type SlackTextractProps struct {
	// The application id of the Slack app.
	AppId *string `field:"required" json:"appId" yaml:"appId"`
	// The **bot** token of the Slack app.
	//
	// The following scopes are required: `chat:write` and `files:read`.
	BotToken awscdk.SecretValue `field:"required" json:"botToken" yaml:"botToken"`
	// The signing secret of the Slack app.
	SigningSecret awscdk.SecretValue `field:"required" json:"signingSecret" yaml:"signingSecret"`
}

