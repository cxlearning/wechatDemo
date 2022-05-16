package api

import (
	"github.com/wechatDemo/app/model"
	"log"
)

// 各类消息处理器
var (
	RecvTextHandler       func(*model.RecvText) model.ReplyMsg
	RecvImageHandler      func(*model.RecvImage) model.ReplyMsg
	RecvVoiceHandler      func(*model.RecvVoice) model.ReplyMsg
	RecvVideoHandler      func(*model.RecvVideo) model.ReplyMsg
	RecvShortVideoHandler func(*model.RecvVideo) model.ReplyMsg
	RecvLocationHandler   func(*model.RecvLocation) model.ReplyMsg
	RecvLinkHandler       func(*model.RecvLink) model.ReplyMsg
)

// 各类事件处理器
var (
	EventSubscribeHandler             func(*model.EventSubscribe) model.ReplyMsg
	EventUnsubscribeHandler           func(*model.EventSubscribe) model.ReplyMsg
	EventLocationHandler              func(*model.EventLocation) model.ReplyMsg
	EventClickHandler                 func(*model.EventClick) model.ReplyMsg
	EventViewHandler                  func(*model.EventView) model.ReplyMsg
	EventTemplateSendJobFinishHandler func(*model.EventTemplateSendJobFinish) model.ReplyMsg

	EventScancodePushHandler    func(*model.EventScancodePush) model.ReplyMsg
	EventScancodeWaitmsgHandler func(*model.EventScancodeWaitmsg) model.ReplyMsg
	EventPicSysphotoHandler     func(*model.EventPicSysphoto) model.ReplyMsg
	EventPicPhotoOrAlbumHandler func(*model.EventPicPhotoOrAlbum) model.ReplyMsg
	EventPicWeixinHandler       func(*model.EventPicWeixin) model.ReplyMsg
	EventLocationSelectHandler  func(*model.EventLocationSelect) model.ReplyMsg

	EventQualificationVerifySuccessHandler func(*model.EventQualificationVerifySuccess) model.ReplyMsg // 资质认证成功
	EventQualificationVerifyFailHandler    func(*model.EventQualificationVerifyFail) model.ReplyMsg    // 资质认证失败
	EventNamingVerifySuccessHandler        func(*model.EventNamingVerifySuccess) model.ReplyMsg        // 名称认证成功（即命名成功）
	EventNamingVerifyFailHandler           func(*model.EventNamingVerifyFail) model.ReplyMsg           // 名称认证失败
	EventAnnualRenewHandler                func(*model.EventAnnualRenew) model.ReplyMsg                // 年审通知
	EventVerifyExpiredHandler              func(*model.EventVerifyExpired) model.ReplyMsg              // 认证过期失效通知
)

// RecvDefaultHandler 如果没有注册某类消息处理器，那么收到这类消息时，使用这个默认处理器
var RecvDefaultHandler = func(msg *model.Message) (reply model.ReplyMsg) {
	log.Printf("unregistered receive message handler %s, use RecvDefaultHandler", msg.MsgType)
	return nil
}

// EventDefaultHandler 如果没有注册某类事件处理器，那么收到这类事件时，使用这个默认处理器
var EventDefaultHandler = func(msg *model.Message) (reply model.ReplyMsg) {
	log.Printf("unregistered receive event handler %s, use EventDefaultHandler", msg.Event)
	return nil
}

// HandleMessage 处理各类消息
func HandleMessage(msg *model.Message) (ret model.ReplyMsg) {
	log.Printf("process `%s` message", msg.MsgType)

	switch msg.MsgType {
	case model.MsgTypeText:
		if RecvTextHandler != nil {
			return RecvTextHandler(model.NewRecvText(msg))
		}
	case model.MsgTypeImage:
		if RecvImageHandler != nil {
			return RecvImageHandler(model.NewRecvImage(msg))
		}
	case model.MsgTypeVoice:
		if RecvVoiceHandler != nil {
			return RecvVoiceHandler(model.NewRecvVoice(msg))
		}
	case model.MsgTypeVideo:
		if RecvVideoHandler != nil {
			return RecvVideoHandler(model.NewRecvVideo(msg))
		}
	case model.MsgTypeShortVideo:
		if RecvShortVideoHandler != nil {
			return RecvShortVideoHandler(model.NewRecvVideo(msg))
		}
	case model.MsgTypeLocation:
		if RecvLocationHandler != nil {
			return RecvLocationHandler(model.NewRecvLocation(msg))
		}
	case model.MsgTypeLink:
		if RecvLinkHandler != nil {
			return RecvLinkHandler(model.NewRecvLink(msg))
		}
	case model.MsgTypeEvent:
		return HandleEvent(msg)
	default:
		log.Printf("unexpected receive MsgType: %s", msg.MsgType)
		return nil
	}

	return RecvDefaultHandler(msg)
}

// HandleEvent 处理各类事件
func HandleEvent(msg *model.Message) (reply model.ReplyMsg) {
	log.Printf("process `%s` event", msg.MsgType)

	switch msg.Event {
	case model.EventTypeSubscribe:
		if EventSubscribeHandler != nil {
			return EventSubscribeHandler(model.NewEventSubscribe(msg))
		}
	case model.EventTypeUnsubscribe:
		if EventUnsubscribeHandler != nil {
			return EventUnsubscribeHandler(model.NewEventSubscribe(msg))
		}
	case model.EventTypeLocation:
		if EventLocationHandler != nil {
			return EventLocationHandler(model.NewEventLocation(msg))
		}
	case model.EventTypeClick:
		if EventClickHandler != nil {
			return EventClickHandler(model.NewEventClick(msg))
		}
	case model.EventTypeView:
		if EventViewHandler != nil {
			return EventViewHandler(model.NewEventView(msg))
		}
	case model.EventTypeTemplateSendJobFinish:
		if EventTemplateSendJobFinishHandler != nil {
			return EventTemplateSendJobFinishHandler(model.NewEventTemplateSendJobFinish(msg))
		}
	case model.EventTypeScancodePush:
		if EventScancodePushHandler != nil {
			return EventScancodePushHandler(model.NewEventScancodePush(msg))
		}
	case model.EventTypeScancodeWaitmsg:
		if EventScancodeWaitmsgHandler != nil {
			return EventScancodeWaitmsgHandler(model.NewEventScancodeWaitmsg(msg))
		}
	case model.EventTypePicSysphoto:
		if EventPicSysphotoHandler != nil {
			return EventPicSysphotoHandler(model.NewEventPicSysphoto(msg))
		}
	case model.EventTypePicPhotoOrAlbum:
		if EventPicPhotoOrAlbumHandler != nil {
			return EventPicPhotoOrAlbumHandler(model.NewEventPicPhotoOrAlbum(msg))
		}
	case model.EventTypePicWeixin:
		if EventPicWeixinHandler != nil {
			return EventPicWeixinHandler(model.NewEventPicWeixin(msg))
		}
	case model.EventTypeLocationSelect:
		if EventLocationSelectHandler != nil {
			return EventLocationSelectHandler(model.NewEventLocationSelect(msg))
		}
	case model.EventTypeQualificationVerifySuccess:
		if EventQualificationVerifySuccessHandler != nil {
			return EventQualificationVerifySuccessHandler(model.NewEventQualificationVerifySuccess(msg))
		}
	case model.EventTypeQualificationVerifyFail:
		if EventQualificationVerifyFailHandler != nil {
			return EventQualificationVerifyFailHandler(model.NewEventQualificationVerifyFail(msg))
		}
	case model.EventTypeNamingVerifySuccess:
		if EventNamingVerifySuccessHandler != nil {
			return EventNamingVerifySuccessHandler(model.NewEventNamingVerifySuccess(msg))
		}
	case model.EventTypeNamingVerifyFail:
		if EventNamingVerifyFailHandler != nil {
			return EventNamingVerifyFailHandler(model.NewEventNamingVerifyFail(msg))
		}
	case model.EventTypeAnnualRenew:
		if EventAnnualRenewHandler != nil {
			return EventAnnualRenewHandler(model.NewEventAnnualRenew(msg))
		}
	case model.EventTypeVerifyExpired:
		if EventVerifyExpiredHandler != nil {
			return EventVerifyExpiredHandler(model.NewEventVerifyExpired(msg))
		}
	default:
		log.Fatalf("unexpected receive EventType: %s", msg.Event)
		return nil
	}

	return EventDefaultHandler(msg)
}
