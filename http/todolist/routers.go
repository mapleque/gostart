package todolist

// Run in console:
//     curl -H'Content-Type: application/json' -d'{"content":"Write the posts about go"}' http://localhost/todo/add
//     #> {"status":0}
//     curl -H'Content-Type: application/json' -d'{}' http://localhost/todo/list
//     #> {"status":0,"data":{"list":[{"id":"1","content":"Write the posts about go","done":false}]}}
//     curl -H'Content-Type: application/json' -d'{"id":"1","content":"Write the posts about go","done":true}' http://localhost/todo/update
//     #> {"status":0}
//     curl -H'Content-Type: application/json' -d'{}' http://localhost/todo/list
//     #> {"status":0,"data":{"list":[{"id":"1","content":"Write the posts about go","done":true}]}}
//     curl -H'Content-Type: application/json' -d'{"id":"1"}' http://localhost/todo/delete
//     #> {"status":0}
//     curl -H'Content-Type: application/json' -d'{}' http://localhost/todo/list
//     #> {"status":0,"data":{"list":[]}}
func (s *Service) initRouter() {
	s.Handle("/todo/add", s.handleTodoAdd)
	s.Handle("/todo/delete", s.handleTodoDelete)
	s.Handle("/todo/update", s.handleTodoUpdate)
	s.Handle("/todo/list", s.handleTodoList)
}
