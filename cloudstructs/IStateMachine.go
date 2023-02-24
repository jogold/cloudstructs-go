// High-level constructs for AWS CDK
package cloudstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A State Machine.
type IStateMachine interface {
	// The ARN of the state machine.
	StateMachineArn() *string
}

// The jsii proxy for IStateMachine
type jsiiProxy_IStateMachine struct {
	_ byte // padding
}

func (j *jsiiProxy_IStateMachine) StateMachineArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateMachineArn",
		&returns,
	)
	return returns
}

