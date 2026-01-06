package cloudstructs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscertificatemanager"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfront"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53"
)

// Properties for a StaticWebsite.
type StaticWebsiteProps struct {
	// The ACM certificate to use for the CloudFront distribution.
	//
	// Must be in us-east-1.
	Certificate awscertificatemanager.ICertificate `field:"required" json:"certificate" yaml:"certificate"`
	// The domain name for this static website.
	//
	// Example:
	//   www.my-static-website.com
	//
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// The hosted zone where records should be added.
	HostedZone awsroute53.IHostedZone `field:"required" json:"hostedZone" yaml:"hostedZone"`
	// A backend configuration that will be saved as `config.json` in the S3 bucket of the static website.
	//
	// The frontend can query this config by doing `fetch('/config.json')`.
	//
	// Example:
	//   { userPoolId: '1234', apiEndoint: 'https://www.my-api.com/api' }
	//
	BackendConfiguration interface{} `field:"optional" json:"backendConfiguration" yaml:"backendConfiguration"`
	// Cache policy for the default behavior.
	// Default: CachePolicy.CACHING_OPTIMIZED
	//
	CachePolicy awscloudfront.ICachePolicy `field:"optional" json:"cachePolicy" yaml:"cachePolicy"`
	// The Lambda@Edge functions to invoke before serving the contents.
	// Default: - no edge Lambdas.
	//
	EdgeLambdas *[]*awscloudfront.EdgeLambda `field:"optional" json:"edgeLambdas" yaml:"edgeLambdas"`
	// A list of domain names that should redirect to `domainName`.
	// Default: - the domain name of the hosted zone.
	//
	Redirects *[]*string `field:"optional" json:"redirects" yaml:"redirects"`
	// Response headers policy for the default behavior.
	// Default: - a new policy is created with best practice security headers.
	//
	ResponseHeadersPolicy awscloudfront.ResponseHeadersPolicy `field:"optional" json:"responseHeadersPolicy" yaml:"responseHeadersPolicy"`
}

