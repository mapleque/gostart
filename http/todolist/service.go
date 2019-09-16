package todolist

/*
 * 启动和基本框架代码
 *
 * 可以修改一些服务配置，与业务无关
 *
 */

import (
	"log"
	"net/http"
	"strings"
	"time"
)

// RunTodoListService run todolist service as an example of gostart/http
//
// Command:
//  `go run github.com/mapleque/gostart/http/todolist`
func RunTodoListService(addr string) {
	todolistService := NewTodoListService(addr)
	todolistService.Run()
}

// Service A service entity of todolist
type Service struct {
	httpServer     *http.Server
	storageService StorageServicer

	h *http.ServeMux
}

// NewTodoListService Create a Service Entity with default settings
func NewTodoListService(addr string) *Service {
	s := &Service{
		h: http.DefaultServeMux,
	}

	// 使用默认设置初始化httpServer
	s.httpServer = &http.Server{
		Addr:           addr,
		Handler:        s.h,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   5 * time.Second,
		MaxHeaderBytes: http.DefaultMaxHeaderBytes,
	}

	// 使用一个模拟的数据持久化服务
	s.storageService = NewMemStorageService()

	// 初始化接口配置
	s.initRouter()
	return s
}

// Run Start the todolist service
func (s *Service) Run() {
	log.Printf("[Info] todolist service listen on %s...\n", s.httpServer.Addr)
	log.Fatal("[Fatal]", s.httpServer.ListenAndServe())
}

// Handle Wrapper the http handle
func (s *Service) Handle(path string, handle http.HandlerFunc) {
	s.h.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		status := "success"

		// 请求结束，打印日志
		defer func(start time.Time) {
			// 打印请求日志
			// 格式样例：
			// 2019/09/16 15:53:04 [Access] success 127.0.0.1 /todo/add 0.233
			log.Printf(
				"[Access] %s %s %s %.3f\n",
				status, // 接口响应状态 success failed
				strings.Split(r.RemoteAddr, ":")[0],
				path, // 接口路径
				float64(time.Now().Sub(start).Nanoseconds())/1e6, // 接口响应时间，毫秒
			)
		}(time.Now())

		// 拦截接口处理逻辑中的panic，并标记请求失败
		defer func() {
			if err := recover(); err != nil {
				log.Printf("[Error] [panic] %v\n", err)
				status = "failed"
				resp(w, 10000, nil)
			}
		}()

		handle(w, r)
	})
}
