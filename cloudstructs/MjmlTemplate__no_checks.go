//go:build no_runtime_type_checking

package cloudstructs

// Building without runtime type checking enabled, so all the below just return nil

func validateMjmlTemplate_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewMjmlTemplateParameters(scope constructs.Construct, id *string, props *MjmlTemplateProps) error {
	return nil
}

