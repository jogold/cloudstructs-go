//go:build !no_runtime_type_checking

package cloudstructs

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

func validateRollTrigger_FromRuleParameters(rule awsevents.Rule) error {
	if rule == nil {
		return fmt.Errorf("parameter rule is required, but nil was provided")
	}

	return nil
}

func validateRollTrigger_FromScheduleParameters(schedule awsevents.Schedule) error {
	if schedule == nil {
		return fmt.Errorf("parameter schedule is required, but nil was provided")
	}

	return nil
}

