// High-level constructs for AWS CDK
package cloudstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/jogold/cloudstructs-go/cloudstructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfront"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsses"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/jogold/cloudstructs-go/cloudstructs/internal"
)

// Mirror a repository to AWS CodeCommit on schedule.
type CodeCommitMirror interface {
	constructs.Construct
	// The tree node.
	Node() constructs.Node
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for CodeCommitMirror
type jsiiProxy_CodeCommitMirror struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_CodeCommitMirror) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


func NewCodeCommitMirror(scope constructs.Construct, id *string, props *CodeCommitMirrorProps) CodeCommitMirror {
	_init_.Initialize()

	j := jsiiProxy_CodeCommitMirror{}

	_jsii_.Create(
		"cloudstructs.CodeCommitMirror",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewCodeCommitMirror_Override(c CodeCommitMirror, scope constructs.Construct, id *string, props *CodeCommitMirrorProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cloudstructs.CodeCommitMirror",
		[]interface{}{scope, id, props},
		c,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func CodeCommitMirror_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cloudstructs.CodeCommitMirror",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodeCommitMirror) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties for a CodeCommitMirror.
type CodeCommitMirrorProps struct {
	// The ECS cluster where to run the mirroring operation.
	Cluster awsecs.ICluster `field:"required" json:"cluster" yaml:"cluster"`
	// The source repository.
	Repository CodeCommitMirrorSourceRepository `field:"required" json:"repository" yaml:"repository"`
	// The schedule for the mirroring operation.
	Schedule awsevents.Schedule `field:"optional" json:"schedule" yaml:"schedule"`
	// Where to run the mirroring Fargate tasks.
	SubnetSelection *awsec2.SubnetSelection `field:"optional" json:"subnetSelection" yaml:"subnetSelection"`
}

// A source repository for AWS CodeCommit mirroring.
type CodeCommitMirrorSourceRepository interface {
	// The name of the repository.
	Name() *string
	// The HTTPS clone URL in plain text, used for a public repository.
	PlainTextUrl() *string
	// The HTTPS clone URL if the repository is private.
	//
	// The secret should contain the username and/or token.
	//
	// Example:
	//   `https://TOKEN@github.com/owner/name`
	//   `https://USERNAME:TOKEN@bitbucket.org/owner/name.git`
	//
	SecretUrl() awsecs.Secret
}

// The jsii proxy struct for CodeCommitMirrorSourceRepository
type jsiiProxy_CodeCommitMirrorSourceRepository struct {
	_ byte // padding
}

func (j *jsiiProxy_CodeCommitMirrorSourceRepository) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodeCommitMirrorSourceRepository) PlainTextUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"plainTextUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodeCommitMirrorSourceRepository) SecretUrl() awsecs.Secret {
	var returns awsecs.Secret
	_jsii_.Get(
		j,
		"secretUrl",
		&returns,
	)
	return returns
}


func NewCodeCommitMirrorSourceRepository_Override(c CodeCommitMirrorSourceRepository) {
	_init_.Initialize()

	_jsii_.Create(
		"cloudstructs.CodeCommitMirrorSourceRepository",
		nil, // no parameters
		c,
	)
}

// Public GitHub repository.
func CodeCommitMirrorSourceRepository_GitHub(owner *string, name *string) CodeCommitMirrorSourceRepository {
	_init_.Initialize()

	var returns CodeCommitMirrorSourceRepository

	_jsii_.StaticInvoke(
		"cloudstructs.CodeCommitMirrorSourceRepository",
		"gitHub",
		[]interface{}{owner, name},
		&returns,
	)

	return returns
}

// Private repository with HTTPS clone URL stored in a AWS Secrets Manager secret or a AWS Systems Manager secure string parameter.
func CodeCommitMirrorSourceRepository_Private(name *string, url awsecs.Secret) CodeCommitMirrorSourceRepository {
	_init_.Initialize()

	var returns CodeCommitMirrorSourceRepository

	_jsii_.StaticInvoke(
		"cloudstructs.CodeCommitMirrorSourceRepository",
		"private",
		[]interface{}{name, url},
		&returns,
	)

	return returns
}

// Roll your ECS service tasks on schedule or with a rule.
type EcsServiceRoller interface {
	constructs.Construct
	// The tree node.
	Node() constructs.Node
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for EcsServiceRoller
type jsiiProxy_EcsServiceRoller struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_EcsServiceRoller) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


func NewEcsServiceRoller(scope constructs.Construct, id *string, props *EcsServiceRollerProps) EcsServiceRoller {
	_init_.Initialize()

	j := jsiiProxy_EcsServiceRoller{}

	_jsii_.Create(
		"cloudstructs.EcsServiceRoller",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewEcsServiceRoller_Override(e EcsServiceRoller, scope constructs.Construct, id *string, props *EcsServiceRollerProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cloudstructs.EcsServiceRoller",
		[]interface{}{scope, id, props},
		e,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func EcsServiceRoller_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cloudstructs.EcsServiceRoller",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsServiceRoller) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties for a EcsServiceRoller.
type EcsServiceRollerProps struct {
	// The ECS cluster where the services run.
	Cluster awsecs.ICluster `field:"required" json:"cluster" yaml:"cluster"`
	// The ECS service for which tasks should be rolled.
	Service awsecs.IService `field:"required" json:"service" yaml:"service"`
	// The rule or schedule that should trigger a roll.
	Trigger RollTrigger `field:"optional" json:"trigger" yaml:"trigger"`
}

// Receive emails through SES, save them to S3 and invokes a Lambda function.
type EmailReceiver interface {
	constructs.Construct
	// The tree node.
	Node() constructs.Node
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for EmailReceiver
type jsiiProxy_EmailReceiver struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_EmailReceiver) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


func NewEmailReceiver(scope constructs.Construct, id *string, props *EmailReceiverProps) EmailReceiver {
	_init_.Initialize()

	j := jsiiProxy_EmailReceiver{}

	_jsii_.Create(
		"cloudstructs.EmailReceiver",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewEmailReceiver_Override(e EmailReceiver, scope constructs.Construct, id *string, props *EmailReceiverProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cloudstructs.EmailReceiver",
		[]interface{}{scope, id, props},
		e,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func EmailReceiver_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cloudstructs.EmailReceiver",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmailReceiver) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties for an EmailReceiver.
type EmailReceiverProps struct {
	// A Lambda function to invoke after the message is saved to S3.
	//
	// The Lambda
	// function will be invoked with a SESMessage as event.
	Function awslambda.IFunction `field:"required" json:"function" yaml:"function"`
	// The SES receipt rule set where a receipt rule will be added.
	ReceiptRuleSet awsses.IReceiptRuleSet `field:"required" json:"receiptRuleSet" yaml:"receiptRuleSet"`
	// The recipients for which emails should be received.
	Recipients *[]*string `field:"required" json:"recipients" yaml:"recipients"`
	// An existing rule after which the new rule will be placed in the rule set.
	AfterRule awsses.IReceiptRule `field:"optional" json:"afterRule" yaml:"afterRule"`
	// A regular expression to whitelist source email addresses.
	SourceWhitelist *string `field:"optional" json:"sourceWhitelist" yaml:"sourceWhitelist"`
}

// A State Machine.
type IStateMachine interface {
	// The ARN of the state machine.
	StateMachineArn() *string
}

// The jsii proxy for IStateMachine
type jsiiProxy_IStateMachine struct {
	_ byte // padding
}

func (j *jsiiProxy_IStateMachine) StateMachineArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateMachineArn",
		&returns,
	)
	return returns
}

// The rule or schedule that should trigger a roll.
type RollTrigger interface {
	// Roll rule.
	Rule() awsevents.Rule
	// Roll schedule.
	Schedule() awsevents.Schedule
}

// The jsii proxy struct for RollTrigger
type jsiiProxy_RollTrigger struct {
	_ byte // padding
}

func (j *jsiiProxy_RollTrigger) Rule() awsevents.Rule {
	var returns awsevents.Rule
	_jsii_.Get(
		j,
		"rule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RollTrigger) Schedule() awsevents.Schedule {
	var returns awsevents.Schedule
	_jsii_.Get(
		j,
		"schedule",
		&returns,
	)
	return returns
}


func NewRollTrigger_Override(r RollTrigger) {
	_init_.Initialize()

	_jsii_.Create(
		"cloudstructs.RollTrigger",
		nil, // no parameters
		r,
	)
}

// Rule that should trigger a roll.
func RollTrigger_FromRule(rule awsevents.Rule) RollTrigger {
	_init_.Initialize()

	var returns RollTrigger

	_jsii_.StaticInvoke(
		"cloudstructs.RollTrigger",
		"fromRule",
		[]interface{}{rule},
		&returns,
	)

	return returns
}

// Schedule that should trigger a roll.
func RollTrigger_FromSchedule(schedule awsevents.Schedule) RollTrigger {
	_init_.Initialize()

	var returns RollTrigger

	_jsii_.StaticInvoke(
		"cloudstructs.RollTrigger",
		"fromSchedule",
		[]interface{}{schedule},
		&returns,
	)

	return returns
}

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
	var returns awsiam.IPrincipal

	_jsii_.Invoke(
		s,
		"withConditions",
		[]interface{}{conditions},
		&returns,
	)

	return returns
}

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

// Properties for a SamlProvider.
// Deprecated: use `SamlProviderProps` from `aws-cdk-lib/aws-iam`.
type SamlIdentityProviderProps struct {
	// An XML document generated by an identity provider (IdP) that supports SAML 2.0.
	//
	// The document includes the issuer's name, expiration information, and keys that
	// can be used to validate the SAML authentication response (assertions) that are
	// received from the IdP. You must generate the metadata document using the identity
	// management software that is used as your organization's IdP.
	// Deprecated: use `SamlProviderProps` from `aws-cdk-lib/aws-iam`.
	MetadataDocument *string `field:"required" json:"metadataDocument" yaml:"metadataDocument"`
	// A name for the SAML identity provider.
	// Deprecated: use `SamlProviderProps` from `aws-cdk-lib/aws-iam`.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

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

// A Slack app manifest.
// See: https://api.slack.com/reference/manifests
//
type SlackAppManifest interface {
	Render(construct constructs.IConstruct) *string
}

// The jsii proxy struct for SlackAppManifest
type jsiiProxy_SlackAppManifest struct {
	_ byte // padding
}

func NewSlackAppManifest(props *SlackAppManifestProps) SlackAppManifest {
	_init_.Initialize()

	j := jsiiProxy_SlackAppManifest{}

	_jsii_.Create(
		"cloudstructs.SlackAppManifest",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewSlackAppManifest_Override(s SlackAppManifest, props *SlackAppManifestProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cloudstructs.SlackAppManifest",
		[]interface{}{props},
		s,
	)
}

func (s *jsiiProxy_SlackAppManifest) Render(construct constructs.IConstruct) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"render",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// App Home configuration.
// See: https://api.slack.com/surfaces/tabs
//
type SlackAppManifestAppHome struct {
	// Wether the Home tab is enabled.
	HomeTab *bool `field:"optional" json:"homeTab" yaml:"homeTab"`
	// Wether the Messages is enabled.
	MessagesTab *bool `field:"optional" json:"messagesTab" yaml:"messagesTab"`
	// Whether the users can send messages to your app in the Messages tab of your App Home.
	MessagesTabReadOnly *bool `field:"optional" json:"messagesTabReadOnly" yaml:"messagesTabReadOnly"`
}

// A Slack app manifest definition.
type SlackAppManifestDefinition interface {
	// Renders the JSON app manifest encoded as a string.
	Render(construct constructs.IConstruct) *string
}

// The jsii proxy struct for SlackAppManifestDefinition
type jsiiProxy_SlackAppManifestDefinition struct {
	_ byte // padding
}

func NewSlackAppManifestDefinition_Override(s SlackAppManifestDefinition) {
	_init_.Initialize()

	_jsii_.Create(
		"cloudstructs.SlackAppManifestDefinition",
		nil, // no parameters
		s,
	)
}

// Creates a Slack app manifest from a file containg a JSON app manifest.
func SlackAppManifestDefinition_FromFile(file *string) SlackAppManifestDefinition {
	_init_.Initialize()

	var returns SlackAppManifestDefinition

	_jsii_.StaticInvoke(
		"cloudstructs.SlackAppManifestDefinition",
		"fromFile",
		[]interface{}{file},
		&returns,
	)

	return returns
}

// Creates a Slack app manifest by specifying properties.
func SlackAppManifestDefinition_FromManifest(props *SlackAppManifestProps) SlackAppManifestDefinition {
	_init_.Initialize()

	var returns SlackAppManifestDefinition

	_jsii_.StaticInvoke(
		"cloudstructs.SlackAppManifestDefinition",
		"fromManifest",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Create a Slack app manifest from a JSON app manifest encoded as a string.
func SlackAppManifestDefinition_FromString(manifest *string) SlackAppManifestDefinition {
	_init_.Initialize()

	var returns SlackAppManifestDefinition

	_jsii_.StaticInvoke(
		"cloudstructs.SlackAppManifestDefinition",
		"fromString",
		[]interface{}{manifest},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SlackAppManifestDefinition) Render(construct constructs.IConstruct) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"render",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Events API configuration for the app.
// See: https://api.slack.com/events-api
//
type SlackAppManifestEventSubscriptions struct {
	// The full https URL that acts as the Events API request URL.
	// See: https://api.slack.com/events-api#the-events-api__subscribing-to-event-types__events-api-request-urls
	//
	RequestUrl *string `field:"required" json:"requestUrl" yaml:"requestUrl"`
	// Event types you want the app to subscribe to.
	//
	// A maximum of 100 event types can be used.
	// See: https://api.slack.com/events
	//
	BotEvents *[]*string `field:"optional" json:"botEvents" yaml:"botEvents"`
	// Event types you want the app to subscribe to on behalf of authorized users.
	//
	// A maximum of 100 event types can be used.
	UserEvents *[]*string `field:"optional" json:"userEvents" yaml:"userEvents"`
}

// Interactivity configuration for the app.
// See: https://api.slack.com/interactivity/handling#setup
//
type SlackAppManifestInteractivity struct {
	// Whether or not interactivity features are enabled.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// The full https URL that acts as th interactive Options Load URL.
	MessageMenuOptionsUrl *string `field:"optional" json:"messageMenuOptionsUrl" yaml:"messageMenuOptionsUrl"`
	// The full https URL that acts as the interactive Request URL.
	RequestUrl *string `field:"optional" json:"requestUrl" yaml:"requestUrl"`
}

// OAuth configuration for the app.
type SlackAppManifestOauthConfig struct {
	// Bot scopes to request upon app installation.
	//
	// A maximum of 255 scopes can be included.
	// See: https://api.slack.com/scopes
	//
	BotScopes *[]*string `field:"optional" json:"botScopes" yaml:"botScopes"`
	// OAuth redirect URLs.
	//
	// A maximum of 1000 redirect URLs can be included.
	// See: https://api.slack.com/authentication/oauth-v2#redirect_urls
	//
	RedirectUrls *[]*string `field:"optional" json:"redirectUrls" yaml:"redirectUrls"`
	// User scopes to request upon app installation.
	//
	// A maximum of 255 scopes can be included.
	// See: https://api.slack.com/scopes
	//
	UserScopes *[]*string `field:"optional" json:"userScopes" yaml:"userScopes"`
}

// Properties for a Slack app manifest.
// See: https://api.slack.com/reference/manifests
//
type SlackAppManifestProps struct {
	// The name of the app.
	//
	// Maximum length is 35 characters.
	Name *string `field:"required" json:"name" yaml:"name"`
	// An array of IP addresses that conform to the Allowed IP Ranges feature.
	// See: https://api.slack.com/authentication/best-practices#ip_allowlisting
	//
	AllowedIpAddressRanges *[]*string `field:"optional" json:"allowedIpAddressRanges" yaml:"allowedIpAddressRanges"`
	// App Home configuration.
	// See: https://api.slack.com/surfaces/tabs
	//
	AppHome *SlackAppManifestAppHome `field:"optional" json:"appHome" yaml:"appHome"`
	// A hex color value that specifies the background color used on hovercards that display information about your app.
	//
	// Can be 3-digit (#000) or 6-digit (#000000) hex values with or without #.
	BackgroundColor *string `field:"optional" json:"backgroundColor" yaml:"backgroundColor"`
	// Bot user configuration.
	// See: https://api.slack.com/bot-users
	//
	BotUser *SlackkAppManifestBotUser `field:"optional" json:"botUser" yaml:"botUser"`
	// A short description of the app for display to users.
	//
	// Maximum length is 140 characters.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Events API configuration for the app.
	// See: https://api.slack.com/events-api
	//
	EventSubscriptions *SlackAppManifestEventSubscriptions `field:"optional" json:"eventSubscriptions" yaml:"eventSubscriptions"`
	// Interactivity configuration for the app.
	Interactivity *SlackAppManifestInteractivity `field:"optional" json:"interactivity" yaml:"interactivity"`
	// A longer version of the description of the app.
	//
	// Maximum length is 4000 characters.
	LongDescription *string `field:"optional" json:"longDescription" yaml:"longDescription"`
	// The major version of the manifest schema to target.
	MajorVersion *float64 `field:"optional" json:"majorVersion" yaml:"majorVersion"`
	// The minor version of the manifest schema to target.
	MinorVersion *float64 `field:"optional" json:"minorVersion" yaml:"minorVersion"`
	// OAuth configuration for the app.
	OauthConfig *SlackAppManifestOauthConfig `field:"optional" json:"oauthConfig" yaml:"oauthConfig"`
	// Whether org-wide deploy is enabled.
	// See: https://api.slack.com/enterprise/apps
	//
	OrgDeploy *bool `field:"optional" json:"orgDeploy" yaml:"orgDeploy"`
	// Shortcuts configuration.
	//
	// A maximum of 5 shortcuts can be included.
	// See: https://api.slack.com/interactivity/shortcuts
	//
	Shortcuts *[]*SlackAppManifestShortcut `field:"optional" json:"shortcuts" yaml:"shortcuts"`
	// Slash commands configuration.
	//
	// A maximum of 5 slash commands can be included.
	// See: https://api.slack.com/interactivity/slash-commands
	//
	SlashCommands *[]*SlackAppManifestSlashCommand `field:"optional" json:"slashCommands" yaml:"slashCommands"`
	// Whether Socket Mode is enabled.
	// See: https://api.slack.com/apis/connections/socket
	//
	SocketMode *bool `field:"optional" json:"socketMode" yaml:"socketMode"`
	// Valid unfurl domains to register.
	//
	// A maximum of 5 unfurl domains can be included.
	// See: https://api.slack.com/reference/messaging/link-unfurling#configuring_domains
	//
	UnfurlDomains *[]*string `field:"optional" json:"unfurlDomains" yaml:"unfurlDomains"`
	// Workflow steps.
	//
	// A maximum of 10 workflow steps can be included.
	// See: https://api.slack.com/workflows/steps
	//
	WorkflowSteps *[]*SlackAppManifestWorkflowStep `field:"optional" json:"workflowSteps" yaml:"workflowSteps"`
}

// Settings section of the app config pages.
type SlackAppManifestSettings struct {
	// An array of IP addresses that conform to the Allowed IP Ranges feature.
	// See: https://api.slack.com/authentication/best-practices#ip_allowlisting
	//
	AllowedIpAddressRanges *[]*string `field:"optional" json:"allowedIpAddressRanges" yaml:"allowedIpAddressRanges"`
	// Events API configuration for the app.
	// See: https://api.slack.com/events-api
	//
	EventSubscriptions *SlackAppManifestEventSubscriptions `field:"optional" json:"eventSubscriptions" yaml:"eventSubscriptions"`
	// Interactivity configuration for the app.
	Interactivity *SlackAppManifestInteractivity `field:"optional" json:"interactivity" yaml:"interactivity"`
	// Whether org-wide deploy is enabled.
	// See: https://api.slack.com/enterprise/apps
	//
	OrgDeploy *bool `field:"optional" json:"orgDeploy" yaml:"orgDeploy"`
	// Whether Socket Mode is enabled.
	// See: https://api.slack.com/apis/connections/socket
	//
	SocketMode *bool `field:"optional" json:"socketMode" yaml:"socketMode"`
}

// Shortcut configuration.
// See: https://api.slack.com/interactivity/shortcuts
//
type SlackAppManifestShortcut struct {
	// The callback ID of the shortcut.
	//
	// Maximum length is 255 characters.
	CallbackId *string `field:"required" json:"callbackId" yaml:"callbackId"`
	// A short description of the shortcut.
	//
	// Maximum length is 150 characters.
	Description *string `field:"required" json:"description" yaml:"description"`
	// The name of the shortcut.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The type of shortcut.
	// See: https://api.slack.com/interactivity/shortcuts
	//
	Type SlackAppManifestShortcutType `field:"required" json:"type" yaml:"type"`
}

// Type of shortcuts.
// See: https://api.slack.com/interactivity/shortcuts
//
type SlackAppManifestShortcutType string

const (
	// Message shortcuts are shown to users in the context menus of messages within Slack.
	// See: https://api.slack.com/interactivity/shortcuts/using#message_shortcuts
	//
	SlackAppManifestShortcutType_MESSAGE SlackAppManifestShortcutType = "MESSAGE"
	// Global shortcuts are available to users via the shortcuts button in the composer, and when using search in Slack.
	// See: https://api.slack.com/interactivity/shortcuts/using#global_shortcuts
	//
	SlackAppManifestShortcutType_GLOBAL SlackAppManifestShortcutType = "GLOBAL"
)

// Slash command configuration.
// See: https://api.slack.com/interactivity/slash-commands
//
type SlackAppManifestSlashCommand struct {
	// The actual slash command.
	//
	// Maximum length is 32 characters.
	Command *string `field:"required" json:"command" yaml:"command"`
	// The description of the slash command.
	//
	// Maximum length is 2000 characters.
	Description *string `field:"required" json:"description" yaml:"description"`
	// Whether channels, users, and links typed with the slash command should be escaped.
	ShouldEscape *bool `field:"optional" json:"shouldEscape" yaml:"shouldEscape"`
	// The full https URL that acts as the slash command's request URL.
	// See: https://api.slack.com/interactivity/slash-commands#creating_commands
	//
	Url *string `field:"optional" json:"url" yaml:"url"`
	// The short usage hint about the slash command for users.
	//
	// Maximum length is 1000 characters.
	UsageHint *string `field:"optional" json:"usageHint" yaml:"usageHint"`
}

// Workflow step.
// See: https://api.slack.com/workflows/steps
//
type SlackAppManifestWorkflowStep struct {
	// The callback ID of the workflow step.
	//
	// Maximum length of 50 characters.
	CallbackId *string `field:"required" json:"callbackId" yaml:"callbackId"`
	// The name of the workflow step.
	//
	// Maximum length of 50 characters.
	Name *string `field:"required" json:"name" yaml:"name"`
}

// Properties for a SlackApp.
type SlackAppProps struct {
	// An AWS Secrets Manager secret containing the app configuration token.
	//
	// Must use the following JSON format:
	//
	// ```
	// {
	//    "refreshToken": "<token>"
	// }
	// ```.
	ConfigurationTokenSecret awssecretsmanager.ISecret `field:"required" json:"configurationTokenSecret" yaml:"configurationTokenSecret"`
	// The definition of the app manifest.
	// See: https://api.slack.com/reference/manifests
	//
	Manifest SlackAppManifestDefinition `field:"required" json:"manifest" yaml:"manifest"`
	// The AWS Secrets Manager secret where to store the app credentials.
	CredentialsSecret awssecretsmanager.ISecret `field:"optional" json:"credentialsSecret" yaml:"credentialsSecret"`
}

// Send Slack events to Amazon EventBridge.
type SlackEvents interface {
	constructs.Construct
	// The custom event bus where Slack events are sent.
	EventBus() awsevents.EventBus
	// The tree node.
	Node() constructs.Node
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for SlackEvents
type jsiiProxy_SlackEvents struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_SlackEvents) EventBus() awsevents.EventBus {
	var returns awsevents.EventBus
	_jsii_.Get(
		j,
		"eventBus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SlackEvents) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


func NewSlackEvents(scope constructs.Construct, id *string, props *SlackEventsProps) SlackEvents {
	_init_.Initialize()

	j := jsiiProxy_SlackEvents{}

	_jsii_.Create(
		"cloudstructs.SlackEvents",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewSlackEvents_Override(s SlackEvents, scope constructs.Construct, id *string, props *SlackEventsProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cloudstructs.SlackEvents",
		[]interface{}{scope, id, props},
		s,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func SlackEvents_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cloudstructs.SlackEvents",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SlackEvents) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties for a SlackEvents.
type SlackEventsProps struct {
	// The signing secret of the Slack app.
	SigningSecret awscdk.SecretValue `field:"required" json:"signingSecret" yaml:"signingSecret"`
	// A name for the API Gateway resource.
	ApiName *string `field:"optional" json:"apiName" yaml:"apiName"`
	// Whether to use a custom event bus.
	CustomEventBus *bool `field:"optional" json:"customEventBus" yaml:"customEventBus"`
}

// Extract text from images posted to Slack using Amazon Textract.
type SlackTextract interface {
	constructs.Construct
	// The tree node.
	Node() constructs.Node
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for SlackTextract
type jsiiProxy_SlackTextract struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_SlackTextract) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


func NewSlackTextract(scope constructs.Construct, id *string, props *SlackTextractProps) SlackTextract {
	_init_.Initialize()

	j := jsiiProxy_SlackTextract{}

	_jsii_.Create(
		"cloudstructs.SlackTextract",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewSlackTextract_Override(s SlackTextract, scope constructs.Construct, id *string, props *SlackTextractProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cloudstructs.SlackTextract",
		[]interface{}{scope, id, props},
		s,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func SlackTextract_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cloudstructs.SlackTextract",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SlackTextract) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties for a SlackTextract.
type SlackTextractProps struct {
	// The application id of the Slack app.
	AppId *string `field:"required" json:"appId" yaml:"appId"`
	// The **bot** token of the Slack app.
	//
	// The following scopes are required: `chat:write` and `files:read`.
	BotToken awscdk.SecretValue `field:"required" json:"botToken" yaml:"botToken"`
	// The signing secret of the Slack app.
	SigningSecret awscdk.SecretValue `field:"required" json:"signingSecret" yaml:"signingSecret"`
}

// Bot user configuration.
// See: https://api.slack.com/bot-users
//
type SlackkAppManifestBotUser struct {
	// The display name of the bot user.
	//
	// Maximum length is 80 characters.
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// Whether the bot user will always appear to be online.
	AlwaysOnline *bool `field:"optional" json:"alwaysOnline" yaml:"alwaysOnline"`
}

// A state machine custom resource provider.
type StateMachineCustomResourceProvider interface {
	constructs.Construct
	// The tree node.
	Node() constructs.Node
	// The service token.
	ServiceToken() *string
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for StateMachineCustomResourceProvider
type jsiiProxy_StateMachineCustomResourceProvider struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_StateMachineCustomResourceProvider) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StateMachineCustomResourceProvider) ServiceToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceToken",
		&returns,
	)
	return returns
}


func NewStateMachineCustomResourceProvider(scope constructs.Construct, id *string, props *StateMachineCustomResourceProviderProps) StateMachineCustomResourceProvider {
	_init_.Initialize()

	j := jsiiProxy_StateMachineCustomResourceProvider{}

	_jsii_.Create(
		"cloudstructs.StateMachineCustomResourceProvider",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewStateMachineCustomResourceProvider_Override(s StateMachineCustomResourceProvider, scope constructs.Construct, id *string, props *StateMachineCustomResourceProviderProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cloudstructs.StateMachineCustomResourceProvider",
		[]interface{}{scope, id, props},
		s,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func StateMachineCustomResourceProvider_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cloudstructs.StateMachineCustomResourceProvider",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StateMachineCustomResourceProvider) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties for a StateMachineCustomResourceProvider.
type StateMachineCustomResourceProviderProps struct {
	// The state machine.
	StateMachine IStateMachine `field:"required" json:"stateMachine" yaml:"stateMachine"`
	// Timeout.
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
}

// A CloudFront static website hosted on S3.
type StaticWebsite interface {
	constructs.Construct
	// The S3 bucket of this static website.
	Bucket() awss3.Bucket
	// The CloudFront distribution of this static website.
	Distribution() awscloudfront.Distribution
	// The tree node.
	Node() constructs.Node
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for StaticWebsite
type jsiiProxy_StaticWebsite struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_StaticWebsite) Bucket() awss3.Bucket {
	var returns awss3.Bucket
	_jsii_.Get(
		j,
		"bucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StaticWebsite) Distribution() awscloudfront.Distribution {
	var returns awscloudfront.Distribution
	_jsii_.Get(
		j,
		"distribution",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StaticWebsite) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


func NewStaticWebsite(scope constructs.Construct, id *string, props *StaticWebsiteProps) StaticWebsite {
	_init_.Initialize()

	j := jsiiProxy_StaticWebsite{}

	_jsii_.Create(
		"cloudstructs.StaticWebsite",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewStaticWebsite_Override(s StaticWebsite, scope constructs.Construct, id *string, props *StaticWebsiteProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cloudstructs.StaticWebsite",
		[]interface{}{scope, id, props},
		s,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func StaticWebsite_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cloudstructs.StaticWebsite",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func StaticWebsite_DefaultSecurityHeadersBehavior() *awscloudfront.ResponseSecurityHeadersBehavior {
	_init_.Initialize()
	var returns *awscloudfront.ResponseSecurityHeadersBehavior
	_jsii_.StaticGet(
		"cloudstructs.StaticWebsite",
		"defaultSecurityHeadersBehavior",
		&returns,
	)
	return returns
}

func StaticWebsite_SetDefaultSecurityHeadersBehavior(val *awscloudfront.ResponseSecurityHeadersBehavior) {
	_init_.Initialize()
	_jsii_.StaticSet(
		"cloudstructs.StaticWebsite",
		"defaultSecurityHeadersBehavior",
		val,
	)
}

func (s *jsiiProxy_StaticWebsite) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties for a StaticWebsite.
type StaticWebsiteProps struct {
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
	// A list of domain names that should redirect to `domainName`.
	Redirects *[]*string `field:"optional" json:"redirects" yaml:"redirects"`
	// Response headers policy for the default behavior.
	ResponseHeadersPolicy awscloudfront.ResponseHeadersPolicy `field:"optional" json:"responseHeadersPolicy" yaml:"responseHeadersPolicy"`
}

// Clean unused S3 and ECR assets from your CDK Toolkit.
type ToolkitCleaner interface {
	constructs.Construct
	// The tree node.
	Node() constructs.Node
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for ToolkitCleaner
type jsiiProxy_ToolkitCleaner struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_ToolkitCleaner) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


func NewToolkitCleaner(scope constructs.Construct, id *string, props *ToolkitCleanerProps) ToolkitCleaner {
	_init_.Initialize()

	j := jsiiProxy_ToolkitCleaner{}

	_jsii_.Create(
		"cloudstructs.ToolkitCleaner",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewToolkitCleaner_Override(t ToolkitCleaner, scope constructs.Construct, id *string, props *ToolkitCleanerProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cloudstructs.ToolkitCleaner",
		[]interface{}{scope, id, props},
		t,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func ToolkitCleaner_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cloudstructs.ToolkitCleaner",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_ToolkitCleaner) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties for a ToolkitCleaner.
type ToolkitCleanerProps struct {
	// Only output number of assets and total size that would be deleted but without actually deleting assets.
	DryRun *bool `field:"optional" json:"dryRun" yaml:"dryRun"`
	// Retain unused assets that were created recently.
	RetainAssetsNewerThan awscdk.Duration `field:"optional" json:"retainAssetsNewerThan" yaml:"retainAssetsNewerThan"`
	// The schedule for the cleaner.
	Schedule awsevents.Schedule `field:"optional" json:"schedule" yaml:"schedule"`
	// Whether to clean on schedule.
	//
	// If you'd like to run the cleanup manually
	// via the console, set to `false`.
	ScheduleEnabled *bool `field:"optional" json:"scheduleEnabled" yaml:"scheduleEnabled"`
}

// URL shortener.
type UrlShortener interface {
	constructs.Construct
	// The underlying API Gateway REST API.
	Api() awsapigateway.LambdaRestApi
	// The endpoint of the URL shortener API.
	ApiEndpoint() *string
	// The tree node.
	Node() constructs.Node
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for UrlShortener
type jsiiProxy_UrlShortener struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_UrlShortener) Api() awsapigateway.LambdaRestApi {
	var returns awsapigateway.LambdaRestApi
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

	var returns *bool

	_jsii_.StaticInvoke(
		"cloudstructs.UrlShortener",
		"isConstruct",
		[]interface{}{x},
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

