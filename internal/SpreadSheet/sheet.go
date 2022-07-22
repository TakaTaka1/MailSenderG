package SpreadSheet

import (
	"os"
	"log"
	// "fmt"
	"context"
	"encoding/json"
	// "github.com/joho/godotenv"
    "google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
	"MailSenderG/internal/StructData"
	// "encoding/base64"
	"io/ioutil"	
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

func ReadSheet(sheetNameRange string) map[int]StructData.SheetData {
	
	// err_read := godotenv.Load("../.env")
	// if err_read != nil {
	// 	log.Fatalf("error: %v", err_read)
	// }

	SHEET_ID := os.Getenv("SHEET_ID")

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
	
	// file, _ := json.MarshalIndent(sheet_credentials, "", "")
	file, _ := json.Marshal(sheet_credentials)	
	_ = ioutil.WriteFile("./secret.json", file, 0644)	

	// byteArray, _ := ioutil.ReadFile("")
	// var jsonObj interface{}
	// test := json.Unmarshal(byteArray, &jsonObj)

	
	// fmt.Print(byteArray)
	
	// jsonエンコード
	// outputJson, _ := json.Marshal("")
	// fmt.Print(outputJson)
	// // outputJson, err := json.Marshal(os.Getenv("SEACRET"))
	// if err != nil {
	// 	panic(err)
	// }
	
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