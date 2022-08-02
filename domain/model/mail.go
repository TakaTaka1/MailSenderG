package model

import (
	"os"
	"strings"
)

type MailInfo struct {
	To	[]string
	From  string
	Subject string
	Header string
}

func RetMailInfoStruct () *MailInfo {
	return &MailInfo {
		To : strings.Split(os.Getenv("TOS"), ","),
		From : os.Getenv("FROM"),
		Subject :os.Getenv("MAIL_SUBJECT"),
		Header :os.Getenv("MAIL_HEADER"),
	}
}



