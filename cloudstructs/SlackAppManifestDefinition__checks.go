//go:build !no_runtime_type_checking

package cloudstructs

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

func (s *jsiiProxy_SlackAppManifestDefinition) validateRenderParameters(construct constructs.IConstruct) error {
	if construct == nil {
		return fmt.Errorf("parameter construct is required, but nil was provided")
	}

	return nil
}

func validateSlackAppManifestDefinition_FromFileParameters(file *string) error {
	if file == nil {
		return fmt.Errorf("parameter file is required, but nil was provided")
	}

	return nil
}

func validateSlackAppManifestDefinition_FromManifestParameters(props *SlackAppManifestProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func validateSlackAppManifestDefinition_FromStringParameters(manifest *string) error {
	if manifest == nil {
		return fmt.Errorf("parameter manifest is required, but nil was provided")
	}

	return nil
}

