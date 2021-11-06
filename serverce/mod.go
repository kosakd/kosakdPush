//这个包用来，定义handler函数
package serverce

import (
	"PushServer/util"
	"fmt"
	"html/template"
	"net/http"
	"time"
)

type Handlers struct {
}

//helloworldweb函数
func (s *Handlers) HandHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello GO Web")
	msg := r.FormValue("msg")
	fmt.Println(msg)
}

//打印消息的msg，web函数服务
func (s *Handlers) MsgHand(w http.ResponseWriter, r *http.Request) {
	if r.FormValue("nameid") == "" || r.FormValue("msgid") == "" {
		fmt.Fprintln(w, "参数错误")
		return
	}
	t, _ := template.ParseFiles("views/index.html")
	nameId := r.FormValue("nameid")
	msgId := r.FormValue("msgid")
	msg, err := util.SelecMsg(msgId, nameId)
	fmt.Println(*msg)
	if err != nil {
		fmt.Println("请求失败: ", err)
		fmt.Fprintln(w, "请求失败，没有这个消息")
		return
	}
	t.Execute(w, msg)
}

//消息推送接口函数，需要参数，nameid，msgtitle，msgcontent
func (s *Handlers) MsgPush(w http.ResponseWriter, r *http.Request) {
	if r.FormValue("nameid") == "" || r.FormValue("msgtitle") == "" {
		fmt.Fprintln(w, "参数错误")
		return
	}
	msg := r.FormValue("msgcontent")
	msgTitle := r.FormValue("msgtitle")
	nameId := r.FormValue("nameid")
	msgTime := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println(msgTitle, msg, nameId, msgTime)
	// fmt.Fprintln(w, msgTitle, msg, nameId, msgTime)
	m := &util.PushMsg{
		MsgTime:  msgTime,
		MsgTitle: msgTitle,
		Msg:      msg,
		NameId:   nameId,
	}
	reb := util.VX.MsgPush(m)
	fmt.Fprintln(w, reb)
}
