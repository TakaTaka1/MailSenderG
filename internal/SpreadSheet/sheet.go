package SpreadSheet

import (
	"os"
	"log"
	_"fmt"
	"context"
	"encoding/json"
    "google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"	
	"MailSenderG/data/StructData"
)

type Credential struct {
	Type   string    `json:"type"`
	Project_id string `json:"project_id"`
	Private_key_id string `json:"private_key_id"`
	Private_key string `json:"private_key"`
	Client_email string `json:"client_email"`
	Client_id string `json:"client_id"`
	Auth_uri string `json:"auth_uri"`
	Token_uri string `json:"token_uri"`
	Auth_provider_x509_cert_url string `json:"auth_provider_x509_cert_url"`
	Client_x509_cert_url string `json:"client_x509_cert_url"`
}

func RetSheet() string {
	return "Sheet"
}

func setCredentials() Credential {
	sheet_credentials := Credential{
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

func ReadSheet(sheetNameRange string) map[int]StructData.SheetData {
	SHEET_ID := os.Getenv("SHEET_ID")
	
	file, _ := json.Marshal(setCredentials())	
	_ = os.WriteFile("./secret.json", file, 0644)	
	
	credential := option.WithCredentialsFile("./secret.json")
    srv, err := sheets.NewService(context.TODO(), credential)
    if err != nil {
        log.Fatal(err)
	}

	_ = os.Remove("./secret.json")

	resp, err := srv.Spreadsheets.Values.Get(SHEET_ID,sheetNameRange).Do()
	if err != nil {
        log.Fatalln(err)
    }
    if len(resp.Values) == 0 {
        log.Fatalln("data not found")
	}

	var dataMap = make(map[int]StructData.SheetData)
	for i, row := range resp.Values {
		dataMap[i] = StructData.SheetData{
						row[0].(string),
						row[1].(string),
						row[2].(string),
						row[3].(string),
						row[4].(string)}
	}
	return dataMap
}