package cloudstructs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsses"
)

// Properties for an EmailReceiver.
type EmailReceiverProps struct {
	// The SES receipt rule set where a receipt rule will be added.
	ReceiptRuleSet awsses.IReceiptRuleSet `field:"required" json:"receiptRuleSet" yaml:"receiptRuleSet"`
	// The recipients for which emails should be received.
	Recipients *[]*string `field:"required" json:"recipients" yaml:"recipients"`
	// An existing rule after which the new rule will be placed in the rule set.
	// Default: - The new rule is inserted at the beginning of the rule list.
	//
	AfterRule awsses.IReceiptRule `field:"optional" json:"afterRule" yaml:"afterRule"`
	// Whether the receiver is active.
	// Default: true.
	//
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// A Lambda function to invoke after the message is saved to S3.
	//
	// The Lambda
	// function will be invoked with a SESMessage as event.
	Function awslambda.IFunction `field:"optional" json:"function" yaml:"function"`
	// A regular expression to whitelist source email addresses.
	// Default: - no whitelisting of source email addresses.
	//
	SourceWhitelist *string `field:"optional" json:"sourceWhitelist" yaml:"sourceWhitelist"`
}

