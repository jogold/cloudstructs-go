//go:build no_runtime_type_checking

package cloudstructs

// Building without runtime type checking enabled, so all the below just return nil

func validateCodeCommitMirrorSourceRepository_GitHubParameters(owner *string, name *string) error {
	return nil
}

func validateCodeCommitMirrorSourceRepository_PrivateParameters(name *string, url awsecs.Secret) error {
	return nil
}

