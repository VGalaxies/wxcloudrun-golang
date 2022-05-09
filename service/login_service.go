package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"wxcloudrun-golang/db/dao"
)

func LoginGetHandler(w http.ResponseWriter, r *http.Request) {
	res := &JsonResult{}
	var model interface{}
	if r.Method == http.MethodPost {
		openid, err := getOpenId(r)
		if err != nil {
			res.Code = -1
			res.ErrorMsg = err.Error()
		} else {
			model, err = dao.UserImp.GetUserInfo(openid)
			if err != nil {
				res.Code = -1
				res.ErrorMsg = err.Error()
			} else {
				res.Data = model
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

func LoginSetHandler(w http.ResponseWriter, r *http.Request) {
	res := &JsonResult{}
	if r.Method == http.MethodPost {
		openid, nickname, avatar, err := getUserInfo(r)
		if err != nil {
			res.Code = -1
			res.ErrorMsg = err.Error()
		} else {
			err = dao.UserImp.SetUserInfo(openid, nickname, avatar)
			if err != nil {
				res.Code = -1
				res.ErrorMsg = err.Error()
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

func code2Session(code string) (resp *http.Response, err error) {
	url := "https://api.weixin.qq.com/sns/jscode2session?appid=wxdcab629e85115972&secret=093bb5adeb959c37e4d225a68123afcb&js_code="
	url += code
	url += "&grant_type=authorization_code"
	return http.Get(url)
}

func LoginInitHandler(w http.ResponseWriter, r *http.Request) {
	res := &JsonResult{}

	var err error
	var code string
	var resp *http.Response
	var body []byte
	var session SessionResult

	if r.Method != http.MethodPost {
		res.Code = -1
		res.ErrorMsg = fmt.Sprintf("请求方法 %s 不支持", r.Method)
		goto FINAL
	}

	code, err = getCode(r)
	if err != nil {
		res.Code = -1
		res.ErrorMsg = err.Error()
		goto FINAL
	}

	resp, err = code2Session(code)
	if err != nil {
		res.Code = -1
		res.ErrorMsg = err.Error()
		goto FINAL
	}

	if resp.Status != "200 OK" {
		res.Code = -1
		res.ErrorMsg = fmt.Sprintf("resp status: %s", resp.Status)
		goto FINAL
	}

	body, err = ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		res.Code = -1
		res.ErrorMsg = err.Error()
		goto FINAL
	}

	fmt.Println(string(body))

	err = json.Unmarshal([]byte(body), &session)
	if err != nil {
		res.Code = -1
		res.ErrorMsg = err.Error()
		goto FINAL
	}

	fmt.Println(session)

	if session.ErrorCode != 0 {
		res.Code = -1
		res.ErrorMsg = session.ErrorMsg
		goto FINAL
	}

	res.Data = session.OpenId

FINAL:
	msg, err := json.Marshal(res)
	if err != nil {
		fmt.Fprint(w, "内部错误")
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(msg)
}
