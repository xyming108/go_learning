package main

import (
	"fmt"
	consulapi "github.com/hashicorp/consul/api"
	"log"
	"net"
	"net/http"
	"strings"
	"time"
)

var (
	appname    = "bb"
	appid      = "bb"
	httpport   = 8284
	consul_url = "127.0.0.1:8500"
)

func RegisterServer(consulAddr, serverId, serverName string, serverPort int) {
	if consulAddr == "" {
		return
	}
	//连接到consul注册中心
	config := consulapi.DefaultConfig()
	config.Address = consulAddr
	client, err := consulapi.NewClient(config)
	if err != nil {
		log.Println("consul client error : ", err)
	}

	//服务器信息
	ip := GetInternalIP() //"10.1.114.254"
	registration := new(consulapi.AgentServiceRegistration)
	registration.ID = serverId       // 服务节点的名称
	registration.Name = serverName   // 服务名称
	registration.Port = serverPort   // 服务端口
	registration.Tags = []string{""} // tag，可以为空
	registration.Address = ip        // 服务 IP
	log.Println("GetInternalIP:", ip)
	log.Println("registration.Address: ", localIP())

	//健康检查
	registration.Check = &consulapi.AgentServiceCheck{
		TTL:     "5s",
		CheckID: appid,
		//HTTP:                           fmt.Sprintf("http://%s:%d", registration.Address, serverPort),
		//Timeout:                        "3s",
		//Interval:                       "5s", // 健康检查间隔
		DeregisterCriticalServiceAfter: "4s", //check失败后30秒删除本服务，注销时间，相当于过期时间
		// GRPC:     fmt.Sprintf("%v:%v/%v", IP, r.Port, r.Service),// grpc 支持，执行健康检查的地址，service 会传到 Health.Check 函数中
	}

	//客户端注册，相对于consul来说，本服务器就是客户端
	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		log.Println("register server error : ", err)
	}

	//用于TTl监听的情况
	for {
		err := client.Agent().UpdateTTL(appid, "哈哈2", "pass")
		if err != nil {
			return
		}
		log.Println("dfkjhdfjkdhjkdghdjksgh")
		time.Sleep(2 * time.Second)
	}

}

//
func localIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

//局域网ip
func GetInternalIP() string {
	// 思路来自于Python版本的内网IP获取，其他版本不准确
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Println("internal IP fetch failed, detail:" + err.Error())
		return ""
	}
	defer conn.Close()
	// udp 面向无连接，所以这些东西只在你本地捣鼓
	res := conn.LocalAddr().String()
	log.Println("res:", res)
	res = strings.Split(res, ":")[0]
	return res
}

func Handle() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("现在在处理健康监测：")
		writer.Write([]byte("兄弟，你监听成功了！"))
	})

	err := http.ListenAndServe("10.0.54.53:8284", nil)
	if err != nil {
		log.Println("err: ", err)
		return
	}
	log.Println("get health check")
}

func main() {
	RegisterServer(consul_url, appname, appid, httpport)

	//用于http连接的情况
	//Handle()
}
