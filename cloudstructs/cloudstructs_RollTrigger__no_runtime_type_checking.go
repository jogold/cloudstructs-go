//go:build no_runtime_type_checking

// High-level constructs for AWS CDK
package cloudstructs

// Building without runtime type checking enabled, so all the below just return nil

func validateRollTrigger_FromRuleParameters(rule awsevents.Rule) error {
	return nil
}

func validateRollTrigger_FromScheduleParameters(schedule awsevents.Schedule) error {
	return nil
}

