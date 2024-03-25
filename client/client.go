package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// Request represents the JSON payload to be sent in the POST request.
type Request struct {
	Timestamp int64   `json:"timestamp"`
	Content   Content `json:"content"`
}

type Content struct {
	ExtTradeNo  string      `json:"extTradeNo"`
	RedirectURL string      `json:"redirectUrl"`
	Company     Company     `json:"company"`
	Customer    Customer    `json:"customer"`
	Product     Product     `json:"product"`
	Installment Installment `json:"installment"`
}

type Company struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Customer struct {
	ExtID  string `json:"extId"`
	Name   string `json:"name"`
	Addr   string `json:"addr"`
	Phone  string `json:"phone"`
	IDCard string `json:"idCard"`
}

type Product struct {
	ExtID   string `json:"extId"`
	Name    string `json:"name"`
	Price   string `json:"price"`
	Content string `json:"Content"`
}

type Installment struct {
	Limit float64 `json:"limit"`
	First float64 `json:"first"`
	Num   int     `json:"num"`
	Type  string  `json:"type"`
}

// ////////////////////////////////////
// Response
type Response struct {
	Code    int             `json:"code"`
	Content ResponseContent `json:"content"`
	Msg     string          `json:"msg"`
}

type ResponseContent struct {
	ContractID string     `json:"contractId"`
	Customer   Customer   `json:"customer"`
	Bill       []BillItem `json:"bill"`
	SignURL    string     `json:"signUrl"`
	Status     int        `json:"status"`
}

type BillItem struct {
	Index  int       `json:"index"`
	Days   time.Time `json:"days"`
	Amount float64   `json:"amount"`
}

// Client is a client for interacting with the Gateway API.
type Client struct {
	AppID string
}

// {"code":0,"content":{"contractId":"co0hdgni2dkrn7or4m8g","customer":{"extId":"user_002","name":"张某","addr":"杭州市西湖区","phone":"13958040000","idCard":"33100219810412251X"},"bill":[{"index":1,"days":"2024-03-25T08:00:00+08:00","amount":0.5},{"index":2,"days":"2024-04-25T08:00:00+08:00","amount":0.5}],"signUrl":"https://u.alipay.cn/_eLriCTVod5djaQX9hEFxd","status":12},"msg":"创建成功"}

// Post sends a POST request to the specified endpoint with the given request data.
func (c *Client) Post(endpoint string, req *Request) (*Response, error) {
	jsonBytes, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}
	// https: //open.xiadandt.com/gate/liteContract/create
	url := fmt.Sprintf("https://open.xiadandt.com%s", endpoint)
	reqBody := bytes.NewBuffer(jsonBytes)

	httpReq, err := http.NewRequest(http.MethodPost, url, reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("X-Gateway-AppId", c.AppID)

	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("failed to send POST request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		body, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("received non-successful status code: %d, body: %s", resp.StatusCode, string(body))
	}

	body, _ := ioutil.ReadAll(resp.Body)
	r := &Response{}
	json.Unmarshal([]byte(body), r)

	return r, nil
}
