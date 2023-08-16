package cloudstructs


// App Home configuration.
// See: https://api.slack.com/surfaces/tabs
//
type SlackAppManifestAppHome struct {
	// Wether the Home tab is enabled.
	// Default: false.
	//
	HomeTab *bool `field:"optional" json:"homeTab" yaml:"homeTab"`
	// Wether the Messages is enabled.
	// Default: false.
	//
	MessagesTab *bool `field:"optional" json:"messagesTab" yaml:"messagesTab"`
	// Whether the users can send messages to your app in the Messages tab of your App Home.
	// Default: false.
	//
	MessagesTabReadOnly *bool `field:"optional" json:"messagesTabReadOnly" yaml:"messagesTabReadOnly"`
}

