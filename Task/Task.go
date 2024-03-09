package Task

import "go_midjourney-api/Models"

// Controller 结构体包含任务队列和相关方法
type Controller struct {
	ActiveTasks []Models.TaskModels
	AllTasks    []Models.TaskModels
}

var TaskControllerInstance = Controller{}

// AddTask 将任务加入队列
func (tc *Controller) AddTask(task Models.TaskModels) {
	tc.AllTasks = append(tc.AllTasks, task)
	tc.ActiveTasks = append(tc.ActiveTasks, task)
}

// GetActiveTasks 查询进行中的任务队列
func (tc *Controller) GetActiveTasks() []Models.TaskModels {
	return tc.ActiveTasks
}

// GetAllTasks 查询所有任务
func (tc *Controller) GetAllTasks() []Models.TaskModels {
	return tc.AllTasks
}

// GetTaskByID 根据指定 ID 获取任务
func (tc *Controller) GetTaskByID(id string) *Models.TaskModels {
	for _, task := range tc.AllTasks {
		if task.ID == id {
			return &task
		}
	}
	return nil
}

// ClearActiveTaskByID 根据任务 ID 清除进行中的任务
func (tc *Controller) ClearActiveTaskByID(id string) {
	var updatedActiveTasks []Models.TaskModels
	for _, task := range tc.ActiveTasks {
		if task.ID != id {
			updatedActiveTasks = append(updatedActiveTasks, task)
		}
	}
	tc.ActiveTasks = updatedActiveTasks
}

// GetActiveTaskByIndex 根据索引值获取进行中任务，默认获取第一个任务
func (tc *Controller) GetActiveTaskByIndex(index ...int) *Models.TaskModels {
	idx := 1
	if len(index) > 0 {
		idx = index[0]
	}
	if idx <= 0 || idx > len(tc.ActiveTasks) {
		return nil
	}
	return &tc.ActiveTasks[idx-1]
}
