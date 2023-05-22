//go:build no_runtime_type_checking

package cloudstructs

// Building without runtime type checking enabled, so all the below just return nil

func validateDmarcReporter_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewDmarcReporterParameters(scope constructs.Construct, id *string, props *DmarcReporterProps) error {
	return nil
}

