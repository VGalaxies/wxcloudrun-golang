package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"wxcloudrun-golang/db/dao"
)

func BookGetHandler(w http.ResponseWriter, r *http.Request) {
	res := &JsonResult{}
	if r.Method == http.MethodPost {
		model, err := BookGetDispatch(r)
		if err != nil {
			res.Code = -1
			res.ErrorMsg = err.Error()
		} else {
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

func BookGetDispatch(r *http.Request) (interface{}, error) {
	tmp, err := getBody(r, []string{"action", "hint"})
	if err != nil {
		return nil, err
	}

	action := tmp[0]
	hint := tmp[1]

	var model interface{}
	if action == "exact" {
		model, err = dao.BookImp.GetBookByName(hint)
		if err != nil {
			return nil, err
		}
	} else if action == "fuzzy" {
		model, err = dao.BookImp.GetBookByNameFzf(hint)
		if err != nil {
			return nil, err
		}
	} else if action == "category" {
		model, err = dao.BookImp.GetBookByCategory(hint)
		if err != nil {
			return nil, err
		}
	} else if action == "id" {
		model, err = dao.BookImp.GetBookById(hint)
		if err != nil {
			return nil, err
		}
	} else {
		err = fmt.Errorf("参数 action : %s 错误", action)
	}

	return model, err
}
