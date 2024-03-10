package Task

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/go-redis/redis/v8"
	"go_midjourney-api/Models"
	"reflect"
	"strings"
	"sync"
)

// Controller 结构体包含 redis 客户端
type RedisController struct {
	RedisClient *redis.Client
}

var (
	instance *RedisController
	once     sync.Once
)

// GetInstance 返回 Controller 类的单例
func GetInstance(redisAddr, redisPassword string) *RedisController {
	once.Do(func() {
		rdb := redis.NewClient(&redis.Options{
			Addr:     redisAddr,
			Password: redisPassword, // Redis 密码
		})
		instance = &RedisController{
			RedisClient: rdb,
		}
	})
	return instance
}

// AddTask 将任务加入 Redis
func (tc *RedisController) AddTask(ctx context.Context, task Models.TaskModels) error {
	taskData, err := json.Marshal(task)
	if err != nil {
		return err
	}
	// 使用 hash 存储任务详细信息
	if err := tc.RedisClient.HSet(ctx, "tasks", task.ID, taskData).Err(); err != nil {
		return err
	}
	// 使用 set 记录所有任务ID
	if err := tc.RedisClient.SAdd(ctx, "allTasks", task.ID).Err(); err != nil {
		return err
	}
	// 使用 set 记录活跃任务ID
	if err := tc.RedisClient.SAdd(ctx, "activeTasks", task.ID).Err(); err != nil {
		return err
	}
	return nil
}

// DeleteActiveTaskByID 根据ID删除进行中的任务
func (tc *RedisController) DeleteActiveTaskByID(ctx context.Context, taskID string) error {
	// 从活跃任务集合中移除任务ID
	if err := tc.RedisClient.SRem(ctx, "activeTasks", taskID).Err(); err != nil {
		return err
	}

	// 可选：如果您也想从所有任务集合中移除此任务ID，则取消注释以下行
	// if err := tc.RedisClient.SRem(ctx, "allTasks", taskID).Err(); err != nil {
	// 	return err
	// }

	// 从任务详细信息哈希中删除任务
	if err := tc.RedisClient.HDel(ctx, "tasks", taskID).Err(); err != nil {
		return err
	}

	return nil
}

// UpdateTaskFieldByID 根据任务 ID 更新任务的指定字段值
func (tc *RedisController) UpdateTaskFieldByID(ctx context.Context, id string, field string, value interface{}) error {
	// 从 Redis 中获取任务
	taskData, err := tc.RedisClient.HGet(ctx, "tasks", id).Result()
	if err != nil {
		return err
	}

	var task Models.TaskModels
	err = json.Unmarshal([]byte(taskData), &task)
	if err != nil {
		return err
	}

	// 使用反射更新任务的指定字段
	taskValue := reflect.ValueOf(&task).Elem()
	fieldValue := taskValue.FieldByName(field)
	if !fieldValue.IsValid() {
		return fmt.Errorf("field %s does not exist in Task struct", field)
	}
	if !fieldValue.CanSet() {
		return fmt.Errorf("cannot set value for field %s", field)
	}
	fieldReflectValue := reflect.ValueOf(value)
	if fieldValue.Type() != fieldReflectValue.Type() {
		return fmt.Errorf("value type does not match field type")
	}
	fieldValue.Set(fieldReflectValue)

	// 将更新后的任务对象写回 Redis
	updatedTaskData, err := json.Marshal(task)
	if err != nil {
		return err
	}
	if err := tc.RedisClient.HSet(ctx, "tasks", id, updatedTaskData).Err(); err != nil {
		return err
	}
	return nil
}

func (tc *RedisController) UpdateTaskProperties(ctx context.Context, id string, components []discordgo.MessageComponent) error {
	// 准备更新后的 Properties
	properties := make(map[string]interface{})

	// 遍历 components 来更新 properties
	for _, component := range components {
		switch comp := component.(type) {
		case *discordgo.ActionsRow:
			for _, innerComp := range comp.Components {
				switch inner := innerComp.(type) {
				case *discordgo.Button:
					label := inner.Label
					// 检查 label 是否为空，并根据需要分配默认值
					if label == "" && inner.Emoji.Name == "🔄" {
						label = "R0"
					}

					if label != "" && inner.CustomID != "" {
						// 分割 CustomID 来获取需要的值
						parts := strings.Split(inner.CustomID, "::")
						if len(parts) > 0 {
							value := parts[len(parts)-1]
							properties[label] = value
						}
					}
					// 处理其他类型的 Component，如 Select Menus 等
				}
			}
		}
	}

	// 获取任务原始数据
	taskData, err := tc.RedisClient.HGet(ctx, "tasks", id).Result()
	if err != nil {
		return err
	}

	var task Models.TaskModels
	err = json.Unmarshal([]byte(taskData), &task)
	if err != nil {
		return err
	}

	// 更新 Properties
	task.Properties = properties

	// 将更新后的任务对象序列化并写回 Redis
	updatedTaskData, err := json.Marshal(task)
	if err != nil {
		return err
	}
	if err := tc.RedisClient.HSet(ctx, "tasks", id, updatedTaskData).Err(); err != nil {
		return err
	}

	return nil
}
