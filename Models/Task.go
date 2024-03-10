package Models

// TaskModels 结构体表示任务的信息
type TaskModels struct {
	ID          string                 `json:"id"`          // Midjourney 任务 ID
	Action      string                 `json:"action"`      // Midjourney 任务类型
	Description string                 `json:"description"` // 任务当前状态描述
	FailReason  string                 `json:"failReason"`  // 任务失败原因
	StartTime   int64                  `json:"startTime"`   // 任务开始时间
	FinishTime  int64                  `json:"finishTime"`  // 任务结束时间
	ImageURL    string                 `json:"imageUrl"`    //
	Progress    string                 `json:"progress"`    // 任务进度
	Prompt      string                 `json:"prompt"`      // 提示词
	PromptEn    string                 `json:"promptEn"`    // 提示词英文
	Properties  map[string]interface{} `json:"properties"`  // 任务属性
	State       string                 `json:"state"`
	Status      string                 `json:"status"`
	SubmitTime  int64                  `json:"submitTime"`
}

type DiscordMessage struct {
	ID              string       `json:"id"`
	ChannelID       string       `json:"channel_id"`
	GuildID         string       `json:"guild_id"`
	Content         string       `json:"content"`
	Timestamp       string       `json:"timestamp"`
	WebhookID       string       `json:"webhook_id"`
	EditedTimestamp interface{}  `json:"edited_timestamp"`
	Interaction     Interaction  `json:"interaction"`
	Attachments     []Attachment `json:"attachments"`
	// Other fields are omitted for brevity
}

type Interaction struct {
	ID   string `json:"id"`
	Type int    `json:"type"`
	Name string `json:"name"`
}

type Attachment struct {
	ID          string `json:"id"`
	URL         string `json:"url"`
	ProxyURL    string `json:"proxy_url"`
	FileName    string `json:"filename"`
	ContentType string `json:"content_type"`
	Width       int    `json:"width"`
	Height      int    `json:"height"`
	Size        int    `json:"size"`
	Ephemeral   bool   `json:"ephemeral"`
}

type Component struct {
	Label    string                 `json:"label"`
	Style    int                    `json:"style"`
	Disabled bool                   `json:"disabled"`
	Emoji    map[string]interface{} `json:"emoji"`
	CustomID string                 `json:"custom_id"`
	Type     int                    `json:"type"`
}

type ActionItem struct {
	Components []Component `json:"components"`
	Type       int         `json:"type"`
}

type Actions []ActionItem
