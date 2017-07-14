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

// Leads returns an array of Leads from the API. The critera to filter leads can
// be configured by passing functional options. See the examples for a look at
// how to do this.
func (c Client) Leads(options ...func(*leadConfig)) (*[]Lead, error) {
	cfg := loadConfig(options)

	if len(c.ID) == 0 {
		return nil, errors.New("Missing Client ID")
	}

	if len(c.Secret) == 0 {
		return nil, errors.New("Missing Client Secret")
	}

	var response *http.Response
	var err error
	if response, err = c.request(apiURL, cfg); err != nil {
		return nil, err
	}

	if response.StatusCode == 403 {
		return nil, errors.New("Forbidden")
	}

	var body []byte
	if body, err = ioutil.ReadAll(response.Body); err != nil {
		return nil, err
	}

	data := leadsResponse{}
	if err = json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	return &data.Leads, nil
}

func (c Client) request(url string, cfg leadConfig) (*http.Response, error) {
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

	req.URL.RawQuery = q.Encode()

	return netClient.Do(req)
}
