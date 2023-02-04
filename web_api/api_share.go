package web_api

import (
	"fmt"
	tools "github.com/sinlov/filebrowser-client/tools/str_tools"
	"strconv"
)

func ApiShares() string {
	return fmt.Sprintf("%s/%s", ApiBase(), "shares")
}

func ApiShare() string {
	return fmt.Sprintf("%s/%s", ApiBase(), "share")
}

const (
	ShareUnitDays    = "days"
	ShareUnitHours   = "hours"
	ShareUnitMinutes = "minutes"
	ShareUnitSeconds = "seconds"
)

// ShareConfig
// @doc https://github.com/filebrowser/filebrowser/blob/master/share/share.go
type ShareConfig struct {
	Password string `json:"password,omitempty"`
	Expires  string `json:"expires,omitempty"`
	Unit     string `json:"unit,omitempty"`
}

var defineShareUnit []string

func ShareUnitDefine() []string {
	if defineShareUnit == nil {
		defineShareUnit = []string{
			ShareUnitDays,
			ShareUnitHours,
			ShareUnitMinutes,
			ShareUnitSeconds,
		}
	}
	return defineShareUnit
}

// CheckFalseShareConfig
// config pass will return false
func CheckFalseShareConfig(config ShareConfig) bool {

	if config.Expires == "" {
		return false
	}

	time, err := strconv.ParseInt(config.Expires, 10, 0)
	if err != nil {
		return true
	}
	if time < 0 {
		return true
	}

	return !tools.StrInArr(config.Unit, ShareUnitDefine())
}

// ShareLink
// Link is the information needed to build a shareable link.
type ShareLink struct {
	// Hash
	// share hash
	Hash string `json:"hash" storm:"id,index"`
	// Path
	// this path start / , is filebrowser user root start.
	Path   string `json:"path" storm:"index"`
	UserID uint   `json:"userID"`
	// Expire
	// time of expire is utc seconds
	Expire int64 `json:"expire"`
	// PasswordHash
	// password hash
	PasswordHash string `json:"password_hash,omitempty"`
	// Token is a random value that will only be set when PasswordHash is set. It is
	// URL-Safe and is used to download links in password-protected shares via a
	// query arg.
	Token string `json:"token,omitempty"`
}
