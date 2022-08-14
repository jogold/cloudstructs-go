// High-level constructs for AWS CDK
package cloudstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/jogold/cloudstructs-go/cloudstructs/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/jogold/cloudstructs-go/cloudstructs/internal"
)

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

