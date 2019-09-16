package todolist

import "net/http"

// /todo/list

// request empty

// response
//  0 data:[]*TodoItem
//  10000
func (s *Service) handleTodoList(w http.ResponseWriter, r *http.Request) {
	list := []*TodoItem{}
	if err := s.storageService.ListTodo(&list); err != nil {
		resp(w, 10000, err)
		return
	}
	resp(w, 0, struct {
		List []*TodoItem `json:"list"`
	}{list})
}
