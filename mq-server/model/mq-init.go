package model

import "github.com/streadway/amqp"

//mq库的引入===》go get github.com/streadway/amqp

var MQ *amqp.Connection

func RabbitMQ(connString string) {
	conn, err := amqp.Dial(connString)
	if err != nil {
		panic(err)
	}
	MQ = conn
}
