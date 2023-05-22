package cloudstructs


// Interactivity configuration for the app.
// See: https://api.slack.com/interactivity/handling#setup
//
type SlackAppManifestInteractivity struct {
	// Whether or not interactivity features are enabled.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// The full https URL that acts as th interactive Options Load URL.
	MessageMenuOptionsUrl *string `field:"optional" json:"messageMenuOptionsUrl" yaml:"messageMenuOptionsUrl"`
	// The full https URL that acts as the interactive Request URL.
	RequestUrl *string `field:"optional" json:"requestUrl" yaml:"requestUrl"`
}

