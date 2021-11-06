package main

import (
	"PushServer/serverce"
	"fmt"
)

func main() {
	s := serverce.NewWebServer()
	err := s.Run()
	if err != nil {
		fmt.Println("web服务开启失败端口被占用了，请换一个端口: ", err)
		return
	}
}
