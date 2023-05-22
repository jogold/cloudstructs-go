//go:build no_runtime_type_checking

package cloudstructs

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SamlFederatedPrincipal) validateAddToPolicyParameters(statement awsiam.PolicyStatement) error {
	return nil
}

func (s *jsiiProxy_SamlFederatedPrincipal) validateAddToPrincipalPolicyParameters(_statement awsiam.PolicyStatement) error {
	return nil
}

func (s *jsiiProxy_SamlFederatedPrincipal) validateWithConditionsParameters(conditions *map[string]interface{}) error {
	return nil
}

func validateNewSamlFederatedPrincipalParameters(identityProvider SamlIdentityProvider) error {
	return nil
}

