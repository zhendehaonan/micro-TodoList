package conf

import (
	"gopkg.in/ini.v1"
	"micro-TodoList/mq-server/model"
	"strings"
)

var (
	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string

	RabbitMQ         string
	RabbitMQUser     string
	RabbitMQPassword string
	RabbitMQHost     string
	RabbitMQPort     string
)

func Init() {
	file, err := ini.Load("./conf/config.ini")
	if err != nil {
		panic(err)
	}

	//连接mysql
	LoadMysql(file)
	connString := strings.Join([]string{DbUser, ":", DbPassword, "@tcp(", DbHost, ":", DbPort, ")/", DbName, "?charset=utf8mb4&parseTime=true"}, "")
	model.Database(connString)

	//连接RabbitMQ
	LoadRabbitMQ(file)
	pathRabbitMQ := strings.Join([]string{RabbitMQ, "://", RabbitMQUser, ":", RabbitMQPassword, "@", RabbitMQHost, ":", RabbitMQPort, "/"}, "")
	model.RabbitMQ(pathRabbitMQ)
}

func LoadMysql(file *ini.File) {
	Db = file.Section("mysql").Key("DB").String()
	DbHost = file.Section("mysql").Key("DbHost").String()
	DbPort = file.Section("mysql").Key("DbPort").String()
	DbUser = file.Section("mysql").Key("DbUser").String()
	DbPassword = file.Section("mysql").Key("DbPassword").String()
	DbName = file.Section("mysql").Key("DbName").String()
}

func LoadRabbitMQ(file *ini.File) {
	RabbitMQ = file.Section("rabbitmq").Key("RabbitMQ").String()
	RabbitMQUser = file.Section("rabbitmq").Key("RabbitMQUser").String()
	RabbitMQPassword = file.Section("rabbitmq").Key("RabbitMQPassword").String()
	RabbitMQHost = file.Section("rabbitmq").Key("RabbitMQHost").String()
	RabbitMQPort = file.Section("rabbitmq").Key("RabbitMQPort").String()

}
