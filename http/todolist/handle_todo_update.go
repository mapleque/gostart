package todolist

import "net/http"

// /todo/update

// request
type todoUpdateParam struct {
	ID      string `json:"id"`      // 必选，非空
	Content string `json:"content"` // 必选，非空
	Done    bool   `json:"done"`    // 可选，默认false
}

// response status:
//  0
//  10000
//  10001
//  10002
func (s *Service) handleTodoUpdate(w http.ResponseWriter, r *http.Request) {
	params := todoUpdateParam{}
	if err := bind(r, &params); err != nil {
		resp(w, 10001, err)
		return
	}
	if params.ID == "" {
		resp(w, 10001, "id不能为空")
		return
	}
	if params.Content == "" {
		resp(w, 10001, "content不能为空")
		return
	}
	if exist, err := s.storageService.UpdateTodo(params.ID, params.Content, params.Done); err != nil {
		resp(w, 10000, err)
		return
	} else if !exist {
		resp(w, 10002, params.ID)
		return
	}
	resp(w, 0, nil)
}
