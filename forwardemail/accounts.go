package forwardemail

import (
	"encoding/json"
	"time"
)

type Account struct {
	Plan           string    `json:"plan"`
	Email          string    `json:"email"`
	FullEmail      string    `json:"full_email"`
	DisplayName    string    `json:"display_name"`
	LastLocale     string    `json:"last_locale"`
	AddressCountry string    `json:"address_country"`
	Id             string    `json:"id"`
	Object         string    `json:"object"`
	Locale         string    `json:"locale"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	AddressHtml    string    `json:"address_html"`
}

func (c *Client) GetAccount() (*Account, error) {
	req, err := c.newRequest("GET", "/v1/account")
	if err != nil {
		return nil, err
	}

	res, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var item Account

	err = json.Unmarshal(res, &item)
	if err != nil {
		return nil, err
	}

	return &item, nil
}
