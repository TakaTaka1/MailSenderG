package SpreadSheet

import (
	"os"
	"log"
	"context"
	"github.com/joho/godotenv"
    "google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
	"MailSenderG/internal/StructData"
)

func ReadSheet(sheetNameRange string) map[int]StructData.SheetData {
	err_read := godotenv.Load("../.env")
	if err_read != nil {
		log.Fatalf("error: %v", err_read)
	}

	SHEET_ID := os.Getenv("SHEET_ID")
    credential := option.WithCredentialsFile("../cd/secret.json")    
    srv, err := sheets.NewService(context.TODO(), credential)
    if err != nil {
        log.Fatal(err)
	}

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