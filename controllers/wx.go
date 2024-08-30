package controllers

import (
	"crypto/sha1"
	"fmt"
	"github.com/astaxie/beego"
	"io"
	"robot-for-fxy/logic/msg"
	"sort"
	"strings"
)

const (
	token = "123456789"
)

type WXController struct {
	beego.Controller
}

func (c *WXController) Get() {

	signature := c.GetString("signature")
	timestamp := c.GetString("timestamp")
	nonce := c.GetString("nonce")
	echostr := c.GetString("echostr")

	// 拼接 token, timestamp 和 nonce
	list := []string{token, timestamp, nonce}
	sort.Strings(list) // 按字典顺序排序

	// 计算 SHA1 哈希
	h := sha1.New()
	io.WriteString(h, strings.Join(list, ""))
	hashcode := fmt.Sprintf("%x", h.Sum(nil))

	// 验证签名
	if hashcode == signature {
		c.Ctx.WriteString(echostr)
	} else {
		c.Ctx.WriteString("hashcode = " + hashcode)
	}
}

func (c *WXController) Post() {

	xmlData := c.Ctx.Input.RequestBody
	fmt.Println(string(xmlData))
	msgSender, err := msg.ParseXML(xmlData)
	if err != nil {
		c.Ctx.WriteString(err.Error())
		return
	}

	// 针对不同的指令，做不同的response
	msgGet := msg.MsgGet{}
	msgGet.SetStrategy(msgSender)
	c.Ctx.WriteString(msgGet.Execute())
}
