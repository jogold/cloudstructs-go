// High-level constructs for AWS CDK
package cloudstructs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
)

// Properties for a SlackApp.
type SlackAppProps struct {
	// An AWS Secrets Manager secret containing the app configuration token.
	//
	// Must use the following JSON format:
	//
	// ```
	// {
	//    "refreshToken": "<token>"
	// }
	// ```.
	ConfigurationTokenSecret awssecretsmanager.ISecret `field:"required" json:"configurationTokenSecret" yaml:"configurationTokenSecret"`
	// The definition of the app manifest.
	// See: https://api.slack.com/reference/manifests
	//
	Manifest SlackAppManifestDefinition `field:"required" json:"manifest" yaml:"manifest"`
	// The AWS Secrets Manager secret where to store the app credentials.
	CredentialsSecret awssecretsmanager.ISecret `field:"optional" json:"credentialsSecret" yaml:"credentialsSecret"`
}

