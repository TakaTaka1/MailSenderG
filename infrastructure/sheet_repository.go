package infrastructure

import (
	"os"
	"log"
	"encoding/json"
	"context"
    "google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
	"MailSenderG/Utils"
	"MailSenderG/data/StructData"
)

type SheetRepository struct {
}
// ポインタ型として返す
func NewSheetRepository() *SheetRepository{
	return &SheetRepository{}
}

// TODO
func (sp *SheetRepository) prepSheet() *sheets.Service{
	// SHEET_ID := os.Getenv("SHEET_ID")
	file, _ := json.Marshal(Utils.SetSheetCredentials())
	_ = os.WriteFile("./secret.json", file, 0644)
	credential := option.WithCredentialsFile("./secret.json")
	srv, err := sheets.NewService(context.TODO(), credential)
	if err != nil {
	    log.Fatal(err)
	}
	_ = os.Remove("./secret.json")
	return srv
}

func (sp *SheetRepository) Rethoge() string {
	return "from sheetRepository"
}

func (sp *SheetRepository) RetSheetData(sheetNameRange string) map[int]StructData.SheetData {
	SHEET_ID := os.Getenv("SHEET_ID")
	srv := sp.prepSheet()
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