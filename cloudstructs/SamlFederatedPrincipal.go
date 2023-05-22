package cloudstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/jogold/cloudstructs-go/cloudstructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/jogold/cloudstructs-go/cloudstructs/internal"
)

// Principal entity that represents a SAML federated identity provider.
// Deprecated: use `SamlPrincipal` from `aws-cdk-lib/aws-iam`.
type SamlFederatedPrincipal interface {
	awsiam.FederatedPrincipal
	// When this Principal is used in an AssumeRole policy, the action to use.
	// Deprecated: use `SamlPrincipal` from `aws-cdk-lib/aws-iam`.
	AssumeRoleAction() *string
	// The conditions under which the policy is in effect.
	//
	// See [the IAM documentation](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_elements_condition.html).
	// Deprecated: use `SamlPrincipal` from `aws-cdk-lib/aws-iam`.
	Conditions() *map[string]interface{}
	// federated identity provider (i.e. 'cognito-identity.amazonaws.com' for users authenticated through Cognito).
	// Deprecated: use `SamlPrincipal` from `aws-cdk-lib/aws-iam`.
	Federated() *string
	// The principal to grant permissions to.
	// Deprecated: use `SamlPrincipal` from `aws-cdk-lib/aws-iam`.
	GrantPrincipal() awsiam.IPrincipal
	// Return the policy fragment that identifies this principal in a Policy.
	// Deprecated: use `SamlPrincipal` from `aws-cdk-lib/aws-iam`.
	PolicyFragment() awsiam.PrincipalPolicyFragment
	// The AWS account ID of this principal.
	//
	// Can be undefined when the account is not known
	// (for example, for service principals).
	// Can be a Token - in that case,
	// it's assumed to be AWS::AccountId.
	// Deprecated: use `SamlPrincipal` from `aws-cdk-lib/aws-iam`.
	PrincipalAccount() *string
	// Add to the policy of this principal.
	// Deprecated: use `SamlPrincipal` from `aws-cdk-lib/aws-iam`.
	AddToPolicy(statement awsiam.PolicyStatement) *bool
	// Add to the policy of this principal.
	// Deprecated: use `SamlPrincipal` from `aws-cdk-lib/aws-iam`.
	AddToPrincipalPolicy(_statement awsiam.PolicyStatement) *awsiam.AddToPrincipalPolicyResult
	// JSON-ify the principal.
	//
	// Used when JSON.stringify() is called
	// Deprecated: use `SamlPrincipal` from `aws-cdk-lib/aws-iam`.
	ToJSON() *map[string]*[]*string
	// Returns a string representation of an object.
	// Deprecated: use `SamlPrincipal` from `aws-cdk-lib/aws-iam`.
	ToString() *string
	// Returns a new PrincipalWithConditions using this principal as the base, with the passed conditions added.
	//
	// When there is a value for the same operator and key in both the principal and the
	// conditions parameter, the value from the conditions parameter will be used.
	//
	// Returns: a new PrincipalWithConditions object.
	// Deprecated: use `SamlPrincipal` from `aws-cdk-lib/aws-iam`.
	WithConditions(conditions *map[string]interface{}) awsiam.IPrincipal
}

// The jsii proxy struct for SamlFederatedPrincipal
type jsiiProxy_SamlFederatedPrincipal struct {
	internal.Type__awsiamFederatedPrincipal
}

func (j *jsiiProxy_SamlFederatedPrincipal) AssumeRoleAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assumeRoleAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlFederatedPrincipal) Conditions() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"conditions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlFederatedPrincipal) Federated() *string {
	var returns *string
	_jsii_.Get(
		j,
		"federated",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlFederatedPrincipal) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlFederatedPrincipal) PolicyFragment() awsiam.PrincipalPolicyFragment {
	var returns awsiam.PrincipalPolicyFragment
	_jsii_.Get(
		j,
		"policyFragment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlFederatedPrincipal) PrincipalAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"principalAccount",
		&returns,
	)
	return returns
}


// Deprecated: use `SamlPrincipal` from `aws-cdk-lib/aws-iam`.
func NewSamlFederatedPrincipal(identityProvider SamlIdentityProvider) SamlFederatedPrincipal {
	_init_.Initialize()

	if err := validateNewSamlFederatedPrincipalParameters(identityProvider); err != nil {
		panic(err)
	}
	j := jsiiProxy_SamlFederatedPrincipal{}

	_jsii_.Create(
		"cloudstructs.SamlFederatedPrincipal",
		[]interface{}{identityProvider},
		&j,
	)

	return &j
}

// Deprecated: use `SamlPrincipal` from `aws-cdk-lib/aws-iam`.
func NewSamlFederatedPrincipal_Override(s SamlFederatedPrincipal, identityProvider SamlIdentityProvider) {
	_init_.Initialize()

	_jsii_.Create(
		"cloudstructs.SamlFederatedPrincipal",
		[]interface{}{identityProvider},
		s,
	)
}

func (s *jsiiProxy_SamlFederatedPrincipal) AddToPolicy(statement awsiam.PolicyStatement) *bool {
	if err := s.validateAddToPolicyParameters(statement); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		s,
		"addToPolicy",
		[]interface{}{statement},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SamlFederatedPrincipal) AddToPrincipalPolicy(_statement awsiam.PolicyStatement) *awsiam.AddToPrincipalPolicyResult {
	if err := s.validateAddToPrincipalPolicyParameters(_statement); err != nil {
		panic(err)
	}
	var returns *awsiam.AddToPrincipalPolicyResult

	_jsii_.Invoke(
		s,
		"addToPrincipalPolicy",
		[]interface{}{_statement},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SamlFederatedPrincipal) ToJSON() *map[string]*[]*string {
	var returns *map[string]*[]*string

	_jsii_.Invoke(
		s,
		"toJSON",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SamlFederatedPrincipal) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SamlFederatedPrincipal) WithConditions(conditions *map[string]interface{}) awsiam.IPrincipal {
	if err := s.validateWithConditionsParameters(conditions); err != nil {
		panic(err)
	}
	var returns awsiam.IPrincipal

	_jsii_.Invoke(
		s,
		"withConditions",
		[]interface{}{conditions},
		&returns,
	)

	return returns
}

