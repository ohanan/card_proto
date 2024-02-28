// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.3
// source: card.proto

package proto

import (
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

type RegisterNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RegisterNotify) Reset() {
	*x = RegisterNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_card_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterNotify) ProtoMessage() {}

func (x *RegisterNotify) ProtoReflect() protoreflect.Message {
	mi := &file_card_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterNotify.ProtoReflect.Descriptor instead.
func (*RegisterNotify) Descriptor() ([]byte, []int) {
	return file_card_proto_rawDescGZIP(), []int{0}
}

type GetPluginInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetPluginInfo) Reset() {
	*x = GetPluginInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_card_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPluginInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPluginInfo) ProtoMessage() {}

func (x *GetPluginInfo) ProtoReflect() protoreflect.Message {
	mi := &file_card_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPluginInfo.ProtoReflect.Descriptor instead.
func (*GetPluginInfo) Descriptor() ([]byte, []int) {
	return file_card_proto_rawDescGZIP(), []int{1}
}

type PluginInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version     *Version `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	Author      string   `protobuf:"bytes,2,opt,name=author,proto3" json:"author,omitempty"`
	Name        string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Description string   `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *PluginInfo) Reset() {
	*x = PluginInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_card_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PluginInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PluginInfo) ProtoMessage() {}

func (x *PluginInfo) ProtoReflect() protoreflect.Message {
	mi := &file_card_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PluginInfo.ProtoReflect.Descriptor instead.
func (*PluginInfo) Descriptor() ([]byte, []int) {
	return file_card_proto_rawDescGZIP(), []int{2}
}

func (x *PluginInfo) GetVersion() *Version {
	if x != nil {
		return x.Version
	}
	return nil
}

func (x *PluginInfo) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *PluginInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PluginInfo) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type Version struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Major uint32 `protobuf:"varint,1,opt,name=major,proto3" json:"major,omitempty"`
	Minor uint32 `protobuf:"varint,2,opt,name=minor,proto3" json:"minor,omitempty"`
	Patch uint32 `protobuf:"varint,3,opt,name=patch,proto3" json:"patch,omitempty"`
}

func (x *Version) Reset() {
	*x = Version{}
	if protoimpl.UnsafeEnabled {
		mi := &file_card_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Version) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Version) ProtoMessage() {}

func (x *Version) ProtoReflect() protoreflect.Message {
	mi := &file_card_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Version.ProtoReflect.Descriptor instead.
func (*Version) Descriptor() ([]byte, []int) {
	return file_card_proto_rawDescGZIP(), []int{3}
}

func (x *Version) GetMajor() uint32 {
	if x != nil {
		return x.Major
	}
	return 0
}

func (x *Version) GetMinor() uint32 {
	if x != nil {
		return x.Minor
	}
	return 0
}

func (x *Version) GetPatch() uint32 {
	if x != nil {
		return x.Patch
	}
	return 0
}

type PlayerInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PlayerInfo) Reset() {
	*x = PlayerInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_card_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerInfo) ProtoMessage() {}

func (x *PlayerInfo) ProtoReflect() protoreflect.Message {
	mi := &file_card_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerInfo.ProtoReflect.Descriptor instead.
func (*PlayerInfo) Descriptor() ([]byte, []int) {
	return file_card_proto_rawDescGZIP(), []int{4}
}

type GetPlayerInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetPlayerInfo) Reset() {
	*x = GetPlayerInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_card_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPlayerInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPlayerInfo) ProtoMessage() {}

func (x *GetPlayerInfo) ProtoReflect() protoreflect.Message {
	mi := &file_card_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPlayerInfo.ProtoReflect.Descriptor instead.
func (*GetPlayerInfo) Descriptor() ([]byte, []int) {
	return file_card_proto_rawDescGZIP(), []int{5}
}

type RegisterNotify_Req struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Event []string `protobuf:"bytes,1,rep,name=event,proto3" json:"event,omitempty"`
}

func (x *RegisterNotify_Req) Reset() {
	*x = RegisterNotify_Req{}
	if protoimpl.UnsafeEnabled {
		mi := &file_card_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterNotify_Req) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterNotify_Req) ProtoMessage() {}

func (x *RegisterNotify_Req) ProtoReflect() protoreflect.Message {
	mi := &file_card_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterNotify_Req.ProtoReflect.Descriptor instead.
func (*RegisterNotify_Req) Descriptor() ([]byte, []int) {
	return file_card_proto_rawDescGZIP(), []int{0, 0}
}

func (x *RegisterNotify_Req) GetEvent() []string {
	if x != nil {
		return x.Event
	}
	return nil
}

type RegisterNotify_Resp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RegisterNotify_Resp) Reset() {
	*x = RegisterNotify_Resp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_card_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterNotify_Resp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterNotify_Resp) ProtoMessage() {}

func (x *RegisterNotify_Resp) ProtoReflect() protoreflect.Message {
	mi := &file_card_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterNotify_Resp.ProtoReflect.Descriptor instead.
func (*RegisterNotify_Resp) Descriptor() ([]byte, []int) {
	return file_card_proto_rawDescGZIP(), []int{0, 1}
}

type GetPluginInfo_Req struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetPluginInfo_Req) Reset() {
	*x = GetPluginInfo_Req{}
	if protoimpl.UnsafeEnabled {
		mi := &file_card_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPluginInfo_Req) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPluginInfo_Req) ProtoMessage() {}

func (x *GetPluginInfo_Req) ProtoReflect() protoreflect.Message {
	mi := &file_card_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPluginInfo_Req.ProtoReflect.Descriptor instead.
func (*GetPluginInfo_Req) Descriptor() ([]byte, []int) {
	return file_card_proto_rawDescGZIP(), []int{1, 0}
}

type GetPluginInfo_Resp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *PluginInfo `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *GetPluginInfo_Resp) Reset() {
	*x = GetPluginInfo_Resp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_card_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPluginInfo_Resp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPluginInfo_Resp) ProtoMessage() {}

func (x *GetPluginInfo_Resp) ProtoReflect() protoreflect.Message {
	mi := &file_card_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPluginInfo_Resp.ProtoReflect.Descriptor instead.
func (*GetPluginInfo_Resp) Descriptor() ([]byte, []int) {
	return file_card_proto_rawDescGZIP(), []int{1, 1}
}

func (x *GetPluginInfo_Resp) GetInfo() *PluginInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

type GetPlayerInfo_Req struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetPlayerInfo_Req) Reset() {
	*x = GetPlayerInfo_Req{}
	if protoimpl.UnsafeEnabled {
		mi := &file_card_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPlayerInfo_Req) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPlayerInfo_Req) ProtoMessage() {}

func (x *GetPlayerInfo_Req) ProtoReflect() protoreflect.Message {
	mi := &file_card_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPlayerInfo_Req.ProtoReflect.Descriptor instead.
func (*GetPlayerInfo_Req) Descriptor() ([]byte, []int) {
	return file_card_proto_rawDescGZIP(), []int{5, 0}
}

func (x *GetPlayerInfo_Req) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type GetPlayerInfo_Resp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *PlayerInfo `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *GetPlayerInfo_Resp) Reset() {
	*x = GetPlayerInfo_Resp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_card_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPlayerInfo_Resp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPlayerInfo_Resp) ProtoMessage() {}

func (x *GetPlayerInfo_Resp) ProtoReflect() protoreflect.Message {
	mi := &file_card_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPlayerInfo_Resp.ProtoReflect.Descriptor instead.
func (*GetPlayerInfo_Resp) Descriptor() ([]byte, []int) {
	return file_card_proto_rawDescGZIP(), []int{5, 1}
}

func (x *GetPlayerInfo_Resp) GetInfo() *PlayerInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_card_proto protoreflect.FileDescriptor

var file_card_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x63, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x35, 0x0a, 0x0e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x79, 0x1a, 0x1b, 0x0a, 0x03, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x1a, 0x06, 0x0a, 0x04, 0x52, 0x65, 0x73, 0x70, 0x22, 0x45, 0x0a, 0x0d, 0x47, 0x65,
	0x74, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x05, 0x0a, 0x03, 0x52,
	0x65, 0x71, 0x1a, 0x2d, 0x0a, 0x04, 0x52, 0x65, 0x73, 0x70, 0x12, 0x25, 0x0a, 0x04, 0x69, 0x6e,
	0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66,
	0x6f, 0x22, 0x84, 0x01, 0x0a, 0x0a, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x28, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x4b, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x61, 0x6a, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x05, 0x6d, 0x61, 0x6a, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x69, 0x6e,
	0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6d, 0x69, 0x6e, 0x6f, 0x72, 0x12,
	0x14, 0x0a, 0x05, 0x70, 0x61, 0x74, 0x63, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05,
	0x70, 0x61, 0x74, 0x63, 0x68, 0x22, 0x0c, 0x0a, 0x0a, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x22, 0x5e, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x1e, 0x0a, 0x03, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x1a, 0x2d, 0x0a, 0x04, 0x52, 0x65, 0x73, 0x70, 0x12, 0x25, 0x0a, 0x04,
	0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69,
	0x6e, 0x66, 0x6f, 0x32, 0x95, 0x01, 0x0a, 0x04, 0x43, 0x61, 0x72, 0x64, 0x12, 0x44, 0x0a, 0x0d,
	0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x18, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x2e, 0x52, 0x65, 0x71, 0x1a, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x47, 0x0a, 0x0e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x79, 0x12, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x52, 0x65, 0x71, 0x1a,
	0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x32, 0x4e, 0x0a, 0x06, 0x50,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x12, 0x44, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47,
	0x65, 0x74, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x52, 0x65, 0x71,
	0x1a, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x42, 0x09, 0x5a, 0x07, 0x2e,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_card_proto_rawDescOnce sync.Once
	file_card_proto_rawDescData = file_card_proto_rawDesc
)

func file_card_proto_rawDescGZIP() []byte {
	file_card_proto_rawDescOnce.Do(func() {
		file_card_proto_rawDescData = protoimpl.X.CompressGZIP(file_card_proto_rawDescData)
	})
	return file_card_proto_rawDescData
}

var file_card_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_card_proto_goTypes = []interface{}{
	(*RegisterNotify)(nil),      // 0: proto.RegisterNotify
	(*GetPluginInfo)(nil),       // 1: proto.GetPluginInfo
	(*PluginInfo)(nil),          // 2: proto.PluginInfo
	(*Version)(nil),             // 3: proto.Version
	(*PlayerInfo)(nil),          // 4: proto.PlayerInfo
	(*GetPlayerInfo)(nil),       // 5: proto.GetPlayerInfo
	(*RegisterNotify_Req)(nil),  // 6: proto.RegisterNotify.Req
	(*RegisterNotify_Resp)(nil), // 7: proto.RegisterNotify.Resp
	(*GetPluginInfo_Req)(nil),   // 8: proto.GetPluginInfo.Req
	(*GetPluginInfo_Resp)(nil),  // 9: proto.GetPluginInfo.Resp
	(*GetPlayerInfo_Req)(nil),   // 10: proto.GetPlayerInfo.Req
	(*GetPlayerInfo_Resp)(nil),  // 11: proto.GetPlayerInfo.Resp
}
var file_card_proto_depIdxs = []int32{
	3,  // 0: proto.PluginInfo.version:type_name -> proto.Version
	2,  // 1: proto.GetPluginInfo.Resp.info:type_name -> proto.PluginInfo
	4,  // 2: proto.GetPlayerInfo.Resp.info:type_name -> proto.PlayerInfo
	10, // 3: proto.Card.GetPlayerInfo:input_type -> proto.GetPlayerInfo.Req
	6,  // 4: proto.Card.RegisterNotify:input_type -> proto.RegisterNotify.Req
	8,  // 5: proto.Plugin.GetPluginInfo:input_type -> proto.GetPluginInfo.Req
	11, // 6: proto.Card.GetPlayerInfo:output_type -> proto.GetPlayerInfo.Resp
	7,  // 7: proto.Card.RegisterNotify:output_type -> proto.RegisterNotify.Resp
	9,  // 8: proto.Plugin.GetPluginInfo:output_type -> proto.GetPluginInfo.Resp
	6,  // [6:9] is the sub-list for method output_type
	3,  // [3:6] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_card_proto_init() }
func file_card_proto_init() {
	if File_card_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_card_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterNotify); i {
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
		file_card_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPluginInfo); i {
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
		file_card_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PluginInfo); i {
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
		file_card_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Version); i {
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
		file_card_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerInfo); i {
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
		file_card_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPlayerInfo); i {
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
		file_card_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterNotify_Req); i {
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
		file_card_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterNotify_Resp); i {
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
		file_card_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPluginInfo_Req); i {
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
		file_card_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPluginInfo_Resp); i {
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
		file_card_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPlayerInfo_Req); i {
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
		file_card_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPlayerInfo_Resp); i {
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
			RawDescriptor: file_card_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_card_proto_goTypes,
		DependencyIndexes: file_card_proto_depIdxs,
		MessageInfos:      file_card_proto_msgTypes,
	}.Build()
	File_card_proto = out.File
	file_card_proto_rawDesc = nil
	file_card_proto_goTypes = nil
	file_card_proto_depIdxs = nil
}
