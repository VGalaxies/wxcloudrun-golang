package service

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// JsonResult 返回结构
type JsonResult struct {
	Code     int         `json:"code"`
	ErrorMsg string      `json:"errorMsg"`
	Data     interface{} `json:"data"`
}

// auth.code2Session 返回结构
type SessionResult struct {
	OpenId     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionId    string `json:"unionid"`
	ErrorCode  int    `json:"errcode"`
	ErrorMsg   string `json:"errmsg"`
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

// get openid
func getOpenId(r *http.Request) (string, error) {
	decoder := json.NewDecoder(r.Body)
	body := make(map[string]interface{})
	if err := decoder.Decode(&body); err != nil {
		return "", err
	}
	defer r.Body.Close()

	openid, ok := body["openid"]
	if !ok {
		return "", fmt.Errorf("缺少 openid 参数")
	}

	return openid.(string), nil
}

// get openid, nickname, avatar
func getUserInfo(r *http.Request) (string, string, string, error) {
	decoder := json.NewDecoder(r.Body)
	body := make(map[string]interface{})
	if err := decoder.Decode(&body); err != nil {
		return "", "", "", err
	}
	defer r.Body.Close()

	openid, ok := body["openid"]
	if !ok {
		return "", "", "", fmt.Errorf("缺少 openid 参数")
	}

	nickname, ok := body["nickname"]
	if !ok {
		return "", "", "", fmt.Errorf("缺少 nickname 参数")
	}

	avatar, ok := body["avatar"]
	if !ok {
		return "", "", "", fmt.Errorf("缺少 avatar 参数")
	}

	return openid.(string), nickname.(string), avatar.(string), nil
}
