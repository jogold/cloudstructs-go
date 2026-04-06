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
	//   "appId": "...",
	//   "clientId": "...",
	//   "clientSecret": "...",
	//   "verificationToken": "...",
	//   "signingSecret": "..."
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
	// Applies one or more mixins to this construct.
	//
	// Mixins are applied in order. The list of constructs is captured at the
	// start of the call, so constructs added by a mixin will not be visited.
	// Use multiple `with()` calls if subsequent mixins should apply to added
	// constructs.
	//
	// Returns: This construct for chaining.
	With(mixins ...constructs.IMixin) constructs.IConstruct
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
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
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

func (s *jsiiProxy_SlackApp) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		s,
		"with",
		args,
		&returns,
	)

	return returns
}

