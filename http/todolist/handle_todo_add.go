package todolist

import "net/http"

// /todo/add
//
// 定义了一个参数格式和校验规则todoAddParam
// 通过bind方法读取参数数据
// 通过自定义的方法进行参数校验

// request
type todoAddParam struct {
	Content string `json:"content"` // 必选，非空
	Done    bool   `json:"bool"`    // 可选，默认false
}

// response status:
//  0
//  10000
//  10001
func (s *Service) handleTodoAdd(w http.ResponseWriter, r *http.Request) {
	params := todoAddParam{}
	if err := bind(r, &params); err != nil {
		resp(w, 10001, err)
		return
	}

	if params.Content == "" {
		resp(w, 10001, "content不能为空")
		return
	}

	if err := s.storageService.SaveTodo(params.Content, params.Done); err != nil {
		resp(w, 10000, err)
		return
	}

	resp(w, 0, nil)
}
