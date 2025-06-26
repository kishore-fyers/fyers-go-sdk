package fyersgosdk

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"net/http"
)

func SetClientData(clientId, appId, appSecret, redirectUrl string) *Client {
	client := &Client{
		clientId:    clientId,
		appId:       appId,
		appSecret:   appSecret,
		redirectUrl: redirectUrl,
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

func (c *Client) GenerateAccessToken(authToken string, fyClient *Client) (string, AccessTokenResponse, error) {
	// Get SHA256 checksum
	h := sha256.New()
	h.Write([]byte(fyClient.appId + ":" + fyClient.appSecret))

	// Create JSON request body
	requestBody := fmt.Sprintf(`{"code":"%s","appIdHash":"%s","grant_type":"authorization_code"}`, authToken, fmt.Sprintf("%x", h.Sum(nil)))

	response, err := c.httpClient.DoRaw(http.MethodPost, ValidateAuthCodeURL, []byte(requestBody), nil)
	if err != nil {
		return "", AccessTokenResponse{}, err
	}

	// Parse the response
	var resp AccessTokenResponse
	if err := json.Unmarshal(response.Body, &resp); err != nil {
		return "", AccessTokenResponse{}, err
	}

	return string(response.Body), resp, nil
}
