package wrappers

import (
	"context"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/micro/go-micro/v2/client"
	"micro-TodoList/api-gateway/service"
	"strconv"
)

func NewTask(id uint64, name string) *service.TaskModel {
	return &service.TaskModel{
		Id:         id,
		Title:      name,
		Content:    "响应超时",
		StartTime:  1000,
		EndTime:    1000,
		Status:     0,
		CreateTime: 1000,
		UpdateTime: 1000,
	}
}

// 降级处理函数
func DefaultTasks(resp any) {
	models := make([]*service.TaskModel, 0)
	var i uint64
	for i = 0; i < 10; i++ {
		models = append(models, NewTask(i, "降级备忘录"+strconv.Itoa(int(20+(i)))))
	}
	result := resp.(*service.TaskListResponse)
	result.TaskList = models
}

type TaskWrapper struct {
	client.Client
}

// 服务降级
func (wrapper *TaskWrapper) Call(ctx context.Context, req client.Request, resp any, opts ...client.CallOption) error {
	cmdName := req.Service() + "." + req.Endpoint()
	config := hystrix.CommandConfig{
		Timeout:                30000,
		RequestVolumeThreshold: 20,   // 熔断器请求阈值，默认20，意思是有20个请求才能进行错误百分比计算
		ErrorPercentThreshold:  50,   // 错误百分比，当错误超过百分比时，直接进行降级处理，直至熔断器再次 开启，默认50%
		SleepWindow:            5000, // 过多长时间，熔断器再次检测是否开启，单位毫秒ms（默认5秒）
	}
	hystrix.ConfigureCommand(cmdName, config)
	return hystrix.Do(cmdName, func() error {
		return wrapper.Client.Call(ctx, req, resp)
	}, func(err error) error {
		return err
	})
}

// NewTaskWrapper 初始化Wrapper
func NewTaskWrapper(c client.Client) client.Client {
	return &TaskWrapper{c}
}
