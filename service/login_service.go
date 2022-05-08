package service

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	res := &JsonResult{}
	if r.Method == http.MethodPost {
		code, err := getCode(r)
		if err != nil {
			res.Code = -1
			res.ErrorMsg = err.Error()
			res.Data = nil
		} else {
			res.Code = 0
			res.ErrorMsg = ""
			res.Data = code
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
