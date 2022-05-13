package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/wechatDemo/util"
	"io/ioutil"
	"log"
	"time"
)

// 与填写的服务器配置中的Token一致
const Token = "***"

func main() {
	router := gin.Default()

	router.GET("/wx", WXCheckSignature)
	router.POST("/wx", WXReceive)

	log.Fatalln(router.Run(":11080"))
}

// WXCheckSignature 微信接入校验
func WXCheckSignature(c *gin.Context) {
	signature := c.Query("signature")
	timestamp := c.Query("timestamp")
	nonce := c.Query("nonce")
	echostr := c.Query("echostr")

	ok := util.CheckSignature(signature, timestamp, nonce, Token)
	if !ok {
		log.Println("微信公众号接入校验失败!")
		return
	}

	log.Println("微信公众号接入校验成功!")
	_, _ = c.Writer.WriteString(echostr)
}

// WXTextMsg 微信文本消息结构体
type WXTextMsg struct {
	ToUserName   string
	FromUserName string
	CreateTime   int64
	MsgType      string
	Content      string
	MsgId        int64
}

// WXMsgReceive 微信消息接收
func WXReceive(c *gin.Context) {
	type Content struct {
		MsgType      string
	}

	reqBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Printf("[消息接收] - c.Request.Body读取失败: %v\n", err)
		return
	}
	// restore Body
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(reqBody))


	var content Content
	err = xml.Unmarshal(reqBody, &content)
	if err != nil {
		log.Printf("[消息接收] - XML数据包解析失败: %v\n", err)
		return
	}

	if content.MsgType == "text" {
		WXMsgReceive(c)
	}else if content.MsgType == "image"{
		WXImgReceive(c)
	}else {
		log.Printf("MsgType:%v not support", content.MsgType)
		return
	}
}

// WXMsgReceive 微信消息接收
func WXMsgReceive(c *gin.Context) {
	var textMsg WXTextMsg
	err := c.ShouldBindXML(&textMsg)
	if err != nil {
		log.Printf("[消息接收] - XML数据包解析失败: %v\n", err)
		return
	}

	log.Printf("[消息接收] - 收到消息, 消息类型为: %s, 消息内容为: %s\n", textMsg.MsgType, textMsg.Content)

	// 对接收的消息进行被动回复
	WXMsgReply(c, textMsg.ToUserName, textMsg.FromUserName)
}

// WXRepTextMsg 微信回复文本消息结构体
type WXRepTextMsg struct {
	ToUserName   string
	FromUserName string
	CreateTime   int64
	MsgType      string
	Content      string
	// 若不标记XMLName, 则解析后的xml名为该结构体的名称
	XMLName xml.Name `xml:"xml"`
}

// WXMsgReply 微信消息回复
func WXMsgReply(c *gin.Context, fromUser, toUser string) {
	repTextMsg := WXRepTextMsg{
		ToUserName:   toUser,
		FromUserName: fromUser,
		CreateTime:   time.Now().Unix(),
		MsgType:      "text",
		Content:      fmt.Sprintf("[消息回复] - %s", time.Now().Format("2006-01-02 15:04:05")),
	}

	msg, err := xml.Marshal(&repTextMsg)
	fmt.Print(string(msg))
	if err != nil {
		log.Printf("[消息回复] - 将对象进行XML编码出错: %v\n", err)
		return
	}
	_, _ = c.Writer.Write(msg)
}


func WXImgReceive(c *gin.Context) {
	var msg WXImgMsg
	err := c.ShouldBindXML(&msg)
	if err != nil {
		log.Printf("[消息接收] - XML数据包解析失败: %v\n", err)
		return
	}

	log.Printf("[消息接收] - 收到消息, 消息类型为: %s, MediaId: %s\n", msg.MsgType, msg.MediaId)

	// 对接收的消息进行被动回复
	WXImgReply(c, msg.ToUserName, msg.FromUserName, msg.MediaId)
}

// WXMsgReply 微信消息回复
func WXImgReply(c *gin.Context, fromUser, toUser, MediaId string) {
	repMsg := WXRepImgMsg{
		ToUserName:   toUser,
		FromUserName: fromUser,
		CreateTime:   time.Now().Unix(),
		MsgType:      "image",
		Image:Image{MediaId:MediaId},
	}

	msg, err := xml.Marshal(&repMsg)
	fmt.Print(string(msg))
	if err != nil {
		log.Printf("[消息回复] - 将对象进行XML编码出错: %v\n", err)
		return
	}
	_, _ = c.Writer.Write(msg)
}


// WXTextMsg 微信文本消息结构体
type WXImgMsg struct {
	ToUserName   string
	FromUserName string
	CreateTime   int64
	MsgType      string
	PicUrl       string
	MediaId      string
	MsgId        int64
}


type Image struct {
	MediaId string
}

type WXRepImgMsg struct {

	ToUserName   string
	FromUserName string
	CreateTime   int64
	MsgType      string
	Image Image
	// 若不标记XMLName, 则解析后的xml名为该结构体的名称
	XMLName xml.Name `xml:"xml"`
}
