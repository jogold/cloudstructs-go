//go:build no_runtime_type_checking

// High-level constructs for AWS CDK
package cloudstructs

// Building without runtime type checking enabled, so all the below just return nil

func validateSamlIdentityProvider_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewSamlIdentityProviderParameters(scope constructs.Construct, id *string, props *SamlIdentityProviderProps) error {
	return nil
}

