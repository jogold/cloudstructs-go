// High-level constructs for AWS CDK
package cloudstructs


// Settings section of the app config pages.
type SlackAppManifestSettings struct {
	// An array of IP addresses that conform to the Allowed IP Ranges feature.
	// See: https://api.slack.com/authentication/best-practices#ip_allowlisting
	//
	AllowedIpAddressRanges *[]*string `field:"optional" json:"allowedIpAddressRanges" yaml:"allowedIpAddressRanges"`
	// Events API configuration for the app.
	// See: https://api.slack.com/events-api
	//
	EventSubscriptions *SlackAppManifestEventSubscriptions `field:"optional" json:"eventSubscriptions" yaml:"eventSubscriptions"`
	// Interactivity configuration for the app.
	Interactivity *SlackAppManifestInteractivity `field:"optional" json:"interactivity" yaml:"interactivity"`
	// Whether org-wide deploy is enabled.
	// See: https://api.slack.com/enterprise/apps
	//
	OrgDeploy *bool `field:"optional" json:"orgDeploy" yaml:"orgDeploy"`
	// Whether Socket Mode is enabled.
	// See: https://api.slack.com/apis/connections/socket
	//
	SocketMode *bool `field:"optional" json:"socketMode" yaml:"socketMode"`
}

