package forwardemail

import "net/http"

const (
	forwardemailApiUrl = "https://api.forwardemail.net"
)

type ClientOptions struct {
	ApiKey string
	ApiUrl string
}

type Client struct {
	ApiKey string
	ApiUrl string

	HttpClient *http.Client
}

// NewClient returns a new Forward Email API Client.
func NewClient(options ClientOptions) *Client {
	apiUrl := forwardemailApiUrl
	if options.ApiUrl != "" {
		apiUrl = options.ApiUrl
	}

	return &Client{
		ApiKey:     options.ApiKey,
		ApiUrl:     apiUrl,
		HttpClient: http.DefaultClient,
	}
}
