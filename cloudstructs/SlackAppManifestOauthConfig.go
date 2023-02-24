// High-level constructs for AWS CDK
package cloudstructs


// OAuth configuration for the app.
type SlackAppManifestOauthConfig struct {
	// Bot scopes to request upon app installation.
	//
	// A maximum of 255 scopes can be included.
	// See: https://api.slack.com/scopes
	//
	BotScopes *[]*string `field:"optional" json:"botScopes" yaml:"botScopes"`
	// OAuth redirect URLs.
	//
	// A maximum of 1000 redirect URLs can be included.
	// See: https://api.slack.com/authentication/oauth-v2#redirect_urls
	//
	RedirectUrls *[]*string `field:"optional" json:"redirectUrls" yaml:"redirectUrls"`
	// User scopes to request upon app installation.
	//
	// A maximum of 255 scopes can be included.
	// See: https://api.slack.com/scopes
	//
	UserScopes *[]*string `field:"optional" json:"userScopes" yaml:"userScopes"`
}

