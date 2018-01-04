package awsutil

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/defaults"
	"github.com/aws/aws-sdk-go/aws/session"
)

var (
	// awsDisableSSL can be set to true by build tag.
	awsDisableSSL = "false"

	// awsEndpoint can be set to a custom endpoint by build tag.
	awsEndpoint = ""
)

// Session returns an AWS session as described http://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/configuring-sdk.html
func Session() (*session.Session, error) {
	cfg := aws.Config{
		DisableSSL:       aws.Bool(awsDisableSSL == "true"),
		S3ForcePathStyle: aws.Bool(true),
		Endpoint:         aws.String(awsEndpoint),
	}
	cfg.Credentials = defaults.CredChain(&cfg, defaults.Handlers())
	return session.NewSessionWithOptions(session.Options{
		Config:                  cfg,
		SharedConfigState:       session.SharedConfigEnable,
		AssumeRoleTokenProvider: stscreds.StdinTokenProvider,
	})
}
