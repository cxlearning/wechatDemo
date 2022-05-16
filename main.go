package main

import (
	"github.com/wechatDemo/app/api"
	"github.com/wechatDemo/app/service"
	"github.com/wechatDemo/config"
	"log"
	"net/http"
)

func main()  {
	addr := ":11080"

	service.Initialize(config.OriginId, config.AppId, config.AppSecret, config.Token, config.EncodingAESKey)


	api.RecvTextHandler = service.EchoMsgText             // 注册文本消息处理器
	api.RecvImageHandler = service.EchoMsgImage           // 注册图片消息处理器
	api.RecvVoiceHandler = service.EchoMsgVoice           // 注册语音消息处理器
	api.RecvVideoHandler = service.EchoMsgVideo           // 注册视频消息处理器
	api.RecvShortVideoHandler = service.EchoMsgShortVideo // 注册小视频消息处理器
	api.RecvLocationHandler = service.EchoMsgLocation     // 注册位置消息处理器
	api.RecvLinkHandler = service.EchoMsgLink             // 注册链接消息处理器
	api.RecvDefaultHandler = service.DefaultHandler       // 注册默认处理器

	api.EventSubscribeHandler = service.EventSubscribeHandler     // 注册关注事件处理器
	api.EventUnsubscribeHandler = service.EventUnsubscribeHandler // 注册取消关注事件处理器
	api.EventLocationHandler = service.EventLocationHandler       // 注册上报地理位置事件处理器
	api.EventClickHandler = service.EventClickHandler             // 注册点击自定义菜单事件处理器
	api.EventViewHandler = service.EventViewHandler               // 注册点击菜单跳转链接时的事件处理器
	// 模版消息发送结果通知事件
	api.EventTemplateSendJobFinishHandler = service.EventTemplateSendJobFinishHandler
	api.EventDefaultHandler = service.EventDefaultHandler // 注册默认处理器



	http.HandleFunc("/wx", service.HandleAccess)


	log.Printf("server is running at %s", addr)
	http.ListenAndServe(addr, nil)

}
