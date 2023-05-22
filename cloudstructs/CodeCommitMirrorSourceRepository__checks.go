//go:build !no_runtime_type_checking

package cloudstructs

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
)

func validateCodeCommitMirrorSourceRepository_GitHubParameters(owner *string, name *string) error {
	if owner == nil {
		return fmt.Errorf("parameter owner is required, but nil was provided")
	}

	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	return nil
}

func validateCodeCommitMirrorSourceRepository_PrivateParameters(name *string, url awsecs.Secret) error {
	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	if url == nil {
		return fmt.Errorf("parameter url is required, but nil was provided")
	}

	return nil
}

