package cloudstructs


// Properties for a Slack app manifest.
// See: https://api.slack.com/reference/manifests
//
type SlackAppManifestProps struct {
	// The name of the app.
	//
	// Maximum length is 35 characters.
	Name *string `field:"required" json:"name" yaml:"name"`
	// An array of IP addresses that conform to the Allowed IP Ranges feature.
	// See: https://api.slack.com/authentication/best-practices#ip_allowlisting
	//
	AllowedIpAddressRanges *[]*string `field:"optional" json:"allowedIpAddressRanges" yaml:"allowedIpAddressRanges"`
	// App Home configuration.
	// See: https://api.slack.com/surfaces/tabs
	//
	AppHome *SlackAppManifestAppHome `field:"optional" json:"appHome" yaml:"appHome"`
	// A hex color value that specifies the background color used on hovercards that display information about your app.
	//
	// Can be 3-digit (#000) or 6-digit (#000000) hex values with or without #.
	BackgroundColor *string `field:"optional" json:"backgroundColor" yaml:"backgroundColor"`
	// Bot user configuration.
	// See: https://api.slack.com/bot-users
	//
	BotUser *SlackkAppManifestBotUser `field:"optional" json:"botUser" yaml:"botUser"`
	// A short description of the app for display to users.
	//
	// Maximum length is 140 characters.
	// Default: - no short description.
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Events API configuration for the app.
	// See: https://api.slack.com/events-api
	//
	EventSubscriptions *SlackAppManifestEventSubscriptions `field:"optional" json:"eventSubscriptions" yaml:"eventSubscriptions"`
	// Interactivity configuration for the app.
	Interactivity *SlackAppManifestInteractivity `field:"optional" json:"interactivity" yaml:"interactivity"`
	// A longer version of the description of the app.
	//
	// Maximum length is 4000 characters.
	LongDescription *string `field:"optional" json:"longDescription" yaml:"longDescription"`
	// The major version of the manifest schema to target.
	// Default: - do not target a specific major version.
	//
	MajorVersion *float64 `field:"optional" json:"majorVersion" yaml:"majorVersion"`
	// The minor version of the manifest schema to target.
	// Default: - do not target a specific minor version.
	//
	MinorVersion *float64 `field:"optional" json:"minorVersion" yaml:"minorVersion"`
	// OAuth configuration for the app.
	OauthConfig *SlackAppManifestOauthConfig `field:"optional" json:"oauthConfig" yaml:"oauthConfig"`
	// Whether org-wide deploy is enabled.
	// See: https://api.slack.com/enterprise/apps
	//
	// Default: false.
	//
	OrgDeploy *bool `field:"optional" json:"orgDeploy" yaml:"orgDeploy"`
	// Shortcuts configuration.
	//
	// A maximum of 5 shortcuts can be included.
	// See: https://api.slack.com/interactivity/shortcuts
	//
	Shortcuts *[]*SlackAppManifestShortcut `field:"optional" json:"shortcuts" yaml:"shortcuts"`
	// Slash commands configuration.
	//
	// A maximum of 5 slash commands can be included.
	// See: https://api.slack.com/interactivity/slash-commands
	//
	SlashCommands *[]*SlackAppManifestSlashCommand `field:"optional" json:"slashCommands" yaml:"slashCommands"`
	// Whether Socket Mode is enabled.
	// See: https://api.slack.com/apis/connections/socket
	//
	// Default: false.
	//
	SocketMode *bool `field:"optional" json:"socketMode" yaml:"socketMode"`
	// Valid unfurl domains to register.
	//
	// A maximum of 5 unfurl domains can be included.
	// See: https://api.slack.com/reference/messaging/link-unfurling#configuring_domains
	//
	UnfurlDomains *[]*string `field:"optional" json:"unfurlDomains" yaml:"unfurlDomains"`
	// Workflow steps.
	//
	// A maximum of 10 workflow steps can be included.
	// See: https://api.slack.com/workflows/steps
	//
	WorkflowSteps *[]*SlackAppManifestWorkflowStep `field:"optional" json:"workflowSteps" yaml:"workflowSteps"`
}

