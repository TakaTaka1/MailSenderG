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

type SheetService interface {
	ReadSheet(sheetNameRange string) map[int]StructData.SheetData
}

type sheetService struct {
}

func NewSheetService() *sheetService {
	return &sheetService{}
}

func (s *sheetService) tmpSetCredentials() credential {
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

func (s *sheetService) TmpReadSheet (sheetNameRange string) map[int]StructData.SheetData{
	SHEET_ID := os.Getenv("SHEET_ID")
	file, _ := json.Marshal(s.tmpSetCredentials())	
	_ = os.WriteFile("./secret.json", file, 0644)
	credential := option.WithCredentialsFile("./secret.json")
	srv, err := sheets.NewService(context.TODO(), credential)
	if err != nil {
	    log.Fatal(err)
	}
	// TODO 別の方法または工夫した方が良さそう
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

func (s *sheetService)SetEachCost(costType string) map[int]StructData.SheetData{
	costRange := "!A2:E13"
	return s.TmpReadSheet(costType + costRange)
}