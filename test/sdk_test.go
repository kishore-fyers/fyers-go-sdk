package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	fyersgosdk "github.com/kishore-fyers/fyers-go-sdk"
)

// Test credentials from env (optional). If not set, tests that require API calls will skip or expect auth errors.
func getTestCreds() (appId, appSecret, redirectUrl, accessToken string) {
	appId = os.Getenv("FYERS_APP_ID")
	appSecret = os.Getenv("FYERS_APP_SECRET")
	redirectUrl = os.Getenv("FYERS_REDIRECT_URL")
	accessToken = os.Getenv("FYERS_ACCESS_TOKEN")
	return
}

func TestSetClientDataAndGetLoginURL(t *testing.T) {
	client := fyersgosdk.SetClientData("test-app", "test-secret", "https://example.com/callback")
	url := client.GetLoginURL()
	if url == "" {
		t.Error("GetLoginURL() should return non-empty URL")
	}
}

func TestNewFyersModel(t *testing.T) {
	_ = fyersgosdk.NewFyersModel("app", "token")
	// no panic
}

func TestGenerateAccessToken_InvalidCode(t *testing.T) {
	client := fyersgosdk.SetClientData("app", "secret", "https://example.com")
	_, err := client.GenerateAccessToken("invalid-code", client)
	if err == nil {
		t.Log("GenerateAccessToken with invalid code should error (may pass if network/API returns error)")
	}
}

func TestGenerateAccessTokenFromRefreshToken_InvalidToken(t *testing.T) {
	client := fyersgosdk.SetClientData("app", "secret", "https://example.com")
	_, err := client.GenerateAccessTokenFromRefreshToken("invalid-refresh", "0000", client)
	if err == nil {
		t.Log("GenerateAccessTokenFromRefreshToken with invalid token should error")
	}
}

func TestGetProfile(t *testing.T) {
	m := fyersgosdk.NewFyersModel("", "")
	_, _ = m.GetProfile() // no panic; may return err or API error body
}

func TestGetFunds(t *testing.T) {
	m := fyersgosdk.NewFyersModel("", "")
	_, _ = m.GetFunds()
}

func TestGetHoldings(t *testing.T) {
	m := fyersgosdk.NewFyersModel("", "")
	_, _ = m.GetHoldings()
}

func TestLogout(t *testing.T) {
	m := fyersgosdk.NewFyersModel("", "")
	_, _ = m.Logout()
}

func TestGetOrderBook(t *testing.T) {
	m := fyersgosdk.NewFyersModel("", "")
	_, _ = m.GetOrderBook()
}

func TestGetOrderBookByTag(t *testing.T) {
	m := fyersgosdk.NewFyersModel("", "")
	_, _ = m.GetOrderBookByTag("tag")
}

func TestGetOrderById(t *testing.T) {
	m := fyersgosdk.NewFyersModel("", "")
	_, _ = m.GetOrderById("1")
}

func TestGetPositions(t *testing.T) {
	m := fyersgosdk.NewFyersModel("", "")
	_, _ = m.GetPositions()
}

func TestGetTradeBook(t *testing.T) {
	m := fyersgosdk.NewFyersModel("", "")
	_, _ = m.GetTradeBook()
}

func TestGetTradeBookByTag(t *testing.T) {
	m := fyersgosdk.NewFyersModel("", "")
	_, _ = m.GetTradeBookByTag("tag")
}

func TestSingleOrderAction(t *testing.T) {
	m := fyersgosdk.NewFyersModel("", "")
	_, _ = m.SingleOrderAction(fyersgosdk.OrderRequest{
		Symbol: "NSE:SBIN-EQ", Qty: 1, Type: 1, Side: 1, ProductType: "CNC", Validity: "DAY",
	})
}

func TestMultiOrderAction(t *testing.T) {
	m := fyersgosdk.NewFyersModel("", "")
	_, _ = m.MultiOrderAction([]fyersgosdk.OrderRequest{{Symbol: "NSE:SBIN-EQ", Qty: 1, Type: 1, Side: 1, ProductType: "CNC", Validity: "DAY"}})
}

func TestModifyOrder(t *testing.T) {
	m := fyersgosdk.NewFyersModel("", "")
	_, _ = m.ModifyOrder(fyersgosdk.ModifyOrderRequest{Id: "1", Qty: 1, Type: 1, Side: 1})
}

func TestModifyMutliOrder_EmptySlice(t *testing.T) {
	m := fyersgosdk.NewFyersModel("", "")
	_, err := m.ModifyMutliOrder([]fyersgosdk.ModifyMultiOrderItem{})
	if err == nil {
		t.Error("ModifyMutliOrder with empty slice should error")
	}
}

func TestCancelOrder(t *testing.T) {
	m := fyersgosdk.NewFyersModel("", "")
	_, _ = m.CancelOrder("1")
}

func TestCancelMutliOrder_EmptySlice(t *testing.T) {
	m := fyersgosdk.NewFyersModel("", "")
	_, err := m.CancelMutliOrder([]string{})
	if err == nil {
		t.Error("CancelMutliOrder with empty slice should error")
	}
}

func TestGTTSingleOrderAction(t *testing.T) {
	m := fyersgosdk.NewFyersModel("", "")
	_, _ = m.GTTSingleOrderAction(fyersgosdk.GTTOrderRequest{
		Side: 1, Symbol: "NSE:SBIN-EQ", ProductType: "CNC",
		OrderInfo: fyersgosdk.OrderInfo{Leg1: fyersgosdk.Leg1{Price: 100, TriggerPrice: 100, Qty: 1}},
	})
}

func TestGTTMultiOrderAction_EmptySlice(t *testing.T) {
	m := fyersgosdk.NewFyersModel("", "")
	_, err := m.GTTMultiOrderAction([]fyersgosdk.GTTOrderRequest{})
	if err == nil {
		t.Error("GTTMultiOrderAction with empty slice should error")
	}
}

func TestModifyGTTOrder_EmptySlice(t *testing.T) {
	m := fyersgosdk.NewFyersModel("", "")
	_, err := m.ModifyGTTOrder([]fyersgosdk.ModifyGTTOrderRequest{})
	if err == nil {
		t.Error("ModifyGTTOrder with empty slice should error")
	}
}

func TestCancelGTTOrder(t *testing.T) {
	m := fyersgosdk.NewFyersModel("", "")
	_, _ = m.CancelGTTOrder("1")
}

func TestGetGTTOrderBook(t *testing.T) {
	m := fyersgosdk.NewFyersModel("", "")
	_, _ = m.GetGTTOrderBook()
}

func TestExitPosition(t *testing.T) {
	m := fyersgosdk.NewFyersModel("", "")
	_, _ = m.ExitPosition()
}

func TestExitPositionById_EmptySlice(t *testing.T) {
	m := fyersgosdk.NewFyersModel("", "")
	_, err := m.ExitPositionById([]string{})
	if err == nil {
		t.Error("ExitPositionById with empty slice should error")
	}
}

func TestExitPositionByProductType(t *testing.T) {
	m := fyersgosdk.NewFyersModel("", "")
	_, _ = m.ExitPositionByProductType(fyersgosdk.ExitPositionByProductTypeRequest{
		Segment: []int{10}, Side: []int{1, -1}, ProductType: []string{"INTRADAY"},
	})
}

func TestCancelPendingOrders(t *testing.T) {
	m := fyersgosdk.NewFyersModel("", "")
	_, _ = m.CancelPendingOrders(fyersgosdk.CancelPendingOrdersRequest{PendingOrdersCancel: 1})
}

func TestConvertPosition(t *testing.T) {
	m := fyersgosdk.NewFyersModel("", "")
	_, _ = m.ConvertPosition(fyersgosdk.ConvertPositionRequest{
		Symbol: "NSE:SBIN-EQ", PositionSide: 1, ConvertQty: 1, ConvertFrom: "INTRADAY", ConvertTo: "CNC", Overnight: 1,
	})
}

func TestGetMarketStatus(t *testing.T) {
	m := fyersgosdk.NewFyersModel("", "")
	_, _ = m.GetMarketStatus()
}

func TestGetHistory(t *testing.T) {
	m := fyersgosdk.NewFyersModel("", "")
	_, _ = m.GetHistory(fyersgosdk.HistoryRequest{
		Symbol: "NSE:SBIN-EQ", Resolution: "30", DateFormat: "1", RangeFrom: "2021-01-01", RangeTo: "2021-01-02",
	})
}

func TestGetStockQuotes_EmptySymbols(t *testing.T) {
	m := fyersgosdk.NewFyersModel("", "")
	_, err := m.GetStockQuotes([]string{})
	if err == nil {
		t.Error("GetStockQuotes with empty symbols should error")
	}
}

func TestGetStockQuotes_MultipleSymbols(t *testing.T) {
	m := fyersgosdk.NewFyersModel("", "")
	_, _ = m.GetStockQuotes([]string{"NSE:SBIN-EQ", "NSE:NIFTY50-INDEX"})
}

func TestGetMarketDepth(t *testing.T) {
	m := fyersgosdk.NewFyersModel("", "")
	_, _ = m.GetMarketDepth(fyersgosdk.MarketDepthRequest{Symbol: "NSE:SBIN-EQ", OHLCV: "1"})
}

func TestGetOptionChain(t *testing.T) {
	m := fyersgosdk.NewFyersModel("", "")
	_, _ = m.GetOptionChain(fyersgosdk.OptionChainRequest{Symbol: "NSE:SBIN-EQ", StrikeCount: 10})
}

func TestGetAlerts(t *testing.T) {
	m := fyersgosdk.NewFyersModel("", "")
	_, _ = m.GetAlerts()
}

func TestCreateAlert(t *testing.T) {
	m := fyersgosdk.NewFyersModel("", "")
	_, _ = m.CreateAlert(fyersgosdk.AlertRequest{})
}

func TestUpdateAlert(t *testing.T) {
	m := fyersgosdk.NewFyersModel("", "")
	_, _ = m.UpdateAlert("1", fyersgosdk.AlertRequest{})
}

func TestDeleteAlert(t *testing.T) {
	m := fyersgosdk.NewFyersModel("", "")
	_, _ = m.DeleteAlert("1")
}

func TestToggleAlert(t *testing.T) {
	m := fyersgosdk.NewFyersModel("", "")
	_, _ = m.ToggleAlert("1")
}

// Integration-style: run with real credentials via env to hit API (optional)
func TestGetProfile_WithToken(t *testing.T) {
	appId, _, _, accessToken := getTestCreds()
	if appId == "" || accessToken == "" {
		t.Skip("FYERS_APP_ID and FYERS_ACCESS_TOKEN not set; skipping integration test")
	}
	m := fyersgosdk.NewFyersModel(appId, accessToken)
	resp, err := m.GetProfile()
	if err != nil {
		t.Logf("GetProfile (with token): %v", err)
		return
	}
	if resp == "" {
		t.Error("GetProfile response should be non-empty")
	}
}

// --- JSON-driven test runner: reads test_cases.json, runs each case, writes test_output.json ---

type testCaseFile struct {
	Config struct {
		AppId       string `json:"appId"`
		AppSecret   string `json:"appSecret"`
		RedirectUrl string `json:"redirectUrl"`
		AccessToken string `json:"accessToken"`
	} `json:"config"`
	ValidationCases []string `json:"validationCases"` // case names to run when validating with token
	Cases           []struct {
		Name  string          `json:"name"`
		Input json.RawMessage `json:"input"`
	} `json:"cases"`
}

type caseResult struct {
	Function string          `json:"function"`
	Input    json.RawMessage `json:"input"`
	Output   string          `json:"output"`
	Error    string          `json:"error,omitempty"`
}

func TestRunAllFromJSON(t *testing.T) {
	testDir := "test"
	if _, err := os.Stat(testDir); os.IsNotExist(err) {
		testDir = "."
	}
	path := filepath.Join(testDir, "test_cases.json")
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read test_cases.json: %v", err)
	}
	var file testCaseFile
	if err := json.Unmarshal(data, &file); err != nil {
		t.Fatalf("parse test_cases.json: %v", err)
	}
	appId := file.Config.AppId
	if e := os.Getenv("FYERS_APP_ID"); e != "" {
		appId = e
	}
	if appId == "" {
		appId = "test-app"
	}
	appSecret := file.Config.AppSecret
	if e := os.Getenv("FYERS_APP_SECRET"); e != "" {
		appSecret = e
	}
	if appSecret == "" {
		appSecret = "test-secret"
	}
	redirectUrl := file.Config.RedirectUrl
	if e := os.Getenv("FYERS_REDIRECT_URL"); e != "" {
		redirectUrl = e
	}
	if redirectUrl == "" {
		redirectUrl = "https://example.com/callback"
	}
	accessToken := file.Config.AccessToken
	if e := os.Getenv("FYERS_ACCESS_TOKEN"); e != "" {
		accessToken = e
	}
	client := fyersgosdk.SetClientData(appId, appSecret, redirectUrl)
	model := fyersgosdk.NewFyersModel(appId, accessToken)

	results := make([]caseResult, 0, len(file.Cases))
	passed := 0
	failed := 0
	var failedCases []struct {
		Name  string
		Error string
	}
	for _, c := range file.Cases {
		out, errStr := runCase(t, c.Name, c.Input, client, model)
		ok := errStr == ""
		if ok {
			passed++
		} else {
			failed++
			failedCases = append(failedCases, struct {
				Name  string
				Error string
			}{c.Name, errStr})
		}
		results = append(results, caseResult{
			Function: c.Name,
			Input:    c.Input,
			Output:   out,
			Error:    errStr,
		})
	}

	outPath := filepath.Join(testDir, "test_output.json")
	outData, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		t.Fatalf("marshal results: %v", err)
	}
	if err := os.WriteFile(outPath, outData, 0644); err != nil {
		t.Fatalf("write test_output.json: %v", err)
	}
	t.Logf("wrote %d results to %s", len(results), outPath)

	// Write text report: report-YYYY-MM-DD.txt
	today := time.Now().Format("2006-01-02")
	reportPath := filepath.Join(testDir, "report-"+today+".txt")
	report := buildReport(today, passed, failed, failedCases, results)
	if err := os.WriteFile(reportPath, []byte(report), 0644); err != nil {
		t.Fatalf("write report: %v", err)
	}
	t.Logf("wrote report to %s", reportPath)
}

// defaultValidationCases used when test_cases.json has no validationCases (all except auth).
var defaultValidationCases = []string{
	"GetProfile", "GetFunds", "GetHoldings", "Logout", "GetOrderBook", "GetOrderBookByTag", "GetOrderById",
	"GetPositions", "GetTradeBook", "GetTradeBookByTag", "SingleOrderAction", "MultiOrderAction", "ModifyOrder",
	"ModifyMutliOrder", "CancelOrder", "CancelMutliOrder", "GTTSingleOrderAction", "GTTMultiOrderAction",
	"ModifyGTTOrder", "CancelGTTOrder", "GetGTTOrderBook", "ExitPosition", "ExitPositionById", "ExitPositionByProductType",
	"CancelPendingOrders", "ConvertPosition", "GetMarketStatus", "GetHistory", "GetStockQuotes", "GetMarketDepth",
	"GetOptionChain", "GetAlerts", "CreateAlert", "UpdateAlert", "DeleteAlert", "ToggleAlert",
}

// TestRunValidationWithToken runs validationCases from test_cases.json with FYERS_APP_ID and FYERS_ACCESS_TOKEN set,
// logs all responses, and fails if any response is not API success (s:"ok" or code 200).
func TestRunValidationWithToken(t *testing.T) {
	appId := os.Getenv("FYERS_APP_ID")
	accessToken := os.Getenv("FYERS_ACCESS_TOKEN")
	if appId == "" || accessToken == "" {
		t.Skip("FYERS_APP_ID and FYERS_ACCESS_TOKEN required for validation test")
	}
	testDir := "test"
	if _, err := os.Stat(testDir); os.IsNotExist(err) {
		testDir = "."
	}
	path := filepath.Join(testDir, "test_cases.json")
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read test_cases.json: %v", err)
	}
	var file testCaseFile
	if err := json.Unmarshal(data, &file); err != nil {
		t.Fatalf("parse test_cases.json: %v", err)
	}
	namesToRun := file.ValidationCases
	if len(namesToRun) == 0 {
		namesToRun = defaultValidationCases
	}
	nameToCase := make(map[string]struct{ Input json.RawMessage })
	for _, c := range file.Cases {
		nameToCase[c.Name] = struct{ Input json.RawMessage }{Input: c.Input}
	}
	appSecret := file.Config.AppSecret
	if appSecret == "" {
		appSecret = "test-secret"
	}
	redirectUrl := file.Config.RedirectUrl
	if redirectUrl == "" {
		redirectUrl = "https://example.com/callback"
	}
	client := fyersgosdk.SetClientData(appId, appSecret, redirectUrl)
	model := fyersgosdk.NewFyersModel(appId, accessToken)

	var failed []string
	for _, name := range namesToRun {
		c, ok := nameToCase[name]
		if !ok {
			t.Logf("skip (not in test_cases.json): %s", name)
			continue
		}
		out, errStr := runCase(t, name, c.Input, client, model)
		t.Logf("[%s] err=%v out_len=%d", name, errStr != "", len(out))
		if errStr != "" {
			failed = append(failed, name+": "+errStr)
			continue
		}
		if !responseIsSuccess(out) {
			failed = append(failed, name+": response not success (s!=ok/code!=200): "+truncate(out, 120))
		}
	}
	if len(failed) > 0 {
		for _, f := range failed {
			t.Error(f)
		}
	}
}

func responseIsSuccess(body string) bool {
	var v struct {
		S    string `json:"s"`
		Code int    `json:"code"`
	}
	if err := json.Unmarshal([]byte(body), &v); err != nil {
		return false
	}
	return v.S == "ok" || v.Code == 200
}

func truncate(s string, max int) string {
	if len(s) <= max {
		return s
	}
	return s[:max] + "..."
}

const reportOutputMaxLen = 2000 // truncate long output in report

func buildReport(date string, passed, failed int, failedCases []struct {
	Name  string
	Error string
}, results []caseResult) string {
	var b strings.Builder
	b.WriteString("FYERS GO SDK - TEST REPORT\n")
	b.WriteString(strings.Repeat("=", 60) + "\n")
	b.WriteString("Date: " + date + "\n")
	b.WriteString(strings.Repeat("-", 60) + "\n\n")

	b.WriteString("SUMMARY\n")
	b.WriteString(strings.Repeat("-", 60) + "\n")
	total := passed + failed
	b.WriteString(fmt.Sprintf("Total cases run:  %d\n", total))
	b.WriteString(fmt.Sprintf("Passed:           %d\n", passed))
	b.WriteString(fmt.Sprintf("Failed:           %d\n", failed))
	b.WriteString("\n")

	b.WriteString("CASES EXECUTED (Function name and status)\n")
	b.WriteString(strings.Repeat("-", 60) + "\n")
	for i, r := range results {
		status := "PASS"
		if r.Error != "" {
			status = "FAIL"
		}
		b.WriteString(fmt.Sprintf("%3d. %-35s [%s]\n", i+1, r.Function, status))
	}
	b.WriteString("\n")

	b.WriteString("DETAILS (Input and Output per case)\n")
	b.WriteString(strings.Repeat("-", 60) + "\n")
	for i, r := range results {
		status := "PASS"
		if r.Error != "" {
			status = "FAIL"
		}
		b.WriteString(fmt.Sprintf("\n--- %d. %s [%s] ---\n", i+1, r.Function, status))
		// Input (pretty-print if JSON)
		inputStr := string(r.Input)
		if len(r.Input) > 0 {
			var buf bytes.Buffer
			if err := json.Indent(&buf, r.Input, "  ", "  "); err == nil {
				inputStr = buf.String()
			}
		}
		if inputStr == "" || inputStr == "{}" {
			b.WriteString("Input:  (none)\n")
		} else {
			b.WriteString("Input:\n  ")
			b.WriteString(strings.ReplaceAll(inputStr, "\n", "\n  "))
			b.WriteString("\n")
		}
		// Output
		out := r.Output
		if len(out) > reportOutputMaxLen {
			out = out[:reportOutputMaxLen] + "\n... (truncated)"
		}
		if out == "" {
			b.WriteString("Output: (empty)\n")
		} else {
			b.WriteString("Output:\n  ")
			b.WriteString(strings.ReplaceAll(out, "\n", "\n  "))
			b.WriteString("\n")
		}
		if r.Error != "" {
			b.WriteString("Error:  " + r.Error + "\n")
		}
	}
	b.WriteString("\n")

	if len(failedCases) > 0 {
		b.WriteString("FAILED CASES (Function name and error)\n")
		b.WriteString(strings.Repeat("-", 60) + "\n")
		for i, fc := range failedCases {
			errShort := fc.Error
			if len(errShort) > 80 {
				errShort = errShort[:77] + "..."
			}
			b.WriteString(fmt.Sprintf("%3d. %s\n     Error: %s\n", i+1, fc.Name, errShort))
		}
	}

	return b.String()
}

func runCase(t *testing.T, name string, input json.RawMessage, client *fyersgosdk.Client, model *fyersgosdk.FyersModel) (output string, errStr string) {
	t.Helper()
	var err error
	defer func() {
		if r := recover(); r != nil {
			errStr = "panic: " + fmtS(r)
		}
	}()
	switch name {
	case "SetClientDataAndGetLoginURL":
		var in struct {
			AppId       string `json:"appId"`
			AppSecret   string `json:"appSecret"`
			RedirectUrl string `json:"redirectUrl"`
		}
		_ = json.Unmarshal(input, &in)
		c := fyersgosdk.SetClientData(in.AppId, in.AppSecret, in.RedirectUrl)
		output = c.GetLoginURL()
		return
	case "GenerateAccessToken":
		var in struct {
			AuthToken string `json:"authToken"`
		}
		_ = json.Unmarshal(input, &in)
		out, err := client.GenerateAccessToken(in.AuthToken, client)
		if err != nil {
			errStr = err.Error()
		}
		output = out
		return
	case "GenerateAccessTokenFromRefreshToken":
		var in struct {
			RefreshToken string `json:"refreshToken"`
			Pin          string `json:"pin"`
		}
		_ = json.Unmarshal(input, &in)
		out, err := client.GenerateAccessTokenFromRefreshToken(in.RefreshToken, in.Pin, client)
		if err != nil {
			errStr = err.Error()
		}
		output = out
		return
	case "GetProfile":
		output, err := model.GetProfile()
		if err != nil {
			errStr = err.Error()
		}
		return output, errStr
	case "GetFunds":
		output, err = model.GetFunds()
		if err != nil {
			errStr = err.Error()
		}
		return
	case "GetHoldings":
		output, err = model.GetHoldings()
		if err != nil {
			errStr = err.Error()
		}
		return
	case "Logout":
		output, err = model.Logout()
		if err != nil {
			errStr = err.Error()
		}
		return
	case "GetOrderBook":
		output, err = model.GetOrderBook()
		if err != nil {
			errStr = err.Error()
		}
		return
	case "GetOrderBookByTag":
		var in struct {
			Tag string `json:"tag"`
		}
		_ = json.Unmarshal(input, &in)
		output, err = model.GetOrderBookByTag(in.Tag)
		if err != nil {
			errStr = err.Error()
		}
		return
	case "GetOrderById":
		var in struct {
			Id string `json:"id"`
		}
		_ = json.Unmarshal(input, &in)
		output, err = model.GetOrderById(in.Id)
		if err != nil {
			errStr = err.Error()
		}
		return
	case "GetPositions":
		output, err = model.GetPositions()
		if err != nil {
			errStr = err.Error()
		}
		return
	case "GetTradeBook":
		output, err = model.GetTradeBook()
		if err != nil {
			errStr = err.Error()
		}
		return
	case "GetTradeBookByTag":
		var in struct {
			Tag string `json:"tag"`
		}
		_ = json.Unmarshal(input, &in)
		output, err = model.GetTradeBookByTag(in.Tag)
		if err != nil {
			errStr = err.Error()
		}
		return
	case "SingleOrderAction":
		var in fyersgosdk.OrderRequest
		_ = json.Unmarshal(input, &in)
		output, err = model.SingleOrderAction(in)
		if err != nil {
			errStr = err.Error()
		}
		return
	case "MultiOrderAction":
		var in struct {
			Orders []fyersgosdk.OrderRequest `json:"orders"`
		}
		_ = json.Unmarshal(input, &in)
		output, err = model.MultiOrderAction(in.Orders)
		if err != nil {
			errStr = err.Error()
		}
		return
	case "ModifyOrder":
		var in fyersgosdk.ModifyOrderRequest
		_ = json.Unmarshal(input, &in)
		output, err = model.ModifyOrder(in)
		if err != nil {
			errStr = err.Error()
		}
		return
	case "ModifyMutliOrder":
		var in struct {
			Requests []fyersgosdk.ModifyMultiOrderItem `json:"requests"`
		}
		_ = json.Unmarshal(input, &in)
		output, err = model.ModifyMutliOrder(in.Requests)
		if err != nil {
			errStr = err.Error()
		}
		return
	case "CancelOrder":
		var in struct {
			Id string `json:"id"`
		}
		_ = json.Unmarshal(input, &in)
		output, err = model.CancelOrder(in.Id)
		if err != nil {
			errStr = err.Error()
		}
		return
	case "CancelMutliOrder":
		var in struct {
			OrderIds []string `json:"orderIds"`
		}
		_ = json.Unmarshal(input, &in)
		output, err = model.CancelMutliOrder(in.OrderIds)
		if err != nil {
			errStr = err.Error()
		}
		return
	case "GTTSingleOrderAction":
		var in fyersgosdk.GTTOrderRequest
		_ = json.Unmarshal(input, &in)
		output, err = model.GTTSingleOrderAction(in)
		if err != nil {
			errStr = err.Error()
		}
		return
	case "GTTMultiOrderAction":
		var in struct {
			Orders []fyersgosdk.GTTOrderRequest `json:"orders"`
		}
		_ = json.Unmarshal(input, &in)
		output, err = model.GTTMultiOrderAction(in.Orders)
		if err != nil {
			errStr = err.Error()
		}
		return
	case "ModifyGTTOrder":
		var in struct {
			Requests []fyersgosdk.ModifyGTTOrderRequest `json:"requests"`
		}
		_ = json.Unmarshal(input, &in)
		output, err = model.ModifyGTTOrder(in.Requests)
		if err != nil {
			errStr = err.Error()
		}
		return
	case "CancelGTTOrder":
		var in struct {
			OrderId string `json:"orderId"`
		}
		_ = json.Unmarshal(input, &in)
		output, err = model.CancelGTTOrder(in.OrderId)
		if err != nil {
			errStr = err.Error()
		}
		return
	case "GetGTTOrderBook":
		output, err = model.GetGTTOrderBook()
		if err != nil {
			errStr = err.Error()
		}
		return
	case "ExitPosition":
		output, err = model.ExitPosition()
		if err != nil {
			errStr = err.Error()
		}
		return
	case "ExitPositionById":
		var in struct {
			OrderIds []string `json:"orderIds"`
		}
		_ = json.Unmarshal(input, &in)
		output, err = model.ExitPositionById(in.OrderIds)
		if err != nil {
			errStr = err.Error()
		}
		return
	case "ExitPositionByProductType":
		var in fyersgosdk.ExitPositionByProductTypeRequest
		_ = json.Unmarshal(input, &in)
		output, err = model.ExitPositionByProductType(in)
		if err != nil {
			errStr = err.Error()
		}
		return
	case "CancelPendingOrders":
		var in fyersgosdk.CancelPendingOrdersRequest
		_ = json.Unmarshal(input, &in)
		output, err = model.CancelPendingOrders(in)
		if err != nil {
			errStr = err.Error()
		}
		return
	case "ConvertPosition":
		var in fyersgosdk.ConvertPositionRequest
		_ = json.Unmarshal(input, &in)
		output, err = model.ConvertPosition(in)
		if err != nil {
			errStr = err.Error()
		}
		return
	case "GetMarketStatus":
		output, err = model.GetMarketStatus()
		if err != nil {
			errStr = err.Error()
		}
		return
	case "GetHistory":
		var in fyersgosdk.HistoryRequest
		_ = json.Unmarshal(input, &in)
		output, err = model.GetHistory(in)
		if err != nil {
			errStr = err.Error()
		}
		return
	case "GetStockQuotes":
		var in struct {
			Symbols []string `json:"symbols"`
		}
		_ = json.Unmarshal(input, &in)
		output, err = model.GetStockQuotes(in.Symbols)
		if err != nil {
			errStr = err.Error()
		}
		return
	case "GetMarketDepth":
		var in fyersgosdk.MarketDepthRequest
		_ = json.Unmarshal(input, &in)
		output, err = model.GetMarketDepth(in)
		if err != nil {
			errStr = err.Error()
		}
		return
	case "GetOptionChain":
		var in fyersgosdk.OptionChainRequest
		_ = json.Unmarshal(input, &in)
		output, err = model.GetOptionChain(in)
		if err != nil {
			errStr = err.Error()
		}
		return
	case "GetAlerts":
		output, err = model.GetAlerts()
		if err != nil {
			errStr = err.Error()
		}
		return
	case "CreateAlert":
		var in fyersgosdk.AlertRequest
		_ = json.Unmarshal(input, &in)
		output, err = model.CreateAlert(in)
		if err != nil {
			errStr = err.Error()
		}
		return
	case "UpdateAlert":
		var in struct {
			AlertId      string                  `json:"alertId"`
			AlertRequest fyersgosdk.AlertRequest `json:"alertRequest"`
		}
		_ = json.Unmarshal(input, &in)
		output, err = model.UpdateAlert(in.AlertId, in.AlertRequest)
		if err != nil {
			errStr = err.Error()
		}
		return
	case "DeleteAlert":
		var in struct {
			AlertId string `json:"alertId"`
		}
		_ = json.Unmarshal(input, &in)
		output, err = model.DeleteAlert(in.AlertId)
		if err != nil {
			errStr = err.Error()
		}
		return
	case "ToggleAlert":
		var in struct {
			AlertId string `json:"alertId"`
		}
		_ = json.Unmarshal(input, &in)
		output, err = model.ToggleAlert(in.AlertId)
		if err != nil {
			errStr = err.Error()
		}
		return
	default:
		errStr = "unknown case: " + name
		return
	}
}

func fmtS(v interface{}) string {
	switch x := v.(type) {
	case string:
		return x
	case error:
		return x.Error()
	default:
		return fmt.Sprintf("%v", v)
	}
}
