package mailInfra

import (
	// "fmt"
	"os"
	// "MailSenderG/domain/model"
	"strconv"
	"MailSenderG/data/StructData"
	"github.com/sendgrid/sendgrid-go/helpers/mail"	
)

type MailRepository struct {
}

func NewMailRepository() *MailRepository {
	return &MailRepository{}
}

func (mp *MailRepository) SetupSendGridMail () *mail.SGMailV3{
	return mail.NewV3Mail()
}

func (mp *MailRepository) SetupMailFrom (sg *mail.SGMailV3, from string) {
	sgfrom := mail.NewEmail("", from)
	sg.SetFrom(sgfrom)
}

func (mp *MailRepository) SetupMailTo (sg *mail.SGMailV3, to []string) {
	p := mail.NewPersonalization()
	to1 := mail.NewEmail("", to[0])
	p.AddTos(to1)
	p.SetSubstitution("%fullname%", os.Getenv("SEND_LIST_1"))
	p.SetSubstitution("%familyname%", os.Getenv("SEND_LIST_1"))
	sg.AddPersonalizations(p)

	// 2つ目の宛先と、対応するSubstitutionタグを指定
	p2 := mail.NewPersonalization()
	to2 := mail.NewEmail("", to[1])
	p2.AddTos(to2)
	p2.SetSubstitution("%fullname%", os.Getenv("SEND_LIST_2"))
	p2.SetSubstitution("%familyname%", os.Getenv("SEND_LIST_2"))
	sg.AddPersonalizations(p2)
}

func (mp *MailRepository) SetupMailSubject (sg *mail.SGMailV3, subject string) {
	sg.Subject = subject
}

func (mp *MailRepository) SetupMailHeader (header string) string {
	return header
}

func (mp *MailRepository) SetupMailBody (sg *mail.SGMailV3, header string , diffPrice int, costs map[string]StructData.SheetData, totalTaPrice int, totalMiPrice int) *mail.SGMailV3{
	var mailTaHtml = "<strong>👨‍💻【" + os.Getenv("SEND_LIST_1") + "】👨‍💻</strong><br>" + "食費: " + costs["食費"].TPrice + "<br>" + "日用品: " + costs["日用品"].TPrice + "<br>" + "雑費: " + costs["雑費"].TPrice + "<br>" + "水道費: " + costs["水道費"].TPrice + "<br>" + "光熱費: " + costs["光熱費"].TPrice + "<br>" + "家賃: " + costs["家賃"].TPrice + "<br>" + "【合計】 : " + strconv.Itoa(totalTaPrice) + "<br><br>"
	var mailMiHtml = "<strong>🤷‍♀【" + os.Getenv("SEND_LIST_2") + "】🤷‍♀️</strong><br>" + "食費: " + costs["食費"].MPrice + "<br>" + "日用品: " + costs["日用品"].MPrice + "<br>" + "雑費: " + costs["雑費"].MPrice + "<br>" + "水道費: " + costs["水道費"].MPrice + "<br>" + "光熱費: " + costs["光熱費"].MPrice + "<br>" + "家賃: " + costs["家賃"].MPrice + "<br>" + "【合計】 : " + strconv.Itoa(totalMiPrice) + "<br><br>"	
	var mailDiffHtml = "差分: 💴" + strconv.Itoa(diffPrice) + "<br><br>"
	var mailPokioCommentHtml = os.Getenv("MAIL_PO_COMMENT_HTML")
	c := mail.NewContent("text/html",header + mailTaHtml + mailMiHtml + mailDiffHtml + mailPokioCommentHtml)
	sg.AddContent(c)

	return sg
}




// 	// // メッセージの構築
// 	// message := mail.NewV3Mail()
// 	// // 送信元を設定
// 	// from := mail.NewEmail("", FROM)
// 	// message.SetFrom(from)

// 	// // 1つ目の宛先と、対応するSubstitutionタグを指定
// 	// p := mail.NewPersonalization()
// 	// to := mail.NewEmail("", TOS[0])
// 	// p.AddTos(to)
// 	// p.SetSubstitution("%fullname%", os.Getenv("SEND_LIST_1"))
// 	// p.SetSubstitution("%familyname%", os.Getenv("SEND_LIST_1"))
// 	// message.AddPersonalizations(p)

// 	// // // 2つ目の宛先と、対応するSubstitutionタグを指定
// 	// p2 := mail.NewPersonalization()
// 	// to2 := mail.NewEmail("", TOS[1])
// 	// p2.AddTos(to2)
// 	// p2.SetSubstitution("%fullname%", os.Getenv("SEND_LIST_2"))
// 	// p2.SetSubstitution("%familyname%", os.Getenv("SEND_LIST_2"))
// 	// message.AddPersonalizations(p2)

// 	// // 件名を設定
// 	// message.Subject = os.Getenv("MAIL_SUBJECT")
// 	// diffPrice := Price.CheckDiffPrice(totalMiPrice, totalTaPrice)

// 	// var mailHeaderHtml = os.Getenv("MAIL_HEADER")
// 	// var mailTaHtml = "<strong>👨‍💻【" + os.Getenv("SEND_LIST_1") + "】👨‍💻</strong><br>" + "食費: " + costs["食費"].TPrice + "<br>" + "日用品: " + costs["日用品"].TPrice + "<br>" + "雑費: " + costs["雑費"].TPrice + "<br>" + "水道費: " + costs["水道費"].TPrice + "<br>" + "光熱費: " + costs["光熱費"].TPrice + "<br>" + "家賃: " + costs["家賃"].TPrice + "<br>" + "【合計】 : " + strconv.Itoa(totalTaPrice) + "<br><br>"
// 	// var mailMiHtml = "<strong>🤷‍♀【" + os.Getenv("SEND_LIST_2") + "】🤷‍♀️</strong><br>" + "食費: " + costs["食費"].MPrice + "<br>" + "日用品: " + costs["日用品"].MPrice + "<br>" + "雑費: " + costs["雑費"].MPrice + "<br>" + "水道費: " + costs["水道費"].MPrice + "<br>" + "光熱費: " + costs["光熱費"].MPrice + "<br>" + "家賃: " + costs["家賃"].MPrice + "<br>" + "【合計】 : " + strconv.Itoa(totalMiPrice) + "<br><br>"	
// 	// var mailDiffHtml = "差分: 💴" + strconv.Itoa(diffPrice) + "<br><br>"
// 	// var mailPokioCommentHtml = os.Getenv("MAIL_PO_COMMENT_HTML")
// 	// c := mail.NewContent("text/html",mailHeaderHtml + mailTaHtml + mailMiHtml + mailDiffHtml + mailPokioCommentHtml)
// 	// message.AddContent(c)

// 	// // メール送信を行い、レスポンスを表示
// 	// client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))

// 	// response, err := client.Send(message)
// 	// if err != nil {
// 	// 	log.Println(err)
// 	// } else {
// 	// 	fmt.Println(response.StatusCode)
// 	// 	fmt.Println(response.Body)
// 	// 	fmt.Println(response.Headers)
// 	// }