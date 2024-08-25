package msg

type MsgSender interface {
	TemplateMsg() string
	//	MsgType() MsgType
}

type MsgGet struct {
	Strategy MsgSender
}

func (m *MsgGet) SetStrategy(strategy MsgSender) {
	m.Strategy = strategy
}

func (m *MsgGet) Execute() string {
	return m.Strategy.TemplateMsg()
}
