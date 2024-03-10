package Router

import (
	"github.com/gin-gonic/gin"
	"go_midjourney-api/Util"
	"net/http"
)

func MidjourneyApiRouter() *gin.Engine {
	r := gin.Default()
	ApiSecret := Util.GetEnvVariable("API_SECRET")
	if ApiSecret != "" {
		// 处理 SUBMIT 请求
		r.POST("/mj/submit/*action", authMiddleware(), HandleSubmitRequest)
		// 处理 TASK 请求
		r.GET("/mj/task/*action", authMiddleware(), HandleTaskRequest)
	} else {
		// 处理 SUBMIT 请求
		r.POST("/mj/submit/*action", HandleSubmitRequest)
		// 处理 TASK 请求
		r.GET("/mj/task/*action", HandleTaskRequest)

	}

	return r
}

// 验证请求
func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizationHeader := c.GetHeader("Authorization")
		ApiSecret := Util.GetEnvVariable("API_SECRET")

		if authorizationHeader == "" || authorizationHeader != ApiSecret {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		c.Next()
	}
}
