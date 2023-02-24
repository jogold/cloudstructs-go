// High-level constructs for AWS CDK
package cloudstructs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53"
)

// Properties for a UrlShortener.
type UrlShortenerProps struct {
	// The hosted zone for the short URLs domain.
	HostedZone awsroute53.IHostedZone `field:"required" json:"hostedZone" yaml:"hostedZone"`
	// Authorizer for API gateway.
	ApiGatewayAuthorizer awsapigateway.IAuthorizer `field:"optional" json:"apiGatewayAuthorizer" yaml:"apiGatewayAuthorizer"`
	// An interface VPC endpoint for API gateway.
	//
	// Specifying this property will
	// make the API private.
	// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-private-apis.html
	//
	ApiGatewayEndpoint awsec2.IInterfaceVpcEndpoint `field:"optional" json:"apiGatewayEndpoint" yaml:"apiGatewayEndpoint"`
	// A name for the bucket saving the redirects.
	BucketName *string `field:"optional" json:"bucketName" yaml:"bucketName"`
	// Allowed origins for CORS.
	CorsAllowOrigins *[]*string `field:"optional" json:"corsAllowOrigins" yaml:"corsAllowOrigins"`
	// Expiration for short urls.
	Expiration awscdk.Duration `field:"optional" json:"expiration" yaml:"expiration"`
	// The record name to use in the hosted zone.
	RecordName *string `field:"optional" json:"recordName" yaml:"recordName"`
}

