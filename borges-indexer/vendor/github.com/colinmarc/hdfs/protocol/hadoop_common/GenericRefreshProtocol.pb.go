// Code generated by protoc-gen-go.
// source: GenericRefreshProtocol.proto
// DO NOT EDIT!

/*
Package hadoop_common is a generated protocol buffer package.

It is generated from these files:
	GenericRefreshProtocol.proto
	GetUserMappingsProtocol.proto
	HAServiceProtocol.proto
	IpcConnectionContext.proto
	ProtobufRpcEngine.proto
	ProtocolInfo.proto
	RefreshAuthorizationPolicyProtocol.proto
	RefreshCallQueueProtocol.proto
	RefreshUserMappingsProtocol.proto
	RpcHeader.proto
	Security.proto
	TraceAdmin.proto
	ZKFCProtocol.proto

It has these top-level messages:
	GenericRefreshRequestProto
	GenericRefreshResponseProto
	GenericRefreshResponseCollectionProto
*/
package hadoop_common

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

// *
//  Refresh request.
type GenericRefreshRequestProto struct {
	Identifier       *string  `protobuf:"bytes,1,opt,name=identifier" json:"identifier,omitempty"`
	Args             []string `protobuf:"bytes,2,rep,name=args" json:"args,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *GenericRefreshRequestProto) Reset()         { *m = GenericRefreshRequestProto{} }
func (m *GenericRefreshRequestProto) String() string { return proto.CompactTextString(m) }
func (*GenericRefreshRequestProto) ProtoMessage()    {}

func (m *GenericRefreshRequestProto) GetIdentifier() string {
	if m != nil && m.Identifier != nil {
		return *m.Identifier
	}
	return ""
}

func (m *GenericRefreshRequestProto) GetArgs() []string {
	if m != nil {
		return m.Args
	}
	return nil
}

// *
// A single response from a refresh handler.
type GenericRefreshResponseProto struct {
	ExitStatus       *int32  `protobuf:"varint,1,opt,name=exitStatus" json:"exitStatus,omitempty"`
	UserMessage      *string `protobuf:"bytes,2,opt,name=userMessage" json:"userMessage,omitempty"`
	SenderName       *string `protobuf:"bytes,3,opt,name=senderName" json:"senderName,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GenericRefreshResponseProto) Reset()         { *m = GenericRefreshResponseProto{} }
func (m *GenericRefreshResponseProto) String() string { return proto.CompactTextString(m) }
func (*GenericRefreshResponseProto) ProtoMessage()    {}

func (m *GenericRefreshResponseProto) GetExitStatus() int32 {
	if m != nil && m.ExitStatus != nil {
		return *m.ExitStatus
	}
	return 0
}

func (m *GenericRefreshResponseProto) GetUserMessage() string {
	if m != nil && m.UserMessage != nil {
		return *m.UserMessage
	}
	return ""
}

func (m *GenericRefreshResponseProto) GetSenderName() string {
	if m != nil && m.SenderName != nil {
		return *m.SenderName
	}
	return ""
}

// *
// Collection of responses from zero or more handlers.
type GenericRefreshResponseCollectionProto struct {
	Responses        []*GenericRefreshResponseProto `protobuf:"bytes,1,rep,name=responses" json:"responses,omitempty"`
	XXX_unrecognized []byte                         `json:"-"`
}

func (m *GenericRefreshResponseCollectionProto) Reset()         { *m = GenericRefreshResponseCollectionProto{} }
func (m *GenericRefreshResponseCollectionProto) String() string { return proto.CompactTextString(m) }
func (*GenericRefreshResponseCollectionProto) ProtoMessage()    {}

func (m *GenericRefreshResponseCollectionProto) GetResponses() []*GenericRefreshResponseProto {
	if m != nil {
		return m.Responses
	}
	return nil
}

func init() {
}
