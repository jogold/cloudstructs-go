// High-level constructs for AWS CDK
package cloudstructs


// Shortcut configuration.
// See: https://api.slack.com/interactivity/shortcuts
//
type SlackAppManifestShortcut struct {
	// The callback ID of the shortcut.
	//
	// Maximum length is 255 characters.
	CallbackId *string `field:"required" json:"callbackId" yaml:"callbackId"`
	// A short description of the shortcut.
	//
	// Maximum length is 150 characters.
	Description *string `field:"required" json:"description" yaml:"description"`
	// The name of the shortcut.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The type of shortcut.
	// See: https://api.slack.com/interactivity/shortcuts
	//
	Type SlackAppManifestShortcutType `field:"required" json:"type" yaml:"type"`
}

