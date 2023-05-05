package utils

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func HttpRequest(methodUrl string, method string, dataBody string, dataHeader map[string]string /* timestamp, token string*/) ([]byte, error) {
	jsonStr := []byte(dataBody)

	//1.获取请求
	client := &http.Client{}
	apiUrl := methodUrl
	request, err := http.NewRequest(method, apiUrl, bytes.NewBuffer(jsonStr))
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
	}

	//2.填充header
	for key, value := range dataHeader {
		request.Header.Set(key, value)
	}

	//3.发送请求
	resp, err := client.Do(request) //发送请求
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		return nil, err
	}
	defer resp.Body.Close() //一定要关闭resp.Body
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		return nil, err
	}

	return content, nil
}
