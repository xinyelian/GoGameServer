package dispatch

import (
	"github.com/funny/link"
	"github.com/funny/binary"
	"protos"
	"protos/dbProto"
	"protos/gameProto"
	"protos/logProto"
	"protos/systemProto"
	. "tools"
	"container/list"
)

type HandleInterface interface {
	DealMsg(session *link.Session, msg []byte)
}

//通用Handle
type Handle map[uint16]func(*link.Session, protos.ProtoMsg)

func (this Handle) DealMsg(session *link.Session, msg []byte) {
	msgID := binary.GetUint16LE(msg[:2])
	var protoMsg protos.ProtoMsg
	if systemProto.IsValidID(msgID) || logProto.IsValidID(msgID) || gameProto.IsValidID(msgID) {
		protoMsg = protos.UnmarshalProtoMsg(msg)
	} else if dbProto.IsValidID(msgID) {
		protoMsg = dbProto.UnmarshalProtoMsg(msg)
	}

	if protoMsg == protos.NullProtoMsg {
		ERR("收到Proto未处理消息：", msgID)
		return
	}

	if f, exists := this[msgID]; exists{
		f(session, protoMsg)
	} else {
		ERR("收到Handle未处理消息：", msgID)
	}
}

//条件Handle
type HandleCondition struct {
	Condition func(uint16) bool
	H Handle
}

func (this HandleCondition) DealMsg(session *link.Session, msg []byte) {
	msgID := binary.GetUint16LE(msg[:2])
	if this.Condition(msgID) {
		this.H.DealMsg(session, msg)
	}
}

//条件函数
type HandleFuncCondition struct {
	Condition func(uint16) bool
	H func(*link.Session, []byte)
}

func (this HandleFuncCondition) DealMsg(session *link.Session, msg []byte) {
	msgID := binary.GetUint16LE(msg[:2])
	if this.Condition(msgID) {
		this.H(session, msg)
	}
}

//函数
type HandleFunc struct {
	H func(*link.Session, []byte)
}

func (this HandleFunc) DealMsg(session *link.Session, msg []byte) {
	this.H(session, msg)
}

//条件Handles
type HandleConditions struct {
	hList *list.List
}

func NewHandleConditions() HandleConditions {
	return HandleConditions{
		hList: list.New(),
	}
}

func (this HandleConditions) Add(handle HandleInterface) {
	this.hList.PushBack(handle)
}

func (this HandleConditions) DealMsg(session *link.Session, msg []byte) {
	for e := this.hList.Front(); e != nil; e = e.Next() {
		e.Value.(HandleInterface).DealMsg(session, msg)
	}
}