package email

import (
	"gopkg.in/gomail.v2"
)

var (
	_my_mail   = "soekchl@163.com"
	_my_passwrod = ""

)

func Config(myMail, pwd string) {
	_my_mail = myMail
	_my_passwrod = pwd
}

func SendCodeInMail(to, account, code string) error {
	m := gomail.NewMessage()
	m.SetAddressHeader("From", _my_mail, "游戏中心机器人") // 发件人
	m.SetHeader("To",                               // 收件人
		m.FormatAddress(to, account+" 验证码"),
	)
	m.SetHeader("Subject", "游戏中心验证码")
	m.SetBody("text/html", "您好！ 感谢注册本网站 账号："+account+" <br> 本次验证码为： <b>"+code+"</b>") // 正文

	//	d := gomail.NewPlainDialer("smtp.163.com", 25, _my_mail, _my_passwrod)
	d := gomail.NewPlainDialer("smtp.163.com", 465, _my_mail, _my_passwrod)
	return d.DialAndSend(m)
}
