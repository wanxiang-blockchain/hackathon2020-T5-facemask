// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        (unknown)
// source: monitor.proto

package proto

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type NodeNewRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Groupid      string `protobuf:"bytes,1,opt,name=groupid,proto3" json:"groupid,omitempty"`
	Chainid      string `protobuf:"bytes,2,opt,name=chainid,proto3" json:"chainid,omitempty"`
	Ip           string `protobuf:"bytes,3,opt,name=ip,proto3" json:"ip,omitempty"`
	Port         string `protobuf:"bytes,4,opt,name=port,proto3" json:"port,omitempty"`
	Rpcport      string `protobuf:"bytes,5,opt,name=rpcport,proto3" json:"rpcport,omitempty"`
	Wsport       string `protobuf:"bytes,6,opt,name=wsport,proto3" json:"wsport,omitempty"`
	Dashport     string `protobuf:"bytes,7,opt,name=dashport,proto3" json:"dashport,omitempty"`
	Password     string `protobuf:"bytes,8,opt,name=password,proto3" json:"password,omitempty"`
	CreatorEnode string `protobuf:"bytes,9,opt,name=creatorEnode,proto3" json:"creatorEnode,omitempty"`
	Bootnodes    string `protobuf:"bytes,10,opt,name=bootnodes,proto3" json:"bootnodes,omitempty"`
}

func (x *NodeNewRequest) Reset() {
	*x = NodeNewRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_monitor_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeNewRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeNewRequest) ProtoMessage() {}

func (x *NodeNewRequest) ProtoReflect() protoreflect.Message {
	mi := &file_monitor_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeNewRequest.ProtoReflect.Descriptor instead.
func (*NodeNewRequest) Descriptor() ([]byte, []int) {
	return file_monitor_proto_rawDescGZIP(), []int{0}
}

func (x *NodeNewRequest) GetGroupid() string {
	if x != nil {
		return x.Groupid
	}
	return ""
}

func (x *NodeNewRequest) GetChainid() string {
	if x != nil {
		return x.Chainid
	}
	return ""
}

func (x *NodeNewRequest) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *NodeNewRequest) GetPort() string {
	if x != nil {
		return x.Port
	}
	return ""
}

func (x *NodeNewRequest) GetRpcport() string {
	if x != nil {
		return x.Rpcport
	}
	return ""
}

func (x *NodeNewRequest) GetWsport() string {
	if x != nil {
		return x.Wsport
	}
	return ""
}

func (x *NodeNewRequest) GetDashport() string {
	if x != nil {
		return x.Dashport
	}
	return ""
}

func (x *NodeNewRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *NodeNewRequest) GetCreatorEnode() string {
	if x != nil {
		return x.CreatorEnode
	}
	return ""
}

func (x *NodeNewRequest) GetBootnodes() string {
	if x != nil {
		return x.Bootnodes
	}
	return ""
}

type NodeNewResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result  string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Pubkey  string `protobuf:"bytes,3,opt,name=pubkey,proto3" json:"pubkey,omitempty"`
}

func (x *NodeNewResponse) Reset() {
	*x = NodeNewResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_monitor_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeNewResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeNewResponse) ProtoMessage() {}

func (x *NodeNewResponse) ProtoReflect() protoreflect.Message {
	mi := &file_monitor_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeNewResponse.ProtoReflect.Descriptor instead.
func (*NodeNewResponse) Descriptor() ([]byte, []int) {
	return file_monitor_proto_rawDescGZIP(), []int{1}
}

func (x *NodeNewResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

func (x *NodeNewResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *NodeNewResponse) GetPubkey() string {
	if x != nil {
		return x.Pubkey
	}
	return ""
}

type NodeStartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Groupid string `protobuf:"bytes,1,opt,name=groupid,proto3" json:"groupid,omitempty"`
}

func (x *NodeStartRequest) Reset() {
	*x = NodeStartRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_monitor_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeStartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeStartRequest) ProtoMessage() {}

func (x *NodeStartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_monitor_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeStartRequest.ProtoReflect.Descriptor instead.
func (*NodeStartRequest) Descriptor() ([]byte, []int) {
	return file_monitor_proto_rawDescGZIP(), []int{2}
}

func (x *NodeStartRequest) GetGroupid() string {
	if x != nil {
		return x.Groupid
	}
	return ""
}

type NodeStartResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result  string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *NodeStartResponse) Reset() {
	*x = NodeStartResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_monitor_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeStartResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeStartResponse) ProtoMessage() {}

func (x *NodeStartResponse) ProtoReflect() protoreflect.Message {
	mi := &file_monitor_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeStartResponse.ProtoReflect.Descriptor instead.
func (*NodeStartResponse) Descriptor() ([]byte, []int) {
	return file_monitor_proto_rawDescGZIP(), []int{3}
}

func (x *NodeStartResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

func (x *NodeStartResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type NodeStopRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Groupid string `protobuf:"bytes,1,opt,name=groupid,proto3" json:"groupid,omitempty"`
}

func (x *NodeStopRequest) Reset() {
	*x = NodeStopRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_monitor_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeStopRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeStopRequest) ProtoMessage() {}

func (x *NodeStopRequest) ProtoReflect() protoreflect.Message {
	mi := &file_monitor_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeStopRequest.ProtoReflect.Descriptor instead.
func (*NodeStopRequest) Descriptor() ([]byte, []int) {
	return file_monitor_proto_rawDescGZIP(), []int{4}
}

func (x *NodeStopRequest) GetGroupid() string {
	if x != nil {
		return x.Groupid
	}
	return ""
}

type NodeStopResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result  string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *NodeStopResponse) Reset() {
	*x = NodeStopResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_monitor_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeStopResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeStopResponse) ProtoMessage() {}

func (x *NodeStopResponse) ProtoReflect() protoreflect.Message {
	mi := &file_monitor_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeStopResponse.ProtoReflect.Descriptor instead.
func (*NodeStopResponse) Descriptor() ([]byte, []int) {
	return file_monitor_proto_rawDescGZIP(), []int{5}
}

func (x *NodeStopResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

func (x *NodeStopResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type NodeRestartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Groupid string `protobuf:"bytes,1,opt,name=groupid,proto3" json:"groupid,omitempty"`
}

func (x *NodeRestartRequest) Reset() {
	*x = NodeRestartRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_monitor_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeRestartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeRestartRequest) ProtoMessage() {}

func (x *NodeRestartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_monitor_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeRestartRequest.ProtoReflect.Descriptor instead.
func (*NodeRestartRequest) Descriptor() ([]byte, []int) {
	return file_monitor_proto_rawDescGZIP(), []int{6}
}

func (x *NodeRestartRequest) GetGroupid() string {
	if x != nil {
		return x.Groupid
	}
	return ""
}

type NodeRestartResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result  string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *NodeRestartResponse) Reset() {
	*x = NodeRestartResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_monitor_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeRestartResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeRestartResponse) ProtoMessage() {}

func (x *NodeRestartResponse) ProtoReflect() protoreflect.Message {
	mi := &file_monitor_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeRestartResponse.ProtoReflect.Descriptor instead.
func (*NodeRestartResponse) Descriptor() ([]byte, []int) {
	return file_monitor_proto_rawDescGZIP(), []int{7}
}

func (x *NodeRestartResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

func (x *NodeRestartResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_monitor_proto protoreflect.FileDescriptor

var file_monitor_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x94, 0x02, 0x0a, 0x0e, 0x4e, 0x6f, 0x64, 0x65, 0x4e, 0x65, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x68, 0x61, 0x69, 0x6e, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63,
	0x68, 0x61, 0x69, 0x6e, 0x69, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x70,
	0x63, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x70, 0x63,
	0x70, 0x6f, 0x72, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x77, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x77, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x64, 0x61, 0x73, 0x68, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x64, 0x61, 0x73, 0x68, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x45,
	0x6e, 0x6f, 0x64, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x6f, 0x72, 0x45, 0x6e, 0x6f, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x62, 0x6f, 0x6f, 0x74,
	0x6e, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62, 0x6f, 0x6f,
	0x74, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x22, 0x5b, 0x0a, 0x0f, 0x4e, 0x6f, 0x64, 0x65, 0x4e, 0x65,
	0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70,
	0x75, 0x62, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x75, 0x62,
	0x6b, 0x65, 0x79, 0x22, 0x2c, 0x0a, 0x10, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x74, 0x61, 0x72, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x69,
	0x64, 0x22, 0x45, 0x0a, 0x11, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x2b, 0x0a, 0x0f, 0x4e, 0x6f, 0x64, 0x65,
	0x53, 0x74, 0x6f, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x69, 0x64, 0x22, 0x44, 0x0a, 0x10, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x74, 0x6f,
	0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x2e, 0x0a, 0x12, 0x4e,
	0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x69, 0x64, 0x22, 0x47, 0x0a, 0x13, 0x4e,
	0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x32, 0xc3, 0x01, 0x0a, 0x04, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x28, 0x0a,
	0x03, 0x4e, 0x65, 0x77, 0x12, 0x0f, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x4e, 0x65, 0x77, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x4e, 0x65, 0x77, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x72, 0x74,
	0x12, 0x11, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x04, 0x53, 0x74, 0x6f, 0x70, 0x12,
	0x10, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x74, 0x6f, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x11, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x74, 0x6f, 0x70, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x07, 0x52, 0x65, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12,
	0x13, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x3b,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_monitor_proto_rawDescOnce sync.Once
	file_monitor_proto_rawDescData = file_monitor_proto_rawDesc
)

func file_monitor_proto_rawDescGZIP() []byte {
	file_monitor_proto_rawDescOnce.Do(func() {
		file_monitor_proto_rawDescData = protoimpl.X.CompressGZIP(file_monitor_proto_rawDescData)
	})
	return file_monitor_proto_rawDescData
}

var file_monitor_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_monitor_proto_goTypes = []interface{}{
	(*NodeNewRequest)(nil),      // 0: NodeNewRequest
	(*NodeNewResponse)(nil),     // 1: NodeNewResponse
	(*NodeStartRequest)(nil),    // 2: NodeStartRequest
	(*NodeStartResponse)(nil),   // 3: NodeStartResponse
	(*NodeStopRequest)(nil),     // 4: NodeStopRequest
	(*NodeStopResponse)(nil),    // 5: NodeStopResponse
	(*NodeRestartRequest)(nil),  // 6: NodeRestartRequest
	(*NodeRestartResponse)(nil), // 7: NodeRestartResponse
}
var file_monitor_proto_depIdxs = []int32{
	0, // 0: Node.New:input_type -> NodeNewRequest
	2, // 1: Node.Start:input_type -> NodeStartRequest
	4, // 2: Node.Stop:input_type -> NodeStopRequest
	6, // 3: Node.ReStart:input_type -> NodeRestartRequest
	1, // 4: Node.New:output_type -> NodeNewResponse
	3, // 5: Node.Start:output_type -> NodeStartResponse
	5, // 6: Node.Stop:output_type -> NodeStopResponse
	7, // 7: Node.ReStart:output_type -> NodeRestartResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_monitor_proto_init() }
func file_monitor_proto_init() {
	if File_monitor_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_monitor_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeNewRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_monitor_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeNewResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_monitor_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeStartRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_monitor_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeStartResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_monitor_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeStopRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_monitor_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeStopResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_monitor_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeRestartRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_monitor_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeRestartResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_monitor_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_monitor_proto_goTypes,
		DependencyIndexes: file_monitor_proto_depIdxs,
		MessageInfos:      file_monitor_proto_msgTypes,
	}.Build()
	File_monitor_proto = out.File
	file_monitor_proto_rawDesc = nil
	file_monitor_proto_goTypes = nil
	file_monitor_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// NodeClient is the client API for Node service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NodeClient interface {
	New(ctx context.Context, in *NodeNewRequest, opts ...grpc.CallOption) (*NodeNewResponse, error)
	Start(ctx context.Context, in *NodeStartRequest, opts ...grpc.CallOption) (*NodeStartResponse, error)
	Stop(ctx context.Context, in *NodeStopRequest, opts ...grpc.CallOption) (*NodeStopResponse, error)
	ReStart(ctx context.Context, in *NodeRestartRequest, opts ...grpc.CallOption) (*NodeRestartResponse, error)
}

type nodeClient struct {
	cc grpc.ClientConnInterface
}

func NewNodeClient(cc grpc.ClientConnInterface) NodeClient {
	return &nodeClient{cc}
}

func (c *nodeClient) New(ctx context.Context, in *NodeNewRequest, opts ...grpc.CallOption) (*NodeNewResponse, error) {
	out := new(NodeNewResponse)
	err := c.cc.Invoke(ctx, "/Node/New", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) Start(ctx context.Context, in *NodeStartRequest, opts ...grpc.CallOption) (*NodeStartResponse, error) {
	out := new(NodeStartResponse)
	err := c.cc.Invoke(ctx, "/Node/Start", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) Stop(ctx context.Context, in *NodeStopRequest, opts ...grpc.CallOption) (*NodeStopResponse, error) {
	out := new(NodeStopResponse)
	err := c.cc.Invoke(ctx, "/Node/Stop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) ReStart(ctx context.Context, in *NodeRestartRequest, opts ...grpc.CallOption) (*NodeRestartResponse, error) {
	out := new(NodeRestartResponse)
	err := c.cc.Invoke(ctx, "/Node/ReStart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NodeServer is the server API for Node service.
type NodeServer interface {
	New(context.Context, *NodeNewRequest) (*NodeNewResponse, error)
	Start(context.Context, *NodeStartRequest) (*NodeStartResponse, error)
	Stop(context.Context, *NodeStopRequest) (*NodeStopResponse, error)
	ReStart(context.Context, *NodeRestartRequest) (*NodeRestartResponse, error)
}

// UnimplementedNodeServer can be embedded to have forward compatible implementations.
type UnimplementedNodeServer struct {
}

func (*UnimplementedNodeServer) New(context.Context, *NodeNewRequest) (*NodeNewResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method New not implemented")
}
func (*UnimplementedNodeServer) Start(context.Context, *NodeStartRequest) (*NodeStartResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Start not implemented")
}
func (*UnimplementedNodeServer) Stop(context.Context, *NodeStopRequest) (*NodeStopResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stop not implemented")
}
func (*UnimplementedNodeServer) ReStart(context.Context, *NodeRestartRequest) (*NodeRestartResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReStart not implemented")
}

func RegisterNodeServer(s *grpc.Server, srv NodeServer) {
	s.RegisterService(&_Node_serviceDesc, srv)
}

func _Node_New_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeNewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).New(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Node/New",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).New(ctx, req.(*NodeNewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeStartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Node/Start",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).Start(ctx, req.(*NodeStartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeStopRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Node/Stop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).Stop(ctx, req.(*NodeStopRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_ReStart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeRestartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).ReStart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Node/ReStart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).ReStart(ctx, req.(*NodeRestartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Node_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Node",
	HandlerType: (*NodeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "New",
			Handler:    _Node_New_Handler,
		},
		{
			MethodName: "Start",
			Handler:    _Node_Start_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _Node_Stop_Handler,
		},
		{
			MethodName: "ReStart",
			Handler:    _Node_ReStart_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "monitor.proto",
}