package main

import (
	// "github.com/go-sql-driver/mysql"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"
	"MailSenderG/internal/Price"
	"MailSenderG/internal/SpreadSheet"
	_"MailSenderG/internal/Mail"
	_"MailSenderG/data/ConstData"
	"MailSenderG/infrastructure"
	"MailSenderG/data/StructData"
	"MailSenderG/usecase"
	"time"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"github.com/joho/godotenv"
	"strconv"
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
	
	// TODO
	sheetRepo := infrastructure.NewSheetRepository()	
	sheet := Usecase.NewSheetService(sheetRepo)
	vStruct := reflect.Indirect(reflect.ValueOf(SpreadSheet.RetSheetNameStruct()))
	vType := vStruct.Type()

	// TODO 先月の日付取得
	t := time.Now() // 現在時刻を実行環境のタイムゾーンで取得
	lastMonth := t.AddDate(0,-1,0).Format("200601")
	
	// TODO 値オブジェクトにする
	totalTaPrice := 0
	totalMiPrice := 0
	costs := map[string]StructData.SheetData{}
	
	for g :=0; g<vType.NumField(); g++ {
		ft := vType.Field(g)		
		fv := vStruct.FieldByName(ft.Name)
		costMap := sheet.Read(sheet.SetCost(fv.String()))
		costLastMonth := Price.GetLastMonthPrice(costMap, lastMonth, fv.String())
		costs[fv.String()] = costLastMonth[fv.String()]
		totalTaPrice += Price.ReturnPrice(costLastMonth[fv.String()].TPrice)
		totalMiPrice += Price.ReturnPrice(costLastMonth[fv.String()].MPrice)
	}
	
	// TODO
	TOS := strings.Split(os.Getenv("TOS"), ",")
	FROM := os.Getenv("FROM")
	
	// メッセージの構築
	message := mail.NewV3Mail()
	// 送信元を設定
	from := mail.NewEmail("", FROM)
	message.SetFrom(from)

	// 1つ目の宛先と、対応するSubstitutionタグを指定
	p := mail.NewPersonalization()
	to := mail.NewEmail("", TOS[0])
	p.AddTos(to)
	p.SetSubstitution("%fullname%", os.Getenv("SEND_LIST_1"))
	p.SetSubstitution("%familyname%", os.Getenv("SEND_LIST_1"))
	message.AddPersonalizations(p)

	// // 2つ目の宛先と、対応するSubstitutionタグを指定
	p2 := mail.NewPersonalization()
	to2 := mail.NewEmail("", TOS[1])
	p2.AddTos(to2)
	p2.SetSubstitution("%fullname%", os.Getenv("SEND_LIST_2"))
	p2.SetSubstitution("%familyname%", os.Getenv("SEND_LIST_2"))
	message.AddPersonalizations(p2)

	// 件名を設定
	message.Subject = os.Getenv("MAIL_SUBJECT")
	diffPrice := Price.CheckDiffPrice(totalMiPrice, totalTaPrice)

	var mailHeaderHtml = os.Getenv("MAIL_HEADER")
	var mailTaHtml = "<strong>👨‍💻【" + os.Getenv("SEND_LIST_1") + "】👨‍💻</strong><br>" + "食費: " + costs["食費"].TPrice + "<br>" + "日用品: " + costs["日用品"].TPrice + "<br>" + "雑費: " + costs["雑費"].TPrice + "<br>" + "水道費: " + costs["水道費"].TPrice + "<br>" + "光熱費: " + costs["光熱費"].TPrice + "<br>" + "家賃: " + costs["家賃"].TPrice + "<br>" + "【合計】 : " + strconv.Itoa(totalTaPrice) + "<br><br>"
	var mailMiHtml = "<strong>🤷‍♀【" + os.Getenv("SEND_LIST_2") + "】🤷‍♀️</strong><br>" + "食費: " + costs["食費"].MPrice + "<br>" + "日用品: " + costs["日用品"].MPrice + "<br>" + "雑費: " + costs["雑費"].MPrice + "<br>" + "水道費: " + costs["水道費"].MPrice + "<br>" + "光熱費: " + costs["光熱費"].MPrice + "<br>" + "家賃: " + costs["家賃"].MPrice + "<br>" + "【合計】 : " + strconv.Itoa(totalMiPrice) + "<br><br>"	
	var mailDiffHtml = "差分: 💴" + strconv.Itoa(diffPrice) + "<br><br>"
	var mailPokioCommentHtml = os.Getenv("MAIL_PO_COMMENT_HTML")
	c := mail.NewContent("text/html",mailHeaderHtml + mailTaHtml + mailMiHtml + mailDiffHtml + mailPokioCommentHtml)
	message.AddContent(c)

	// メール送信を行い、レスポンスを表示
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))

	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}

	// go mapでのループ処理はランダム出力される。。
	// https://free-engineer.life/golang-map-for-loops/
}