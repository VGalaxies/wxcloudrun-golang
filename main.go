package main

import (
	"fmt"
	"log"
	"net/http"

	"wxcloudrun-golang/db"
	"wxcloudrun-golang/service"
)

func main() {
	if err := db.Init(); err != nil {
		panic(fmt.Sprintf("mysql init failed with %+v", err))
	}

	http.HandleFunc("/", service.IndexHandler)
	http.HandleFunc("/api/book", service.BookGetHandler)
	http.HandleFunc("/api/category", service.CategoryGetHandler)
	http.HandleFunc("/api/loginInit", service.LoginInitHandler)
	http.HandleFunc("/api/loginGet", service.LoginGetHandler)
	http.HandleFunc("/api/loginSet", service.LoginSetHandler)
	http.HandleFunc("/api/commentSet", service.CommentSetHandler)
	http.HandleFunc("/api/commentGet", service.CommentGetHandler)

	log.Fatal(http.ListenAndServe(":80", nil))
}
