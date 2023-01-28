package web_api

import "fmt"

type LoginRequest struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Recaptcha string `json:"recaptcha"`
}

func ApiLogin() string {
	return fmt.Sprintf("%s/%s", ApiBase(), "login")
}
