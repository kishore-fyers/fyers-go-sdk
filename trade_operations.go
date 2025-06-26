package fyersgosdk

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) ModifyOrder(fyClient *Client, orderRequest ModifyOrderRequest) (string, OrderResponse, error) {
	token := fmt.Sprintf("%s:%s", fyClient.appId, fyClient.accessToken)
	headers := http.Header{}
	headers.Add("Authorization", token)
	requestBody, err := json.Marshal(orderRequest)
	if err != nil {
		return "", OrderResponse{}, fmt.Errorf("failed to marshal order request: %w", err)
	}
	response, err := c.httpClient.DoRaw(http.MethodPatch, SingleOrderActionURL, requestBody, headers)
	if err != nil {
		return "", OrderResponse{}, err
	}
	var orderBook OrderResponse
	if err := json.Unmarshal(response.Body, &orderBook); err != nil {
		return "", OrderResponse{}, err
	}
	return string(response.Body), orderBook, nil
}

func (c *Client) ModifyMutliOrder(fyClient *Client) (string, ModifyMutliOrderRequest, error) {
	token := fmt.Sprintf("%s:%s", fyClient.appId, fyClient.accessToken)
	headers := http.Header{}
	headers.Add("Authorization", token)
	response, err := c.httpClient.DoRaw(http.MethodPost, MultipleOrderActionURL, nil, headers)
	if err != nil {
		return "", ModifyMutliOrderRequest{}, err
	}
	var orderBook ModifyMutliOrderRequest
	if err := json.Unmarshal(response.Body, &orderBook); err != nil {
		return "", ModifyMutliOrderRequest{}, err
	}
	return string(response.Body), orderBook, nil
}

func (c *Client) CancelOrder(fyClient *Client) (string, OrderResponse, error) {
	token := fmt.Sprintf("%s:%s", fyClient.appId, fyClient.accessToken)
	headers := http.Header{}
	headers.Add("Authorization", token)
	response, err := c.httpClient.DoRaw(http.MethodPost, MultiLegOrderURL, nil, headers)
	if err != nil {
		return "", OrderResponse{}, err
	}
	var orderBook OrderResponse
	if err := json.Unmarshal(response.Body, &orderBook); err != nil {
		return "", OrderResponse{}, err
	}
	return string(response.Body), orderBook, nil
}

func (c *Client) CancelMutliOrder(fyClient *Client) (string, ModifyMutliOrderRequest, error) {
	token := fmt.Sprintf("%s:%s", fyClient.appId, fyClient.accessToken)
	headers := http.Header{}
	headers.Add("Authorization", token)
	response, err := c.httpClient.DoRaw(http.MethodPost, MultiLegOrderURL, nil, headers)
	if err != nil {
		return "", ModifyMutliOrderRequest{}, err
	}
	var orderBook ModifyMutliOrderRequest
	if err := json.Unmarshal(response.Body, &orderBook); err != nil {
		return "", ModifyMutliOrderRequest{}, err
	}
	return string(response.Body), orderBook, nil
}

func (c *Client) ExitPosition(fyClient *Client) (string, OrderResponse, error) {
	token := fmt.Sprintf("%s:%s", fyClient.appId, fyClient.accessToken)
	headers := http.Header{}
	headers.Add("Authorization", token)
	response, err := c.httpClient.DoRaw(http.MethodPost, PositionURL, nil, headers)
	if err != nil {
		return "", OrderResponse{}, err
	}
	var orderBook OrderResponse
	if err := json.Unmarshal(response.Body, &orderBook); err != nil {
		return "", OrderResponse{}, err
	}
	return string(response.Body), orderBook, nil
}

func (c *Client) ExitPositionById(fyClient *Client) (string, ModifyMutliOrderRequest, error) {
	token := fmt.Sprintf("%s:%s", fyClient.appId, fyClient.accessToken)
	headers := http.Header{}
	headers.Add("Authorization", token)
	response, err := c.httpClient.DoRaw(http.MethodPost, PositionURL, nil, headers)
	if err != nil {
		return "", ModifyMutliOrderRequest{}, err
	}
	var orderBook ModifyMutliOrderRequest
	if err := json.Unmarshal(response.Body, &orderBook); err != nil {
		return "", ModifyMutliOrderRequest{}, err
	}
	return string(response.Body), orderBook, nil
}

func (c *Client) ExitPositionByProductType(fyClient *Client) (string, ModifyMutliOrderRequest, error) {
	token := fmt.Sprintf("%s:%s", fyClient.appId, fyClient.accessToken)
	headers := http.Header{}
	headers.Add("Authorization", token)
	response, err := c.httpClient.DoRaw(http.MethodPost, PositionURL, nil, headers)
	if err != nil {
		return "", ModifyMutliOrderRequest{}, err
	}
	var orderBook ModifyMutliOrderRequest
	if err := json.Unmarshal(response.Body, &orderBook); err != nil {
		return "", ModifyMutliOrderRequest{}, err
	}
	return string(response.Body), orderBook, nil
}

func (c *Client) CancelPendingOrders(fyClient *Client) (string, ModifyMutliOrderRequest, error) {
	token := fmt.Sprintf("%s:%s", fyClient.appId, fyClient.accessToken)
	headers := http.Header{}
	headers.Add("Authorization", token)
	response, err := c.httpClient.DoRaw(http.MethodPost, PositionURL, nil, headers)
	if err != nil {
		return "", ModifyMutliOrderRequest{}, err
	}
	var orderBook ModifyMutliOrderRequest
	if err := json.Unmarshal(response.Body, &orderBook); err != nil {
		return "", ModifyMutliOrderRequest{}, err
	}
	return string(response.Body), orderBook, nil
}

func (c *Client) ConvertPosition(fyClient *Client) (string, ModifyMutliOrderRequest, error) {
	token := fmt.Sprintf("%s:%s", fyClient.appId, fyClient.accessToken)
	headers := http.Header{}
	headers.Add("Authorization", token)
	response, err := c.httpClient.DoRaw(http.MethodPost, PositionURL, nil, headers)
	if err != nil {
		return "", ModifyMutliOrderRequest{}, err
	}
	var orderBook ModifyMutliOrderRequest
	if err := json.Unmarshal(response.Body, &orderBook); err != nil {
		return "", ModifyMutliOrderRequest{}, err
	}
	return string(response.Body), orderBook, nil
}
