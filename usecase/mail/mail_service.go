package mailService

import (
	mailInfra "MailSenderG/infrastructure/mail"
)

type mailService struct {
	MailRepo *mailInfra.MailRepository
}

func NewMailService(repos *mailInfra.MailRepository) *mailService {
	return &mailService{
		MailRepo: repos,
	}
}
