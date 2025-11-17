//go:build !no_runtime_type_checking

package cloudstructs

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

func (s *jsiiProxy_SamlFederatedPrincipal) validateAddToAssumeRolePolicyParameters(document awsiam.PolicyDocument) error {
	if document == nil {
		return fmt.Errorf("parameter document is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_SamlFederatedPrincipal) validateAddToPolicyParameters(statement awsiam.PolicyStatement) error {
	if statement == nil {
		return fmt.Errorf("parameter statement is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_SamlFederatedPrincipal) validateAddToPrincipalPolicyParameters(statement awsiam.PolicyStatement) error {
	if statement == nil {
		return fmt.Errorf("parameter statement is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_SamlFederatedPrincipal) validateWithConditionsParameters(conditions *map[string]interface{}) error {
	if conditions == nil {
		return fmt.Errorf("parameter conditions is required, but nil was provided")
	}

	return nil
}

func validateNewSamlFederatedPrincipalParameters(identityProvider SamlIdentityProvider) error {
	if identityProvider == nil {
		return fmt.Errorf("parameter identityProvider is required, but nil was provided")
	}

	return nil
}

