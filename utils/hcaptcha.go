package utils

import (
	"encoding/json"

	"github.com/go-resty/resty/v2"
	"pentag.kr/api-server/configs"
)

const VERIFY_URL = "https://api.hcaptcha.com/siteverify"


type CaptchaResponse struct {
	Success bool `json:"success"`
}

func CheckCaptcha(token string) bool {
	resp, err := resty.New().R().
		SetFormData(map[string]string{
			"secret": configs.Env.CaptchaSecret,
			"response": token,
		}).Post(VERIFY_URL)	

	if err != nil {
		return false
	}
	if resp.StatusCode() != 200 {
		return false
	}

	var captchaResponse CaptchaResponse
	body := resp.Body()

	if err := json.Unmarshal(body, &captchaResponse); err != nil {
		return false
	}

	return captchaResponse.Success
}