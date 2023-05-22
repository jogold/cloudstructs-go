package cloudstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/jogold/cloudstructs-go/cloudstructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// The rule or schedule that should trigger a roll.
type RollTrigger interface {
	// Roll rule.
	Rule() awsevents.Rule
	// Roll schedule.
	Schedule() awsevents.Schedule
}

// The jsii proxy struct for RollTrigger
type jsiiProxy_RollTrigger struct {
	_ byte // padding
}

func (j *jsiiProxy_RollTrigger) Rule() awsevents.Rule {
	var returns awsevents.Rule
	_jsii_.Get(
		j,
		"rule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RollTrigger) Schedule() awsevents.Schedule {
	var returns awsevents.Schedule
	_jsii_.Get(
		j,
		"schedule",
		&returns,
	)
	return returns
}


func NewRollTrigger_Override(r RollTrigger) {
	_init_.Initialize()

	_jsii_.Create(
		"cloudstructs.RollTrigger",
		nil, // no parameters
		r,
	)
}

// Rule that should trigger a roll.
func RollTrigger_FromRule(rule awsevents.Rule) RollTrigger {
	_init_.Initialize()

	if err := validateRollTrigger_FromRuleParameters(rule); err != nil {
		panic(err)
	}
	var returns RollTrigger

	_jsii_.StaticInvoke(
		"cloudstructs.RollTrigger",
		"fromRule",
		[]interface{}{rule},
		&returns,
	)

	return returns
}

// Schedule that should trigger a roll.
func RollTrigger_FromSchedule(schedule awsevents.Schedule) RollTrigger {
	_init_.Initialize()

	if err := validateRollTrigger_FromScheduleParameters(schedule); err != nil {
		panic(err)
	}
	var returns RollTrigger

	_jsii_.StaticInvoke(
		"cloudstructs.RollTrigger",
		"fromSchedule",
		[]interface{}{schedule},
		&returns,
	)

	return returns
}

