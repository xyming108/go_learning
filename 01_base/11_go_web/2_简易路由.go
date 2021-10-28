package main

import (
	"fmt"
	"net/http"
)

/**
Author: xym
Date: 2021/5/16 17:41
Project: myproject2_go_learning
Description:
*/

type MyMux struct {
}

func (p *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		sayhelloName1(w, r)
		return
	}
	http.NotFound(w, r)
	return
}

func sayhelloName1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello myroute!")
}

func main() {
	mux := &MyMux{}
	http.ListenAndServe(":9090", mux)
}
