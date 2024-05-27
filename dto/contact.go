package dto

import (
	"errors"
	"fmt"
	"strings"

	"pentag.kr/api-server/configs"
	"pentag.kr/api-server/utils"
)

type ContactReq struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Category  int   `json:"category"`
	Message  string `json:"message"`
	Token    string `json:"token"`
}

func (c *ContactReq) Validate() error {

	if c.FirstName == "" || len(c.FirstName) > 50 {
		return errors.New("first name cannot be empty or more than 50 characters")
	}
	if c.LastName == "" || len(c.LastName) > 50 {
		return errors.New("last name cannot be empty or more than 50 characters")
	}
	if c.Email == "" || len(c.Email) > 500 {
		return errors.New("email cannot be empty or more than 500 characters")
	}
	if c.Phone == "" || len(c.Phone) > 30 {
		return errors.New("phone cannot be empty or more than 30 characters")
	}
	if c.Category < 0 || c.Category > 2 {
		return errors.New("category must be between 0 and 2")
	}
	if c.Message == "" || len(c.Message) > 5000 {

		return errors.New("message cannot be empty or more than 5000 characters")
	}
	if c.Token == "" {
		return errors.New("token cannot be empty")
	}
	return nil
}

func (c *ContactReq) ReplaceSymbol() {
	var replace = func (value string) string {
		
		value = strings.Replace(value, "<", "&lt;", -1)
		value = strings.Replace(value, ">", "&gt;", -1)
		value = strings.Replace(value, "(", "&#40;", -1)
		value = strings.Replace(value, ")", "&#41;", -1)
		value = strings.Replace(value, "'", "&#x27;", -1)
		value = strings.Replace(value, "\"", "&#x22;", -1)
		value = strings.Replace(value, "&", "&amp;", -1)
		value = strings.Replace(value, "+", "&#43;", -1)
		value = strings.Replace(value, "%", "&#37;", -1)
		value = strings.Replace(value, ";", "&#59;", -1)
		value = strings.Replace(value, ":", "&#58;", -1)
		value = strings.Replace(value, "!", "&#33;", -1)
		value = strings.Replace(value, "?", "&#63;", -1)
		value = strings.Replace(value, "=", "&#61;", -1)
		value = strings.Replace(value, "*", "&#42;", -1)
		value = strings.Replace(value, "#", "&#35;", -1)
		value = strings.Replace(value, "$", "&#36;", -1)
		value = strings.Replace(value, "{", "&#123;", -1)
		value = strings.Replace(value, "}", "&#125;", -1)
		value = strings.Replace(value, "[", "&#91;", -1)
		value = strings.Replace(value, "]", "&#93;", -1)
		value = strings.Replace(value, "/", "&#47;", -1)
		value = strings.Replace(value, "\\", "&#92;", -1)
		value = strings.Replace(value, "|", "&#124;", -1)
		value = strings.Replace(value, "`", "&#96;", -1)
		value = strings.Replace(value, "~", "&#126;", -1)
		value = strings.Replace(value, "^", "&#94;", -1)
		value = strings.Replace(value, "@", "&#64;", -1)
		value = strings.Replace(value, ",", "&#44;", -1)
		value = strings.Replace(value, ".", "&#46;", -1)
		value = strings.Replace(value, "_", "&#95;", -1)
		value = strings.Replace(value, "-", "&#45;", -1)
		value = strings.Replace(value, " ", "&nbsp;", -1)
		return value
	}
	c.FirstName = replace(c.FirstName)
	c.LastName = replace(c.LastName)
	c.Email = replace(c.Email)
	c.Phone = replace(c.Phone)
	c.Message = replace(c.Message)
}

func (c *ContactReq) SendVerifyEmail(originalEmail string, verifyCode string) error {
	body := fmt.Sprintf(VERIFT_EMAIL_TEMPLATE, verifyCode, c.FirstName, c.LastName, c.Email, c.Phone, c.CategoryToString(), c.Message)
	return utils.SendEmail(configs.Env.SMTPUser, []string{originalEmail}, "pentag.kr 문의를 위한 이메일 인증", body)
}

const VERIFT_EMAIL_TEMPLATE = `
<!DOCTYPE html>
<html>
<head>
	<title>pentag.kr 이메일 인증</title>
</head>
<body>
	<h1>이메일 인증 및 사본</h1>
	<p>Click the link below to verify your email address</p>
	<a href="https://pentag.kr/verify/%s">이메일 인증(Verify Email)</a>

	<h2>문의 사본(Copy)</h2>
	<p>이름(Name): %s %s</p>
	<p>이메일(Email): %s</p>
	<p>전화번호(Phone): %s</p>
	<p>문의 유형(Category): %s</p>
	<p>메시지(Message): %s</p>
</body>
</html>
`

func(c *ContactReq) CategoryToString() string {
	switch c.Category {
	case 0:
		return "협업 관련(Collaboration)"
	case 1:
		return "채용 관련(Hiring)"
	case 2:
		return "기타 문의(Etc)"
	default:
		return "Unknown"
	}
}