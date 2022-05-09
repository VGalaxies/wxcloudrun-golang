package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"wxcloudrun-golang/db/dao"
)

func CategoryGetHandler(w http.ResponseWriter, r *http.Request) {
	res := &JsonResult{}
	if r.Method == http.MethodPost {
		model, err := CategoryGetDispatch(r)
		if err != nil {
			res.Code = -1
			res.ErrorMsg = err.Error()
			res.Data = nil
		} else {
			res.Code = 0
			res.ErrorMsg = ""
			res.Data = model
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

func CategoryGetDispatch(r *http.Request) (interface{}, error) {
	action, hint, err := getBody(r)
	if err != nil {
		return nil, err
	}

	var model interface{}
	if action == "single" {
		model, err = dao.CategoryImp.GetCategory(hint)
		if err != nil {
			return nil, err
		}
	} else if action == "all" {
		model, err = dao.CategoryImp.GetCategoryAll()
		if err != nil {
			return nil, err
		}
	} else {
		err = fmt.Errorf("参数 action : %s 错误", action)
	}

	return model, err
}