package fyersgosdk

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) GetProfile(fyClient *Client) (string, Profile, error) {
	token := fmt.Sprintf("%s:%s", fyClient.appId, fyClient.accessToken)
	headers := http.Header{}
	headers.Add("Authorization", token)
	response, err := c.httpClient.DoRaw(http.MethodGet, ProfileURL, nil, headers)
	if err != nil {
		return "", Profile{}, err
	}
	var profile Profile
	if err := json.Unmarshal(response.Body, &profile); err != nil {
		return "", Profile{}, err
	}
	return string(response.Body), profile, nil
}

func (c *Client) GetFunds(fyClient *Client) (string, Funds, error) {
	token := fmt.Sprintf("%s:%s", fyClient.appId, fyClient.accessToken)
	headers := http.Header{}
	headers.Add("Authorization", token)
	response, err := c.httpClient.DoRaw(http.MethodGet, FundURL, nil, headers)
	if err != nil {
		return "", Funds{}, err
	}
	var funds Funds
	if err := json.Unmarshal(response.Body, &funds); err != nil {
		return "", Funds{}, err
	}
	return string(response.Body), funds, nil
}

func (c *Client) GetHoldings(fyClient *Client) (string, Holdings, error) {
	token := fmt.Sprintf("%s:%s", fyClient.appId, fyClient.accessToken)
	headers := http.Header{}
	headers.Add("Authorization", token)
	response, err := c.httpClient.DoRaw(http.MethodGet, HoldingsURL, nil, headers)
	if err != nil {
		return "", Holdings{}, err
	}
	var holdings Holdings
	if err := json.Unmarshal(response.Body, &holdings); err != nil {
		return "", Holdings{}, err
	}
	return string(response.Body), holdings, nil
}

func (c *Client) Logout(fyClient *Client) (string, APIResponse, error) {
	token := fmt.Sprintf("%s:%s", fyClient.appId, fyClient.accessToken)
	headers := http.Header{}
	headers.Add("Authorization", token)
	response, err := c.httpClient.DoRaw(http.MethodPost, LogoutURL, nil, headers)
	if err != nil {
		return "", APIResponse{}, err
	}
	var apiResponse APIResponse
	if err := json.Unmarshal(response.Body, &apiResponse); err != nil {
		return "", APIResponse{}, err
	}
	
	return string(response.Body), apiResponse, nil
}