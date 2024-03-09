package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

const (
	// Discord WebSocket地址
	DISCORD_GATEWAY = "wss://gateway.discord.gg/?v=9&encoding=json"

	// 你的 Discord Bot Token
	BOT_TOKEN = "MjU5OTQ3NzMyNzA4NDkxMjY1.GBGOSG.MwcbOKjROQNY6RyVYgcxfSOsEjtSM9vTBRbJVI"
	TEST_TOKE = "MTIxNTUzMTg5NDU1NzY0Mjc2Mg.GW2myz.drmctqSpipVl9ZyNO5Esk_rToBhpRMsLLqHd7U"
)

func main2() {
	//chennalID := "1086912786426241114"
	//guildID := "858638164059095060"
	//dialSocksProxy, _ := proxy.SOCKS5("tcp", "127.0.0.1:7890", nil, proxy.Direct)

	discord, err := discordgo.New("Bot " + TEST_TOKE)
	if err != nil {
		err.Error()
	}

	err = discord.Open()
	if err != nil {
		fmt.Println("Error opening Discord session: ", err)
		return
	}
	defer discord.Close()

	discord.AddHandler(newMessage)
	discord.AddHandler(updateMessage)

	//discord.ChannelMessageSend(chennalID, "test")
	//discord.Identify.Intents = discordgo.IntentsAllWithoutPrivileged
	/*discord.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == discord.State.User.ID {
			return
		}
		println(m.Content)
		// Respond to messages
		switch {
		case strings.Contains(m.Content, "weather"):
			discord.ChannelMessageSend(m.ChannelID, "I can help with that!")
		case strings.Contains(m.Content, "bot"):
			discord.ChannelMessageSend(m.ChannelID, "Hi there!")
		}
	})*/
	/*discord.Identify.Intents = discordgo.IntentsAllWithoutPrivileged
	discord.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		s.ChannelMessageSend(chennalID, "HelloDiscord")
	})


	send, err := discord.ChannelMessageSend(chennalID, "TestMessage")
	println(send)*/

	/*err = discord.Open()
	if err != nil {
		println("cuowu")
		println(err.Error())
	}*/

	println("the bot is online")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	/*// 建立WebSocket连接
	conn, _, err := websocket.DefaultDialer.Dial(DISCORD_GATEWAY, nil)
	if err != nil {
		log.Fatal("Error connecting to Discord:", err)
	}
	defer conn.Close()

	// 处理中断信号，以便在退出时正常关闭连接
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	// 发送身份验证信息
	authData := map[string]interface{}{
		"op": 2,
		"d": map[string]interface{}{
			"token":   BOT_TOKEN,
			"intents": 513, // 你的 Bot 的意图
			"properties": map[string]string{
				"$os":      "linux",
				"$browser": "my_library",
				"$device":  "my_library",
			},
		},
	}
	if err := conn.WriteJSON(authData); err != nil {
		log.Fatal("Error sending auth data:", err)
	}

	// 启动一个goroutine用于接收消息
	go func() {
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Println("Error reading message:", err)
				return
			}
			fmt.Println("Received message:", string(message))
		}
	}()

	// 等待中断信号
	select {
	case <-interrupt:
		log.Println("Received interrupt signal, closing connection...")
		// 关闭连接
		err := conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
		if err != nil {
			log.Println("Error closing connection:", err)
		}
		time.Sleep(time.Second)
		return
	}*/
}

func newMessage(discord *discordgo.Session, message *discordgo.MessageCreate) {

	// Ignore bot messaage
	if message.Author.ID == discord.State.User.ID {
		return
	}
	println(message.Content)
	println("ATT:")
	println("..........................................")
	//println(updateMessage.Content)
	// Respond to messages
	switch {
	case strings.Contains(message.Content, "weather"):
		discord.ChannelMessageSend(message.ChannelID, "I can help with that!")
	case strings.Contains(message.Content, "bot"):
		discord.ChannelMessageSend(message.ChannelID, "Hi there!")
	}
}

func updateMessage(discord *discordgo.Session, updateMessage *discordgo.MessageUpdate) {
	println(updateMessage.Content)
}

func userSendMessage() {

}

/*func ReRoll(messageId string, messageHash string) error {
	requestBody := ReqResetDiscord{
		Type:          3,
		GuildId:       config.GetConfig().DISCORD_SERVER_ID,
		ChannelId:     config.GetConfig().DISCORD_CHANNEL_ID,
		MessageFlags:  0,
		MessageId:     messageId,
		ApplicationId: appId,
		SessionId:     "45bc04dd4da37141a5f73dfbfaf5bdcf",
		Data: UpscaleData{
			ComponentType: 2,
			CustomId:      fmt.Sprintf("MJ::JOB::reroll::0::%s::SOLO", messageHash),
		},
	}
	res, err := request(requestBody, url)
	return err
}*/
