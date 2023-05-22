package cloudstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/jogold/cloudstructs-go/cloudstructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/jogold/cloudstructs-go/cloudstructs/internal"
)

// A Slack application deployed with a manifest.
// See: https://api.slack.com/reference/manifests
//
type SlackApp interface {
	constructs.Construct
	// The ID of the application.
	AppId() *string
	// A dynamic reference to the client ID of the app.
	ClientId() *string
	// A dynamic reference to the client secret of the app.
	ClientSecret() *string
	// An AWS Secrets Manager secret containing the credentials of the application.
	//
	// ```
	// {
	//    "appId": "...",
	//    "clientId": "...",
	//    "clientSecret": "...",
	//    "verificationToken": "...",
	//    "signingSecret": "..."
	// }
	// ```.
	Credentials() awssecretsmanager.ISecret
	// The tree node.
	Node() constructs.Node
	// A dynamic reference to the signing secret of the app.
	SigningSecret() *string
	// A dynamic reference to the verification token of the app.
	VerificationToken() *string
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for SlackApp
type jsiiProxy_SlackApp struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_SlackApp) AppId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SlackApp) ClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SlackApp) ClientSecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SlackApp) Credentials() awssecretsmanager.ISecret {
	var returns awssecretsmanager.ISecret
	_jsii_.Get(
		j,
		"credentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SlackApp) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SlackApp) SigningSecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signingSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SlackApp) VerificationToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"verificationToken",
		&returns,
	)
	return returns
}


func NewSlackApp(scope constructs.Construct, id *string, props *SlackAppProps) SlackApp {
	_init_.Initialize()

	if err := validateNewSlackAppParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_SlackApp{}

	_jsii_.Create(
		"cloudstructs.SlackApp",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewSlackApp_Override(s SlackApp, scope constructs.Construct, id *string, props *SlackAppProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cloudstructs.SlackApp",
		[]interface{}{scope, id, props},
		s,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func SlackApp_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSlackApp_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cloudstructs.SlackApp",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SlackApp) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

