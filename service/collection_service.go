package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"wxcloudrun-golang/db/dao"
)

func CollectionUnsetHandler(w http.ResponseWriter, r *http.Request) {
	res := &JsonResult{}
	if r.Method == http.MethodPost {
		tmp, err := getBody(r, []string{"userid", "bookid"})
		if err != nil {
			res.Code = -1
			res.ErrorMsg = err.Error()
		} else {
			userid := tmp[0]
			bookid := tmp[1]
			err = dao.CollectionImp.UnsetCollectionInfo(userid, bookid)
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

func CollectionSetHandler(w http.ResponseWriter, r *http.Request) {
	res := &JsonResult{}
	if r.Method == http.MethodPost {
		tmp, err := getBody(r, []string{"userid", "bookid"})
		if err != nil {
			res.Code = -1
			res.ErrorMsg = err.Error()
		} else {
			userid := tmp[0]
			bookid := tmp[1]
			err = dao.CollectionImp.SetCollectionInfo(userid, bookid)
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

func CollectionGetHandler(w http.ResponseWriter, r *http.Request) {
	res := &JsonResult{}
	if r.Method == http.MethodPost {
		model, err := CollectionGetDispatch(r)
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

func CollectionGetDispatch(r *http.Request) (interface{}, error) {
	tmp, err := getBody(r, []string{"action", "hint"})
	if err != nil {
		return nil, err
	}

	action := tmp[0]
	hint := tmp[1]

	var model interface{}
	if action == "user" {
		model, err = dao.CollectionImp.GetCollectionInfoByUser(hint)
		if err != nil {
			return nil, err
		}
	} else if action == "book" {
		model, err = dao.CollectionImp.GetCollectionInfoByBook(hint)
		if err != nil {
			return nil, err
		}
	} else {
		err = fmt.Errorf("参数 action : %s 错误", action)
	}

	return model, err
}
