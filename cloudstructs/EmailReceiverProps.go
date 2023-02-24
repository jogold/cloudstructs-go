// High-level constructs for AWS CDK
package cloudstructs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsses"
)

// Properties for an EmailReceiver.
type EmailReceiverProps struct {
	// A Lambda function to invoke after the message is saved to S3.
	//
	// The Lambda
	// function will be invoked with a SESMessage as event.
	Function awslambda.IFunction `field:"required" json:"function" yaml:"function"`
	// The SES receipt rule set where a receipt rule will be added.
	ReceiptRuleSet awsses.IReceiptRuleSet `field:"required" json:"receiptRuleSet" yaml:"receiptRuleSet"`
	// The recipients for which emails should be received.
	Recipients *[]*string `field:"required" json:"recipients" yaml:"recipients"`
	// An existing rule after which the new rule will be placed in the rule set.
	AfterRule awsses.IReceiptRule `field:"optional" json:"afterRule" yaml:"afterRule"`
	// A regular expression to whitelist source email addresses.
	SourceWhitelist *string `field:"optional" json:"sourceWhitelist" yaml:"sourceWhitelist"`
}

