package cloudstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/jogold/cloudstructs-go/cloudstructs/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/jogold/cloudstructs-go/cloudstructs/internal"
)

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

	if err := validateNewEcsServiceRollerParameters(scope, id, props); err != nil {
		panic(err)
	}
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

	if err := validateEcsServiceRoller_IsConstructParameters(x); err != nil {
		panic(err)
	}
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

