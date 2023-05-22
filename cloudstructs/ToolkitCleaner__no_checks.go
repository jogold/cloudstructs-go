//go:build no_runtime_type_checking

package cloudstructs

// Building without runtime type checking enabled, so all the below just return nil

func validateToolkitCleaner_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewToolkitCleanerParameters(scope constructs.Construct, id *string, props *ToolkitCleanerProps) error {
	return nil
}

