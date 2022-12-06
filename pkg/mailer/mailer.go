package pkg

import (
	"fmt"
	"net/smtp"

	"github.com/tuongnguyen1209/poly-career-back/config"
)

type DataMailer struct {
	Subject   string `json:"subject"`
	Body      string `json:"body"`
	Recipient string `json:"recipient" required:"biding"`
}

func mailer(title string, data *DataMailer, templateId int) {
	config := config.GetConfig()

	go func() {
		from := config.Mailer.Sender
		password := config.Mailer.Pass
		toEmailAddress := data.Recipient
		to := []string{toEmailAddress}
		host := config.Mailer.Host
		port := config.Mailer.Port
		smtp.Dial(host)
		address := host + ":" + port
		fmt.Println(from, password, host)
		auth := smtp.PlainAuth("", from, password, host)
		//----------------------------------------------------------------------------------------------------
		subject := fmt.Sprintf("Subject:%s\nFrom: PolyCareer<%s>\nTo: <%s>\n", title, config.Mailer.Sender, toEmailAddress)
		mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
		body := templateGender(data.Subject, data.Body, templateId)

		//----------------------------------------------------------------------------------------------------

		err := smtp.SendMail(address, auth, config.Mailer.Sender, to, []byte(subject+mime+body))
		if err != nil {
			fmt.Println("Error email: ", err.Error())
		}
	}()
}

func MailVerifyAccount(data *DataMailer) {
	mailer("Xác nhận đăng ký tài khoản", data, 1)
}

func MailChangeEmail(data *DataMailer) {
	mailer("Xác nhận thay đổi email", data, 2)
}
