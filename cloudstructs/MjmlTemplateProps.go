package cloudstructs


// Properties for a MjmlTemplate.
type MjmlTemplateProps struct {
	// The MJML for the email.
	Mjml *string `field:"required" json:"mjml" yaml:"mjml"`
	// The subject line of the email.
	Subject *string `field:"required" json:"subject" yaml:"subject"`
	// The name of the template.
	// Default: - a CloudFormation generated name.
	//
	TemplateName *string `field:"optional" json:"templateName" yaml:"templateName"`
}

