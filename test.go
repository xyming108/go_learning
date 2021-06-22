//brief_intro/echo.go
package main

import (
	"log"
	"net/http"
	"os"
	"time"
)

var logger = log.New(os.Stdout, "", 0)

func echo(wr http.ResponseWriter, r *http.Request) {
	timeStart := time.Now()
	_, err := wr.Write([]byte("hello"))
	timeEnd := time.Since(timeStart)
	logger.Println(timeEnd)
	if err != nil {
		return
	}
}

func main() {
	http.HandleFunc("/", echo)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
