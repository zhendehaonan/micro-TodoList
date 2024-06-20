package service

import (
	"encoding/json"
	"log"
	"micro-TodoList/mq-server/model"
)

// 从RabbitMQ中接受数据，写入数据库
func CreateTask() {
	//mq的连接
	ch, err := model.MQ.Channel()
	if err != nil {
		panic(err)
	}
	//队列的声明
	q, _ := ch.QueueDeclare("task_queue", true, false, false, false, nil)
	//消息消费
	err = ch.Qos(1, 0, false) //每次消费一个消息
	msgs, err := ch.Consume(q.Name, "", false, false, false, false, nil)
	if err != nil {
		panic(err)
	}
	//消费端始终要处于一个监听状态，一直监听生产端的生产，所以这里我们要阻塞主进程
	go func() {
		for d := range msgs {
			var t model.Task
			err := json.Unmarshal(d.Body, &t)
			if err != nil {
				panic(err)
			}
			DB := model.NewDBClient()
			DB.Create(&t)
			log.Println("Done")
			_ = d.Ack(false)
		}
	}()
}
