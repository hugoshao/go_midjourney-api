package main

import (
	"flag"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"go_midjourney-api/DiscordService"
	"go_midjourney-api/Router"
	"go_midjourney-api/Task"
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
	// 注册消息处理函数
	s.AddHandler(DiscordService.MessageCreate)
	// 注册消息更新处理函数
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
		fmt.Println("加载环境变量时出错")
		return
	}

	// 初始化DiscordGo
	go initDiscordGo()

	// 初始化Redis处理器
	Task.GetInstance(Util.GetEnvVariable("REDIS_HOST"), Util.GetEnvVariable("REDIS_PASSWORD"))

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
