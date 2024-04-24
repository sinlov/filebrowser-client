package file_browser_client

type FileBrowserBasicFunc interface {
	Debug(isDebug bool)

	GetBaseUrl() string

	GetUsername() string

	SetRecaptcha(recaptcha string)

	GetRecaptcha() string

	IsLogin() bool

	Login() error
}
