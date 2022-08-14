// High-level constructs for AWS CDK
package cloudstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/jogold/cloudstructs-go/cloudstructs/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/jogold/cloudstructs-go/cloudstructs/internal"
)

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

