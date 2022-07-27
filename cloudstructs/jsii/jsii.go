// Package jsii contains the functionaility needed for jsii packages to
// initialize their dependencies and themselves. Users should never need to use this package
// directly. If you find you need to - please report a bug at
// https://github.com/aws/jsii/issues/new/choose
package jsii

import (
	_                                   "embed"

	_jsii_                              "github.com/aws/jsii-runtime-go/runtime"

	awscdk                              "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	awscdkapigatewayv2alpha             "github.com/aws/aws-cdk-go/awscdkapigatewayv2alpha/v2/jsii"
	awscdkapigatewayv2integrationsalpha "github.com/aws/aws-cdk-go/awscdkapigatewayv2integrationsalpha/v2/jsii"
	constructs                          "github.com/aws/constructs-go/constructs/v10/jsii"
)

//go:embed cloudstructs-0.5.9.tgz
var tarball []byte

// Initialize loads the necessary packages in the @jsii/kernel to support the enclosing module.
// The implementation is idempotent (and hence safe to be called over and over).
func Initialize() {
	// Ensure all dependencies are initialized
	awscdkapigatewayv2alpha.Initialize()
	awscdkapigatewayv2integrationsalpha.Initialize()
	awscdk.Initialize()
	constructs.Initialize()

	// Load this library into the kernel
	_jsii_.Load("cloudstructs", "0.5.9", tarball)
}
