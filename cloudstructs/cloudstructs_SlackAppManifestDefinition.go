// High-level constructs for AWS CDK
package cloudstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/jogold/cloudstructs-go/cloudstructs/jsii"

	"github.com/aws/constructs-go/constructs/v10"
)

// A Slack app manifest definition.
type SlackAppManifestDefinition interface {
	// Renders the JSON app manifest encoded as a string.
	Render(construct constructs.IConstruct) *string
}

// The jsii proxy struct for SlackAppManifestDefinition
type jsiiProxy_SlackAppManifestDefinition struct {
	_ byte // padding
}

func NewSlackAppManifestDefinition_Override(s SlackAppManifestDefinition) {
	_init_.Initialize()

	_jsii_.Create(
		"cloudstructs.SlackAppManifestDefinition",
		nil, // no parameters
		s,
	)
}

// Creates a Slack app manifest from a file containg a JSON app manifest.
func SlackAppManifestDefinition_FromFile(file *string) SlackAppManifestDefinition {
	_init_.Initialize()

	var returns SlackAppManifestDefinition

	_jsii_.StaticInvoke(
		"cloudstructs.SlackAppManifestDefinition",
		"fromFile",
		[]interface{}{file},
		&returns,
	)

	return returns
}

// Creates a Slack app manifest by specifying properties.
func SlackAppManifestDefinition_FromManifest(props *SlackAppManifestProps) SlackAppManifestDefinition {
	_init_.Initialize()

	var returns SlackAppManifestDefinition

	_jsii_.StaticInvoke(
		"cloudstructs.SlackAppManifestDefinition",
		"fromManifest",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Create a Slack app manifest from a JSON app manifest encoded as a string.
func SlackAppManifestDefinition_FromString(manifest *string) SlackAppManifestDefinition {
	_init_.Initialize()

	var returns SlackAppManifestDefinition

	_jsii_.StaticInvoke(
		"cloudstructs.SlackAppManifestDefinition",
		"fromString",
		[]interface{}{manifest},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SlackAppManifestDefinition) Render(construct constructs.IConstruct) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"render",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

