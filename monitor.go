package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/",index)
	http.HandleFunc("/push/v1.0/add",add)
	//http://127.0.0.1/json ,不推荐
	http.HandleFunc("/json",func(rw http.ResponseWriter, req *http.Request) {
		_json := ResultJson{Code: 200,Msg: "使用系统自带的json编码操作成功"}
		err := json.NewEncoder(rw).Encode(&_json)
		if err != nil {
			_, _ = fmt.Fprint(rw, "错误")
		}
	})
	//http://127.0.0.1/zml
	http.HandleFunc("/zml",func(rw http.ResponseWriter, req *http.Request) {
		_json := ResultZml(200, "使用字面量构建json数据格式")
		responseJson(_json,rw)
	})
	_ = http.ListenAndServe(":80", nil)
}

func index(rw http.ResponseWriter, req *http.Request) {
	responseJson("欢迎访问本服务平台",rw)
}
//http://127.0.0.1/push/v1.0/add?name=yinlz&age=3
func add(rw http.ResponseWriter, req *http.Request) {
	name := req.FormValue("name")
	age := req.FormValue("age")
	msg := "操作成功,name->"+name+",age->"+age
	_json := Result(200,msg)
	responseJson(_json,rw)
}

func responseJson(json string,response http.ResponseWriter)  {
	response.Header().Set("Content-Type", "text/html;charset=utf-8")
	_, _ = fmt.Fprint(response, json)
}

type ResultJson struct {
	Msg string `json:"msg"`
	Code       int    `json:"code"`
}
//使用字面量构建json
func ResultZml(code int,msg string) string {
	return `{"msg":"`+msg+`","code":`+strconv.Itoa(code)+"}"
}

func Result(code int,msg string) string {
	return "{\"msg\":\""+msg+"\",\"code\":"+strconv.Itoa(code)+"}"
}