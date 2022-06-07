package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

const requestIDKey int = 0

func WithRequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(rw http.ResponseWriter, req *http.Request) {
			// 从 header 中提取 request-id
			reqID := req.Header.Get("X-Request-ID")
			// 创建 valueCtx。使用自定义的类型，不容易冲突
			ctx := context.WithValue(
				req.Context(), requestIDKey, reqID)

			// 创建新的请求
			req = req.WithContext(ctx)

			// 调用 HTTP 处理函数
			next.ServeHTTP(rw, req)
		},
	)
}

// GetRequestID 获取 request-id
func GetRequestID(ctx context.Context) string {
	r := ctx.Value(requestIDKey).(string)
	log.Println("requestID", r)
	return r
}

func Handle(rw http.ResponseWriter, req *http.Request) {
	// 拿到 reqId，后面可以记录日志等等
	reqID := GetRequestID(req.Context())

	log.Printf("连接成功: %+v", req.RemoteAddr)
	log.Println("requestID:",reqID)
	fmt.Println("url:", req.URL)
	fmt.Println("header:", req.Header)
	fmt.Println("body:", req.Body)
	fmt.Println("========================")

	rw.Write([]byte("我是服务端"))
}

func main() {
	handlerFunc := http.HandlerFunc(Handle)
	handler := WithRequestID(handlerFunc)
	err := http.ListenAndServe("127.0.0.1:8000", handler)
	if err != nil {
		return
	}
}
