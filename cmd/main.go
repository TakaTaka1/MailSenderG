package main

import (	
	"fmt"
	"log"
	"os"
	"reflect"
	"MailSenderG/internal/Price"
	"MailSenderG/internal/SpreadSheet"
	"MailSenderG/infrastructure/mail"
	"MailSenderG/infrastructure/sheet"
	"MailSenderG/data/StructData"
	"MailSenderG/usecase/sheet"
	"MailSenderG/domain/model/mail"
	"time"
	"github.com/joho/godotenv"
)


func main() {
	// if file exists	
	f := "./.env"
	if _, err := os.Stat(f); err == nil {
		err_read := godotenv.Load(f)
		if err_read != nil {
			log.Fatalf("error: %v", err_read)
		}					
		fmt.Println(".env is existed")
    // 存在します
	} else {
		fmt.Println(".env is not existed")
	}		
	
	// TODO 先月の日付取得
	t := time.Now() // 現在時刻を実行環境のタイムゾーンで取得
	lastMonth := t.AddDate(0,-1,0).Format("200601")
	
	// TODO 値オブジェクトにする
	totalTaPrice := 0
	totalMiPrice := 0
	costs := map[string]StructData.SheetData{}
	
	// TODO
	sheetRepo := sheetInfra.NewSheetRepository()	
	sheet := sheetService.NewSheetService(sheetRepo)	
	vStruct := reflect.Indirect(reflect.ValueOf(spreadSheet.RetSheetNameStruct()))
	vType := vStruct.Type() // Typeインターフェース
	vFieldNum := vType.NumField()

	// reflectでの処理 , NumField()でフィールド数の取得
	for g :=0; g<vFieldNum; g++ {
		// Field(index)でft.Nameで構造体のフィールド名を取得できる
		ft := vType.Field(g)
		fv := vStruct.FieldByName(ft.Name) // フィールドの値を取得できる
		costMap := sheet.Read(sheet.SetCost(fv.String()))
		costLastMonth := Price.GetLastMonthPrice(costMap, lastMonth, fv.String())
		costs[fv.String()] = costLastMonth[fv.String()]
		totalTaPrice += Price.ReturnPrice(costLastMonth[fv.String()].TPrice)
		totalMiPrice += Price.ReturnPrice(costLastMonth[fv.String()].MPrice)
	}
	
	diffPrice := Price.CheckDiffPrice(totalMiPrice, totalTaPrice)
	// domain mailで定数取得
	mailInfo := mailModel.CreateMailInfo()
	// mail repositoryの各メソッドにdomain mailの定数を渡す
	mailRepo := mailInfra.NewMailRepository()
	sgContents := mailRepo.SetupSendGridMail()
	mailRepo.SetupMailFrom(sgContents, mailInfo.From)
	mailRepo.SetupMailTo(sgContents, mailInfo.To)
	mailRepo.SetupMailSubject(sgContents,mailInfo.Subject)
	header := mailRepo.SetupMailHeader(mailInfo.Header)
	readySgContents := mailRepo.SetupMailBody(
		sgContents, 
		header, 
		diffPrice, 
		costs,
		totalTaPrice,
		totalMiPrice,
	)
	mailModel.SendMail(readySgContents)

	// go mapでのループ処理はランダム出力される。。
	// https://free-engineer.life/golang-map-for-loops/
}