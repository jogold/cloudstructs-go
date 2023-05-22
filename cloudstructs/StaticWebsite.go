package cloudstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/jogold/cloudstructs-go/cloudstructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfront"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/jogold/cloudstructs-go/cloudstructs/internal"
)

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

	if err := validateNewStaticWebsiteParameters(scope, id, props); err != nil {
		panic(err)
	}
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

	if err := validateStaticWebsite_IsConstructParameters(x); err != nil {
		panic(err)
	}
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
	if err := validateStaticWebsite_SetDefaultSecurityHeadersBehaviorParameters(val); err != nil {
		panic(err)
	}
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

