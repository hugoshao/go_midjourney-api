package DiscordService

import (
	"encoding/json"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"go_midjourney-api/Models"
	"go_midjourney-api/Util"
	"os"
	"os/signal"
	"syscall"
)

func init() {
	BOT_TOKEN := Util.GetEnvVariable("BOT_TOKEN")
	discord, err := discordgo.New("Bot " + BOT_TOKEN)
	if err != nil {
		println("Discordgo New err: " + err.Error())
	}

	discord.Open()
	defer discord.Close()

	//TODO 添加消息监听

	println("discord bot is online")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}

func DcNewMessageHandler(session *discordgo.Session, message *discordgo.MessageCreate) {
	//TODO Ignore bot message
	if message.Author.ID == session.State.User.ID {
		return
	}

	//TODO 监听新消息进行处理
	println(message.Content)
}

func DcUpdateMessageHandler(session *discordgo.Session, message *discordgo.MessageUpdate) {
	//TODO 监听消息更新进行处理
	println(message.Content)
}

func Imagine(body []byte) {
	discordAPIURL := "http://discord.com"
	var req Models.ImagineRequest
	json.Unmarshal(body, req)

	//敏感词处理

	data, err := json.Marshal(req)
	if err != nil {
		fmt.Println("序列化出错,错误原因: ", err)
		return
	}
	res, err := Util.HTTPPost(discordAPIURL, data, nil, Util.GetEnvVariable("USER_TOKEN"))
	if err != nil {
		fmt.Println("POST 错误：", err)
		return
	}
	fmt.Printf("POST响应", res)
}
