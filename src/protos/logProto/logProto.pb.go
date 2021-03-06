// Code generated by protoc-gen-go.
// source: logProto.proto
// DO NOT EDIT!

/*
Package logProto is a generated protocol buffer package.

It is generated from these files:
	logProto.proto

It has these top-level messages:
	Log_CommonLogC2S
*/
package logProto

import proto "code.google.com/p/goprotobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type Log_CommonLogC2S struct {
	Dir              *string `protobuf:"bytes,1,req" json:"Dir,omitempty"`
	Type             *uint32 `protobuf:"varint,2,req" json:"Type,omitempty"`
	Content          *string `protobuf:"bytes,3,req" json:"Content,omitempty"`
	Time             *int64  `protobuf:"varint,4,req" json:"Time,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Log_CommonLogC2S) Reset()         { *m = Log_CommonLogC2S{} }
func (m *Log_CommonLogC2S) String() string { return proto.CompactTextString(m) }
func (*Log_CommonLogC2S) ProtoMessage()    {}

func (m *Log_CommonLogC2S) GetDir() string {
	if m != nil && m.Dir != nil {
		return *m.Dir
	}
	return ""
}

func (m *Log_CommonLogC2S) GetType() uint32 {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return 0
}

func (m *Log_CommonLogC2S) GetContent() string {
	if m != nil && m.Content != nil {
		return *m.Content
	}
	return ""
}

func (m *Log_CommonLogC2S) GetTime() int64 {
	if m != nil && m.Time != nil {
		return *m.Time
	}
	return 0
}

func init() {
}
