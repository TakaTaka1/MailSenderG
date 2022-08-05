package Utils

import (
	"os"
)

type credential struct {
	Type                        string `json:"type"`
	Project_id                  string `json:"project_id"`
	Private_key_id              string `json:"private_key_id"`
	Private_key                 string `json:"private_key"`
	Client_email                string `json:"client_email"`
	Client_id                   string `json:"client_id"`
	Auth_uri                    string `json:"auth_uri"`
	Token_uri                   string `json:"token_uri"`
	Auth_provider_x509_cert_url string `json:"auth_provider_x509_cert_url"`
	Client_x509_cert_url        string `json:"client_x509_cert_url"`
}

func SetSheetCredentials() credential {
	sheet_credentials := credential{
		os.Getenv("TYPE"),
		os.Getenv("PROJECT_ID"),
		os.Getenv("PRIVATE_KEY_ID"),
		os.Getenv("PRIVATE_KEY"),
		os.Getenv("CLIENT_EMAIL"),
		os.Getenv("CLIENT_ID"),
		os.Getenv("AUTH_URI"),
		os.Getenv("TOKEN_URI"),
		os.Getenv("AUTH_PROVIDER_CERT_URL"),
		os.Getenv("CLIENT_CERT_URL"),
	}
	return sheet_credentials
}
