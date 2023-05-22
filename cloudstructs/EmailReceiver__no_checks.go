//go:build no_runtime_type_checking

package cloudstructs

// Building without runtime type checking enabled, so all the below just return nil

func validateEmailReceiver_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewEmailReceiverParameters(scope constructs.Construct, id *string, props *EmailReceiverProps) error {
	return nil
}

