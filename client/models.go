package client

type Data struct {
	Fields []string        `json:"fields"`
	Items  [][]interface{} `json:"items"`
}

type APIResponse struct {
	RequestID string      `json:"request_id"`
	Code      int         `json:"code"`
	Msg       interface{} `json:"msg"`
	Data      Data        `json:"data"`
}
