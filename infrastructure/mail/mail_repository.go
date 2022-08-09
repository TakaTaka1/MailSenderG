package mailInfra

import (
	"MailSenderG/data/StructData"
	mailModel "MailSenderG/domain/model/mail"
	"os"
	"strconv"

	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type MailRepository struct {
}

func NewMailRepository() *MailRepository {
	return &MailRepository{}
}

func (mp *MailRepository) SetupSendGridMail() *mail.SGMailV3 {
	return mail.NewV3Mail()
}

func (mp *MailRepository) SetupMailFrom(sg *mail.SGMailV3, mailInfo mailModel.MailInfo) {
	sgfrom := mail.NewEmail("", mailInfo.From)
	sg.SetFrom(sgfrom)
}

func (mp *MailRepository) SetupMailTo(sg *mail.SGMailV3, mailInfo mailModel.MailInfo) {
	p := mail.NewPersonalization()
	to1 := mail.NewEmail("", mailInfo.To[0])
	p.AddTos(to1)
	p.SetSubstitution("%fullname%", os.Getenv("SEND_LIST_1"))
	p.SetSubstitution("%familyname%", os.Getenv("SEND_LIST_1"))
	sg.AddPersonalizations(p)

	// 2つ目の宛先と、対応するSubstitutionタグを指定
	p2 := mail.NewPersonalization()
	to2 := mail.NewEmail("", mailInfo.To[1])
	p2.AddTos(to2)
	p2.SetSubstitution("%fullname%", os.Getenv("SEND_LIST_2"))
	p2.SetSubstitution("%familyname%", os.Getenv("SEND_LIST_2"))
	sg.AddPersonalizations(p2)
}

func (mp *MailRepository) SetupMailSubject(sg *mail.SGMailV3, mailInfo mailModel.MailInfo) {
	sg.Subject = mailInfo.Subject
}

func (mp *MailRepository) SetupMailHeader(mailInfo mailModel.MailInfo) string {
	return mailInfo.Header
}

func (mp *MailRepository) SetupMailBody(sg *mail.SGMailV3, header string, diffPrice int, costs map[string]StructData.SheetData, totalTaPrice int, totalMiPrice int) *mail.SGMailV3 {
	var mailTaHtml = "<strong>👨‍💻【" + os.Getenv("SEND_LIST_1") + "】👨‍💻</strong><br>" + "食費: " + costs["食費"].TPrice + "<br>" + "日用品: " + costs["日用品"].TPrice + "<br>" + "雑費: " + costs["雑費"].TPrice + "<br>" + "水道費: " + costs["水道費"].TPrice + "<br>" + "光熱費: " + costs["光熱費"].TPrice + "<br>" + "家賃: " + costs["家賃"].TPrice + "<br>" + "【合計】 : " + strconv.Itoa(totalTaPrice) + "<br><br>"
	var mailMiHtml = "<strong>🤷‍♀【" + os.Getenv("SEND_LIST_2") + "】🤷‍♀️</strong><br>" + "食費: " + costs["食費"].MPrice + "<br>" + "日用品: " + costs["日用品"].MPrice + "<br>" + "雑費: " + costs["雑費"].MPrice + "<br>" + "水道費: " + costs["水道費"].MPrice + "<br>" + "光熱費: " + costs["光熱費"].MPrice + "<br>" + "家賃: " + costs["家賃"].MPrice + "<br>" + "【合計】 : " + strconv.Itoa(totalMiPrice) + "<br><br>"
	var mailDiffHtml = "差分: 💴" + strconv.Itoa(diffPrice) + "<br><br>"
	var mailPokioCommentHtml = os.Getenv("MAIL_PO_COMMENT_HTML")
	c := mail.NewContent("text/html", header+mailTaHtml+mailMiHtml+mailDiffHtml+mailPokioCommentHtml)
	sg.AddContent(c)

	return sg
}
