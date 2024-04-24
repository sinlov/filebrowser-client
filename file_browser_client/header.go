package file_browser_client

const (
	baseUserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36"
)

var baseHeaderSave map[string]string

func BaseHeader() map[string]string {
	if baseHeaderSave != nil {
		return baseHeaderSave
	}
	header := map[string]string{
		"User-Agent": baseUserAgent + " github.com/sinlov/filebrowser-client",
	}
	baseHeaderSave = header
	return baseHeaderSave
}
