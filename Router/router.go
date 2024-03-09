package Router

import (
	"github.com/gin-gonic/gin"
)

func MidjourneyApiRouter() *gin.Engine {
	r := gin.Default()
	// 处理 SUBMIT 请求
	r.POST("/mj/submit/*action", HandleSubmitRequest)
	// 处理 TASK 请求
	r.GET("/mj/task/*action", HandleTaskRequest)

	return r
}
