//go:build no_runtime_type_checking

package cloudstructs

// Building without runtime type checking enabled, so all the below just return nil

func validateEcsServiceRoller_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewEcsServiceRollerParameters(scope constructs.Construct, id *string, props *EcsServiceRollerProps) error {
	return nil
}

