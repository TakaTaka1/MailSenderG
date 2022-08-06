package sheetInfra

import (
	"MailSenderG/data/StructData"
	Utils "MailSenderG/utils"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	_ "golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

type SheetRepository struct {
}

// ポインタ型として返す
func NewSheetRepository() *SheetRepository {
	return &SheetRepository{}
}

// TODO
func (sp *SheetRepository) prepCredentials() (*sheets.Service, error) {
	file, _ := json.Marshal(Utils.SetSheetCredentials())

	err := os.WriteFile("./cd/secret.json", file, 0644)
	if err != nil {
		return nil, err
	}

	// TODO credentialの読み込みをWithCredentialsで実行する方法
	// ctx := context.Background()
	// credential, _ := google.CredentialsFromJSON(ctx, file)
	// optcredential := option.WithCredentials(credential)
	credential := option.WithCredentialsFile("./cd/secret.json")
	srv, err := sheets.NewService(context.TODO(), credential)
	if err != nil {
		errRemove := os.Remove("./cd/secret.json")
		if errRemove != nil {
			return nil, errRemove
		}
		return nil, err
	}

	errRemove := os.Remove("./cd/secret.json")
	if errRemove != nil {
		fmt.Println("removed!!")
		return nil, errRemove
	}
	return srv, nil
}

func (sp *SheetRepository) RetSheetData(sheetNameRange string) (map[int]StructData.SheetData, error) {
	SHEET_ID := os.Getenv("SHEET_ID")
	srv, err := sp.prepCredentials()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	resp, err := srv.Spreadsheets.Values.Get(SHEET_ID, sheetNameRange).Do()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if len(resp.Values) == 0 {
		log.Println("data not found")
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
	return dataMap, nil
}
