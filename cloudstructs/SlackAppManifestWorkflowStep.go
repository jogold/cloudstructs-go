package cloudstructs


// Workflow step.
// See: https://api.slack.com/workflows/steps
//
type SlackAppManifestWorkflowStep struct {
	// The callback ID of the workflow step.
	//
	// Maximum length of 50 characters.
	CallbackId *string `field:"required" json:"callbackId" yaml:"callbackId"`
	// The name of the workflow step.
	//
	// Maximum length of 50 characters.
	Name *string `field:"required" json:"name" yaml:"name"`
}

