package fyersgosdk

import (
	"crypto/sha256"
	"fmt"
	"net/http"
)

func SetClientData(clientId, appId, appSecret, redirectUrl, pin string) *Client {
	client := &Client{
		clientId:    clientId,
		appId:       appId,
		appSecret:   appSecret,
		redirectUrl: redirectUrl,
		pin:         pin,
		httpClient:  NewHTTPClient(nil, nil, false),
	}

	// Create a default http handler with default timeout.
	// client.SetHTTPClient(&http.Client{
	// 	Timeout: requestTimeout,
	// })

	return client
}

// func NewClient(options ClientOptions) *Client {
// 		httpClient := NewHTTPClient(options.HTTPClient, options.Logger, options.Debug)

// }

func (c *Client) SetAccessToken(accessToken string) *Client {
	c.accessToken = accessToken
	return c
}

func (c *Client) SetRefreshToken(refreshToken string) *Client {
	c.refreshToken = refreshToken
	return c
}
func (c *Client) GetLoginURL() string {
	return fmt.Sprintf("%s&client_id=%s&redirect_uri=%s&response_type=%s&state=%s", GenerateAuthCodeURL, c.appId, c.redirectUrl, "code", "sample_state")
}

// NewFyersModel creates an API client with client_id and access token.
// Use this for all API calls after obtaining the access token via Client.GenerateAccessToken.
func NewFyersModel(appId, accessToken string) *FyersModel {
	return &FyersModel{
		appId:       appId,
		accessToken: accessToken,
		httpClient:  NewHTTPClient(nil, nil, false),
	}
}

func (m *FyersModel) authHeader() http.Header {
	h := http.Header{}
	h.Set("Authorization", fmt.Sprintf("%s:%s", m.appId, m.accessToken))
	return h
}

func (c *Client) GenerateAccessToken(authToken string, fyClient *Client) (string, error) {
	// Get SHA256 checksum
	h := sha256.New()
	h.Write([]byte(fyClient.appId + ":" + fyClient.appSecret))

	// Create JSON request body
	requestBody := fmt.Sprintf(`{"code":"%s","appIdHash":"%s","grant_type":"authorization_code"}`, authToken, fmt.Sprintf("%x", h.Sum(nil)))

	response, err := c.httpClient.DoRaw(http.MethodPost, ValidateAuthCodeURL, []byte(requestBody), nil)
	if err != nil {
		return "", err
	}

	return string(response.Body), nil
}

func (c *Client) GenerateAccessTokenFromRefreshToken(fyClient *Client) (string, error) {
	// Get SHA256 checksum
	h := sha256.New()
	h.Write([]byte(fyClient.appId + ":" + fyClient.appSecret))

	// Create JSON request body
	requestBody := fmt.Sprintf(`{"refresh_token":"%s","appIdHash":"%s","grant_type":"refresh_token","pin":"%s"}`, fyClient.refreshToken, fmt.Sprintf("%x", h.Sum(nil)), fyClient.pin)

	response, err := c.httpClient.DoRaw(http.MethodPost, ValidateAuthCodeURL, []byte(requestBody), nil)
	if err != nil {
		return "", err
	}

	return string(response.Body), nil
}
