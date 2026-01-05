package cloudstructs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscertificatemanager"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53"
)

// Properties for a UrlShortener.
type UrlShortenerProps struct {
	// The ACM certificate to use for the CloudFront distribution.
	//
	// Must be in us-east-1.
	Certificate awscertificatemanager.ICertificate `field:"required" json:"certificate" yaml:"certificate"`
	// The hosted zone for the short URLs domain.
	HostedZone awsroute53.IHostedZone `field:"required" json:"hostedZone" yaml:"hostedZone"`
	// Authorizer for API gateway.
	// Default: - do not use an authorizer for the API.
	//
	ApiGatewayAuthorizer awsapigateway.IAuthorizer `field:"optional" json:"apiGatewayAuthorizer" yaml:"apiGatewayAuthorizer"`
	// An interface VPC endpoint for API gateway.
	//
	// Specifying this property will
	// make the API private.
	// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-private-apis.html
	//
	// Default: - API is public.
	//
	ApiGatewayEndpoint awsec2.IInterfaceVpcEndpoint `field:"optional" json:"apiGatewayEndpoint" yaml:"apiGatewayEndpoint"`
	// A name for the bucket saving the redirects.
	// Default: - derived from short link domain name.
	//
	BucketName *string `field:"optional" json:"bucketName" yaml:"bucketName"`
	// Allowed origins for CORS.
	// Default: - CORS is not enabled.
	//
	CorsAllowOrigins *[]*string `field:"optional" json:"corsAllowOrigins" yaml:"corsAllowOrigins"`
	// Expiration for short urls.
	// Default: cdk.Duration.days(365)
	//
	Expiration awscdk.Duration `field:"optional" json:"expiration" yaml:"expiration"`
	// Whether to use IAM authorization.
	// Default: - do not use IAM authorization.
	//
	IamAuthorization *bool `field:"optional" json:"iamAuthorization" yaml:"iamAuthorization"`
	// The record name to use in the hosted zone.
	// Default: - zone root.
	//
	RecordName *string `field:"optional" json:"recordName" yaml:"recordName"`
}

