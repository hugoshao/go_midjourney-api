package DiscordService

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"go_midjourney-api/Models"
	"go_midjourney-api/Util"
	"go_midjourney-api/handlers"
	"regexp"
)

// ProcessStartContent 处理新绘画消息内容，并且加入任务列表
func ProcessImagineStart(m *discordgo.MessageCreate) {
	content := m.Content

	// 获取提示词
	re := regexp.MustCompile(`\*\*(.*?)\*\*`)
	result := re.FindStringSubmatch(content)

	if len(result) > 1 {
		fmt.Println(result[1])
	} else {
		fmt.Println("没有找到提示词")
		return
	}
	TaskID := Util.GenerateID(result[1])

	task := Models.TaskModels{
		ID:          TaskID,
		Action:      m.Interaction.Name,
		Description: "任务已创建等待生成",
		FailReason:  "",
		StartTime:   0,
		Progress:    "Waiting to start",
		Prompt:      result[1],
		PromptEn:    result[1],
	}

	handlers.AddTask(task)

}

// UpdateImagineProgress 更新绘画任务的进度
func UpdateImagineProgress(c string) {
	re := regexp.MustCompile(`\((\d+)%\)`)
	matches := re.FindStringSubmatch(c)
	percentage := matches[1]

	// 获取提示词
	po := regexp.MustCompile(`\*\*(.*?)\*\*`)
	poRes := po.FindStringSubmatch(c)

	if len(poRes) > 1 {
		fmt.Println(poRes[1])
	} else {
		fmt.Println("没有找到提示词")
		return
	}
	TaskID := Util.GenerateID(poRes[1])
	handlers.UpdateTask(TaskID, "Progress", percentage)
	handlers.UpdateTask(TaskID, "Description", "任务进行中")
}

func ImagineSuccess(m *discordgo.MessageCreate) {
	content := m.Content

	// 获取提示词
	re := regexp.MustCompile(`\*\*(.*?)\*\*`)
	result := re.FindStringSubmatch(content)

	if len(result) > 1 {
		fmt.Println(result[1])
	} else {
		fmt.Println("没有找到提示词")
		return
	}
	TaskID := Util.GenerateID(result[1])

	imageUrl := m.Attachments[0].URL
	handlers.UpdateTask(TaskID, "Progress", "Success")
	handlers.UpdateTask(TaskID, "ImageURL", imageUrl)
	handlers.UpdateTask(TaskID, "Description", "任务已完成")
	handlers.ClearActiveTaskByID(TaskID)
	handlers.UpdateTaskProperties(TaskID, m.Components)
}
