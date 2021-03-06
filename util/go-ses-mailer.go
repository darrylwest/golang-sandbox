package main

import (
	"fmt"
	"github.com/sourcegraph/go-ses"
)

func main() {

	// Change the From address to a sender address that is verified in your Amazon SES account.
	from := "dpw@raincitysoftware.com"
	to := "darryl.west@raincitysoftware.com"

	// ses.EnvConfig uses the AWS credentials in the environment variables $AWS_ACCESS_KEY_ID and $AWS_SECRET_KEY.
	res, err := ses.EnvConfig.SendEmail(from, to, "Hello, world from go-ses!", "Here is the message body.")

	if err == nil {
		fmt.Printf("Sent email: %s...\n", res[:32])
	} else {
		fmt.Printf("Error sending email: %s\n", err)
	}

	// output:
	// Sent email: <SendEmailResponse xmlns="http:/...
}
