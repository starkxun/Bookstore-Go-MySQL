package utils

import (
	"encoding/json"
	"io"	//io/ioutil已经被新的go版本替代了
	"net/http"
)
//将http请求中的参数解析到指定的数据结构中
func ParseBody(r *http.Request,x interface{}){
	if Body,err := io.ReadAll(r.Body);err ==nil{
		if err := json.Unmarshal([]byte(Body),x);err!=nil{
			return
		}
	}
}