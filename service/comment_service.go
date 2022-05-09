package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"wxcloudrun-golang/db/dao"
)

func CommentSetHandler(w http.ResponseWriter, r *http.Request) {
	res := &JsonResult{}
	if r.Method == http.MethodPost {
		userid, bookid, comment, err := getBodyComment(r)
		if err != nil {
			res.Code = -1
			res.ErrorMsg = err.Error()
		} else {
			err = dao.CommentImp.SetCommentInfo(userid, bookid, comment)
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

func CommentGetHandler(w http.ResponseWriter, r *http.Request) {
	res := &JsonResult{}
	if r.Method == http.MethodPost {
		model, err := CommentGetDispatch(r)
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

func CommentGetDispatch(r *http.Request) (interface{}, error) {
	action, hint, err := getBodyActionAndHint(r)
	if err != nil {
		return nil, err
	}

	var model interface{}
	if action == "user" {
		model, err = dao.CommentImp.GetCommentInfoByUser(hint)
		if err != nil {
			return nil, err
		}
	} else if action == "book" {
		model, err = dao.CommentImp.GetCommentInfoByBook(hint)
		if err != nil {
			return nil, err
		}
	} else {
		err = fmt.Errorf("参数 action : %s 错误", action)
	}

	return model, err
}
