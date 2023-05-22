package cloudstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/jogold/cloudstructs-go/cloudstructs/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/jogold/cloudstructs-go/cloudstructs/internal"
)

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

	if err := validateNewEmailReceiverParameters(scope, id, props); err != nil {
		panic(err)
	}
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

	if err := validateEmailReceiver_IsConstructParameters(x); err != nil {
		panic(err)
	}
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

