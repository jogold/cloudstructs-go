package cloudstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/jogold/cloudstructs-go/cloudstructs/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/jogold/cloudstructs-go/cloudstructs/internal"
)

// SES email template from [MJML](https://mjml.io/).
type MjmlTemplate interface {
	constructs.Construct
	// The tree node.
	Node() constructs.Node
	// The name of the template.
	TemplateName() *string
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for MjmlTemplate
type jsiiProxy_MjmlTemplate struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_MjmlTemplate) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MjmlTemplate) TemplateName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"templateName",
		&returns,
	)
	return returns
}


func NewMjmlTemplate(scope constructs.Construct, id *string, props *MjmlTemplateProps) MjmlTemplate {
	_init_.Initialize()

	if err := validateNewMjmlTemplateParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_MjmlTemplate{}

	_jsii_.Create(
		"cloudstructs.MjmlTemplate",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewMjmlTemplate_Override(m MjmlTemplate, scope constructs.Construct, id *string, props *MjmlTemplateProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cloudstructs.MjmlTemplate",
		[]interface{}{scope, id, props},
		m,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func MjmlTemplate_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMjmlTemplate_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cloudstructs.MjmlTemplate",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MjmlTemplate) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

