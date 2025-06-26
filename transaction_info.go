package fyersgosdk

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) GetOrderBook(fyClient *Client) (string, OrderBook, error) {
	token := fmt.Sprintf("%s:%s", fyClient.appId, fyClient.accessToken)
	headers := http.Header{}
	headers.Add("Authorization", token)
	response, err := c.httpClient.DoRaw(http.MethodGet, OrderBookURL, nil, headers)
	if err != nil {
		return "", OrderBook{}, err
	}
	var orderBook OrderBook
	if err := json.Unmarshal(response.Body, &orderBook); err != nil {
		return "", OrderBook{}, err
	}
	return string(response.Body), orderBook, nil
}

func (c *Client) GetOrderBookByTag(fyClient *Client, tag string) (string, OrderBook, error) {
	token := fmt.Sprintf("%s:%s", fyClient.appId, fyClient.accessToken)
	headers := http.Header{}
	headers.Add("Authorization", token)
	response, err := c.httpClient.DoRaw(http.MethodGet, OrdersByTagURL, nil, headers)
	if err != nil {
		return "", OrderBook{}, err
	}
	var orderBook OrderBook
	if err := json.Unmarshal(response.Body, &orderBook); err != nil {
		return "", OrderBook{}, err
	}
	return string(response.Body), orderBook, nil
}

func (c *Client) GetOrderById(fyClient *Client, id string) (string, OrderBook, error) {
	token := fmt.Sprintf("%s:%s", fyClient.appId, fyClient.accessToken)
	headers := http.Header{}
	headers.Add("Authorization", token)
	url := OrderByIdURL + id
	response, err := c.httpClient.DoRaw(http.MethodGet, url, nil, headers)
	if err != nil {
		return "", OrderBook{}, err
	}
	var orderBook OrderBook
	if err := json.Unmarshal(response.Body, &orderBook); err != nil {
		return "", OrderBook{}, err
	}
	return string(response.Body), orderBook, nil
}

func (c *Client) GetPositions(fyClient *Client) (string, Position, error) {
	token := fmt.Sprintf("%s:%s", fyClient.appId, fyClient.accessToken)
	headers := http.Header{}
	headers.Add("Authorization", token)
	response, err := c.httpClient.DoRaw(http.MethodGet, OrderCheckMarginURL, nil, headers)
	if err != nil {
		return "", Position{}, err
	}
	var orderCheckMargin Position
	if err := json.Unmarshal(response.Body, &orderCheckMargin); err != nil {
		return "", Position{}, err
	}
	return string(response.Body), orderCheckMargin, nil
}

func (c *Client) GetTradeBook(fyClient *Client) (string, TradeBook, error) {
	token := fmt.Sprintf("%s:%s", fyClient.appId, fyClient.accessToken)
	headers := http.Header{}
	headers.Add("Authorization", token)
	response, err := c.httpClient.DoRaw(http.MethodGet, TradeBookURL, nil, headers)
	if err != nil {
		return "", TradeBook{}, err
	}
	var tradeBook TradeBook
	if err := json.Unmarshal(response.Body, &tradeBook); err != nil {
		return "", TradeBook{}, err
	}
	return string(response.Body), tradeBook, nil
}

func (c *Client) GetTradeBookByTag(fyClient *Client, tag string) (string, TradeBook, error) {
	token := fmt.Sprintf("%s:%s", fyClient.appId, fyClient.accessToken)
	headers := http.Header{}
	headers.Add("Authorization", token)
	url := TradeBookByTagURL + tag
	response, err := c.httpClient.DoRaw(http.MethodGet, url, nil, headers)
	if err != nil {
		return "", TradeBook{}, err
	}
	var tradeBook TradeBook
	if err := json.Unmarshal(response.Body, &tradeBook); err != nil {
		return "", TradeBook{}, err
	}
	return string(response.Body), tradeBook, nil
}
