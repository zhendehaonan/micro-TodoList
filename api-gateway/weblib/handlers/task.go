package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	"micro-TodoList/api-gateway/pkg/utils"
	"micro-TodoList/api-gateway/service"
	"net/http"
	"strconv"
)

// 获取备忘录列表
func GetTaskList(ginCtx *gin.Context) {
	var taskReq service.TaskRequest
	PanicIfTaskError(ginCtx.Bind(&taskReq))
	//从gin.Keys中取出服务实例
	taskService := ginCtx.Keys["taskService"].(service.TaskService)
	claim, _ := utils.ParseToken(ginCtx.GetHeader("Authorization"))
	taskReq.Uid = uint64(claim.ID)
	//调用服务端的函数
	taskResp, err := taskService.GetTaskList(context.Background(), &taskReq)
	PanicIfTaskError(err)
	ginCtx.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"task":  taskResp.TaskList,
			"count": taskResp.Count,
		},
	})
}

// 获取备忘录详细信息
func GetTaskDetail(ginCtx *gin.Context) {
	var taskReq service.TaskRequest
	PanicIfTaskError(ginCtx.Bind(&taskReq))
	//从gin.Keys中取出服务实例
	taskService := ginCtx.Keys["taskService"].(service.TaskService)
	claim, _ := utils.ParseToken(ginCtx.GetHeader("Authorization"))
	taskReq.Uid = uint64(claim.ID)
	id, _ := strconv.Atoi(ginCtx.Param("id"))
	taskReq.Id = uint64(id)
	//调用服务端的函数
	taskResp, err := taskService.GetTask(context.Background(), &taskReq)
	PanicIfTaskError(err)
	ginCtx.JSON(http.StatusOK, gin.H{
		"data": taskResp.TaskDetail,
	})
}

// 创建备忘录
func CreateTask(ginCtx *gin.Context) {
	var taskReq service.TaskRequest
	PanicIfTaskError(ginCtx.Bind(&taskReq))
	//从gin.Keys中取出服务实例
	taskService := ginCtx.Keys["taskService"].(service.TaskService)
	claim, _ := utils.ParseToken(ginCtx.GetHeader("Authorization"))
	taskReq.Uid = uint64(claim.ID)
	//调用服务端的函数
	taskResp, err := taskService.CreateTask(context.Background(), &taskReq)
	PanicIfTaskError(err)
	ginCtx.JSON(http.StatusOK, gin.H{
		"data": taskResp.TaskDetail,
	})
}

// 更新备忘录
func UpdateTask(ginCtx *gin.Context) {
	var taskReq service.TaskRequest
	PanicIfTaskError(ginCtx.Bind(&taskReq))
	//从gin.Keys中取出服务实例
	taskService := ginCtx.Keys["taskService"].(service.TaskService)
	claim, _ := utils.ParseToken(ginCtx.GetHeader("Authorization"))
	taskReq.Uid = uint64(claim.ID)
	id, _ := strconv.Atoi(ginCtx.Param("id"))
	taskReq.Id = uint64(id)
	//调用服务端的函数
	taskResp, err := taskService.UpdateTask(context.Background(), &taskReq)
	PanicIfTaskError(err)
	ginCtx.JSON(http.StatusOK, gin.H{
		"data": taskResp.GetTaskDetail(),
	})
}

// 删除备忘录
func DeleteTask(ginCtx *gin.Context) {
	var taskReq service.TaskRequest
	PanicIfTaskError(ginCtx.Bind(&taskReq))
	//从gin.Keys中取出服务实例
	taskService := ginCtx.Keys["taskService"].(service.TaskService)
	claim, _ := utils.ParseToken(ginCtx.GetHeader("Authorization"))
	taskReq.Uid = uint64(claim.ID)
	id, _ := strconv.Atoi(ginCtx.Param("id"))
	taskReq.Id = uint64(id)
	//调用服务端的函数
	taskResp, err := taskService.DeleteTask(context.Background(), &taskReq)
	PanicIfTaskError(err)
	ginCtx.JSON(http.StatusOK, gin.H{
		"data": taskResp.GetTaskDetail(),
	})
}
