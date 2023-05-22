package cloudstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/jogold/cloudstructs-go/cloudstructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
)

// A source repository for AWS CodeCommit mirroring.
type CodeCommitMirrorSourceRepository interface {
	// The name of the repository.
	Name() *string
	// The HTTPS clone URL in plain text, used for a public repository.
	PlainTextUrl() *string
	// The HTTPS clone URL if the repository is private.
	//
	// The secret should contain the username and/or token.
	//
	// Example:
	//   `https://TOKEN@github.com/owner/name`
	//   `https://USERNAME:TOKEN@bitbucket.org/owner/name.git`
	//
	SecretUrl() awsecs.Secret
}

// The jsii proxy struct for CodeCommitMirrorSourceRepository
type jsiiProxy_CodeCommitMirrorSourceRepository struct {
	_ byte // padding
}

func (j *jsiiProxy_CodeCommitMirrorSourceRepository) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodeCommitMirrorSourceRepository) PlainTextUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"plainTextUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodeCommitMirrorSourceRepository) SecretUrl() awsecs.Secret {
	var returns awsecs.Secret
	_jsii_.Get(
		j,
		"secretUrl",
		&returns,
	)
	return returns
}


func NewCodeCommitMirrorSourceRepository_Override(c CodeCommitMirrorSourceRepository) {
	_init_.Initialize()

	_jsii_.Create(
		"cloudstructs.CodeCommitMirrorSourceRepository",
		nil, // no parameters
		c,
	)
}

// Public GitHub repository.
func CodeCommitMirrorSourceRepository_GitHub(owner *string, name *string) CodeCommitMirrorSourceRepository {
	_init_.Initialize()

	if err := validateCodeCommitMirrorSourceRepository_GitHubParameters(owner, name); err != nil {
		panic(err)
	}
	var returns CodeCommitMirrorSourceRepository

	_jsii_.StaticInvoke(
		"cloudstructs.CodeCommitMirrorSourceRepository",
		"gitHub",
		[]interface{}{owner, name},
		&returns,
	)

	return returns
}

// Private repository with HTTPS clone URL stored in a AWS Secrets Manager secret or a AWS Systems Manager secure string parameter.
func CodeCommitMirrorSourceRepository_Private(name *string, url awsecs.Secret) CodeCommitMirrorSourceRepository {
	_init_.Initialize()

	if err := validateCodeCommitMirrorSourceRepository_PrivateParameters(name, url); err != nil {
		panic(err)
	}
	var returns CodeCommitMirrorSourceRepository

	_jsii_.StaticInvoke(
		"cloudstructs.CodeCommitMirrorSourceRepository",
		"private",
		[]interface{}{name, url},
		&returns,
	)

	return returns
}

