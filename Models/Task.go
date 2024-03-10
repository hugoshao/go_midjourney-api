package Models

// TaskModels 结构体表示任务的信息
type TaskModels struct {
	ID          string                 `json:"id"`
	Action      string                 `json:"action"`
	Description string                 `json:"description"`
	FailReason  string                 `json:"failReason"`
	StartTime   int64                  `json:"startTime"`
	FinishTime  int64                  `json:"finishTime"`
	ImageURL    string                 `json:"imageUrl"`
	Progress    string                 `json:"progress"`
	Prompt      string                 `json:"prompt"`
	PromptEn    string                 `json:"promptEn"`
	Properties  map[string]interface{} `json:"properties"`
	State       string                 `json:"state"`
	Status      string                 `json:"status"`
	SubmitTime  int64                  `json:"submitTime"`
}

type DiscordMessage struct {
	ID              string      `json:"id"`
	ChannelID       string      `json:"channel_id"`
	GuildID         string      `json:"guild_id"`
	Content         string      `json:"content"`
	Timestamp       string      `json:"timestamp"`
	WebhookID       string      `json:"webhook_id"`
	EditedTimestamp interface{} `json:"edited_timestamp"`
	Interaction     Interaction `json:"interaction"`
	// Other fields are omitted for brevity
}

type Interaction struct {
	ID   string `json:"id"`
	Type int    `json:"type"`
	Name string `json:"name"`
}
