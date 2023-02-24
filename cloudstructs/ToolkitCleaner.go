// High-level constructs for AWS CDK
package cloudstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/jogold/cloudstructs-go/cloudstructs/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/jogold/cloudstructs-go/cloudstructs/internal"
)

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

	if err := validateNewToolkitCleanerParameters(scope, id, props); err != nil {
		panic(err)
	}
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

	if err := validateToolkitCleaner_IsConstructParameters(x); err != nil {
		panic(err)
	}
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

