// High-level constructs for AWS CDK
package cloudstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/jogold/cloudstructs-go/cloudstructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/jogold/cloudstructs-go/cloudstructs/internal"
)

// Perform SSL server test for a hostname.
type SslServerTest interface {
	constructs.Construct
	// The topic to which the SSL test results are sent when the grade is below the minimum grade.
	AlarmTopic() awssns.ITopic
	// The tree node.
	Node() constructs.Node
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for SslServerTest
type jsiiProxy_SslServerTest struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_SslServerTest) AlarmTopic() awssns.ITopic {
	var returns awssns.ITopic
	_jsii_.Get(
		j,
		"alarmTopic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SslServerTest) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


func NewSslServerTest(scope constructs.Construct, id *string, props *SslServerTestProps) SslServerTest {
	_init_.Initialize()

	if err := validateNewSslServerTestParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_SslServerTest{}

	_jsii_.Create(
		"cloudstructs.SslServerTest",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewSslServerTest_Override(s SslServerTest, scope constructs.Construct, id *string, props *SslServerTestProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cloudstructs.SslServerTest",
		[]interface{}{scope, id, props},
		s,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func SslServerTest_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSslServerTest_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cloudstructs.SslServerTest",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SslServerTest) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

