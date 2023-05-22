//go:build no_runtime_type_checking

package cloudstructs

// Building without runtime type checking enabled, so all the below just return nil

func validateSlackEvents_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewSlackEventsParameters(scope constructs.Construct, id *string, props *SlackEventsProps) error {
	return nil
}

