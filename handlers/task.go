package handlers

import (
	"context"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/gin-gonic/gin"
	"go_midjourney-api/Models"
	"go_midjourney-api/Task"
	"go_midjourney-api/Util"
)

// GetAllTask 获取所有任务的处理函数
func GetAllTask(c *gin.Context) {
	//// 获取所有任务
	//tasks := Task.TaskControllerInstance.GetAllTasks()
	//if len(tasks) == 0 {
	//	c.JSON(http.StatusOK, gin.H{"message": "No tasks"})
	//	return
	//}
	//// 返回所有任务
	//c.JSON(http.StatusOK, tasks)
}

// GetNextTask 获取下一个任务的处理函数
func GetNextTask(c *gin.Context) {
	//// 获取所有任务
	//tasks := Task.TaskControllerInstance.GetActiveTaskByIndex()
	//if tasks == nil {
	//	c.JSON(http.StatusOK, gin.H{"message": "No tasks"})
	//	return
	//}
	//// 返回所有任务
	//c.JSON(http.StatusOK, tasks)
}

// GetActiveTask 获取进行中的任务的处理函数
func GetActiveTask(c *gin.Context) {
	//// 获取所有任务
	//tasks := Task.TaskControllerInstance.GetActiveTasks()
	//if len(tasks) == 0 {
	//	c.JSON(http.StatusOK, gin.H{"message": "No tasks"})
	//	return
	//}
	//// 返回所有任务
	//c.JSON(http.StatusOK, tasks)
}

// GetTaskByID 根据 ID 获取任务的处理函数
func GetTaskByID(c *gin.Context, id string) {
	//// 根据 ID 获取任务
	//task := Task.TaskControllerInstance.GetTaskByID(id)
	//if task == nil {
	//	c.JSON(http.StatusOK, gin.H{"message": "No task"})
	//	return
	//}
	//// 返回任务
	//c.JSON(http.StatusOK, task)
}

// ClearActiveTaskByID 从队列中清除进行中的任务
func ClearActiveTaskByID(id string) {
	//Task.TaskControllerInstance.ClearActiveTaskByID(id)
}

func AddTask(task Models.TaskModels) {
	// 获取所有任务
	//Task.TaskControllerInstance.AddTask(task)
	Redis := Task.GetInstance(Util.GetEnvVariable("REDIS_HOST"), Util.GetEnvVariable("REDIS_PASSWORD"))
	ctx := context.Background()
	Redis.AddTask(ctx, task)
}

func UpdateTask(id string, field string, value interface{}) {
	Redis := Task.GetInstance(Util.GetEnvVariable("REDIS_HOST"), Util.GetEnvVariable("REDIS_PASSWORD"))
	ctx := context.Background()
	err := Redis.UpdateTaskFieldByID(ctx, id, field, value)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func UpdateTaskProperties(id string, components []discordgo.MessageComponent) {
	Redis := Task.GetInstance(Util.GetEnvVariable("REDIS_HOST"), Util.GetEnvVariable("REDIS_PASSWORD"))
	ctx := context.Background()
	err := Redis.UpdateTaskProperties(ctx, id, components)
	if err != nil {
		fmt.Println(err)
		return
	}
}
