// High-level constructs for AWS CDK
package cloudstructs


// Bot user configuration.
// See: https://api.slack.com/bot-users
//
type SlackkAppManifestBotUser struct {
	// The display name of the bot user.
	//
	// Maximum length is 80 characters.
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// Whether the bot user will always appear to be online.
	AlwaysOnline *bool `field:"optional" json:"alwaysOnline" yaml:"alwaysOnline"`
}
