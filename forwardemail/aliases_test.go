package forwardemail

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestClient_GetAlias(t *testing.T) {
	type request struct {
		domain string
		alias  string
	}

	tests := []struct {
		name     string
		req      request
		domain   string
		response string
		want     *Alias
	}{
		{
			name: "no data",
		},
		{
			name: "ok",
			req: request{
				domain: "stark.com",
				alias:  "tony",
			},
			response: `{
				"user": {
				  "email": "tony@stark.com",
				  "display_name": "tony@stark.com",
				  "id": "59ad551ae6fb4a4c53427ca38079f029"
				},
				"domain": {
				  "name": "stark.com",
				  "id": "15ff615b6180f1fc7faf40e6"
				},
				"name": "tony",
				"description": "main email",
				"labels": [
				  "catch-all"
				],
				"is_enabled": true,
				"has_recipient_verification": true,
				"recipients": [
				  "james@rhodes.com"
				],
				"id": "6525b03e0bde8f333ace5824",
				"object": "alias",
				"created_at": "2023-10-10T20:12:46.588Z",
				"updated_at": "2023-10-10T20:12:46.588Z"
			}`,
			want: &Alias{
				Account: Account{
					Email:       "tony@stark.com",
					DisplayName: "tony@stark.com",
					Id:          "59ad551ae6fb4a4c53427ca38079f029",
				},
				Domain: Domain{
					Name: "stark.com",
					Id:   "15ff615b6180f1fc7faf40e6",
				},
				Name:                     "tony",
				Description:              "main email",
				Labels:                   []string{"catch-all"},
				IsEnabled:                true,
				HasRecipientVerification: true,
				Recipients:               []string{"james@rhodes.com"},
				Id:                       "6525b03e0bde8f333ace5824",
				Object:                   "alias",
				CreatedAt:                parseTime("2023-10-10T20:12:46.588Z"),
				UpdatedAt:                parseTime("2023-10-10T20:12:46.588Z"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintf(w, tt.response)
			}))
			defer svr.Close()

			c := NewClient(ClientOptions{
				ApiUrl: svr.URL,
			})

			got, _ := c.GetAlias(tt.req.domain, tt.req.alias)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Fatalf("values are not the same %s", diff)
			}
		})
	}
}

func TestClient_GetAliases(t *testing.T) {
	type request struct {
		domain string
	}

	tests := []struct {
		name string
		req  request
		res  string
		want []Alias
	}{
		{
			name: "no data",
		},
		{
			name: "ok",
			req: request{
				domain: "stark.com",
			},
			res: `[
				{
					"user": {
					  "email": "tony@stark.com",
					  "display_name": "tony@stark.com",
					  "id": "59ad551ae6fb4a4c53427ca38079f029"
					},
					"domain": {
					  "name": "stark.com",
					  "id": "15ff615b6180f1fc7faf40e6"
					},
					"name": "tony",
					"description": "main email",
					"labels": [
					  "catch-all"
					],
					"is_enabled": true,
					"has_recipient_verification": true,
					"recipients": [
					  "james@rhodes.com"
					],
					"id": "6525b03e0bde8f333ace5824",
					"object": "alias",
					"created_at": "2023-10-10T20:12:46.588Z",
					"updated_at": "2023-10-10T20:12:46.588Z"
				},
				{
					"user": {
					  "email": "tony@stark.com",
					  "display_name": "tony@stark.com",
					  "id": "59ad551ae6fb4a4c53427ca38079f029"
					},
					"domain": {
					  "name": "stark.com",
					  "id": "15ff615b6180f1fc7faf40e6"
					},
					"name": "james",
					"labels": [
					  "catch-all"
					],
					"is_enabled": true,
					"has_recipient_verification": true,
					"recipients": [
					  "james@rhodes.com"
					],
					"id": "b078f60f2636c4d6cf668d9b36a3e42e",
					"object": "alias",
					"created_at": "2023-10-12T18:11:22.123Z",
					"updated_at": "2023-10-12T19:55:56.534Z"
				}
			]`,
			want: []Alias{
				{
					Account: Account{
						Email:       "tony@stark.com",
						DisplayName: "tony@stark.com",
						Id:          "59ad551ae6fb4a4c53427ca38079f029",
					},
					Domain: Domain{
						Name: "stark.com",
						Id:   "15ff615b6180f1fc7faf40e6",
					},
					Name:                     "tony",
					Description:              "main email",
					Labels:                   []string{"catch-all"},
					IsEnabled:                true,
					HasRecipientVerification: true,
					Recipients:               []string{"james@rhodes.com"},
					Id:                       "6525b03e0bde8f333ace5824",
					Object:                   "alias",
					CreatedAt:                parseTime("2023-10-10T20:12:46.588Z"),
					UpdatedAt:                parseTime("2023-10-10T20:12:46.588Z"),
				},
				{
					Account: Account{
						Email:       "tony@stark.com",
						DisplayName: "tony@stark.com",
						Id:          "59ad551ae6fb4a4c53427ca38079f029",
					},
					Domain: Domain{
						Name: "stark.com",
						Id:   "15ff615b6180f1fc7faf40e6",
					},
					Name:                     "james",
					Labels:                   []string{"catch-all"},
					IsEnabled:                true,
					HasRecipientVerification: true,
					Recipients:               []string{"james@rhodes.com"},
					Id:                       "b078f60f2636c4d6cf668d9b36a3e42e",
					Object:                   "alias",
					CreatedAt:                parseTime("2023-10-12T18:11:22.123Z"),
					UpdatedAt:                parseTime("2023-10-12T19:55:56.534Z"),
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintf(w, tt.res)
			}))
			defer svr.Close()

			c := NewClient(ClientOptions{
				ApiUrl: svr.URL,
			})

			got, _ := c.GetAliases(tt.req.domain)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Fatalf("values are not the same %s", diff)
			}
		})
	}
}

func TestClient_CreateAlias(t *testing.T) {
	type request struct {
		domain string
		alias  string
		params AliasParameters
	}

	tests := []struct {
		name string
		req  request
		res  string
		want *Alias
	}{
		{
			name: "no data",
		},
		{
			name: "ok",
			req: request{
				domain: "stark.com",
				alias:  "*",
				params: AliasParameters{
					Recipients:               pointSliceOfStrings([]string{"james@rhodes.com"}),
					Description:              "main email",
					Labels:                   pointSliceOfStrings([]string{"catch-all"}),
					IsEnabled:                pointBool(true),
					HasRecipientVerification: pointBool(true),
				},
			},
			res: `{
				"user": {
				  "email": "tony@stark.com",
				  "display_name": "tony@stark.com",
				  "id": "59ad551ae6fb4a4c53427ca38079f029"
				},
				"domain": {
				  "name": "stark.com",
				  "id": "15ff615b6180f1fc7faf40e6"
				},
				"name": "*",
				"description": "main email",
				"labels": [
				  "catch-all"
				],
				"is_enabled": true,
				"has_recipient_verification": true,
				"recipients": [
				  "james@rhodes.com"
				],
				"id": "6525b03e0bde8f333ace5824",
				"object": "alias",
				"created_at": "2023-10-10T20:12:46.588Z",
				"updated_at": "2023-11-11T22:12:42.533Z"
			}`,
			want: &Alias{
				Account: Account{
					Email:       "tony@stark.com",
					DisplayName: "tony@stark.com",
					Id:          "59ad551ae6fb4a4c53427ca38079f029",
				},
				Domain: Domain{
					Name: "stark.com",
					Id:   "15ff615b6180f1fc7faf40e6",
				},
				Name:                     "*",
				Description:              "main email",
				Labels:                   []string{"catch-all"},
				IsEnabled:                true,
				HasRecipientVerification: true,
				Recipients:               []string{"james@rhodes.com"},
				Id:                       "6525b03e0bde8f333ace5824",
				Object:                   "alias",
				CreatedAt:                parseTime("2023-10-10T20:12:46.588Z"),
				UpdatedAt:                parseTime("2023-11-11T22:12:42.533Z"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				logRequestBody(r, t)

				fmt.Fprintf(w, tt.res)
			}))
			defer svr.Close()

			c := NewClient(ClientOptions{
				ApiUrl: svr.URL,
			})

			got, _ := c.CreateAlias(tt.req.domain, tt.req.alias, tt.req.params)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Fatalf("values are not the same %s", diff)
			}
		})
	}
}

func TestClient_UpdateAlias(t *testing.T) {
	type request struct {
		domain string
		alias  string
		params AliasParameters
	}

	tests := []struct {
		name string
		req  request
		res  string
		want *Alias
	}{
		{
			name: "no data",
		},
		{
			name: "ok",
			req: request{
				domain: "stark.com",
				alias:  "james",
				params: AliasParameters{
					Description:              "main email",
					Recipients:               pointSliceOfStrings([]string{"james@rhodes.com"}),
					Labels:                   pointSliceOfStrings([]string{"catch-all", "friends"}),
					IsEnabled:                pointBool(true),
					HasRecipientVerification: pointBool(true),
				},
			},
			res: `{
				"user": {
				  "email": "tony@stark.com",
				  "display_name": "tony@stark.com",
				  "id": "59ad551ae6fb4a4c53427ca38079f029"
				},
				"domain": {
				  "name": "stark.com",
				  "id": "15ff615b6180f1fc7faf40e6"
				},
				"name": "james",
				"description": "main email",
				"labels": [
				  "catch-all",
				  "friends"
				],
				"is_enabled": true,
				"has_recipient_verification": true,
				"recipients": [
				  "james@rhodes.com"
				],
				"id": "6525b03e0bde8f333ace5824",
				"object": "alias",
				"created_at": "2023-10-10T20:12:46.588Z",
				"updated_at": "2023-11-11T22:12:42.533Z"
			}`,
			want: &Alias{
				Account: Account{
					Email:       "tony@stark.com",
					DisplayName: "tony@stark.com",
					Id:          "59ad551ae6fb4a4c53427ca38079f029",
				},
				Domain: Domain{
					Name: "stark.com",
					Id:   "15ff615b6180f1fc7faf40e6",
				},
				Name:                     "james",
				Description:              "main email",
				Labels:                   []string{"catch-all", "friends"},
				IsEnabled:                true,
				HasRecipientVerification: true,
				Recipients:               []string{"james@rhodes.com"},
				Id:                       "6525b03e0bde8f333ace5824",
				Object:                   "alias",
				CreatedAt:                parseTime("2023-10-10T20:12:46.588Z"),
				UpdatedAt:                parseTime("2023-11-11T22:12:42.533Z"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				logRequestBody(r, t)

				fmt.Fprintf(w, tt.res)
			}))
			defer svr.Close()

			c := NewClient(ClientOptions{
				ApiUrl: svr.URL,
			})

			got, _ := c.UpdateAlias(tt.req.domain, tt.req.alias, tt.req.params)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Fatalf("values are not the same %s", diff)
			}
		})
	}
}

func TestClient_DeleteAlias(t *testing.T) {
	type request struct {
		domain string
		alias  string
	}
	type response struct {
		code int
		body string
	}

	tests := []struct {
		name string
		req  request
		res  response
		want error
	}{
		{
			name: "ok",
			res: response{
				code: http.StatusNoContent,
			},
		},
		{
			name: "not ok",
			res: response{
				code: http.StatusInternalServerError,
				body: "oh no",
			},
			want: fmt.Errorf("status: 500, body: oh no"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(tt.res.code)
				fmt.Fprintf(w, tt.res.body)
			}))
			defer svr.Close()

			c := NewClient(ClientOptions{
				ApiUrl: svr.URL,
			})

			got := c.DeleteAlias(tt.req.domain, tt.req.alias)
			if diff := cmp.Diff(tt.want, got, cmp.Comparer(equateErrorMessage)); diff != "" {
				t.Fatalf("values are not the same %s", diff)
			}
		})
	}
}

func pointSliceOfStrings(s []string) *[]string {
	return &s
}

func logRequestBody(r *http.Request, t *testing.T) {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}
	t.Logf("body encoded fields: %s", b)
}
