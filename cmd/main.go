package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"MailSenderG/internal/Price" // go install MailSenderG/internal/Price
	"MailSenderG/internal/SpreadSheet"	
	"time"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	// "github.com/joho/godotenv"
	"strconv"
)


func main() {

	// err_read := godotenv.Load("../.env")
	// if err_read != nil {
	// 	log.Fatalf("error: %v", err_read)
	// }
		
	TOS := strings.Split(os.Getenv("TOS"), ",")
	FROM := os.Getenv("FROM")   

	// 用に支払い金額分を取得して、シート名と金額をメールで送信する
	foodCostMap := SpreadSheet.ReadSheet("食費!A2:E13")
	houseHoldItemMap := SpreadSheet.ReadSheet("日用品!A2:E13")
	generalCostMap := SpreadSheet.ReadSheet("雑費!A2:E13")
	waterCostMap := SpreadSheet.ReadSheet("水道費!A2:E13")
	gasCostMap := SpreadSheet.ReadSheet("光熱費!A2:E13")
	rentCostMap := SpreadSheet.ReadSheet("家賃!A2:E13")
    
	// 先月の日付取得
	t := time.Now() // 現在時刻を実行環境のタイムゾーンで取得
	lastMonth := t.AddDate(0,-1,0).Format("200601")
	
	// 先月の分のレコードを取得する
	foodCostLastMonth := Price.GetLastMonthPrice(foodCostMap, lastMonth, "食費")	
	houseHoldItemLastMonth := Price.GetLastMonthPrice(houseHoldItemMap, lastMonth, "日用品")
	generalCostLastMonth := Price.GetLastMonthPrice(generalCostMap, lastMonth, "雑費")
	waterCostLastMonth := Price.GetLastMonthPrice(waterCostMap, lastMonth, "水道費")
	gasCostLastMonth := Price.GetLastMonthPrice(gasCostMap, lastMonth, "光熱費")
	rentCostLastMonth := Price.GetLastMonthPrice(rentCostMap, lastMonth, "家賃")
	
	totalTaPrice := 0
	totalMiPrice := 0
	
	totalTaPrice += Price.ReturnPrice(foodCostLastMonth["食費"].TPrice)
	totalTaPrice += Price.ReturnPrice(houseHoldItemLastMonth["日用品"].TPrice)
	totalTaPrice += Price.ReturnPrice(generalCostLastMonth["雑費"].TPrice)
	totalTaPrice += Price.ReturnPrice(waterCostLastMonth["水道費"].TPrice)
	totalTaPrice += Price.ReturnPrice(gasCostLastMonth["光熱費"].TPrice)
	totalTaPrice += Price.ReturnPrice(rentCostLastMonth["家賃"].TPrice)
	
	totalMiPrice += Price.ReturnPrice(foodCostLastMonth["食費"].MPrice)
	totalMiPrice += Price.ReturnPrice(houseHoldItemLastMonth["日用品"].MPrice)
	totalMiPrice += Price.ReturnPrice(generalCostLastMonth["雑費"].MPrice)
	totalMiPrice += Price.ReturnPrice(waterCostLastMonth["水道費"].MPrice)
	totalMiPrice += Price.ReturnPrice(gasCostLastMonth["光熱費"].MPrice)
	totalMiPrice += Price.ReturnPrice(rentCostLastMonth["家賃"].MPrice)

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
	var diffPrice int
	if totalMiPrice >= totalTaPrice {
		diffPrice = totalMiPrice - totalTaPrice
	} else if totalMiPrice <= totalTaPrice {
		diffPrice = totalTaPrice - totalMiPrice
	}

	var mailHeaderHtml = os.Getenv("MAIL_HEADER")
	var mailTaHtml = "<strong>👨‍💻【" + os.Getenv("SEND_LIST_1") + "】👨‍💻</strong><br>" + "食費: " + foodCostLastMonth["食費"].TPrice + "<br>" + "日用品: " + houseHoldItemLastMonth["日用品"].TPrice + "<br>" + "雑費: " + generalCostLastMonth["雑費"].TPrice + "<br>" + "水道費: " + waterCostLastMonth["水道費"].TPrice + "<br>" + "光熱費: " + gasCostLastMonth["光熱費"].TPrice + "<br>" + "家賃: " + rentCostLastMonth["家賃"].TPrice + "<br>" + "【合計】 : " + strconv.Itoa(totalTaPrice) + "<br><br>"
	var mailMiHtml = "<strong>🤷‍♀【" + os.Getenv("SEND_LIST_2") + "】🤷‍♀️</strong><br>" + "食費: " + foodCostLastMonth["食費"].MPrice + "<br>" + "日用品: " + houseHoldItemLastMonth["日用品"].MPrice + "<br>" + "雑費: " + generalCostLastMonth["雑費"].MPrice + "<br>" + "水道費: " + waterCostLastMonth["水道費"].MPrice + "<br>" + "光熱費: " + gasCostLastMonth["光熱費"].MPrice + "<br>" + "家賃: " + rentCostLastMonth["家賃"].MPrice + "<br>" + "【合計】 : " + strconv.Itoa(totalMiPrice) + "<br><br>"
	var mailDiffHtml = "差分: 💴" + strconv.Itoa(diffPrice) + "<br><br>"
	var mailPokioCommentHtml = os.Getenv("MAIL_PO_COMMENT_HTML")
	c := mail.NewContent("text/html",mailHeaderHtml + mailTaHtml + mailMiHtml + mailDiffHtml + mailPokioCommentHtml)
	message.AddContent(c)

	// カテゴリ情報を付加
	// message.AddCategories("category1")
	// カスタムヘッダを指定
	// message.SetHeader("X-Sent-Using", "SendGrid-API")
	// 画像ファイルを添付
	// a := mail.NewAttachment()
	// file, _ := os.OpenFile("./gif.gif", os.O_RDONLY, 0600)
	// defer file.Close()
	// data, _ := ioutil.ReadAll(file)
	// data_enc := base64.StdEncoding.EncodeToString(data)
	// a.SetContent(data_enc)
	// a.SetType("image/gif")
	// a.SetFilename("owl.gif")
	// a.SetDisposition("attachment")
	// message.AddAttachment(a)

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