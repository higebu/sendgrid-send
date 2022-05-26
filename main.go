package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

var (
	from        string
	fromName    string
	to          string
	toName      string
	subject     string
	content     string
	htmlContent string
)

func init() {
	flag.StringVar(&from, "from", "test@example.com", "from address")
	flag.StringVar(&fromName, "from-name", "Example User", "from name")
	flag.StringVar(&to, "to", "test@example.com", "from address")
	flag.StringVar(&toName, "to-name", "Example User", "from name")
	flag.StringVar(&subject, "subject", "Hello SendGrid", "subject")
	flag.StringVar(&content, "content", "Hello Sendgrid", "text content")
	flag.StringVar(&htmlContent, "html-content", "<strong>Hello Sendgrid</strong>", "html content")
}

func main() {
	flag.Parse()
	from := mail.NewEmail(fromName, from)
	to := mail.NewEmail(toName, to)
	message := mail.NewSingleEmail(from, subject, to, content, htmlContent)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}
