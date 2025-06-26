package fyersgosdk

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

func (c *Client) GetHistory(fyClient *Client, historyRequest HistoryRequest) (string, HistoryResponse, error) {
	token := fmt.Sprintf("%s:%s", fyClient.appId, fyClient.accessToken)
	// url := fmt.Sprintf("%s%s", StockHistoryURL, url.Values{
	// 	"symbol":      {historyRequest.Symbol},
	// 	"resolution":  {historyRequest.Resolution},
	// 	"date_format": {historyRequest.DateFormat},
	// 	"range_from":  {historyRequest.RangeFrom},
	// 	"range_to":    {historyRequest.RangeTo},
	// 	"cont_flag":   {historyRequest.ContFlag},
	// })
	headers := http.Header{}
	headers.Add("Authorization", token)
	queryParams := url.Values{}
	queryParams.Add("symbol", historyRequest.Symbol)
	queryParams.Add("resolution", historyRequest.Resolution)
	queryParams.Add("date_format", historyRequest.DateFormat)
	queryParams.Add("range_from", historyRequest.RangeFrom)
	queryParams.Add("range_to", historyRequest.RangeTo)
	queryParams.Add("cont_flag", historyRequest.ContFlag)
	response, err := c.httpClient.Do(http.MethodGet, StockHistoryURL, queryParams, headers)
	if err != nil {
		return "", HistoryResponse{}, err
	}
	var historyResponse HistoryResponse
	if err := json.Unmarshal(response.Body, &historyResponse); err != nil {
		return "", HistoryResponse{}, err
	}
	return string(response.Body), historyResponse, nil
}

func (c *Client) GetStockQuotes(fyClient *Client, symbol string) (string, StockQuotesResponse, error) {
	token := fmt.Sprintf("%s:%s", fyClient.appId, fyClient.accessToken)

	headers := http.Header{}
	headers.Add("Authorization", token)
	queryParams := url.Values{}
	queryParams.Add("symbol", symbol)
	response, err := c.httpClient.DoRaw(http.MethodGet, StockQuotesURL, nil, headers)
	if err != nil {
		return "", StockQuotesResponse{}, err
	}
	var stockQuotesResponse StockQuotesResponse
	if err := json.Unmarshal(response.Body, &stockQuotesResponse); err != nil {
		return "", StockQuotesResponse{}, err
	}
	return string(response.Body), stockQuotesResponse, nil
}

func (c *Client) GetMarketDepth(fyClient *Client, marketDepthRequest MarketDepthRequest) (string, MarketDepthResponse, error) {
	token := fmt.Sprintf("%s:%s", fyClient.appId, fyClient.accessToken)
	headers := http.Header{}
	headers.Add("Authorization", token)
	queryParams := url.Values{}
	queryParams.Add("symbol", marketDepthRequest.Symbol)
	queryParams.Add("ohlcv_flag", marketDepthRequest.OHLCV)
	response, err := c.httpClient.DoRaw(http.MethodGet, MarketDepthURL, nil, headers)
	if err != nil {
		return "", MarketDepthResponse{}, err
	}
	var marketDepthResponse MarketDepthResponse
	if err := json.Unmarshal(response.Body, &marketDepthResponse); err != nil {
		return "", MarketDepthResponse{}, err
	}
	return string(response.Body), marketDepthResponse, nil
}

func (c *Client) GetOptionChain(fyClient *Client, optionChainRequest OptionChainRequest) (string, OptionChainResponse, error) {
	token := fmt.Sprintf("%s:%s", fyClient.appId, fyClient.accessToken)
	headers := http.Header{}
	headers.Add("Authorization", token)
	queryParams := url.Values{}
	queryParams.Add("symbol", optionChainRequest.Symbol)
	queryParams.Add("strikecount", optionChainRequest.StrikeCount)
	queryParams.Add("timestamp", optionChainRequest.Timestamp)
	response, err := c.httpClient.DoRaw(http.MethodGet, OptionChainURl, nil, headers)
	if err != nil {
		return "", OptionChainResponse{}, err
	}
	var optionChainResponse OptionChainResponse
	if err := json.Unmarshal(response.Body, &optionChainResponse); err != nil {
		return "", OptionChainResponse{}, err
	}
	return string(response.Body), optionChainResponse, nil
}
