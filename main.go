package main

import (
	"fmt"
	"go_midjourney-api/Router"
	"go_midjourney-api/Util"
	"log"
)

func main() {
	// 加载环境变量
	err := Util.LoadEnv()
	if err != nil {
		fmt.Println("Error loading environment variables")
		return
	}

	// 初始化任务队列
	//Task.TaskControllerInstance.AddTask(Models.TaskModels{ID: "1", Description: "测试任务", State: "running", Status: "running", Progress: "0", Prompt: "测试任务", PromptEn: "Test Task", Properties: map[string]interface{}{"test": "test"}})

	r := Router.MidjourneyApiRouter()

	err = r.Run(":8080")
	if err != nil {
		log.Fatalln("Failed to start server: ", err)
		return
	}
}
