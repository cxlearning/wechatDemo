package service

import (
	"fmt"
	"github.com/wechatDemo/app/model"
	"log"
)

// EventDefaultHandler 注册默认处理器
func EventDefaultHandler(m *model.Message) model.ReplyMsg {
	log.Printf("%+v", m)

	// echo message
	ret := &model.ReplyText{
		ToUserName:   m.FromUserName,
		FromUserName: m.ToUserName,
		CreateTime:   m.CreateTime,
		Content:      fmt.Sprintf("Event=%s", m.Event),
	}

	log.Printf("replay message: %+v", ret)
	return ret
}

// EventSubscribeHandler 注册关注事件处理器
func EventSubscribeHandler(m *model.EventSubscribe) model.ReplyMsg {
	log.Printf("%+v", m)

	// echo message
	ret := &model.ReplyText{
		ToUserName:   m.FromUserName,
		FromUserName: m.ToUserName,
		CreateTime:   m.CreateTime,
		Content:      fmt.Sprintf("Event=%s, EventKey=%s, Ticket=%s", m.Event, m.EventKey, m.Ticket),
	}

	log.Printf("replay message: %+v", ret)
	return ret
}

// EventUnsubscribeHandler 注册取消关注事件处理器
func EventUnsubscribeHandler(m *model.EventSubscribe) model.ReplyMsg {
	log.Printf("someone gone")
	return nil
}

// EventLocationHandler 注册上报地理位置事件处理器
func EventLocationHandler(m *model.EventLocation) model.ReplyMsg {
	log.Printf("%+v", m)

	// echo message
	ret := &model.ReplyText{
		ToUserName:   m.FromUserName,
		FromUserName: m.ToUserName,
		CreateTime:   m.CreateTime,
		Content: fmt.Sprintf("Latitude=%.6f, Longitude=%.6f, Precision=%.6f",
			m.Latitude, m.Longitude, m.Precision),
	}

	log.Printf("replay message: %+v", ret)
	return ret
}

// EventClickHandler 注册点击自定义菜单事件处理器
func EventClickHandler(m *model.EventClick) model.ReplyMsg {
	log.Printf("%+v", m)

	// echo message
	ret := &model.ReplyText{
		ToUserName:   m.FromUserName,
		FromUserName: m.ToUserName,
		CreateTime:   m.CreateTime,
		Content:      fmt.Sprintf("Event=%s, EventKey=%s", m.Event, m.EventKey),
	}

	log.Printf("replay message: %+v", ret)
	return ret
}

// EventViewHandler 注册点击菜单跳转链接时的事件处理器
func EventViewHandler(m *model.EventView) model.ReplyMsg {
	log.Printf("%+v", m)

	// echo message 貌似用户收不到回复的消息？？？
	ret := &model.ReplyText{
		ToUserName:   m.FromUserName,
		FromUserName: m.ToUserName,
		CreateTime:   m.CreateTime,
		Content:      fmt.Sprintf("Event=%s, EventKey=%s", m.Event, m.EventKey),
	}

	log.Printf("replay message: %+v", ret)
	return ret
}

// EventTemplateSendJobFinishHandler 模版消息发送结果通知事件
func EventTemplateSendJobFinishHandler(m *model.EventTemplateSendJobFinish) model.ReplyMsg {
	log.Printf("%+v", m)

	// echo message
	ret := &model.ReplyText{
		ToUserName:   m.FromUserName,
		FromUserName: m.ToUserName,
		CreateTime:   m.CreateTime,
		Content:      fmt.Sprintf("Event=%s, MsgID=%d, Status=%s", m.Event, m.MsgID, m.Status),
	}

	log.Printf("replay message: %+v", ret)
	return ret
}
