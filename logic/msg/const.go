package msg

const (
	MsgTypeText  = "text"
	MsgTypeImage = "image"
	MsgTypeEvent = "event"
)

const (
	EventSubscribe = "subscribe"
)

const (
	TextMsgXMLTemplate = `<xml>
						 	<ToUserName><![CDATA[%s]]></ToUserName>
						 	<FromUserName><![CDATA[%s]]></FromUserName>
						 	<CreateTime>%d</CreateTime>
						 	<MsgType><![CDATA[text]]></MsgType>
						 	<Content><![CDATA[%s]]></Content>
						 </xml>`
	ImageMsgXMLTemplate = `<xml>
						  	<ToUserName><![CDATA[%s]]></ToUserName>
						  	<FromUserName><![CDATA[%s]]></FromUserName>
						  	<CreateTime>%d</CreateTime>
						  	<MsgType><![CDATA[image]]></MsgType>
						  	<Image>
						  	<MediaId><![CDATA[%s]]></MediaId>
						  	</Image>
						  </xml>`
)
