package fyersgosdk

import (
	"net/http"
	"net/url"
	"strconv"
)

func (m *FyersModel) GetHistory(historyRequest HistoryRequest) (string, error) {
	params := url.Values{}
	params.Set("symbol", historyRequest.Symbol)
	params.Set("resolution", historyRequest.Resolution)
	params.Set("date_format", historyRequest.DateFormat)
	params.Set("range_from", historyRequest.RangeFrom)
	params.Set("range_to", historyRequest.RangeTo)
	params.Set("cont_flag", historyRequest.ContFlag)
	resp, err := m.httpClient.Do(http.MethodGet, StockHistoryURL, params, m.authHeader())
	if err != nil {
		return "", err
	}
	return string(resp.Body), nil
}

func (m *FyersModel) GetStockQuotes(symbol string) (string, error) {
	params := url.Values{}
	params.Set("symbols", symbol)
	resp, err := m.httpClient.Do(http.MethodGet, StockQuotesURL, params, m.authHeader())
	if err != nil {
		return "", err
	}
	return string(resp.Body), nil
}

func (m *FyersModel) GetMarketDepth(marketDepthRequest MarketDepthRequest) (string, error) {
	params := url.Values{}
	params.Set("symbol", marketDepthRequest.Symbol)
	params.Set("ohlcv_flag", marketDepthRequest.OHLCV)
	resp, err := m.httpClient.Do(http.MethodGet, MarketDepthURL, params, m.authHeader())
	if err != nil {
		return "", err
	}
	return string(resp.Body), nil
}

func (m *FyersModel) GetOptionChain(optionChainRequest OptionChainRequest) (string, error) {
	params := url.Values{}
	params.Set("symbol", optionChainRequest.Symbol)
	params.Set("strikecount", strconv.Itoa(optionChainRequest.StrikeCount))
	params.Set("timestamp", optionChainRequest.Timestamp)
	resp, err := m.httpClient.Do(http.MethodGet, OptionChainURl, params, m.authHeader())
	if err != nil {
		return "", err
	}
	return string(resp.Body), nil
}
