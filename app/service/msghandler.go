package service

import (
	"encoding/json"
	"github.com/wechatDemo/app/model"
	"log"
)

func DefaultHandler(msg *model.Message) model.ReplyMsg {
	log.Printf("%+v", msg)

	event := model.NewRecvEvent(msg)
	js, _ := json.Marshal(event)

	// echo message
	ret := &model.ReplyText{
		ToUserName:   msg.FromUserName,
		FromUserName: msg.ToUserName,
		CreateTime:   msg.CreateTime,
		Content:      string(js),
	}

	log.Printf("replay message: %+v", ret)
	return ret
}

func EchoMsgText(m *model.RecvText) model.ReplyMsg {
	log.Printf("receive message: %+v", m)

	// echo message
	ret := &model.ReplyText{
		ToUserName:   m.FromUserName,
		FromUserName: m.ToUserName,
		CreateTime:   m.CreateTime,
		Content:      m.FromUserName + ", " + m.Content,
	}

	log.Printf("replay message: %+v", ret)
	return ret
}

func EchoMsgImage(m *model.RecvImage) model.ReplyMsg {
	log.Printf("%+v", m)

	// echo message
	ret := &model.ReplyImage{
		ToUserName:   m.FromUserName,
		FromUserName: m.ToUserName,
		CreateTime:   m.CreateTime,
		PicUrl:       m.PicUrl,
		MediaId:      m.MediaId,
	}

	log.Printf("%+v", ret)
	return ret
}

func EchoMsgVoice(m *model.RecvVoice) model.ReplyMsg {
	log.Printf("%+v", m)

	// echo message
	ret := &model.ReplyVoice{
		ToUserName:   m.FromUserName,
		FromUserName: m.ToUserName,
		CreateTime:   m.CreateTime,
		MediaId:      m.MediaId,
	}

	log.Printf("%+v", ret)
	return ret
}

func EchoMsgVideo(m *model.RecvVideo) model.ReplyMsg {
	log.Printf("%+v", m)

	// MediaId ???
	ret := &model.ReplyVideo{
		ToUserName:   m.FromUserName,
		FromUserName: m.ToUserName,
		CreateTime:   m.CreateTime,
		MediaId:      m.MediaId,
		Title:        "video",
		Description:  "thist is a test desc...",
	}

	log.Printf("%+v", ret)
	return ret
}

func EchoMsgShortVideo(m *model.RecvVideo) model.ReplyMsg {
	log.Printf("%+v", m)

	// MediaId ???
	ret := &model.ReplyVideo{
		ToUserName:   m.FromUserName,
		FromUserName: m.ToUserName,
		CreateTime:   m.CreateTime,
		MediaId:      m.ThumbMediaId,
		Title:        "shortvideo",
		Description:  "thist is a test desc...",
	}

	log.Printf("%+v", ret)
	return ret
}

func EchoMsgLocation(m *model.RecvLocation) model.ReplyMsg {
	log.Printf("%+v", m)

	// echo message
	ret := &model.ReplyText{
		ToUserName:   m.FromUserName,
		FromUserName: m.ToUserName,
		CreateTime:   m.CreateTime,
		Content:      AccessToken(),
	}

	log.Printf("replay message: %+v", ret)
	return ret
}

func EchoMsgLink(m *model.RecvLink) model.ReplyMsg {
	log.Printf("%+v", m)

	// 回复图文消息

	return nil
}
