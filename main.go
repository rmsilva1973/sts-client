package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {
	// cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("sa-east-1"))
	// if err != nil {
	// 	log.Fatalf("unable to load SDK config, %v", err)
	// }

	stsSession := session.Must(session.NewSession())

	creds := stscreds.NewCredentials(stsSession, "myRoleArn")

	svc := s3.New(stsSession, &aws.Config{Credentials: creds})
}
