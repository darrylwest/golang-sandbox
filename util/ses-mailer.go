package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

func GenerateRandomBytes() ([]byte, error) {
	b := make([]byte, 8)
	_, err := rand.Read(b)

	if err != nil {
		return nil, err
	}

	return b, nil
}

// return in ####-####-####-####
func GenerateRandomHex() (string, error) {
	b, err := GenerateRandomBytes()

	if err != nil {
		return "", err
	}

	s := hex.EncodeToString(b)

	return fmt.Sprintf("%s-%s-%s-%s", s[0:4], s[4:8], s[8:12], s[12:16]), nil
}

func listVerifiedEmailAddresses(svc *ses.SES) {

	var params *ses.ListVerifiedEmailAddressesInput
	resp, err := svc.ListVerifiedEmailAddresses(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func listIdentities(svc *ses.SES) {

	params := &ses.ListIdentitiesInput{
	// MaxItems:     aws.Int64(100),
	}

	resp, err := svc.ListIdentities(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func sendStatistics(svc *ses.SES) {

	var params *ses.GetSendStatisticsInput
	resp, err := svc.GetSendStatistics(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func showQuota(svc *ses.SES) {
	var params *ses.GetSendQuotaInput
	resp, err := svc.GetSendQuota(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
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

func sendEmail(svc *ses.SES) {
	to := "7752508168@messaging.sprintpcs.com"
	// to := "darryl.west@raincitysoftware.com"
	code, _ := GenerateRandomHex()

	params := &ses.SendEmailInput{
		Destination: &ses.Destination{ // Required
			ToAddresses: []*string{
				aws.String(to), //
				// More values...
			},
		},
		Message: &ses.Message{ // Required
			Body: &ses.Body{ // Required
				Text: &ses.Content{
					Data:    aws.String(code), // Required
					Charset: aws.String("utf-8"),
				},
			},
			Subject: &ses.Content{ // Required
				Data:    aws.String("code: "), // Required
				Charset: aws.String("utf-8"),
			},
		},
		Source: aws.String("dpw@raincitysoftware.com"), // Required
		ReplyToAddresses: []*string{
			aws.String("dpw@raincitysoftware.com"), // Required
			// More values...
		},
		// ReturnPath:    aws.String("Address"),
		// ReturnPathArn: aws.String("AmazonResourceName"),
		// SourceArn:     aws.String("AmazonResourceName"),
	}

	resp, err := svc.SendEmail(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

// not very useful
func verifyEmailAddress(svc *ses.SES) {
	params := &ses.VerifyEmailIdentityInput{
		EmailAddress: aws.String("7752508168@messaging.sprintpcs.com"), // Required
	}

	resp, err := svc.VerifyEmailIdentity(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func main() {
	sess := session.New()
	creds := createCredentials()

	// config := aws.NewConfig().WithRegion("us-east-1").WithCredentials(creds)
	config := aws.NewConfig().WithRegion("us-west-2").WithCredentials(creds)

	svc := ses.New(sess, config)

	// showQuota(svc)
	// sendStatistics( svc )
	listIdentities(svc)
	listVerifiedEmailAddresses(svc)
	sendEmail(svc)
	// verifyEmailAddress( svc )
}
