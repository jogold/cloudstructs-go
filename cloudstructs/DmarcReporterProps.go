// High-level constructs for AWS CDK
package cloudstructs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsses"
)

// Properties for a DmarcReporter.
type DmarcReporterProps struct {
	// The DMARC policy to apply to messages that fail DMARC compliance.
	//
	// This can be one of the following values:
	// - none: Do not apply any special handling to messages that fail DMARC compliance.
	// - quarantine: Quarantine messages that fail DMARC compliance.
	// - reject: Reject messages that fail DMARC compliance.
	DmarcPolicy DmarcPolicy `field:"required" json:"dmarcPolicy" yaml:"dmarcPolicy"`
	// A Lambda function to invoke after the message is saved to S3.
	//
	// The Lambda
	// function will be invoked with a SESMessage as event.
	Function awslambda.IFunction `field:"required" json:"function" yaml:"function"`
	// The Route 53 hosted zone to create the DMARC record in.
	HostedZone awsroute53.IHostedZone `field:"required" json:"hostedZone" yaml:"hostedZone"`
	// The SES receipt rule set where a receipt rule will be added.
	ReceiptRuleSet awsses.IReceiptRuleSet `field:"required" json:"receiptRuleSet" yaml:"receiptRuleSet"`
	// Additional email addresses to send DMARC reports to.
	AdditionalEmailAddresses *[]*string `field:"optional" json:"additionalEmailAddresses" yaml:"additionalEmailAddresses"`
	// An existing rule after which the new rule will be placed in the rule set.
	AfterRule awsses.IReceiptRule `field:"optional" json:"afterRule" yaml:"afterRule"`
	// The alignment mode to use for DKIM signatures.
	//
	// This can be one of the following values:
	// - relaxed: Use relaxed alignment mode.
	// - strict: Use strict alignment mode.
	DmarcDkimAlignment DmarcAlignment `field:"optional" json:"dmarcDkimAlignment" yaml:"dmarcDkimAlignment"`
	// The percentage of messages that should be checked for DMARC compliance.
	//
	// This is a value between 0 and 100.
	DmarcPercentage *float64 `field:"optional" json:"dmarcPercentage" yaml:"dmarcPercentage"`
	// The alignment mode to use for SPF signatures.
	//
	// This can be one of the following values:
	// - relaxed: Use relaxed alignment mode.
	// - strict: Use strict alignment mode.
	DmarcSpfAlignment DmarcAlignment `field:"optional" json:"dmarcSpfAlignment" yaml:"dmarcSpfAlignment"`
	// The DMARC policy to apply to messages that fail DMARC compliance for subdomains.
	//
	// This can be one of the following values:
	// - none: Do not apply any special handling to messages that fail DMARC compliance.
	// - quarantine: Quarantine messages that fail DMARC compliance.
	// - reject: Reject messages that fail DMARC compliance.
	DmarcSubdomainPolicy DmarcPolicy `field:"optional" json:"dmarcSubdomainPolicy" yaml:"dmarcSubdomainPolicy"`
	// The email address to send DMARC reports to.
	//
	// This email address must be verified in SES.
	EmailAddress *string `field:"optional" json:"emailAddress" yaml:"emailAddress"`
}

