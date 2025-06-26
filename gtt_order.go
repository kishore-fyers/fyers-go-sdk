package fyersgosdk

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) GTTSingleOrderAction(fyClient *Client, orderRequest GTTOrderRequest) (string, OrderResponse, error) {
	token := fmt.Sprintf("%s:%s", fyClient.appId, fyClient.accessToken)
	headers := http.Header{}
	headers.Add("Authorization", token)
	requestBody, err := json.Marshal(orderRequest)
	if err != nil {
		return "", OrderResponse{}, fmt.Errorf("failed to marshal order request: %w", err)
	}
	response, err := c.httpClient.DoRaw(http.MethodPost, GTTOrderURL, requestBody, headers)
	if err != nil {
		return "", OrderResponse{}, err
	}
	var orderResp OrderResponse
	if err := json.Unmarshal(response.Body, &orderResp); err != nil {
		return "", OrderResponse{}, err
	}
	return string(response.Body), orderResp, nil
}

func (c *Client) GTTMultiOrderAction(fyClient *Client, orderRequests []GTTOrderRequest) (string, OrderResponse, error) {
	token := fmt.Sprintf("%s:%s", fyClient.appId, fyClient.accessToken)
	headers := http.Header{}
	headers.Add("Authorization", token)

	// Marshal the order requests to JSON
	requestBody, err := json.Marshal(orderRequests)
	if err != nil {
		return "", OrderResponse{}, fmt.Errorf("failed to marshal order requests: %w", err)
	}

	response, err := c.httpClient.DoRaw(http.MethodPost, GTTOrderURL, requestBody, headers)
	if err != nil {
		return "", OrderResponse{}, err
	}
	var orderResp OrderResponse
	if err := json.Unmarshal(response.Body, &orderResp); err != nil {
		return "", OrderResponse{}, err
	}
	return string(response.Body), orderResp, nil
}

func (c *Client) ModifyGTTOrder(fyClient *Client, orderRequest ModifyGTTOrderRequest) (string, OrderResponse, error) {
	token := fmt.Sprintf("%s:%s", fyClient.appId, fyClient.accessToken)
	headers := http.Header{}
	headers.Add("Authorization", token)
	requestBody, err := json.Marshal(orderRequest)
	if err != nil {
		return "", OrderResponse{}, fmt.Errorf("failed to marshal order request: %w", err)
	}
	response, err := c.httpClient.DoRaw(http.MethodPatch, GTTOrderURL, requestBody, headers)
	if err != nil {
		return "", OrderResponse{}, err
	}
	var orderResp OrderResponse
	if err := json.Unmarshal(response.Body, &orderResp); err != nil {
		return "", OrderResponse{}, err
	}
	return string(response.Body), orderResp, nil
}

func (c *Client) CancelGTTOrder(fyClient *Client, orderId string) (string, OrderResponse, error) {
	token := fmt.Sprintf("%s:%s", fyClient.appId, fyClient.accessToken)
	headers := http.Header{}
	headers.Add("Authorization", token)

	// Create cancel request with orderId
	cancelRequest := CancelGTTOrderRequest{Id: orderId}
	requestBody, err := json.Marshal(cancelRequest)
	if err != nil {
		return "", OrderResponse{}, fmt.Errorf("failed to marshal cancel request: %w", err)
	}

	response, err := c.httpClient.DoRaw(http.MethodDelete, GTTOrderURL, requestBody, headers)
	if err != nil {
		return "", OrderResponse{}, err
	}
	var orderResp OrderResponse
	if err := json.Unmarshal(response.Body, &orderResp); err != nil {
		return "", OrderResponse{}, err
	}
	return string(response.Body), orderResp, nil
}

func (c *Client) GetGTTOrderBook(fyClient *Client) (string, GTTOrderBookItem, error) {
	token := fmt.Sprintf("%s:%s", fyClient.appId, fyClient.accessToken)
	headers := http.Header{}
	headers.Add("Authorization", token)
	response, err := c.httpClient.DoRaw(http.MethodGet, GTTOrderBookURL, nil, headers)
	if err != nil {
		return "", GTTOrderBookItem{}, err
	}
	var orderResp GTTOrderBookItem
	if err := json.Unmarshal(response.Body, &orderResp); err != nil {
		return "", GTTOrderBookItem{}, err
	}
	return string(response.Body), orderResp, nil
}
