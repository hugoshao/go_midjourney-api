package DiscordService

import (
	"encoding/json"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"go_midjourney-api/Util"
	"regexp"
	"strings"
)

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	Util.SendLog(Util.ToJson(m))
	//if m.Components != nil {
	//	Util.SendLog("Components: " + Util.ToJson(m.Components))
	//	return
	//}
	s.UpdateGameStatus(0, "Midjourney")

	println("messageCreate | " + m.Content)
	switch adjustTaskType(m.Content) {
	case "Waiting to start":
		ProcessImagineStart(m.Content)
		return
	default:
		return

	}
}

func MessageUpdate(s *discordgo.Session, m *discordgo.MessageUpdate) {
	//if m.ChannelID != Util.GetEnvVariable("CHANNEL_ID") {
	//	return
	//}

	Util.SendLog(Util.ToJson(m))

	//if m.Author == nil {
	//	if len(m.Embeds) > 0 && len(m.Attachments) == 0 && m.Embeds[0].Type == "rich" {
	//		app.handleDescribeUpdateEvent(m)
	//	}
	//
	//	return
	//}
	//
	//if m.Author.Username == "Midjourney Bot" {
	//	if len(m.Attachments) > 0 && len(m.Content) > 0 {
	//		app.handleRateEvent(m)
	//		return
	//	}
	//}
}

func adjustTaskType(content string) string {
	err := json.Unmarshal([]byte(content), &content)
	if err != nil {
		fmt.Println("解析 JSON 出错:", err)
		return ""
	}

	// 判断是否包含(relaxed, stealth)并输出百分比数字
	if strings.Contains(content, "(Waiting to start)") {
		return "Waiting to start"
	}

	// 判断是否包含(relaxed, stealth)并输出百分比数字
	if strings.Contains(content, "(relaxed, stealth)") {
		re := regexp.MustCompile(`\((\d+)%\)`)
		matches := re.FindStringSubmatch(content)
		if len(matches) > 1 {
			percentage := matches[1]
			// TODO 处理正在进行的任务的进度
			fmt.Println("百分比数字:", percentage)
			return ""
		} else {
			// TODO 处理已完成消息
			return ""
		}
	}

	return ""

}
