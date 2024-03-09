package main

import (
	"flag"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"go_midjourney-api/DiscordService"
	"go_midjourney-api/Router"
	"go_midjourney-api/Util"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var Token string

// 初始化DiscordGo
func initDiscordGo() {
	botToken := Util.GetEnvVariable("BOT_TOKEN")
	flag.StringVar(&Token, "t", botToken, "Bot Token")
	flag.Parsed()
	s, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("Error creating Discord session: ", err)
		return
	}

	bot := s.Open()

	if bot != nil {
		fmt.Println("Error opening connection to Discord: ", bot)
		return
	}
	s.AddHandler(DiscordService.MessageCreate)
	s.AddHandler(DiscordService.MessageUpdate)

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
	defer s.Close()

}

func main() {
	// 加载环境变量
	err := Util.LoadEnv()
	if err != nil {
		fmt.Println("Error loading environment variables")
		return
	}

	go initDiscordGo()

	// 初始化任务队列
	/*Task.TaskControllerInstance.AddTask(Models.TaskModels{ID: "1", Description: "测试任务1", State: "running", Status: "running", Progress: "0", Prompt: "测试任务", PromptEn: "Test Task", Properties: map[string]interface{}{"test": "test"}})
	Task.TaskControllerInstance.AddTask(Models.TaskModels{ID: "2", Description: "测试任务2", State: "running", Status: "running", Progress: "0", Prompt: "测试任务", PromptEn: "Test Task", Properties: map[string]interface{}{"test": "test"}})
	Task.TaskControllerInstance.AddTask(Models.TaskModels{ID: "3", Description: "测试任务3", State: "running", Status: "running", Progress: "0", Prompt: "测试任务", PromptEn: "Test Task", Properties: map[string]interface{}{"test": "test"}})
	*/
	r := Router.MidjourneyApiRouter()

	err = r.Run(":8080")
	if err != nil {
		log.Fatalln("Failed to start server: ", err)
		return
	}

	// 无限循环，保持 initDiscordGo 函数一直运行
	for {
		time.Sleep(time.Second)
	}

}
