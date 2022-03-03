package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

/*
绝大多数请求失败的情况下，resp 的值为 nil 且 err 为 non-nil。
但如果你得到的是重定向错误，那它俩的值都是 non-nil，最后依旧可能发生内存泄露。

2 个解决办法：
可以直接在处理 HTTP 响应错误的代码块中，直接关闭非 nil 的响应体。
手动调用 defer 来关闭响应体：
*/

// 正确示例
func main() {
	resp, err := http.Get("http://www.baidu.co")
	fmt.Println("-----", resp)

	// 关闭 resp.Body 的正确姿势
	if resp != nil {
		defer resp.Body.Close()
	}

	checkError2(err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	checkError2(err)

	fmt.Println(string(body))
}

func checkError2(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

/*
resp.Body.Close() 早先版本的实现是读取响应体的数据之后丢弃，
保证了 keep-alive 的 HTTP 连接能重用处理不止一个请求。
但 Go 的最新版本将读取并丢弃数据的任务交给了用户，如果你不处理，
HTTP 连接可能会直接关闭而非重用，参考在 Go 1.5 版本文档。

如果程序大量重用 HTTP 长连接，你可能要在处理响应的逻辑代码中加入：

_, err = io.Copy(ioutil.Discard, resp.Body)    // 手动丢弃读取完毕的数据
如果你需要完整读取响应，上边的代码是需要写的。比如在解码 API 的 JSON 响应数据：

json.NewDecoder(resp.Body).Decode(&data)

*/
