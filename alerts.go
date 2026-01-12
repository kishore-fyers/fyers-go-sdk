package fyersgosdk

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) GetAlerts(fyClient *Client) (string, AlertsResponse, error) {
	token := fmt.Sprintf("%s:%s", fyClient.appId, fyClient.accessToken)
	headers := http.Header{}
	headers.Add("Authorization", token)
	response, err := c.httpClient.DoRaw(http.MethodGet, AlertsURL, nil, headers)
	if err != nil {
		return "", AlertsResponse{}, err
	}

	var alertsResponse AlertsResponse
	if err := json.Unmarshal(response.Body, &alertsResponse); err != nil {
		return string(response.Body), AlertsResponse{}, err
	}
	return string(response.Body), alertsResponse, nil
}

func (c *Client) ToggleAlert(fyClient *Client, alertId string) (string, APIResponse, error) {
	token := fmt.Sprintf("%s:%s", fyClient.appId, fyClient.accessToken)
	headers := http.Header{}
	headers.Add("Authorization", token)
	headers.Add("Content-Type", "application/json")

	body := map[string]string{"alertId": alertId, "agent": "fyers-api"}
	jsonBody, _ := json.Marshal(body)

	response, err := c.httpClient.DoRaw(http.MethodPut, ToggleAlertURL, jsonBody, headers)
	if err != nil {
		return "", APIResponse{}, err
	}
	var apiResponse APIResponse
	if err := json.Unmarshal(response.Body, &apiResponse); err != nil {
		return "", APIResponse{}, err
	}
	return string(response.Body), apiResponse, nil
}

func (c *Client) CreateAlert(fyClient *Client, alertRequest AlertRequest) (string, APIResponse, error) {
	token := fmt.Sprintf("%s:%s", fyClient.appId, fyClient.accessToken)
	headers := http.Header{}
	headers.Add("Authorization", token)
	headers.Add("Content-Type", "application/json")

	if alertRequest.Agent == "" {
		alertRequest.Agent = "fyers-api"
	}

	jsonBody, _ := json.Marshal(alertRequest)

	response, err := c.httpClient.DoRaw(http.MethodPost, AlertsURL, jsonBody, headers)
	if err != nil {
		return "", APIResponse{}, err
	}
	var apiResponse APIResponse
	if err := json.Unmarshal(response.Body, &apiResponse); err != nil {
		return "", APIResponse{}, err
	}
	return string(response.Body), apiResponse, nil
}

func (c *Client) DeleteAlert(fyClient *Client, alertId string) (string, APIResponse, error) {
	token := fmt.Sprintf("%s:%s", fyClient.appId, fyClient.accessToken)
	headers := http.Header{}
	headers.Add("Authorization", token)
	headers.Add("Content-Type", "application/json")

	body := map[string]string{"alertId": alertId, "agent": "fyers-api"}
	jsonBody, _ := json.Marshal(body)

	response, err := c.httpClient.DoRaw(http.MethodDelete, AlertsURL, jsonBody, headers)
	if err != nil {
		return "", APIResponse{}, err
	}
	var apiResponse APIResponse
	if err := json.Unmarshal(response.Body, &apiResponse); err != nil {
		return "", APIResponse{}, err
	}
	return string(response.Body), apiResponse, nil
}

func (c *Client) UpdateAlert(fyClient *Client, alertId string, alertRequest AlertRequest) (string, APIResponse, error) {
	token := fmt.Sprintf("%s:%s", fyClient.appId, fyClient.accessToken)
	headers := http.Header{}
	headers.Add("Authorization", token)
	headers.Add("Content-Type", "application/json")

	if alertRequest.Agent == "" {
		alertRequest.Agent = "fyers-api"
	}

	type updateRequest struct {
		AlertId string `json:"alertId"`
		AlertRequest
	}
	req := updateRequest{AlertId: alertId, AlertRequest: alertRequest}
	jsonBody, _ := json.Marshal(req)

	response, err := c.httpClient.DoRaw(http.MethodPut, AlertsURL, jsonBody, headers)
	if err != nil {
		return "", APIResponse{}, err
	}
	var apiResponse APIResponse
	if err := json.Unmarshal(response.Body, &apiResponse); err != nil {
		return "", APIResponse{}, err
	}
	return string(response.Body), apiResponse, nil
}
