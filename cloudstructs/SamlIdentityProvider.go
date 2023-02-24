// High-level constructs for AWS CDK
package cloudstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/jogold/cloudstructs-go/cloudstructs/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/jogold/cloudstructs-go/cloudstructs/internal"
)

// Create a SAML identity provider.
// Deprecated: use `SamlProvider` from `aws-cdk-lib/aws-iam`.
type SamlIdentityProvider interface {
	constructs.Construct
	// The tree node.
	// Deprecated: use `SamlProvider` from `aws-cdk-lib/aws-iam`.
	Node() constructs.Node
	// The ARN of the SAML identity provider.
	// Deprecated: use `SamlProvider` from `aws-cdk-lib/aws-iam`.
	SamlIdentityProviderArn() *string
	// Returns a string representation of this construct.
	// Deprecated: use `SamlProvider` from `aws-cdk-lib/aws-iam`.
	ToString() *string
}

// The jsii proxy struct for SamlIdentityProvider
type jsiiProxy_SamlIdentityProvider struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_SamlIdentityProvider) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlIdentityProvider) SamlIdentityProviderArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"samlIdentityProviderArn",
		&returns,
	)
	return returns
}


// Deprecated: use `SamlProvider` from `aws-cdk-lib/aws-iam`.
func NewSamlIdentityProvider(scope constructs.Construct, id *string, props *SamlIdentityProviderProps) SamlIdentityProvider {
	_init_.Initialize()

	if err := validateNewSamlIdentityProviderParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_SamlIdentityProvider{}

	_jsii_.Create(
		"cloudstructs.SamlIdentityProvider",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Deprecated: use `SamlProvider` from `aws-cdk-lib/aws-iam`.
func NewSamlIdentityProvider_Override(s SamlIdentityProvider, scope constructs.Construct, id *string, props *SamlIdentityProviderProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cloudstructs.SamlIdentityProvider",
		[]interface{}{scope, id, props},
		s,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func SamlIdentityProvider_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSamlIdentityProvider_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cloudstructs.SamlIdentityProvider",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SamlIdentityProvider) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

