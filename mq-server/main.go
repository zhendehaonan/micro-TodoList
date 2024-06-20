package main

import (
	"micro-TodoList/mq-server/conf"
	"micro-TodoList/mq-server/service"
)

func main() {
	conf.Init()
	forever := make(chan bool)
	service.CreateTask()
	<-forever
}
