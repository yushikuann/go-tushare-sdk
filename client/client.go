package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type TuShare struct {
	token  string
	client *http.Client
}

func New(token string) *TuShare {
	return NewWithClient(token, http.DefaultClient)
}

func NewWithClient(token string, httpClient *http.Client) *TuShare {
	return &TuShare{
		token:  token,
		client: httpClient,
	}
}

func (api *TuShare) request(method, path string, body interface{}) (*http.Request, error) {
	bodyJSON, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(method, path, bytes.NewBuffer(bodyJSON))
	if err != nil {
		return nil, err
	}
	return req, nil
}

func (api *TuShare) doRequest(req *http.Request) (*APIResponse, error) {
	req.Header.Set("Content-Type", "application/json")

	var resp *http.Response
	var err error
	if resp, err = api.client.Do(req); err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("net work error: %d", resp.StatusCode)
	}
	defer resp.Body.Close()

	var body []byte
	if body, err = ioutil.ReadAll(resp.Body); err != nil {
		return nil, err
	}

	var result *APIResponse
	if err = json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	switch result.Code {
	case -2001:
		return result, ERR_ARGUEMENT
	case -2002:
		return result, ERR_PERMISSION
	default:
		return result, nil
	}
}

func (api *TuShare) postData(body map[string]interface{}) (*APIResponse, error) {
	req, err := api.request(PostMethod, Domain, body)
	if err != nil {
		return nil, err
	}
	resp, err := api.doRequest(req)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
