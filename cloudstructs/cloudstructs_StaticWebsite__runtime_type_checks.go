//go:build !no_runtime_type_checking

// High-level constructs for AWS CDK
package cloudstructs

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfront"
	"github.com/aws/constructs-go/constructs/v10"
)

func validateStaticWebsite_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateStaticWebsite_SetDefaultSecurityHeadersBehaviorParameters(val *awscloudfront.ResponseSecurityHeadersBehavior) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
		return err
	}

	return nil
}

func validateNewStaticWebsiteParameters(scope constructs.Construct, id *string, props *StaticWebsiteProps) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}
