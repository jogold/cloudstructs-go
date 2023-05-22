//go:build no_runtime_type_checking

package cloudstructs

// Building without runtime type checking enabled, so all the below just return nil

func validateStaticWebsite_IsConstructParameters(x interface{}) error {
	return nil
}

func validateStaticWebsite_SetDefaultSecurityHeadersBehaviorParameters(val *awscloudfront.ResponseSecurityHeadersBehavior) error {
	return nil
}

func validateNewStaticWebsiteParameters(scope constructs.Construct, id *string, props *StaticWebsiteProps) error {
	return nil
}

