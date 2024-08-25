package msg

import (
	"fmt"
	"time"
)

type MsgType struct {
	MsgType string `xml:"MsgType"`
}

type Msg struct {
	//	MsgType
	ToUserName   string `xml:"ToUserName"`
	FromUserName string `xml:"FromUserName"`
	CreateTime   int64  `xml:"CreateTime"`
	MsgId        int64  `xml:"MsgId"`
}

// TextMsg 文本消息
type TextMsg struct {
	Msg
	Content string `xml:"Content"`
}

func (msg TextMsg) TemplateMsg() string {

	reply := msg
	reply.ToUserName = msg.FromUserName
	reply.FromUserName = msg.ToUserName
	reply.CreateTime = time.Now().Unix()
	reply.Content = "test"

	return fmt.Sprintf(TextMsgXMLTemplate, reply.ToUserName, reply.FromUserName, reply.CreateTime, reply.Content)
}

func (msg TextMsg) MsgType() MsgType {
	return MsgType{MsgTypeText}
}

// ImageMsg 图片消息
type ImageMsg struct {
	Msg
	PicUrl  string `xml:"PicUrl"`
	MediaId string `xml:"MediaId"`
}

func (msg ImageMsg) TemplateMsg() string {

	reply := msg
	reply.ToUserName = msg.FromUserName
	reply.FromUserName = msg.ToUserName
	reply.CreateTime = time.Now().Unix()

	return fmt.Sprintf(ImageMsgXMLTemplate, reply.ToUserName, reply.FromUserName, reply.CreateTime, reply.MediaId)
}

func (msg ImageMsg) MsgType() MsgType {
	return MsgType{MsgTypeImage}
}

type EventMsg struct {
	Msg
	Event string `xml:"Event"`
}

func (msg EventMsg) TemplateMsg() string {

	if msg.Event != EventSubscribe {
		return "success"
	}

	reply := TextMsg{}
	reply.ToUserName = msg.FromUserName
	reply.FromUserName = msg.ToUserName
	reply.CreateTime = time.Now().Unix()
	reply.Content = "哔哔哔哔"
	return fmt.Sprintf(TextMsgXMLTemplate, reply.ToUserName, reply.FromUserName, reply.CreateTime, reply.Content)
}
