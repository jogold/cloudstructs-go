//go:build no_runtime_type_checking

package cloudstructs

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SlackAppManifest) validateRenderParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewSlackAppManifestParameters(props *SlackAppManifestProps) error {
	return nil
}

