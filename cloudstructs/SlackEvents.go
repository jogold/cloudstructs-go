// High-level constructs for AWS CDK
package cloudstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/jogold/cloudstructs-go/cloudstructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/jogold/cloudstructs-go/cloudstructs/internal"
)

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

	if err := validateNewSlackEventsParameters(scope, id, props); err != nil {
		panic(err)
	}
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

	if err := validateSlackEvents_IsConstructParameters(x); err != nil {
		panic(err)
	}
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

