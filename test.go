package main

import (
	"bytes"
	"compress/zlib"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
)

type TestItem struct {
	Name   string      `json:"name"`
	Result interface{} `json:"result"`
	Val    interface{} `json:"val"`
}

type TestItems struct {
	Item []TestItem `json:"items"`
}

func uZipBytes(data []byte) []byte {
	var out bytes.Buffer
	var in bytes.Buffer
	in.Write(data)
	r, _ := zlib.NewReader(&in)
	r.Close()
	io.Copy(&out, r)
	return out.Bytes()
}

func getBase64(data string) ([]byte, error) {
	a, err := base64.StdEncoding.DecodeString(data)

	return a, err
}

func UnCompress(s string, jsonFlag bool) string {
	// 开启 了 json 验证 并且 是有效的 json ,则 直接返回 原来的值
	if jsonFlag && json.Valid([]byte(s)) || len(s) == 0 {
		return s
	}
	data, err := getBase64(s)
	if err != nil {
		log.Println("解析报错")
		return ""
	}
	data = uZipBytes(data)
	fmt.Println(string(data))
	return string(data)
}

func main() {
	ss := "eJy8lU9rwkAQxe+Ffodlzyq7k9j8ubUx0h4swbV4KD0sOpTFJNrN1raI370EEeyl7BzWUzaT8H7zeOHlcHvDGGPcOGw6nrPX0z1jh/OBMd7qBnnO+MOEDy7GFrvP2vGcycvpXtf9u2ValkUBE0hTyc+Pj4P/5IsNrjYURnWvlKd22Tq008UsjPi3capG3M22awxCeCye5tihCyW+1Bt82QVRn08X2xbLxoRZfo56DbGAavuF1p8Qj9KISFCmfa9RYdsZZ/bG/fjTxEiQYLEk2olGKRAJV7STCnI6ckwkXM1O4BLsEVOLH4Tsx4Kyf1URSlD6flXq2V/0LssAADIRJVHmq7/SLYVBaBCFrtBhuk85bd3CNISfAgiAoZRDAUyKHJI8Fvk4SXwTXlrjkBTxXzunw1t/Of4GAAD//xNOpJw="
	//ss := "eJzM1c9rE0EUB/B7/4plTgomvJ2Z3SR7S8IchK2UbBRBZGnrHgJtAmkiSCkUFNo0hkZKaMW0tVAJiCQFfxSzVv+Zzk7yX8jGgoQoHuYFnOObt8uH7z7ebi4YN4eMw6NR/1yFL2S7RRzj983ktlQL1oljPJqqxmdzpjLpLy+vB8QxiLx8J4f70duz8fuX5M6fe6vBRn2tRhyD/qXh6fJa/K6lrOeRmY6t2Yf+YfLu6UrSwDIpDnkOHCwKDEM16n+X5zty/1IXJxhPWcA5Bmp89lW96ct2S/UudF1ZlrFYwk2o5z151UGJbG+gGjvypImQWk541AbwC0vFh0UMnPoYqvBUNXaj7gddXN70H7AkTdKkbftmrWKCv1Kvlv0sAw636G0Mb/S5ORp0rodD+epI13vfEwUTJcTJrFxftXRFJuBlJL9ty14zOvwyPvyk66JAaQIyCZoxTOpY4ICNB42JnR+6RJsl0yjL5BdKhQfRSVcXFe225d4pyhI5OFaN17J7IY+3tVdvWoh8nqXSFk9h2HKu+D9hN6Gh/Be8Sr38ZLVSDQy3tBJUa88MlGmLs5uXz8+5AsNYLNxd1N5sFOWLLgpPVxLvxfYAA+OWVoPyRjA30FTl8cJ0fetnAAAA//9X/kp2"

	//var results map[string]TestItems
	var results TestItems
	err := json.Unmarshal([]byte(UnCompress(ss, true)), &results)
	if err != nil {
		log.Println("------", err)
	}
	fmt.Println(results)
}
