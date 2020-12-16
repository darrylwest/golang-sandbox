package main

import (
	"fmt"

	"crypto/rand"
	"encoding/hex"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

func generateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)

	if err != nil {
		return nil, err
	}

	return b, nil
}

func generateRandomHex(n int) string {
	b, _ := generateRandomBytes(n)

	return hex.EncodeToString(b)
}

type PasswordReset struct {
	To   string
	From string
	Key  string
}

func sendResetKey(svc *ses.SES, reset *PasswordReset) error {

	params := &ses.SendEmailInput{
		Destination: &ses.Destination{ // Required
			ToAddresses: []*string{
				aws.String(reset.To),
			},
		},
		Message: &ses.Message{ // Required
			Body: &ses.Body{ // Required
				Text: &ses.Content{
					Data:    aws.String(reset.Key), // Required
					Charset: aws.String("utf-8"),
				},
			},
			Subject: &ses.Content{ // Required
				Data:    aws.String(reset.Key),
				Charset: aws.String("utf-8"),
			},
		},
		Source: aws.String(reset.From),
		ReplyToAddresses: []*string{
			aws.String(reset.From),
			// More values...
		},
	}

	resp, err := svc.SendEmail(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return err
	} else {
		// Pretty-print the response data.
		fmt.Println(resp)

		return nil
	}
}

func createCredentials() *credentials.Credentials {
	// these are AMA creds
	id := "AKIAJTNJ346N3XWJLUEQ"
	secret := "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"

	creds := credentials.NewStaticCredentials(id, secret, "")
	// value, _ := creds.Get()
	// fmt.Println( value )

	return creds
}

func main() {
	sess := session.New()
	creds := createCredentials()
	from := "info@roundpoint.com"
	to := "darryl.west@raincitysoftware.com"

	config := aws.NewConfig().WithRegion("us-east-1").WithCredentials(creds)

	svc := ses.New(sess, config)
	// fmt.Println( svc )

	key := generateRandomHex(5)

	// fmt.Println( key )
	reset := &PasswordReset{
		To:   to,
		From: from,
		Key:  key,
	}

	sendResetKey(svc, reset)
}
