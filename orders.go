package fyersgosdk

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) SingleOrderAction(fyClient *Client, orderRequest OrderRequest) (string, OrderResponse, error) {
	token := fmt.Sprintf("%s:%s", fyClient.appId, fyClient.accessToken)
	headers := http.Header{}
	headers.Add("Authorization", token)

	// Marshal the order request to JSON
	requestBody, err := json.Marshal(orderRequest)
	if err != nil {
		return "", OrderResponse{}, fmt.Errorf("failed to marshal order request: %w", err)
	}

	response, err := c.httpClient.DoRaw(http.MethodPost, SingleOrderActionURL, requestBody, headers)
	if err != nil {
		return "", OrderResponse{}, err
	}
	var orderResp OrderResponse
	if err := json.Unmarshal(response.Body, &orderResp); err != nil {
		return "", OrderResponse{}, err
	}
	return string(response.Body), orderResp, nil
}

func (c *Client) MultiOrderAction(fyClient *Client, orderRequests []OrderRequest) (string, OrderResponse, error) {
	token := fmt.Sprintf("%s:%s", fyClient.appId, fyClient.accessToken)
	headers := http.Header{}
	headers.Add("Authorization", token)

	// Marshal the order requests to JSON
	requestBody, err := json.Marshal(orderRequests)
	if err != nil {
		return "", OrderResponse{}, fmt.Errorf("failed to marshal order requests: %w", err)
	}

	response, err := c.httpClient.DoRaw(http.MethodPost, MultipleOrderActionURL, requestBody, headers)
	if err != nil {
		return "", OrderResponse{}, err
	}
	var orderResp OrderResponse
	if err := json.Unmarshal(response.Body, &orderResp); err != nil {
		return "", OrderResponse{}, err
	}
	return string(response.Body), orderResp, nil
}

func (c *Client) MultiLegOrderAction(fyClient *Client, orderRequests []MultiLegOrderRequest) (string, OrderResponse, error) {
	token := fmt.Sprintf("%s:%s", fyClient.appId, fyClient.accessToken)
	headers := http.Header{}
	headers.Add("Authorization", token)

	// Marshal the order requests to JSON
	requestBody, err := json.Marshal(orderRequests)
	if err != nil {
		return "", OrderResponse{}, fmt.Errorf("failed to marshal order requests: %w", err)
	}

	response, err := c.httpClient.DoRaw(http.MethodPost, MultiLegOrderURL, requestBody, headers)
	if err != nil {
		return "", OrderResponse{}, err
	}
	var orderResp OrderResponse
	if err := json.Unmarshal(response.Body, &orderResp); err != nil {
		return "", OrderResponse{}, err
	}
	return string(response.Body), orderResp, nil
}
