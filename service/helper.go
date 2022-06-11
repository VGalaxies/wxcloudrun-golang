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

func getBody(r *http.Request, argv []string) ([]string, error) {
	argc := len(argv)
	decoder := json.NewDecoder(r.Body)
	body := make(map[string]interface{})
	if err := decoder.Decode(&body); err != nil {
		return make([]string, argc), err
	}
	defer r.Body.Close()

	res := make([]string, argc)
	for i, v := range argv {
		tmp, ok := body[v]
		if !ok {
			return make([]string, argc), fmt.Errorf("缺少 %s 参数", v)
		}
		res[i] = tmp.(string)
	}

	return res, nil
}
