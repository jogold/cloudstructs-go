//go:build no_runtime_type_checking

package cloudstructs

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SlackAppManifestDefinition) validateRenderParameters(construct constructs.IConstruct) error {
	return nil
}

func validateSlackAppManifestDefinition_FromFileParameters(file *string) error {
	return nil
}

func validateSlackAppManifestDefinition_FromManifestParameters(props *SlackAppManifestProps) error {
	return nil
}

func validateSlackAppManifestDefinition_FromStringParameters(manifest *string) error {
	return nil
}

