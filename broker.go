package fyersgosdk

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) GetMarketStatus(fyClient *Client) (string, MarketStatus, error) {
	token := fmt.Sprintf("%s:%s", fyClient.appId, fyClient.accessToken)
	headers := http.Header{}
	headers.Add("Authorization", token)
	response, err := c.httpClient.DoRaw(http.MethodGet, MarketStatusURL, nil, headers)
	if err != nil {
		return "", MarketStatus{}, err
	}
	var marketStatus MarketStatus
	if err := json.Unmarshal(response.Body, &marketStatus); err != nil {
		return "", MarketStatus{}, err
	}
	return string(response.Body), marketStatus, nil
}
