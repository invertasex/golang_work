// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: proto/telegram.proto

package server

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CheckUserRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	BotToken      string                 `protobuf:"bytes,1,opt,name=bot_token,json=botToken,proto3" json:"bot_token,omitempty"`
	ChannelLink   string                 `protobuf:"bytes,2,opt,name=channel_link,json=channelLink,proto3" json:"channel_link,omitempty"`
	UserId        int64                  `protobuf:"varint,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CheckUserRequest) Reset() {
	*x = CheckUserRequest{}
	mi := &file_proto_telegram_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CheckUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckUserRequest) ProtoMessage() {}

func (x *CheckUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_telegram_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckUserRequest.ProtoReflect.Descriptor instead.
func (*CheckUserRequest) Descriptor() ([]byte, []int) {
	return file_proto_telegram_proto_rawDescGZIP(), []int{0}
}

func (x *CheckUserRequest) GetBotToken() string {
	if x != nil {
		return x.BotToken
	}
	return ""
}

func (x *CheckUserRequest) GetChannelLink() string {
	if x != nil {
		return x.ChannelLink
	}
	return ""
}

func (x *CheckUserRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type CheckUserResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	IsInGroup     bool                   `protobuf:"varint,1,opt,name=is_in_group,json=isInGroup,proto3" json:"is_in_group,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CheckUserResponse) Reset() {
	*x = CheckUserResponse{}
	mi := &file_proto_telegram_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CheckUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckUserResponse) ProtoMessage() {}

func (x *CheckUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_telegram_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckUserResponse.ProtoReflect.Descriptor instead.
func (*CheckUserResponse) Descriptor() ([]byte, []int) {
	return file_proto_telegram_proto_rawDescGZIP(), []int{1}
}

func (x *CheckUserResponse) GetIsInGroup() bool {
	if x != nil {
		return x.IsInGroup
	}
	return false
}

func (x *CheckUserResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_proto_telegram_proto protoreflect.FileDescriptor

const file_proto_telegram_proto_rawDesc = "" +
	"\n" +
	"\x14proto/telegram.proto\x12\btelegram\"k\n" +
	"\x10CheckUserRequest\x12\x1b\n" +
	"\tbot_token\x18\x01 \x01(\tR\bbotToken\x12!\n" +
	"\fchannel_link\x18\x02 \x01(\tR\vchannelLink\x12\x17\n" +
	"\auser_id\x18\x03 \x01(\x03R\x06userId\"M\n" +
	"\x11CheckUserResponse\x12\x1e\n" +
	"\vis_in_group\x18\x01 \x01(\bR\tisInGroup\x12\x18\n" +
	"\amessage\x18\x02 \x01(\tR\amessage2W\n" +
	"\x0fTelegramService\x12D\n" +
	"\tCheckUser\x12\x1a.telegram.CheckUserRequest\x1a\x1b.telegram.CheckUserResponseB1Z/github.com/your-repo/grpc-service/server;serverb\x06proto3"

var (
	file_proto_telegram_proto_rawDescOnce sync.Once
	file_proto_telegram_proto_rawDescData []byte
)

func file_proto_telegram_proto_rawDescGZIP() []byte {
	file_proto_telegram_proto_rawDescOnce.Do(func() {
		file_proto_telegram_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_telegram_proto_rawDesc), len(file_proto_telegram_proto_rawDesc)))
	})
	return file_proto_telegram_proto_rawDescData
}

var file_proto_telegram_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_telegram_proto_goTypes = []any{
	(*CheckUserRequest)(nil),  // 0: telegram.CheckUserRequest
	(*CheckUserResponse)(nil), // 1: telegram.CheckUserResponse
}
var file_proto_telegram_proto_depIdxs = []int32{
	0, // 0: telegram.TelegramService.CheckUser:input_type -> telegram.CheckUserRequest
	1, // 1: telegram.TelegramService.CheckUser:output_type -> telegram.CheckUserResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_telegram_proto_init() }
func file_proto_telegram_proto_init() {
	if File_proto_telegram_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_telegram_proto_rawDesc), len(file_proto_telegram_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_telegram_proto_goTypes,
		DependencyIndexes: file_proto_telegram_proto_depIdxs,
		MessageInfos:      file_proto_telegram_proto_msgTypes,
	}.Build()
	File_proto_telegram_proto = out.File
	file_proto_telegram_proto_goTypes = nil
	file_proto_telegram_proto_depIdxs = nil
}
