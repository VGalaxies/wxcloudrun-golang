package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"wxcloudrun-golang/db/dao"
	"wxcloudrun-golang/db/model"
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

func BookGetDispatch(r *http.Request) (*[]model.BookModel, error) {
	action, hint, err := getBody(r)
	if err != nil {
		return nil, err
	}

	var model *[]model.BookModel
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
		model, err = dao.BookImp.GetBookByNameCate(hint)
		if err != nil {
			return nil, err
		}
	} else {
		err = fmt.Errorf("参数 action : %s 错误", action)
	}

	return model, err
}
