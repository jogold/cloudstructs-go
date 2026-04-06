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

func (s *jsiiProxy_SlackEvents) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

