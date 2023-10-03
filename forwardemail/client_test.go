package forwardemail

import (
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNewClient(t *testing.T) {
	tests := []struct {
		name    string
		options ClientOptions
		want    *Client
	}{
		{
			name:    "empty options",
			options: ClientOptions{},
			want: &Client{
				ApiUrl:     "https://api.forwardemail.net",
				HttpClient: &http.Client{},
			},
		},
		{
			name: "with api key",
			options: ClientOptions{
				ApiKey: "4e4d6c332b6fe62a63afe56171fd3725",
			},
			want: &Client{
				ApiKey:     "4e4d6c332b6fe62a63afe56171fd3725",
				ApiUrl:     "https://api.forwardemail.net",
				HttpClient: &http.Client{},
			},
		},
		{
			name: "with api url",
			options: ClientOptions{
				ApiUrl: "https://google.com",
			},
			want: &Client{
				ApiUrl:     "https://google.com",
				HttpClient: &http.Client{},
			},
		},
		{
			name: "with everything at once",
			options: ClientOptions{
				ApiKey: "4e4d6c332b6fe62a63afe56171fd3725",
				ApiUrl: "https://google.com",
			},
			want: &Client{
				ApiKey:     "4e4d6c332b6fe62a63afe56171fd3725",
				ApiUrl:     "https://google.com",
				HttpClient: &http.Client{},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewClient(tt.options)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Fatalf("values are not the same %s", diff)
			}
		})
	}
}
