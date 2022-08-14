// High-level constructs for AWS CDK
package cloudstructs


// Slash command configuration.
// See: https://api.slack.com/interactivity/slash-commands
//
type SlackAppManifestSlashCommand struct {
	// The actual slash command.
	//
	// Maximum length is 32 characters.
	Command *string `field:"required" json:"command" yaml:"command"`
	// The description of the slash command.
	//
	// Maximum length is 2000 characters.
	Description *string `field:"required" json:"description" yaml:"description"`
	// Whether channels, users, and links typed with the slash command should be escaped.
	ShouldEscape *bool `field:"optional" json:"shouldEscape" yaml:"shouldEscape"`
	// The full https URL that acts as the slash command's request URL.
	// See: https://api.slack.com/interactivity/slash-commands#creating_commands
	//
	Url *string `field:"optional" json:"url" yaml:"url"`
	// The short usage hint about the slash command for users.
	//
	// Maximum length is 1000 characters.
	UsageHint *string `field:"optional" json:"usageHint" yaml:"usageHint"`
}

