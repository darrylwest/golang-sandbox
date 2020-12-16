package main

import (
	"crypto/rand"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func GenerateRandomBytes() ([]byte, error) {
	b := make([]byte, 8)
	_, err := rand.Read(b)

	if err != nil {
		return nil, err
	}

	return b, nil
}

func createCredentials() *credentials.Credentials {
	// these are AMA creds
	// id := "AKIAJTNJ346N3XWJLUEQ"
	// secret := "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"
	id := "AKIAIJYB2DZZT46NJUUA"
	secret := "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"
	// id := "AKIAJ2JXPQQ5ULI4KXGA"
	// secret := "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"

	creds := credentials.NewStaticCredentials(id, secret, "")
	// value, _ := creds.Get()
	// fmt.Println( value )

	return creds
}

func main() {
	fmt.Println("read the creds, start the session...")

	sess := session.New()
	creds := createCredentials()

	config := aws.NewConfig().WithRegion("us-east-1").WithCredentials(creds)
	// config := aws.NewConfig().WithRegion("us-west-2").WithCredentials(creds)

	svc := s3.New(sess, config)

	var params *s3.ListBucketsInput
	fmt.Println("list the buckets...")
	resp, err := svc.ListBuckets(params)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(resp)
}
