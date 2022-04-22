package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"wxcloudrun-golang/db/dao"
	"wxcloudrun-golang/db/model"
)

func BookGetHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hit BookGetHandler")

	res := &JsonResult{}
	if r.Method == http.MethodGet {
		fmt.Println("hit BookGetHandler MethodGet")
		model, err := BookGetDispatch(r)
    fmt.Println(model)
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

func BookGetDispatch(r *http.Request) (*model.BookModel, error) {
	action, err := getAction(r)
	if err != nil {
		return nil, err
	}

	hint, err := getHint(r)
	if err != nil {
		return nil, err
	}

	fmt.Println(action, hint)

	var model *model.BookModel
	if action == "exact" {
		model, err = dao.BookImp.GetBookByName(hint)
		if err != nil {
			return nil, err
		}
	} else if action == "fuzzy" {

	} else if action == "category" {

	} else {
		err = fmt.Errorf("参数 action : %s 错误", action)
	}

	return model, err
}

func getHint(r *http.Request) (string, error) {
	decoder := json.NewDecoder(r.Body)
	body := make(map[string]interface{})
	if err := decoder.Decode(&body); err != nil {
		return "", err
	}
	defer r.Body.Close()

	action, ok := body["hint"]
	if !ok {
		return "", fmt.Errorf("缺少 hint 参数")
	}

	return action.(string), nil
}
