// Package kaeufer is an API Client for the KäuferPortal Lead API
package kaeufer

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net"
	"net/http"
	"strconv"
	"time"
)

const apiURL = "https://leads.kaeuferportal.de/v201609/leads"

type leadsResponse struct {
	Leads []Lead `json:"leads"`
}

// Client contains the authentication information for the API
type Client struct {
	ID     string
	Secret string
}

// Validate that the client has a Client ID and Secret.
func (c Client) Validate() error {
	if len(c.ID) == 0 {
		return errors.New("Missing Client ID")
	}

	if len(c.Secret) == 0 {
		return errors.New("Missing Client Secret")
	}

	return nil
}

func (c Client) fetchPage(cfg leadConfig, page int) (*[]Lead, error) {
	r, err := c.request(apiURL, cfg, page)
	if err != nil {
		return nil, err
	}

	defer r.Body.Close()

	if r.StatusCode == 403 {
		return nil, errors.New("Forbidden")
	}

	return parseLeadsFromResponse(r)
}

func parseLeadsFromResponse(r *http.Response) (*[]Lead, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	data := leadsResponse{}
	if err = json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	return &data.Leads, nil
}

// Leads returns an array of Leads from the API.
func (c Client) Leads(options ...func(*leadConfig)) (*[]Lead, error) {
	if err := c.Validate(); err != nil {
		return nil, err
	}

	cfg := loadConfig(options)

	var leads []Lead

	currentPage := 1
	for {
		pageLeads, err := c.fetchPage(cfg, currentPage)

		if err != nil {
			return nil, err
		}

		for _, l := range *pageLeads {
			leads = append(leads, l)
		}

		if len(*pageLeads) < cfg.PerPage {
			break
		}

		currentPage++
	}

	return &leads, nil
}

func (c Client) request(url string, cfg leadConfig, page int) (*http.Response, error) {
	timeout := 5 * time.Second
	dialer := net.Dialer{Timeout: timeout}

	netTransport := &http.Transport{
		Dial:                (&dialer).Dial,
		TLSHandshakeTimeout: timeout,
	}
	netClient := &http.Client{
		Timeout:   2 * timeout,
		Transport: netTransport,
	}

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("X-CLIENT-ID", c.ID)
	req.Header.Set("X-CLIENT-SECRET", c.Secret)

	q := req.URL.Query()
	if cfg.PerPage > -1 {
		q.Add("per_page", strconv.Itoa(cfg.PerPage))
	}

	if cfg.FromTimestamp > -1 {
		q.Add("from_timestamp", strconv.Itoa(cfg.FromTimestamp))
	}

	q.Add("page", strconv.Itoa(page))

	req.URL.RawQuery = q.Encode()

	return netClient.Do(req)
}
