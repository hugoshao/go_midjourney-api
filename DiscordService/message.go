package DiscordService

import (
	"github.com/bwmarrin/discordgo"
	"go_midjourney-api/Util"
)

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	Util.SendLog(Util.ToJson(m))
	if m.Components != nil {
		Util.SendLog("Components: " + Util.ToJson(m.Components))
		return
	}
	s.UpdateGameStatus(0, "Midjourney")
	println("messageCreate | " + m.Content)
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
