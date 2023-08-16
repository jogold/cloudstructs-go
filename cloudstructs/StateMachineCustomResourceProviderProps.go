package cloudstructs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for a StateMachineCustomResourceProvider.
type StateMachineCustomResourceProviderProps struct {
	// The state machine.
	StateMachine IStateMachine `field:"required" json:"stateMachine" yaml:"stateMachine"`
	// Timeout.
	// Default: Duration.minutes(30)
	//
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
}

