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

	r := Router.MidjourneyApiRouter()

	err = r.Run(":8080")
	if err != nil {
		log.Fatalln("Failed to start server: ", err)
		return
	}
}
