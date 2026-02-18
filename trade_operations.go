package fyersgosdk

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (m *FyersModel) ModifyOrder(orderRequest ModifyOrderRequest) (string, error) {
	body, err := json.Marshal(orderRequest)
	if err != nil {
		return "", fmt.Errorf("marshal order request: %w", err)
	}
	resp, err := m.httpClient.DoRaw(http.MethodPatch, SingleOrderActionURL, body, m.authHeader())
	if err != nil {
		return "", err
	}
	return string(resp.Body), nil
}

func (m *FyersModel) ModifyMutliOrder(requests []ModifyMultiOrderItem) (string, error) {
	if len(requests) == 0 {
		return "", fmt.Errorf("at least one order modification required")
	}
	body, err := json.Marshal(requests)
	if err != nil {
		return "", fmt.Errorf("marshal multi order request: %w", err)
	}
	headers := m.authHeader()
	headers.Set("Content-Type", "application/json")
	resp, err := m.httpClient.DoRaw(http.MethodPatch, MultipleOrderActionURL, body, headers)
	if err != nil {
		return "", err
	}
	return string(resp.Body), nil
}

func (m *FyersModel) CancelOrder() (string, error) {
	resp, err := m.httpClient.DoRaw(http.MethodDelete, OrderBookURL, []byte("{}"), m.authHeader())
	if err != nil {
		return "", err
	}
	return string(resp.Body), nil
}

func (m *FyersModel) CancelMutliOrder() (string, error) {
	resp, err := m.httpClient.DoRaw(http.MethodPost, MultiLegOrderURL, nil, m.authHeader())
	if err != nil {
		return "", err
	}
	return string(resp.Body), nil
}

func (m *FyersModel) ExitPosition() (string, error) {
	body, _ := json.Marshal(map[string]int{"exit_all": 1})
	resp, err := m.httpClient.DoRaw(http.MethodDelete, PositionURL, body, m.authHeader())
	if err != nil {
		return "", err
	}
	return string(resp.Body), nil
}

func (m *FyersModel) ExitPositionById() (string, error) {
	resp, err := m.httpClient.DoRaw(http.MethodPost, PositionURL, nil, m.authHeader())
	if err != nil {
		return "", err
	}
	return string(resp.Body), nil
}

func (m *FyersModel) ExitPositionByProductType() (string, error) {
	resp, err := m.httpClient.DoRaw(http.MethodPost, PositionURL, nil, m.authHeader())
	if err != nil {
		return "", err
	}
	return string(resp.Body), nil
}

func (m *FyersModel) CancelPendingOrders() (string, error) {
	resp, err := m.httpClient.DoRaw(http.MethodPost, PositionURL, nil, m.authHeader())
	if err != nil {
		return "", err
	}
	return string(resp.Body), nil
}

func (m *FyersModel) ConvertPosition() (string, error) {
	resp, err := m.httpClient.DoRaw(http.MethodPost, PositionURL, nil, m.authHeader())
	if err != nil {
		return "", err
	}
	return string(resp.Body), nil
}
