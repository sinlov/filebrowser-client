package file_browser_client

type (
	FileBrowserClient struct {
		isDebug bool

		recaptcha string

		username string
		password string
		baseUrl  string

		timeoutSecond     uint
		timeoutFileSecond uint

		authHeadVal string

		FileBrowserBasicFunc FileBrowserBasicFunc `json:"-"`

		FileBrowserResourceFunc FileBrowserResourceFunc `json:"-"`

		FileBrowserDirectoryFunc FileBrowserDirectoryFunc `json:"-"`

		FileBrowserShareFunc FileBrowserShareFunc `json:"-"`
	}
)

// NewClient
// new client for filebrowser
func NewClient(
	username, password, baseUrl string,
	timeoutSecond, timeoutFileSecond uint,
) (FileBrowserClient, error) {
	client := FileBrowserClient{}
	if timeoutSecond < 10 {
		timeoutSecond = 10
	}

	if timeoutFileSecond < 30 {
		timeoutFileSecond = 30
	}
	return client.client(
		username, password, baseUrl,
		timeoutSecond, timeoutFileSecond,
	)
}
