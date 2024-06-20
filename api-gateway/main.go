package main

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/micro/go-micro/v2/web"
	"micro-TodoList/api-gateway/service"
	"micro-TodoList/api-gateway/weblib"
	"micro-TodoList/api-gateway/wrappers"
	"time"
)

func main() {
	//etcd注册件
	etcdReg := etcd.NewRegistry(
		registry.Addrs("127.0.0.1:2379"), //etcd地址
	)
	//用户
	userMicroService := micro.NewService(
		micro.Name("userService.client "), //微服务名字
		micro.WrapClient(wrappers.NewUserWrapper),
	)
	//用户服务调用实例
	userService := service.NewUserService("rpcUserService", userMicroService.Client())
	// 创建微服务实例，使用gin暴露http接口并注册到etcd
	server := web.NewService(
		web.Name("httpService"),
		web.Address("127.0.0.1:4000"),
		// 将服务调用实例使用gin处理
		web.Handler(weblib.NewRouter(userService)),
		web.Registry(etcdReg),
		web.RegisterTTL(time.Second*30),
		web.RegisterInterval(time.Second*15),
		web.Metadata(map[string]string{"protocol": "http"}),
	)
	//接受命令行参数
	_ = server.Init()
	_ = server.Run()
}
