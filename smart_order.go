package fyersgosdk

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

// CreateSmartOrderLimit creates a smart order limit.
// Returns the API response body as string or an error.
func (m *FyersModel) CreateSmartOrderLimit(req CreateSmartOrderLimitRequest) (string, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return "", fmt.Errorf("marshal create smart order limit request: %w", err)
	}
	headers := m.authHeader()
	headers.Set("Content-Type", "application/json")
	resp, err := m.httpClient.DoRaw(http.MethodPost, CreateSmartorderLimit, body, headers)
	if err != nil {
		return "", err
	}
	return string(resp.Body), nil
}

// CreateSmartOrderStep creates a smart order step.
// Returns the API response body as string or an error.
func (m *FyersModel) CreateSmartOrderStep(req CreateSmartOrderStepRequest) (string, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return "", fmt.Errorf("marshal create smart order step request: %w", err)
	}
	headers := m.authHeader()
	headers.Set("Content-Type", "application/json")
	resp, err := m.httpClient.DoRaw(http.MethodPost, CreateSmartorderStep, body, headers)
	if err != nil {
		return "", err
	}
	return string(resp.Body), nil
}

// CreateSmartOrderSIP creates a smart order SIP.
// Returns the API response body as string or an error.
func (m *FyersModel) CreateSmartOrderSIP(req CreateSmartOrderSIPRequest) (string, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return "", fmt.Errorf("marshal create smart order SIP request: %w", err)
	}
	headers := m.authHeader()
	headers.Set("Content-Type", "application/json")
	resp, err := m.httpClient.DoRaw(http.MethodPost, CreateSmartorderSip, body, headers)
	if err != nil {
		return "", err
	}
	return string(resp.Body), nil
}

// CreateSmartOrderTrail creates a smart order trail.
// Returns the API response body as string or an error.
func (m *FyersModel) CreateSmartOrderTrail(req CreateSmartOrderTrailRequest) (string, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return "", fmt.Errorf("marshal create smart order trail request: %w", err)
	}
	headers := m.authHeader()
	headers.Set("Content-Type", "application/json")
	resp, err := m.httpClient.DoRaw(http.MethodPost, CreateSmartorderTrail, body, headers)
	if err != nil {
		return "", err
	}
	return string(resp.Body), nil
}

// ModifySmartOrder modifies a smart order. Pass the flowId (or "id" from create response).
// PATCH /smart-order/modify with body {"flowId": "...", ...optional fields}.
// Returns the API response body as string or an error.
func (m *FyersModel) ModifySmartOrder(req ModifySmartOrderRequest) (string, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return "", fmt.Errorf("marshal modify smart order request: %w", err)
	}
	headers := m.authHeader()
	headers.Set("Content-Type", "application/json")
	resp, err := m.httpClient.DoRaw(http.MethodPatch, ModifySmartorder, body, headers)
	if err != nil {
		return "", err
	}
	return string(resp.Body), nil
}

// CancelSmartOrder cancels a smart order. Pass the flowId (or "id" from create response).
// DELETE /smart-order/cancel with body {"flowId": "..."}.
func (m *FyersModel) CancelSmartOrder(req FlowIdRequest) (string, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return "", fmt.Errorf("marshal cancel smart order request: %w", err)
	}
	headers := m.authHeader()
	headers.Set("Content-Type", "application/json")
	resp, err := m.httpClient.DoRaw(http.MethodDelete, CancelSmartorder, body, headers)
	if err != nil {
		return "", err
	}
	return string(resp.Body), nil
}

// PauseSmartOrder pauses a smart order. Pass the flowId (or "id" from create response).
// PATCH /smart-order/pause with body {"flowId": "..."}.
func (m *FyersModel) PauseSmartOrder(req FlowIdRequest) (string, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return "", fmt.Errorf("marshal pause smart order request: %w", err)
	}
	headers := m.authHeader()
	headers.Set("Content-Type", "application/json")
	resp, err := m.httpClient.DoRaw(http.MethodPatch, PauseSmartorder, body, headers)
	if err != nil {
		return "", err
	}
	return string(resp.Body), nil
}

// ResumeSmartOrder resumes a smart order. Pass the flowId (or "id" from create response).
// PATCH /smart-order/resume with body {"flowId": "..."}.
func (m *FyersModel) ResumeSmartOrder(req FlowIdRequest) (string, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return "", fmt.Errorf("marshal resume smart order request: %w", err)
	}
	headers := m.authHeader()
	headers.Set("Content-Type", "application/json")
	resp, err := m.httpClient.DoRaw(http.MethodPatch, ResumeSmartorder, body, headers)
	if err != nil {
		return "", err
	}
	return string(resp.Body), nil
}

// GetSmartOrderBookWithFilter returns the smart order book with optional filters.
// Pass nil for no filters. Returns the API response body as string or an error.
func (m *FyersModel) GetSmartOrderBookWithFilter(req *GetSmartOrderBookFilter) (string, error) {
	params := url.Values{}
	if req != nil {
		for _, e := range req.Exchange {
			params.Add("exchange", e)
		}
		for _, s := range req.Side {
			params.Add("side", strconv.Itoa(s))
		}
		for _, f := range req.Flowtype {
			params.Add("flowtype", strconv.Itoa(f))
		}
		for _, p := range req.Product {
			params.Add("product", p)
		}
		for _, mt := range req.MessageType {
			params.Add("messageType", strconv.Itoa(mt))
		}
		if req.Search != "" {
			params.Set("search", req.Search)
		}
		if req.SortBy != "" {
			params.Set("sort_by", req.SortBy)
		}
		if req.OrdBy != 0 {
			params.Set("ord_by", strconv.Itoa(req.OrdBy))
		}
		if req.PageNo != 0 {
			params.Set("page_no", strconv.Itoa(req.PageNo))
		}
		if req.PageSize != 0 {
			params.Set("page_size", strconv.Itoa(req.PageSize))
		}
	}
	resp, err := m.httpClient.Do(http.MethodGet, SmartorderOrderBook, params, m.authHeader())
	if err != nil {
		return "", err
	}
	return string(resp.Body), nil
}

// CreateSmartExitTrigger creates a smart exit trigger.
// Type: 1 = Only Alert, 2 = Exit with Alert, 3 = Exit with Alert + Wait for Recovery. WaitTime is required for type 3.
// Returns the API response body as string or an error.
func (m *FyersModel) CreateSmartExitTrigger(req CreateSmartExitTriggerRequest) (string, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return "", fmt.Errorf("marshal create smart exit trigger request: %w", err)
	}
	headers := m.authHeader()
	headers.Set("Content-Type", "application/json")
	resp, err := m.httpClient.DoRaw(http.MethodPost, SmartExitTrigger, body, headers)
	if err != nil {
		return "", err
	}
	return string(resp.Body), nil
}

// GetSmartExitTrigger returns smart exit trigger details. Pass nil or filter with FlowId set for a specific trigger.
// Returns the API response body as string or an error.
func (m *FyersModel) GetSmartExitTrigger(req *GetSmartExitTriggerFilter) (string, error) {
	params := url.Values{}
	if req != nil && req.FlowId != "" {
		params.Set("flowId", req.FlowId)
	}
	resp, err := m.httpClient.Do(http.MethodGet, SmartExitTrigger, params, m.authHeader())
	if err != nil {
		return "", err
	}
	return string(resp.Body), nil
}

// UpdateSmartExitTrigger updates a smart exit trigger by flowId.
// Returns the API response body as string or an error.
func (m *FyersModel) UpdateSmartExitTrigger(req UpdateSmartExitTriggerRequest) (string, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return "", fmt.Errorf("marshal update smart exit trigger request: %w", err)
	}
	headers := m.authHeader()
	headers.Set("Content-Type", "application/json")
	resp, err := m.httpClient.DoRaw(http.MethodPut, SmartExitTrigger, body, headers)
	if err != nil {
		return "", err
	}
	return string(resp.Body), nil
}

// ActivateDeactivateSmartExitTrigger activates or deactivates a smart exit trigger by flowId.
// Returns the API response body as string or an error.
func (m *FyersModel) ActivateDeactivateSmartExitTrigger(req FlowIdRequest) (string, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return "", fmt.Errorf("marshal activate/deactivate smart exit trigger request: %w", err)
	}
	headers := m.authHeader()
	headers.Set("Content-Type", "application/json")
	resp, err := m.httpClient.DoRaw(http.MethodPost, ActivateSmartExitTrigger, body, headers)
	if err != nil {
		return "", err
	}
	return string(resp.Body), nil
}
