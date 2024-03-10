package Task

import "go_midjourney-api/Models"

// Controller 结构体包含任务队列和相关方法
type Controller struct {
	AllTasks    map[string]Models.TaskModels
	ActiveTasks map[string]Models.TaskModels
}

var TaskControllerInstance = Controller{
	AllTasks:    make(map[string]Models.TaskModels),
	ActiveTasks: make(map[string]Models.TaskModels),
}

// AddTask 将任务加入队列
func (tc *Controller) AddTask(task Models.TaskModels) {
	tc.AllTasks[task.ID] = task
	tc.ActiveTasks[task.ID] = task
}

// GetActiveTasks 查询进行中的任务队列
func (tc *Controller) GetActiveTasks() []Models.TaskModels {
	var activeTasks []Models.TaskModels
	for _, task := range tc.ActiveTasks {
		activeTasks = append(activeTasks, task)
	}
	return activeTasks
}

// GetAllTasks 查询所有任务
func (tc *Controller) GetAllTasks() []Models.TaskModels {
	var allTasks []Models.TaskModels
	for _, task := range tc.AllTasks {
		allTasks = append(allTasks, task)
	}
	return allTasks
}

// GetTaskByID 根据指定 ID 获取任务
func (tc *Controller) GetTaskByID(id string) *Models.TaskModels {
	task, ok := tc.AllTasks[id]
	if !ok {
		return nil
	}
	return &task
}

// ClearActiveTaskByID 根据任务 ID 清除进行中的任务
func (tc *Controller) ClearActiveTaskByID(id string) {
	delete(tc.ActiveTasks, id)
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

	i := 0
	for _, task := range tc.ActiveTasks {
		if i == idx-1 {
			return &task
		}
		i++
	}

	return nil
}
