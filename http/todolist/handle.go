package todolist

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// 全局通用的返回数据结构以及状态码定义
//
// 通过resp方法，包装返回数据和对应的状态码

var responseMessage = map[int]string{
	0:     "成功",
	10000: "内部错误",
	10001: "参数不正确",
	10002: "数据不存在",
}

type response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func resp(w http.ResponseWriter, status int, data interface{}) {
	msg, exist := responseMessage[status]
	if !exist {
		msg = "未知错误类型"
	}
	ret := response{
		status,
		msg,
		data,
	}
	rt, _ := json.Marshal(ret)
	w.Header().Set("Content-Type", "application/json")
	w.Write(rt)
}

func bind(r *http.Request, tar interface{}) interface{} {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return "读取请体求错误"
	}
	defer r.Body.Close()
	if err := json.Unmarshal(body, tar); err != nil {
		return "请求体为非法JSON"
	}
	return nil
}
