// High-level constructs for AWS CDK
package cloudstructs


// App Home configuration.
// See: https://api.slack.com/surfaces/tabs
//
type SlackAppManifestAppHome struct {
	// Wether the Home tab is enabled.
	HomeTab *bool `field:"optional" json:"homeTab" yaml:"homeTab"`
	// Wether the Messages is enabled.
	MessagesTab *bool `field:"optional" json:"messagesTab" yaml:"messagesTab"`
	// Whether the users can send messages to your app in the Messages tab of your App Home.
	MessagesTabReadOnly *bool `field:"optional" json:"messagesTabReadOnly" yaml:"messagesTabReadOnly"`
}

