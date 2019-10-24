package services

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

const S3_REGION = "eu-west-1"

func GetConfig() *session.Session {
	sess, err := session.NewSession(&aws.Config{Region: aws.String(S3_REGION)})
	if err != nil {
		panic("unable to load SDK config, " + err.Error())
	}
	return sess
}
