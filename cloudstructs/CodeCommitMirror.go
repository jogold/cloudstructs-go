// High-level constructs for AWS CDK
package cloudstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/jogold/cloudstructs-go/cloudstructs/jsii"

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

	if err := validateNewCodeCommitMirrorParameters(scope, id, props); err != nil {
		panic(err)
	}
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

	if err := validateCodeCommitMirror_IsConstructParameters(x); err != nil {
		panic(err)
	}
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

