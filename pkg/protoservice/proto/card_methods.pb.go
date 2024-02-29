// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.1
// source: proto/card_methods.proto

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

type Hello struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Hello) Reset() {
	*x = Hello{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_card_methods_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Hello) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Hello) ProtoMessage() {}

func (x *Hello) ProtoReflect() protoreflect.Message {
	mi := &file_proto_card_methods_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Hello.ProtoReflect.Descriptor instead.
func (*Hello) Descriptor() ([]byte, []int) {
	return file_proto_card_methods_proto_rawDescGZIP(), []int{0}
}

type GetPlayerInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetPlayerInfo) Reset() {
	*x = GetPlayerInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_card_methods_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPlayerInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPlayerInfo) ProtoMessage() {}

func (x *GetPlayerInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_card_methods_proto_msgTypes[1]
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
	return file_proto_card_methods_proto_rawDescGZIP(), []int{1}
}

type RegisterNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RegisterNotify) Reset() {
	*x = RegisterNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_card_methods_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterNotify) ProtoMessage() {}

func (x *RegisterNotify) ProtoReflect() protoreflect.Message {
	mi := &file_proto_card_methods_proto_msgTypes[2]
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
	return file_proto_card_methods_proto_rawDescGZIP(), []int{2}
}

type Hello_Req struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Abcd string `protobuf:"bytes,1,opt,name=abcd,proto3" json:"abcd,omitempty"`
}

func (x *Hello_Req) Reset() {
	*x = Hello_Req{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_card_methods_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Hello_Req) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Hello_Req) ProtoMessage() {}

func (x *Hello_Req) ProtoReflect() protoreflect.Message {
	mi := &file_proto_card_methods_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Hello_Req.ProtoReflect.Descriptor instead.
func (*Hello_Req) Descriptor() ([]byte, []int) {
	return file_proto_card_methods_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Hello_Req) GetAbcd() string {
	if x != nil {
		return x.Abcd
	}
	return ""
}

type Hello_Resp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Hello_Resp) Reset() {
	*x = Hello_Resp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_card_methods_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Hello_Resp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Hello_Resp) ProtoMessage() {}

func (x *Hello_Resp) ProtoReflect() protoreflect.Message {
	mi := &file_proto_card_methods_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Hello_Resp.ProtoReflect.Descriptor instead.
func (*Hello_Resp) Descriptor() ([]byte, []int) {
	return file_proto_card_methods_proto_rawDescGZIP(), []int{0, 1}
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
		mi := &file_proto_card_methods_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPlayerInfo_Req) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPlayerInfo_Req) ProtoMessage() {}

func (x *GetPlayerInfo_Req) ProtoReflect() protoreflect.Message {
	mi := &file_proto_card_methods_proto_msgTypes[5]
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
	return file_proto_card_methods_proto_rawDescGZIP(), []int{1, 0}
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
		mi := &file_proto_card_methods_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPlayerInfo_Resp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPlayerInfo_Resp) ProtoMessage() {}

func (x *GetPlayerInfo_Resp) ProtoReflect() protoreflect.Message {
	mi := &file_proto_card_methods_proto_msgTypes[6]
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
	return file_proto_card_methods_proto_rawDescGZIP(), []int{1, 1}
}

func (x *GetPlayerInfo_Resp) GetInfo() *PlayerInfo {
	if x != nil {
		return x.Info
	}
	return nil
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
		mi := &file_proto_card_methods_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterNotify_Req) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterNotify_Req) ProtoMessage() {}

func (x *RegisterNotify_Req) ProtoReflect() protoreflect.Message {
	mi := &file_proto_card_methods_proto_msgTypes[7]
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
	return file_proto_card_methods_proto_rawDescGZIP(), []int{2, 0}
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
		mi := &file_proto_card_methods_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterNotify_Resp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterNotify_Resp) ProtoMessage() {}

func (x *RegisterNotify_Resp) ProtoReflect() protoreflect.Message {
	mi := &file_proto_card_methods_proto_msgTypes[8]
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
	return file_proto_card_methods_proto_rawDescGZIP(), []int{2, 1}
}

var File_proto_card_methods_proto protoreflect.FileDescriptor

var file_proto_card_methods_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x6d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2a, 0x0a, 0x05, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x1a, 0x19,
	0x0a, 0x03, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x62, 0x63, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x62, 0x63, 0x64, 0x1a, 0x06, 0x0a, 0x04, 0x52, 0x65, 0x73,
	0x70, 0x22, 0x5e, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x1a, 0x1e, 0x0a, 0x03, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x1a, 0x2d, 0x0a, 0x04, 0x52, 0x65, 0x73, 0x70, 0x12, 0x25, 0x0a, 0x04, 0x69, 0x6e,
	0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66,
	0x6f, 0x22, 0x35, 0x0a, 0x0e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x79, 0x1a, 0x1b, 0x0a, 0x03, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x1a, 0x06, 0x0a, 0x04, 0x52, 0x65, 0x73, 0x70, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x68, 0x61, 0x6e, 0x61, 0x6e, 0x2f, 0x63, 0x61,
	0x72, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_card_methods_proto_rawDescOnce sync.Once
	file_proto_card_methods_proto_rawDescData = file_proto_card_methods_proto_rawDesc
)

func file_proto_card_methods_proto_rawDescGZIP() []byte {
	file_proto_card_methods_proto_rawDescOnce.Do(func() {
		file_proto_card_methods_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_card_methods_proto_rawDescData)
	})
	return file_proto_card_methods_proto_rawDescData
}

var file_proto_card_methods_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_proto_card_methods_proto_goTypes = []interface{}{
	(*Hello)(nil),               // 0: proto.Hello
	(*GetPlayerInfo)(nil),       // 1: proto.GetPlayerInfo
	(*RegisterNotify)(nil),      // 2: proto.RegisterNotify
	(*Hello_Req)(nil),           // 3: proto.Hello.Req
	(*Hello_Resp)(nil),          // 4: proto.Hello.Resp
	(*GetPlayerInfo_Req)(nil),   // 5: proto.GetPlayerInfo.Req
	(*GetPlayerInfo_Resp)(nil),  // 6: proto.GetPlayerInfo.Resp
	(*RegisterNotify_Req)(nil),  // 7: proto.RegisterNotify.Req
	(*RegisterNotify_Resp)(nil), // 8: proto.RegisterNotify.Resp
	(*PlayerInfo)(nil),          // 9: proto.PlayerInfo
}
var file_proto_card_methods_proto_depIdxs = []int32{
	9, // 0: proto.GetPlayerInfo.Resp.info:type_name -> proto.PlayerInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_card_methods_proto_init() }
func file_proto_card_methods_proto_init() {
	if File_proto_card_methods_proto != nil {
		return
	}
	file_proto_player_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_card_methods_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Hello); i {
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
		file_proto_card_methods_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_card_methods_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_card_methods_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Hello_Req); i {
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
		file_proto_card_methods_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Hello_Resp); i {
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
		file_proto_card_methods_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_card_methods_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_card_methods_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_card_methods_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_card_methods_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_card_methods_proto_goTypes,
		DependencyIndexes: file_proto_card_methods_proto_depIdxs,
		MessageInfos:      file_proto_card_methods_proto_msgTypes,
	}.Build()
	File_proto_card_methods_proto = out.File
	file_proto_card_methods_proto_rawDesc = nil
	file_proto_card_methods_proto_goTypes = nil
	file_proto_card_methods_proto_depIdxs = nil
}
