//此包用于封装web访问的函数
package serverce

import (
	"PushServer/conf"
	"net/http"
	"time"
)

//web服务结构体
type WebServer struct {
	Addr         string
	Hand         *Handlers
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

//开启WebServer
func (s *WebServer) Run() error {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	http.HandleFunc("/", s.Hand.HandHello)
	http.HandleFunc("/msg", s.Hand.MsgHand)
	http.HandleFunc("/msgpush", s.Hand.MsgPush)
	err := http.ListenAndServe(s.Addr, nil)
	return err
}

//用于调用，初始化一个WebServer结构体
func NewWebServer() *WebServer {
	return &WebServer{
		Addr:         conf.Conf.ServerYaml.Ip + ":" + conf.Conf.ServerYaml.Port,
		ReadTimeout:  time.Duration(conf.Conf.ServerYaml.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(conf.Conf.ServerYaml.WriteTimeout) * time.Second,
	}

}
