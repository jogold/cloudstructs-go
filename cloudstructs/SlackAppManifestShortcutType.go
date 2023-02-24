// High-level constructs for AWS CDK
package cloudstructs


// Type of shortcuts.
// See: https://api.slack.com/interactivity/shortcuts
//
type SlackAppManifestShortcutType string

const (
	// Message shortcuts are shown to users in the context menus of messages within Slack.
	// See: https://api.slack.com/interactivity/shortcuts/using#message_shortcuts
	//
	SlackAppManifestShortcutType_MESSAGE SlackAppManifestShortcutType = "MESSAGE"
	// Global shortcuts are available to users via the shortcuts button in the composer, and when using search in Slack.
	// See: https://api.slack.com/interactivity/shortcuts/using#global_shortcuts
	//
	SlackAppManifestShortcutType_GLOBAL SlackAppManifestShortcutType = "GLOBAL"
)

