// High-level constructs for AWS CDK
package cloudstructs


// Events API configuration for the app.
// See: https://api.slack.com/events-api
//
type SlackAppManifestEventSubscriptions struct {
	// The full https URL that acts as the Events API request URL.
	// See: https://api.slack.com/events-api#the-events-api__subscribing-to-event-types__events-api-request-urls
	//
	RequestUrl *string `field:"required" json:"requestUrl" yaml:"requestUrl"`
	// Event types you want the app to subscribe to.
	//
	// A maximum of 100 event types can be used.
	// See: https://api.slack.com/events
	//
	BotEvents *[]*string `field:"optional" json:"botEvents" yaml:"botEvents"`
	// Event types you want the app to subscribe to on behalf of authorized users.
	//
	// A maximum of 100 event types can be used.
	UserEvents *[]*string `field:"optional" json:"userEvents" yaml:"userEvents"`
}

