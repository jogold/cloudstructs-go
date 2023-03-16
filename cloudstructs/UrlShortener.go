// High-level constructs for AWS CDK
package cloudstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/jogold/cloudstructs-go/cloudstructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/jogold/cloudstructs-go/cloudstructs/internal"
)

// URL shortener.
type UrlShortener interface {
	constructs.Construct
	// The underlying API Gateway REST API.
	Api() awsapigateway.RestApi
	// The endpoint of the URL shortener API.
	ApiEndpoint() *string
	// The tree node.
	Node() constructs.Node
	// Grant access to invoke the URL shortener if protected by IAM authorization.
	GrantInvoke(grantee awsiam.IGrantable) awsiam.Grant
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for UrlShortener
type jsiiProxy_UrlShortener struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_UrlShortener) Api() awsapigateway.RestApi {
	var returns awsapigateway.RestApi
	_jsii_.Get(
		j,
		"api",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UrlShortener) ApiEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UrlShortener) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


func NewUrlShortener(scope constructs.Construct, id *string, props *UrlShortenerProps) UrlShortener {
	_init_.Initialize()

	if err := validateNewUrlShortenerParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_UrlShortener{}

	_jsii_.Create(
		"cloudstructs.UrlShortener",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewUrlShortener_Override(u UrlShortener, scope constructs.Construct, id *string, props *UrlShortenerProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cloudstructs.UrlShortener",
		[]interface{}{scope, id, props},
		u,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func UrlShortener_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateUrlShortener_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cloudstructs.UrlShortener",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UrlShortener) GrantInvoke(grantee awsiam.IGrantable) awsiam.Grant {
	if err := u.validateGrantInvokeParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		u,
		"grantInvoke",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UrlShortener) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		u,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

