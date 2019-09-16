package todolist

/*
 * 用于定义数据结构
 */

// TodoItem 任务数据格式定义
type TodoItem struct {
	ID      string `json:"id"`
	Content string `json:"content"`
	Done    bool   `json:"done"`
}
