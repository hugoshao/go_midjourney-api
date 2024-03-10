package DiscordService

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go_midjourney-api/Models"
	"go_midjourney-api/Util"
	"net/http"
	"strings"
)

func Simple(c *gin.Context, bodyString string) {
	// 创建一个结构体实例，用于存储解析后的数据
	var requestBody Models.SimpleRequest

	// 解析 JSON 字符串到结构体
	err := json.Unmarshal([]byte(bodyString), &requestBody)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": "请求体错误", "body": err})
		return
	}
	parts := strings.Split(requestBody.Content, " ")

	if len(parts) != 2 {
		c.JSON(http.StatusOK, gin.H{"message": "请求体错误", "body": "内容错误"})
		return
	}

	//beforeSpace := parts[0]
	afterSpace := parts[1]
	afterSpaceParts := strings.Split(afterSpace, "")
	if len(afterSpaceParts) != 2 {
		c.JSON(http.StatusOK, gin.H{"message": "请求体错误", "body": "内容错误"})
	}

	//ID := beforeSpace
	// 操作
	//firstChar := afterSpaceParts[0]
	// 操作的数字
	//secondChar := afterSpaceParts[1]
	// TODO 处理绘画变换任务请求,向Discord发送请求

	ApplicationID := Util.GetEnvVariable("APPLICATION_ID")
	GuildID := Util.GetEnvVariable("GUILD_ID")
	ChannelID := Util.GetEnvVariable("CHANNEL_ID")
	SessionID := Util.GetEnvVariable("SESSION_ID")

	interactionsReq := &Models.InteractionsRequest{
		Type:          2,
		ApplicationID: ApplicationID,
		GuildID:       GuildID,
		ChannelID:     ChannelID,
		SessionID:     SessionID,
		Data: map[string]any{
			"version": "1166847114203123795",
			"id":      "938956540159881230",
			"name":    "imagine",
			"type":    "1",
			"options": []map[string]any{
				{
					"type":  3,
					"name":  "prompt",
					"value": "requestBody.Prompt",
				},
			},
			"application_command": map[string]any{
				"id":                         "938956540159881230",
				"application_id":             ApplicationID,
				"version":                    "1166847114203123795",
				"default_permission":         true,
				"default_member_permissions": nil,
				"type":                       1,
				"nsfw":                       false,
				"name":                       "imagine",
				"description":                "Create images with Midjourney",
				"dm_permission":              true,
				"options": []map[string]any{
					{
						"type":           3,
						"name":           "prompt",
						"description":    "The prompt to imagine",
						"required":       true,
						"name_localized": "prompt",
					},
				},
				"attachments": []any{},
			},
		},
	}
	b, _ := json.Marshal(interactionsReq)
	if Util.HttpToDiscord(b) {
		c.JSON(http.StatusOK, gin.H{"message": "Success", "body": "请求成功"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Error", "body": "错误的请求"})
		return
	}

}
