package DiscordService

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"go_midjourney-api/Util"
	"regexp"
	"strings"
)

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.ChannelID != Util.GetEnvVariable("CHANNEL_ID") {
		return
	}
	if m.GuildID != Util.GetEnvVariable("GUILD_ID") {
		return
	}

	Util.SendLog(Util.ToJson(m))
	//if m.Components != nil {
	//	Util.SendLog("Components: " + Util.ToJson(m.Components))
	//	return
	//}
	s.UpdateGameStatus(0, "Midjourney")

	mjID := Util.GetEnvVariable("APPLICATION_ID")
	if m.Author.ID == mjID {
		MissionType := adjustTaskType(m.Content)
		switch MissionType {
		case "waiting to start":
			ProcessImagineStart(m)
			return
		case "in progress":
			UpdateImagineProgress(m.Content)
			return
		case "success":
			ImagineSuccess(m)
			return
		default:
			return

		}
	}

}

func MessageUpdate(s *discordgo.Session, m *discordgo.MessageUpdate) {
	if m.ChannelID != Util.GetEnvVariable("CHANNEL_ID") {
		return
	}
	if m.GuildID != Util.GetEnvVariable("GUILD_ID") {
		return
	}

	mjID := Util.GetEnvVariable("APPLICATION_ID")
	if m.Author.ID == mjID {
		MissionType := adjustTaskType(m.Content)
		switch MissionType {
		case "in progress":
			UpdateImagineProgress(m.Content)
			return
		default:
			return

		}
	}
}

func adjustTaskType(content string) string {

	// 判断是否包含 "Waiting to start"
	if strings.Contains(content, "(Waiting to start)") {
		return "waiting to start"
	}
	if strings.Contains(content, "(relaxed, stealth)") {
		re := regexp.MustCompile(`\((\d+)%\)`)
		matches := re.FindStringSubmatch(content)
		if len(matches) > 1 {
			percentage := matches[1]
			fmt.Println("进度: ", percentage)
			return "in progress"
		} else {
			return "success"
		}
	}
	return ""
}
