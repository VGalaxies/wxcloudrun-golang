package service

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// JsonResult 返回结构
type JsonResult struct {
	Code     int         `json:"code"`
	ErrorMsg string      `json:"errorMsg,omitempty"`
	Data     interface{} `json:"data"`
}

// get action, hint
func getBody(r *http.Request) (string, string, error) {
	decoder := json.NewDecoder(r.Body)
	body := make(map[string]interface{})
	if err := decoder.Decode(&body); err != nil {
		return "", "", err
	}
	defer r.Body.Close()

	action, ok := body["action"]
	if !ok {
		return "", "", fmt.Errorf("缺少 action 参数")
	}

	hint, ok := body["hint"]
	if !ok {
		return "", "", fmt.Errorf("缺少 hint 参数")
	}

	return action.(string), hint.(string), nil
}

// get code
func getCode(r *http.Request) (string, error) {
	decoder := json.NewDecoder(r.Body)
	body := make(map[string]interface{})
	if err := decoder.Decode(&body); err != nil {
		return "", err
	}
	defer r.Body.Close()

	code, ok := body["code"]
	if !ok {
		return "", fmt.Errorf("缺少 code 参数")
	}

	return code.(string), nil
}
