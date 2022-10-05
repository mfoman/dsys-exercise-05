// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: api/v1/time.proto

package api

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type NTPTime struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Settings       uint32 `protobuf:"fixed32,1,opt,name=settings,proto3" json:"settings,omitempty"`
	Stratum        uint32 `protobuf:"fixed32,2,opt,name=stratum,proto3" json:"stratum,omitempty"`
	Poll           int32  `protobuf:"varint,3,opt,name=poll,proto3" json:"poll,omitempty"`
	Precision      int32  `protobuf:"varint,4,opt,name=precision,proto3" json:"precision,omitempty"`
	Rootdelay      int32  `protobuf:"varint,5,opt,name=rootdelay,proto3" json:"rootdelay,omitempty"`
	Rootdispersion int32  `protobuf:"varint,6,opt,name=rootdispersion,proto3" json:"rootdispersion,omitempty"`
	Referenceid    int32  `protobuf:"varint,7,opt,name=referenceid,proto3" json:"referenceid,omitempty"`
	Reftimesec     int32  `protobuf:"varint,8,opt,name=reftimesec,proto3" json:"reftimesec,omitempty"`
	Reftimefrac    int32  `protobuf:"varint,9,opt,name=reftimefrac,proto3" json:"reftimefrac,omitempty"`
	Origtimesec    int32  `protobuf:"varint,10,opt,name=origtimesec,proto3" json:"origtimesec,omitempty"`
	Orgtimefrac    int32  `protobuf:"varint,11,opt,name=orgtimefrac,proto3" json:"orgtimefrac,omitempty"`
	Rxtimesec      int32  `protobuf:"varint,12,opt,name=rxtimesec,proto3" json:"rxtimesec,omitempty"`
	Rxtimefrac     int32  `protobuf:"varint,13,opt,name=rxtimefrac,proto3" json:"rxtimefrac,omitempty"`
	Txtimesec      int32  `protobuf:"varint,14,opt,name=txtimesec,proto3" json:"txtimesec,omitempty"`
	Txtimefrac     int32  `protobuf:"varint,15,opt,name=txtimefrac,proto3" json:"txtimefrac,omitempty"`
}

func (x *NTPTime) Reset() {
	*x = NTPTime{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_time_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NTPTime) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NTPTime) ProtoMessage() {}

func (x *NTPTime) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_time_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NTPTime.ProtoReflect.Descriptor instead.
func (*NTPTime) Descriptor() ([]byte, []int) {
	return file_api_v1_time_proto_rawDescGZIP(), []int{0}
}

func (x *NTPTime) GetSettings() uint32 {
	if x != nil {
		return x.Settings
	}
	return 0
}

func (x *NTPTime) GetStratum() uint32 {
	if x != nil {
		return x.Stratum
	}
	return 0
}

func (x *NTPTime) GetPoll() int32 {
	if x != nil {
		return x.Poll
	}
	return 0
}

func (x *NTPTime) GetPrecision() int32 {
	if x != nil {
		return x.Precision
	}
	return 0
}

func (x *NTPTime) GetRootdelay() int32 {
	if x != nil {
		return x.Rootdelay
	}
	return 0
}

func (x *NTPTime) GetRootdispersion() int32 {
	if x != nil {
		return x.Rootdispersion
	}
	return 0
}

func (x *NTPTime) GetReferenceid() int32 {
	if x != nil {
		return x.Referenceid
	}
	return 0
}

func (x *NTPTime) GetReftimesec() int32 {
	if x != nil {
		return x.Reftimesec
	}
	return 0
}

func (x *NTPTime) GetReftimefrac() int32 {
	if x != nil {
		return x.Reftimefrac
	}
	return 0
}

func (x *NTPTime) GetOrigtimesec() int32 {
	if x != nil {
		return x.Origtimesec
	}
	return 0
}

func (x *NTPTime) GetOrgtimefrac() int32 {
	if x != nil {
		return x.Orgtimefrac
	}
	return 0
}

func (x *NTPTime) GetRxtimesec() int32 {
	if x != nil {
		return x.Rxtimesec
	}
	return 0
}

func (x *NTPTime) GetRxtimefrac() int32 {
	if x != nil {
		return x.Rxtimefrac
	}
	return 0
}

func (x *NTPTime) GetTxtimesec() int32 {
	if x != nil {
		return x.Txtimesec
	}
	return 0
}

func (x *NTPTime) GetTxtimefrac() int32 {
	if x != nil {
		return x.Txtimefrac
	}
	return 0
}

type RoundTime struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	T2 *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=t2,proto3" json:"t2,omitempty"`
	T3 *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=t3,proto3" json:"t3,omitempty"`
}

func (x *RoundTime) Reset() {
	*x = RoundTime{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_time_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoundTime) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoundTime) ProtoMessage() {}

func (x *RoundTime) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_time_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoundTime.ProtoReflect.Descriptor instead.
func (*RoundTime) Descriptor() ([]byte, []int) {
	return file_api_v1_time_proto_rawDescGZIP(), []int{1}
}

func (x *RoundTime) GetT2() *timestamppb.Timestamp {
	if x != nil {
		return x.T2
	}
	return nil
}

func (x *RoundTime) GetT3() *timestamppb.Timestamp {
	if x != nil {
		return x.T3
	}
	return nil
}

type Time struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Time *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
}

func (x *Time) Reset() {
	*x = Time{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_time_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Time) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Time) ProtoMessage() {}

func (x *Time) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_time_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Time.ProtoReflect.Descriptor instead.
func (*Time) Descriptor() ([]byte, []int) {
	return file_api_v1_time_proto_rawDescGZIP(), []int{2}
}

func (x *Time) GetTime() *timestamppb.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

// The request message containing the user's name.
type HelloRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *HelloRequest) Reset() {
	*x = HelloRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_time_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloRequest) ProtoMessage() {}

func (x *HelloRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_time_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloRequest.ProtoReflect.Descriptor instead.
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_time_proto_rawDescGZIP(), []int{3}
}

func (x *HelloRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// The response message containing the greetings
type HelloReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *HelloReply) Reset() {
	*x = HelloReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_time_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloReply) ProtoMessage() {}

func (x *HelloReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_time_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloReply.ProtoReflect.Descriptor instead.
func (*HelloReply) Descriptor() ([]byte, []int) {
	return file_api_v1_time_proto_rawDescGZIP(), []int{4}
}

func (x *HelloReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_api_v1_time_proto protoreflect.FileDescriptor

var file_api_v1_time_proto_rawDesc = []byte{
	0x0a, 0x11, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdb, 0x03, 0x0a, 0x07, 0x4e,
	0x54, 0x50, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e,
	0x67, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x07, 0x52, 0x08, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e,
	0x67, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x74, 0x72, 0x61, 0x74, 0x75, 0x6d, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x07, 0x52, 0x07, 0x73, 0x74, 0x72, 0x61, 0x74, 0x75, 0x6d, 0x12, 0x12, 0x0a, 0x04,
	0x70, 0x6f, 0x6c, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x6f, 0x6c, 0x6c,
	0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x09, 0x70, 0x72, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1c,
	0x0a, 0x09, 0x72, 0x6f, 0x6f, 0x74, 0x64, 0x65, 0x6c, 0x61, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x09, 0x72, 0x6f, 0x6f, 0x74, 0x64, 0x65, 0x6c, 0x61, 0x79, 0x12, 0x26, 0x0a, 0x0e,
	0x72, 0x6f, 0x6f, 0x74, 0x64, 0x69, 0x73, 0x70, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x72, 0x6f, 0x6f, 0x74, 0x64, 0x69, 0x73, 0x70, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63,
	0x65, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x72, 0x65, 0x66, 0x65, 0x72,
	0x65, 0x6e, 0x63, 0x65, 0x69, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x66, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x65, 0x63, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x72, 0x65, 0x66, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x65, 0x63, 0x12, 0x20, 0x0a, 0x0b, 0x72, 0x65, 0x66, 0x74, 0x69, 0x6d,
	0x65, 0x66, 0x72, 0x61, 0x63, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x72, 0x65, 0x66,
	0x74, 0x69, 0x6d, 0x65, 0x66, 0x72, 0x61, 0x63, 0x12, 0x20, 0x0a, 0x0b, 0x6f, 0x72, 0x69, 0x67,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x65, 0x63, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x6f,
	0x72, 0x69, 0x67, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x65, 0x63, 0x12, 0x20, 0x0a, 0x0b, 0x6f, 0x72,
	0x67, 0x74, 0x69, 0x6d, 0x65, 0x66, 0x72, 0x61, 0x63, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0b, 0x6f, 0x72, 0x67, 0x74, 0x69, 0x6d, 0x65, 0x66, 0x72, 0x61, 0x63, 0x12, 0x1c, 0x0a, 0x09,
	0x72, 0x78, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x65, 0x63, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x09, 0x72, 0x78, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x65, 0x63, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x78,
	0x74, 0x69, 0x6d, 0x65, 0x66, 0x72, 0x61, 0x63, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a,
	0x72, 0x78, 0x74, 0x69, 0x6d, 0x65, 0x66, 0x72, 0x61, 0x63, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x78,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x65, 0x63, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x74,
	0x78, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x65, 0x63, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x78, 0x74, 0x69,
	0x6d, 0x65, 0x66, 0x72, 0x61, 0x63, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x78,
	0x74, 0x69, 0x6d, 0x65, 0x66, 0x72, 0x61, 0x63, 0x22, 0x63, 0x0a, 0x09, 0x52, 0x6f, 0x75, 0x6e,
	0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x02, 0x74, 0x32, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x02, 0x74,
	0x32, 0x12, 0x2a, 0x0a, 0x02, 0x74, 0x33, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x02, 0x74, 0x33, 0x22, 0x36, 0x0a,
	0x04, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x04, 0x74, 0x69, 0x6d, 0x65, 0x22, 0x22, 0x0a, 0x0c, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x26, 0x0a, 0x0a, 0x48, 0x65, 0x6c,
	0x6c, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x32, 0xc4, 0x01, 0x0a, 0x07, 0x47, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x12, 0x32, 0x0a,
	0x08, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x12, 0x2e, 0x74, 0x69, 0x6d, 0x65,
	0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e,
	0x74, 0x69, 0x6d, 0x65, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x00, 0x12, 0x23, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x0a, 0x2e, 0x74,
	0x69, 0x6d, 0x65, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x1a, 0x0a, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x22, 0x00, 0x12, 0x2c, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x4e, 0x54, 0x50,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x0d, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x4e, 0x54, 0x50, 0x54,
	0x69, 0x6d, 0x65, 0x1a, 0x0d, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x4e, 0x54, 0x50, 0x54, 0x69,
	0x6d, 0x65, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x75, 0x6e, 0x64,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x0f, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x52, 0x6f, 0x75, 0x6e,
	0x64, 0x54, 0x69, 0x6d, 0x65, 0x1a, 0x0f, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x52, 0x6f, 0x75,
	0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x00, 0x42, 0x10, 0x5a, 0x0e, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_api_v1_time_proto_rawDescOnce sync.Once
	file_api_v1_time_proto_rawDescData = file_api_v1_time_proto_rawDesc
)

func file_api_v1_time_proto_rawDescGZIP() []byte {
	file_api_v1_time_proto_rawDescOnce.Do(func() {
		file_api_v1_time_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1_time_proto_rawDescData)
	})
	return file_api_v1_time_proto_rawDescData
}

var file_api_v1_time_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_api_v1_time_proto_goTypes = []interface{}{
	(*NTPTime)(nil),               // 0: time.NTPTime
	(*RoundTime)(nil),             // 1: time.RoundTime
	(*Time)(nil),                  // 2: time.Time
	(*HelloRequest)(nil),          // 3: time.HelloRequest
	(*HelloReply)(nil),            // 4: time.HelloReply
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
}
var file_api_v1_time_proto_depIdxs = []int32{
	5, // 0: time.RoundTime.t2:type_name -> google.protobuf.Timestamp
	5, // 1: time.RoundTime.t3:type_name -> google.protobuf.Timestamp
	5, // 2: time.Time.time:type_name -> google.protobuf.Timestamp
	3, // 3: time.Greeter.SayHello:input_type -> time.HelloRequest
	2, // 4: time.Greeter.GetTime:input_type -> time.Time
	0, // 5: time.Greeter.GetNTPTime:input_type -> time.NTPTime
	1, // 6: time.Greeter.GetRoundTime:input_type -> time.RoundTime
	4, // 7: time.Greeter.SayHello:output_type -> time.HelloReply
	2, // 8: time.Greeter.GetTime:output_type -> time.Time
	0, // 9: time.Greeter.GetNTPTime:output_type -> time.NTPTime
	1, // 10: time.Greeter.GetRoundTime:output_type -> time.RoundTime
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_api_v1_time_proto_init() }
func file_api_v1_time_proto_init() {
	if File_api_v1_time_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_v1_time_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NTPTime); i {
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
		file_api_v1_time_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoundTime); i {
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
		file_api_v1_time_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Time); i {
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
		file_api_v1_time_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloRequest); i {
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
		file_api_v1_time_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloReply); i {
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
			RawDescriptor: file_api_v1_time_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v1_time_proto_goTypes,
		DependencyIndexes: file_api_v1_time_proto_depIdxs,
		MessageInfos:      file_api_v1_time_proto_msgTypes,
	}.Build()
	File_api_v1_time_proto = out.File
	file_api_v1_time_proto_rawDesc = nil
	file_api_v1_time_proto_goTypes = nil
	file_api_v1_time_proto_depIdxs = nil
}