package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func code2Session(code string) (resp *http.Response, err error) {
	url := "https://api.weixin.qq.com/sns/jscode2session?appid=wxdcab629e85115972&secret=093bb5adeb959c37e4d225a68123afcb&js_code="
	url += code
	url += "&grant_type=authorization_code"
	return http.Get(url)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	res := &JsonResult{}
	if r.Method == http.MethodPost {
		code, err := getCode(r)
		if err != nil {
			res.Code = -1
			res.ErrorMsg = err.Error()
			res.Data = nil
		} else {
			resp, err := code2Session(code)
			if err != nil {
				res.Code = -1
				res.ErrorMsg = err.Error()
				res.Data = nil
			} else {
				var session SessionResult

				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					res.Code = -1
					res.ErrorMsg = err.Error()
					res.Data = nil
				}

				err = json.Unmarshal([]byte(body), &session)
				if err != nil {
					res.Code = -1
					res.ErrorMsg = err.Error()
					res.Data = nil
				}

				fmt.Println(session)

				if session.ErrorMsg != "" {
					res.Code = -1
					res.ErrorMsg = session.ErrorMsg
					res.Data = nil
				} else {
					res.Code = 0
					res.ErrorMsg = ""
					res.Data = session.OpenId
				}
			}
		}
	} else {
		res.Code = -1
		res.ErrorMsg = fmt.Sprintf("请求方法 %s 不支持", r.Method)
	}

	msg, err := json.Marshal(res)
	if err != nil {
		fmt.Fprint(w, "内部错误")
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(msg)
}
