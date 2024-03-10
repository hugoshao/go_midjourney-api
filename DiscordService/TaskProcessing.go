package DiscordService

import (
	"encoding/json"
	"fmt"
	"go_midjourney-api/Models"
	"go_midjourney-api/handlers"
	"regexp"
	"time"
)

// ProcessStartContent 处理新绘画消息内容，并且加入任务列表
func ProcessImagineStart(cont string) {
	var content Models.DiscordMessage
	err := json.Unmarshal([]byte(cont), &content)
	if err != nil {
		return
	}

	t, err := time.Parse(time.RFC3339Nano, content.Timestamp)
	if err != nil {
		fmt.Println("解析时间戳出错:", err)
		return
	}
	// 将时间转换为 Unix 时间戳（以秒为单位）
	timestamp := t.Unix()

	// 获取提示词
	var middleContent string
	re := regexp.MustCompile(`\*\*(.*?)\*\*`)
	matches := re.FindStringSubmatch(content.Content)
	if len(matches) > 1 {
		middleContent = matches[1]
	}

	task := Models.TaskModels{
		ID:          content.Interaction.ID,
		Action:      content.Interaction.Name,
		Description: "任务已创建等待生成",
		FailReason:  "",
		StartTime:   timestamp,
		Progress:    "Waiting to start",
		Prompt:      middleContent,
		PromptEn:    middleContent,
	}

	handlers.AddTask(task)

}
