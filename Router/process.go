package Router

import (
	"github.com/gin-gonic/gin"
	"go_midjourney-api/Util"
	"go_midjourney-api/handlers"
	"io"
	"net/http"
	"strings"
)

func HandleSubmitRequest(c *gin.Context) {
	DebugMode := Util.GetEnvVariable("DEBUG_MODE")
	bodyBytes, _ := io.ReadAll(c.Request.Body)
	bodyString := string(bodyBytes)

	if DebugMode == "True" {
		Util.SendLog("收到Submit请求：" + bodyString)
	}

	// 获取请求的路径
	path := c.Request.URL.Path
	// 根据不同的路径进行不同的处理
	switch path {
	case "/mj/submit/imagine":
		// TODO 处理绘画请求,向Discord发送请求
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error", "body": "错误的请求"})
		return
	case "/mj/submit/simple-change":
		// TODO 处理绘画变换任务请求,向Discord发送请求
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error", "body": "错误的请求"})
		return
	case "/mj/submit/describe":
		// TODO 处理描述（Describe）任务请求,向Discord发送请求
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error", "body": "错误的请求"})
		return
	case "/mj/submit/blend":
		// TODO 处理图片混合（blend）任务请求,向Discord发送请求
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error", "body": "错误的请求"})
		return
	case "/mj/submit/change":
		// TODO 处理绘图（change）任务请求,向Discord发送请求
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error", "body": "错误的请求"})
		return
	default:
		// 处理其他路径的逻辑
		c.JSON(404, gin.H{"error": "Not Found"})
		return
	}

}

func HandleTaskRequest(c *gin.Context) {
	DebugMode := Util.GetEnvVariable("DEBUG_MODE")
	bodyBytes, _ := io.ReadAll(c.Request.Body)
	bodyString := string(bodyBytes)

	if DebugMode == "True" {
		Util.SendLog("收到TASK请求：" + bodyString)
	}

	// 获取请求的路径
	path := c.Request.URL.Path
	// 根据不同的路径进行不同的处理
	switch path {
	case "/mj/task/list-by-condition":
		// TODO 处理根据条件获取任务列表请求
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error", "body": "错误的请求"})
		return
	case "/mj/task/queue":
		handlers.GetActiveTask(c)
		return
	case "/mj/task/queue/next":
		handlers.GetNextTask(c)
		return
	case "/mj/task/list":
		handlers.GetAllTask(c)
		return
	default:
		// 处理其他路径的逻辑
		// 判断是否是动态路径
		if strings.HasPrefix(path, "/mj/task/") && strings.HasSuffix(path, "/fetch") {
			// 提取动态路径参数 {id}
			id := strings.TrimPrefix(path, "/mj/task/")
			id = strings.TrimSuffix(id, "/fetch")

			handlers.GetTaskByID(c, id)
			return
		} else {
			c.JSON(404, gin.H{"error": "Not Found"})
			return
		}
	}
}
