package core

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/streadway/amqp"
	"micro-TodoList/task/model"
	"micro-TodoList/task/service"
)

// 创建备忘录，将备忘录信息生产,放到rabbitMQ消息队列中
func (*TaskService) CreateTask(ctx context.Context, req *service.TaskRequest, resp *service.TaskDetailResponse) error {
	//mq的连接
	ch, err := model.MQ.Channel()
	if err != nil {
		err = errors.New("rabbitMQ channel err:" + err.Error())
	}
	//队列的声明
	q, _ := ch.QueueDeclare("task_queue", true, false, false, false, nil)
	body, _ := json.Marshal(req) //请求信息序列化
	//发布消息到队列
	err = ch.Publish("", q.Name, false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "application/json",
		Body:         body,
	})
	if err != nil {
		err = errors.New("rabbitMQ Publish err:" + err.Error())
	}
	return nil
}

// 实现备忘录服务接口，获取备忘录列表
func (*TaskService) GetTaskList(ctx context.Context, req *service.TaskRequest, resp *service.TaskListResponse) error {
	if req.Limit == 0 {
		req.Limit = 10
	}
	var taskData []*model.Task
	var count int64
	DB := model.NewDBClient()
	err := DB.Offset(int(req.Start)).Limit(int(req.Limit)).Where("uid=?", req.Uid).Find(&taskData).Count(&count).Error
	if err != nil {
		return errors.New("GetTaskList err:" + err.Error())
	}
	var taskRes []*service.TaskModel
	for _, item := range taskData {
		taskRes = append(taskRes, BuildTask(item))
	}
	resp.TaskList = taskRes
	resp.Count = uint32(count)
	return nil
}

// 获取备忘录的详细信息
func (*TaskService) GetTask(ctx context.Context, req *service.TaskRequest, resp *service.TaskDetailResponse) error {
	task := model.Task{}
	DB := model.NewDBClient()
	err := DB.Model(&model.Task{}).Where("id = ?", req.Id).First(&task).Error
	if err != nil {
		return errors.New(" GetTask err:" + err.Error())
	}
	taskRes := BuildTask(&task)
	resp.TaskDetail = taskRes
	return nil
}

// 修改备忘录
func (*TaskService) UpdateTask(ctx context.Context, req *service.TaskRequest, resp *service.TaskDetailResponse) error {
	task := model.Task{}
	DB := model.NewDBClient()
	err := DB.Model(&model.Task{}).Where("id = ? AND uid=?", req.Id, req.Uid).First(&task).Error
	if err != nil {
		return errors.New(" UpdateTask err:" + err.Error())
	}
	task.Title = req.Title
	task.Status = int(req.Status)
	task.Content = req.Content
	err = DB.Model(&model.Task{}).Where("id = ? AND uid=?", req.Id, req.Uid).Updates(&task).Error
	if err != nil {
		return errors.New(" UpdateTask err:" + err.Error())
	}
	taskRes := BuildTask(&task)
	resp.TaskDetail = taskRes
	return nil
}

// 修改备忘录
func (*TaskService) DeleteTask(ctx context.Context, req *service.TaskRequest, resp *service.TaskDetailResponse) error {
	DB := model.NewDBClient()
	err := DB.Model(&model.Task{}).Where("id = ? AND uid=?", req.Id, req.Uid).Delete(&model.Task{}).Error
	if err != nil {
		return errors.New(" DeleteTask err:" + err.Error())
	}
	return nil
}
