package handlers

import (
	"fmt"
	"github.com/silenceper/wechat"
	"github.com/silenceper/wechat/message"
	"log"
	"net/http"
	config2 "remind-go/config"
)

func Message(w http.ResponseWriter, r *http.Request) {
	wechatConfig := config2.LoadConfig()
	log.Println(wechatConfig.Wechat)
	//配置微信
	config := &wechat.Config{
		AppID:          wechatConfig.Wechat.AppID,
		AppSecret:      wechatConfig.Wechat.AppSecret,
		Token:          wechatConfig.Wechat.Token,
		EncodingAESKey: wechatConfig.Wechat.EncodingAESKey,
	}
	wc := wechat.NewWechat(config)
	server := wc.GetServer(r, w)
	//设置接收消息的处理方法
	server.SetMessageHandler(func(msg message.MixMessage) *message.Reply {
		switch msg.MsgType {
		case message.MsgTypeText:
			//回复消息：演示回复用户发送的消息
			//text := message.NewText(msg.Content)
			res := message.NewText(HandleMessage(msg.Content))
			return &message.Reply{MsgType: message.MsgTypeText, MsgData: res}
		case message.MsgTypeVoice:
			text := message.NewVoice(msg.Content)
			return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}
		default:
			return &message.Reply{MsgType: message.MsgTypeText, MsgData: message.NewText("我睡着了，听不懂你在说啥")}
		}
		return nil
	})

	//处理消息接收以及回复
	err := server.Serve()
	if err != nil {
		fmt.Println(err)
		return
	}
	//发送回复的消息
	err = server.Send()
	if err != nil {
		fmt.Println(err.Error())
	}
}
