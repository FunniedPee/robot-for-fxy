package msg

import (
	"encoding/xml"
	"fmt"
)

func ParseXML(xmlData []byte) (MsgSender, error) {

	if len(xmlData) == 0 {
		return nil, fmt.Errorf("no data provided")
	}

	var msgType MsgType
	if err := xml.Unmarshal(xmlData, &msgType); err != nil {
		return nil, err
	}

	switch msgType.MsgType {
	case MsgTypeText:

		var textMsg TextMsg
		if err := xml.Unmarshal(xmlData, &textMsg); err != nil {
			return nil, err
		}
		return &textMsg, nil
	case MsgTypeImage:

		var imageMsg ImageMsg
		if err := xml.Unmarshal(xmlData, &imageMsg); err != nil {
			return nil, err
		}
		return &imageMsg, nil
	case MsgTypeEvent:

		var eventMsg EventMsg
		if err := xml.Unmarshal(xmlData, &eventMsg); err != nil {
			return nil, err
		}
		return &eventMsg, nil
	default:
		return nil, fmt.Errorf("unsupported message type: %s", msgType.MsgType)
	}
}
