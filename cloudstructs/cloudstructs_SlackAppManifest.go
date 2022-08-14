// High-level constructs for AWS CDK
package cloudstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/jogold/cloudstructs-go/cloudstructs/jsii"

	"github.com/aws/constructs-go/constructs/v10"
)

// A Slack app manifest.
// See: https://api.slack.com/reference/manifests
//
type SlackAppManifest interface {
	Render(construct constructs.IConstruct) *string
}

// The jsii proxy struct for SlackAppManifest
type jsiiProxy_SlackAppManifest struct {
	_ byte // padding
}

func NewSlackAppManifest(props *SlackAppManifestProps) SlackAppManifest {
	_init_.Initialize()

	j := jsiiProxy_SlackAppManifest{}

	_jsii_.Create(
		"cloudstructs.SlackAppManifest",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewSlackAppManifest_Override(s SlackAppManifest, props *SlackAppManifestProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cloudstructs.SlackAppManifest",
		[]interface{}{props},
		s,
	)
}

func (s *jsiiProxy_SlackAppManifest) Render(construct constructs.IConstruct) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"render",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

