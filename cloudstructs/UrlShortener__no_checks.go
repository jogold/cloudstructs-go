//go:build no_runtime_type_checking

// High-level constructs for AWS CDK
package cloudstructs

// Building without runtime type checking enabled, so all the below just return nil

func validateUrlShortener_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewUrlShortenerParameters(scope constructs.Construct, id *string, props *UrlShortenerProps) error {
	return nil
}

