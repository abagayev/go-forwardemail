package forwardemail

import (
	"encoding/json"
	"fmt"
	"time"
)

type Alias struct {
	Account                  Account   `json:"user"`
	Domain                   Domain    `json:"domain"`
	Name                     string    `json:"name"`
	Labels                   []string  `json:"labels"`
	IsEnabled                bool      `json:"is_enabled"`
	HasRecipientVerification bool      `json:"has_recipient_verification"`
	Recipients               []string  `json:"recipients"`
	Id                       string    `json:"id"`
	Object                   string    `json:"object"`
	CreatedAt                time.Time `json:"created_at"`
	UpdatedAt                time.Time `json:"updated_at"`
}

func (c *Client) GetAliases(domain string) ([]Alias, error) {
	req, err := c.newRequest("GET", fmt.Sprintf("/v1/domains/%s/aliases", domain))
	if err != nil {
		return nil, err
	}

	res, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var items []Alias

	err = json.Unmarshal(res, &items)
	if err != nil {
		return nil, err
	}

	return items, nil
}

func (c *Client) GetAlias(domain string, alias string) (*Alias, error) {
	req, err := c.newRequest("GET", fmt.Sprintf("/v1/domains/%s/aliases/%s", domain, alias))
	if err != nil {
		return nil, err
	}

	res, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var item Alias

	err = json.Unmarshal(res, &item)
	if err != nil {
		return nil, err
	}

	return &item, nil
}

func (c *Client) DeleteAlias(domain string, alias string) error {
	req, err := c.newRequest("DELETE", fmt.Sprintf("/v1/domains/%s/aliases/%s", domain, alias))
	if err != nil {
		return err
	}

	_, err = c.doRequest(req)
	if err != nil {
		return err
	}

	return nil
}
