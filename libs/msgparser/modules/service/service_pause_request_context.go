package service

import (
	"github.com/bianjieai/cosmos-sync/libs/msgparser/modules"
	. "github.com/bianjieai/cosmos-sync/libs/msgparser/modules"
	"strings"
)

type (
	DocMsgPauseRequestContext struct {
		RequestContextID string `bson:"request_context_id" yaml:"request_context_id"`
		Consumer         string `bson:"consumer" yaml:"consumer"`
	}
)

func (m *DocMsgPauseRequestContext) GetType() string {
	return MsgTypePauseRequestContext
}

func (m *DocMsgPauseRequestContext) BuildMsg(v interface{}) {
	msg := v.(*MsgPauseRequestContext)

	m.RequestContextID = strings.ToUpper(msg.RequestContextId)
	m.Consumer = msg.Consumer
}

func (m *DocMsgPauseRequestContext) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var addrs []string

	msg := v.(*MsgPauseRequestContext)
	addrs = append(addrs, msg.Consumer)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return modules.CreateMsgDocInfo(v, handler)
}