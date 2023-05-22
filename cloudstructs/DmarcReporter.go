package cloudstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/jogold/cloudstructs-go/cloudstructs/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/jogold/cloudstructs-go/cloudstructs/internal"
)

// Creates a DMARC record in Route 53 and invokes a Lambda function to process incoming reports.
type DmarcReporter interface {
	constructs.Construct
	// The tree node.
	Node() constructs.Node
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for DmarcReporter
type jsiiProxy_DmarcReporter struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_DmarcReporter) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


func NewDmarcReporter(scope constructs.Construct, id *string, props *DmarcReporterProps) DmarcReporter {
	_init_.Initialize()

	if err := validateNewDmarcReporterParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_DmarcReporter{}

	_jsii_.Create(
		"cloudstructs.DmarcReporter",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewDmarcReporter_Override(d DmarcReporter, scope constructs.Construct, id *string, props *DmarcReporterProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cloudstructs.DmarcReporter",
		[]interface{}{scope, id, props},
		d,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func DmarcReporter_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDmarcReporter_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cloudstructs.DmarcReporter",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmarcReporter) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

