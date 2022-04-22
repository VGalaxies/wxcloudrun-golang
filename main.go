package main

import (
	"fmt"
	"log"
	"net/http"

  "wxcloudrun-golang/db/model"
	"wxcloudrun-golang/db"
	"wxcloudrun-golang/service"
)

func main() {
	if err := db.Init(); err != nil {
		panic(fmt.Sprintf("mysql init failed with %+v", err))
	}

  if err := model.InitBook(); err != nil {
		panic(fmt.Sprintf("init book failed with %+v", err))
  }

  if err := model.InitCategory(); err != nil {
		panic(fmt.Sprintf("init book failed with %+v", err))
  }

	http.HandleFunc("/", service.IndexHandler)
	http.HandleFunc("/api/count", service.CounterHandler)

  http.HandleFunc("/book/get", service.BookGetHandler)

	log.Fatal(http.ListenAndServe(":80", nil))
}
