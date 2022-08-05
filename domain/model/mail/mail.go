package mailModel

import (
	"fmt"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"log"
	"os"
	"strings"
)

type MailInfo struct {
	To      []string
	From    string
	Subject string
	Header  string
}

func CreateMailInfo() *MailInfo {
	return &MailInfo{
		To:      strings.Split(os.Getenv("TOS"), ","),
		From:    os.Getenv("FROM"),
		Subject: os.Getenv("MAIL_SUBJECT"),
		Header:  os.Getenv("MAIL_HEADER"),
	}
}

func SendMail(sgContents *mail.SGMailV3) {
	// メール送信を行い、レスポンスを表示
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(sgContents)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Succeed!!")
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}

}
