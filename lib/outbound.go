package lib

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/Investly-id/common-go/v3/payload"
)

type Outbound struct {
	secret  string
	client  *http.Client
	baseUrl string
}

func NewOutboundService(client *http.Client, secret string, baseUrl string) *Outbound {
	return &Outbound{
		secret:  secret,
		baseUrl: baseUrl,
		client:  client,
	}
}

func (m *Outbound) ConstructJsonRequest(method string, url string, body io.Reader) (*payload.Response, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-internal-token", m.secret)

	res, err := m.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	var data *payload.Response

	err = json.NewDecoder(res.Body).Decode(&data)

	if err != nil {
		return nil, err
	}

	return data, nil
}
