package core

import (
	"micro-TodoList/task/model"
	"micro-TodoList/task/service"
)

func BuildTask(item *model.Task) *service.TaskModel {
	return &service.TaskModel{
		Id:         uint64(item.ID),
		Uid:        uint64(item.Uid),
		Title:      item.Title,
		Content:    item.Content,
		StartTime:  item.StartTime,
		EndTime:    item.EndTime,
		Status:     int64(item.Status),
		CreateTime: item.CreatedAt.Unix(),
		UpdateTime: item.UpdatedAt.Unix(),
	}
}
