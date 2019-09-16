package todolist

import "net/http"

// /todo/delete

// request
type todoDeleteParam struct {
	ID string `json:"id"` // 必选，非空
}

// response status:
//  0
//  10000
//  10001
//  10002
func (s *Service) handleTodoDelete(w http.ResponseWriter, r *http.Request) {
	params := todoDeleteParam{}
	if err := bind(r, &params); err != nil {
		resp(w, 10001, err)
		return
	}
	if params.ID == "" {
		resp(w, 10001, "id不能为空")
		return
	}

	if exist, err := s.storageService.DeleteTodo(params.ID); err != nil {
		resp(w, 10000, err)
		return
	} else if !exist {
		resp(w, 10002, params.ID)
		return
	}
	resp(w, 0, nil)
}
