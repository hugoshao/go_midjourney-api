package Models

// ImagineRequest 提交绘画（Imagine）任务
type ImagineRequest struct {
	Prompt string  `json:"prompt"`
	Base64 *string `json:"base64"`
}

// SimpleRequest 提交绘画变换（Simple）任务
type SimpleRequest struct {
	Content string `json:"content"`
}

// DescribeRequest 提交描述（Describe）任务
type DescribeRequest struct {
	Base64 string `json:"base64"`
}
