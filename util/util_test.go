package util

import (
	"fmt"
	"testing"
	"time"
)

func TestRun(T *testing.T) {

	time.Sleep(time.Second * 20)
	msg := &PushMsg{
		NameId:   "kosakd",
		MsgTime:  time.Now().Format("2006-01-02 15:04:05"),
		MsgTitle: "文章的标题",
		Msg:      "文章的内容，啊哈哈哈哈哈哈哈哈哈哈哈哈",
	}
	test := VX.MsgPush(msg)
	fmt.Println(test)
	time.Sleep(time.Hour * 1)
}
