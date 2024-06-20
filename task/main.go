package main

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"micro-TodoList/task/conf"
	"micro-TodoList/task/core"
	"micro-TodoList/task/service"
)

func main() {
	conf.Init()
	//etcd注册件
	etcdReg := etcd.NewRegistry(
		registry.Addrs("127.0.0.1:2379"), //etcd地址
	)
	//得到一个微服务实例
	microService := micro.NewService(
		micro.Name("rpcTaskService"),    //微服务名字
		micro.Address("127.0.0.1:8083"), //微服务地址
		micro.Registry(etcdReg),         // etcd注册件
	)
	//结构命令行参数.初始化
	microService.Init()
	//将服务注册到服务发现中
	_ = service.RegisterTaskServiceHandler(microService.Server(), new(core.TaskService))
	//启动微服务
	_ = microService.Run()
}
