// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: chat/v1/channels.proto

package v1

import (
	proto "github.com/golang/protobuf/proto"
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

// Channel Kinds:
//
// Channel kinds specified in an official Harmony protocol will start with a
// "h." prefix. Third-party extensions should not use the "h." prefix. If no
// kind is specified, the channel is a text channel.
//
// Kinds indicate additional functionality a channel may have: for example,
// h.voice can indicate that a channel has voice functionalities alongside
// the usual text fare.
//
type CreateChannelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GuildId     uint64 `protobuf:"varint,1,opt,name=guild_id,json=guildId,proto3" json:"guild_id,omitempty"`
	ChannelName string `protobuf:"bytes,2,opt,name=channel_name,json=channelName,proto3" json:"channel_name,omitempty"` //  [(validate.rules).string.min_len = 1]
	IsCategory  bool   `protobuf:"varint,3,opt,name=is_category,json=isCategory,proto3" json:"is_category,omitempty"`
	PreviousId  uint64 `protobuf:"varint,5,opt,name=previous_id,json=previousId,proto3" json:"previous_id,omitempty"`
	NextId      uint64 `protobuf:"varint,4,opt,name=next_id,json=nextId,proto3" json:"next_id,omitempty"`
	ChannelKind string `protobuf:"bytes,6,opt,name=channel_kind,json=channelKind,proto3" json:"channel_kind,omitempty"`
}

func (x *CreateChannelRequest) Reset() {
	*x = CreateChannelRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_v1_channels_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateChannelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateChannelRequest) ProtoMessage() {}

func (x *CreateChannelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chat_v1_channels_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateChannelRequest.ProtoReflect.Descriptor instead.
func (*CreateChannelRequest) Descriptor() ([]byte, []int) {
	return file_chat_v1_channels_proto_rawDescGZIP(), []int{0}
}

func (x *CreateChannelRequest) GetGuildId() uint64 {
	if x != nil {
		return x.GuildId
	}
	return 0
}

func (x *CreateChannelRequest) GetChannelName() string {
	if x != nil {
		return x.ChannelName
	}
	return ""
}

func (x *CreateChannelRequest) GetIsCategory() bool {
	if x != nil {
		return x.IsCategory
	}
	return false
}

func (x *CreateChannelRequest) GetPreviousId() uint64 {
	if x != nil {
		return x.PreviousId
	}
	return 0
}

func (x *CreateChannelRequest) GetNextId() uint64 {
	if x != nil {
		return x.NextId
	}
	return 0
}

func (x *CreateChannelRequest) GetChannelKind() string {
	if x != nil {
		return x.ChannelKind
	}
	return ""
}

type CreateChannelResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChannelId uint64 `protobuf:"varint,1,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
}

func (x *CreateChannelResponse) Reset() {
	*x = CreateChannelResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_v1_channels_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateChannelResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateChannelResponse) ProtoMessage() {}

func (x *CreateChannelResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chat_v1_channels_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateChannelResponse.ProtoReflect.Descriptor instead.
func (*CreateChannelResponse) Descriptor() ([]byte, []int) {
	return file_chat_v1_channels_proto_rawDescGZIP(), []int{1}
}

func (x *CreateChannelResponse) GetChannelId() uint64 {
	if x != nil {
		return x.ChannelId
	}
	return 0
}

type GetGuildChannelsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GuildId uint64 `protobuf:"varint,1,opt,name=guild_id,json=guildId,proto3" json:"guild_id,omitempty"`
}

func (x *GetGuildChannelsRequest) Reset() {
	*x = GetGuildChannelsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_v1_channels_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGuildChannelsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGuildChannelsRequest) ProtoMessage() {}

func (x *GetGuildChannelsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chat_v1_channels_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGuildChannelsRequest.ProtoReflect.Descriptor instead.
func (*GetGuildChannelsRequest) Descriptor() ([]byte, []int) {
	return file_chat_v1_channels_proto_rawDescGZIP(), []int{2}
}

func (x *GetGuildChannelsRequest) GetGuildId() uint64 {
	if x != nil {
		return x.GuildId
	}
	return 0
}

type GetGuildChannelsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Channels []*GetGuildChannelsResponse_Channel `protobuf:"bytes,1,rep,name=channels,proto3" json:"channels,omitempty"`
}

func (x *GetGuildChannelsResponse) Reset() {
	*x = GetGuildChannelsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_v1_channels_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGuildChannelsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGuildChannelsResponse) ProtoMessage() {}

func (x *GetGuildChannelsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chat_v1_channels_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGuildChannelsResponse.ProtoReflect.Descriptor instead.
func (*GetGuildChannelsResponse) Descriptor() ([]byte, []int) {
	return file_chat_v1_channels_proto_rawDescGZIP(), []int{3}
}

func (x *GetGuildChannelsResponse) GetChannels() []*GetGuildChannelsResponse_Channel {
	if x != nil {
		return x.Channels
	}
	return nil
}

type UpdateChannelNameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GuildId        uint64 `protobuf:"varint,1,opt,name=guild_id,json=guildId,proto3" json:"guild_id,omitempty"`
	ChannelId      uint64 `protobuf:"varint,2,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	NewChannelName string `protobuf:"bytes,3,opt,name=new_channel_name,json=newChannelName,proto3" json:"new_channel_name,omitempty"`
}

func (x *UpdateChannelNameRequest) Reset() {
	*x = UpdateChannelNameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_v1_channels_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateChannelNameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateChannelNameRequest) ProtoMessage() {}

func (x *UpdateChannelNameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chat_v1_channels_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateChannelNameRequest.ProtoReflect.Descriptor instead.
func (*UpdateChannelNameRequest) Descriptor() ([]byte, []int) {
	return file_chat_v1_channels_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateChannelNameRequest) GetGuildId() uint64 {
	if x != nil {
		return x.GuildId
	}
	return 0
}

func (x *UpdateChannelNameRequest) GetChannelId() uint64 {
	if x != nil {
		return x.ChannelId
	}
	return 0
}

func (x *UpdateChannelNameRequest) GetNewChannelName() string {
	if x != nil {
		return x.NewChannelName
	}
	return ""
}

type UpdateChannelOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GuildId    uint64 `protobuf:"varint,1,opt,name=guild_id,json=guildId,proto3" json:"guild_id,omitempty"`
	ChannelId  uint64 `protobuf:"varint,2,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	PreviousId uint64 `protobuf:"varint,3,opt,name=previous_id,json=previousId,proto3" json:"previous_id,omitempty"`
	NextId     uint64 `protobuf:"varint,4,opt,name=next_id,json=nextId,proto3" json:"next_id,omitempty"`
}

func (x *UpdateChannelOrderRequest) Reset() {
	*x = UpdateChannelOrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_v1_channels_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateChannelOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateChannelOrderRequest) ProtoMessage() {}

func (x *UpdateChannelOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chat_v1_channels_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateChannelOrderRequest.ProtoReflect.Descriptor instead.
func (*UpdateChannelOrderRequest) Descriptor() ([]byte, []int) {
	return file_chat_v1_channels_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateChannelOrderRequest) GetGuildId() uint64 {
	if x != nil {
		return x.GuildId
	}
	return 0
}

func (x *UpdateChannelOrderRequest) GetChannelId() uint64 {
	if x != nil {
		return x.ChannelId
	}
	return 0
}

func (x *UpdateChannelOrderRequest) GetPreviousId() uint64 {
	if x != nil {
		return x.PreviousId
	}
	return 0
}

func (x *UpdateChannelOrderRequest) GetNextId() uint64 {
	if x != nil {
		return x.NextId
	}
	return 0
}

type DeleteChannelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GuildId   uint64 `protobuf:"varint,1,opt,name=guild_id,json=guildId,proto3" json:"guild_id,omitempty"`
	ChannelId uint64 `protobuf:"varint,2,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
}

func (x *DeleteChannelRequest) Reset() {
	*x = DeleteChannelRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_v1_channels_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteChannelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteChannelRequest) ProtoMessage() {}

func (x *DeleteChannelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chat_v1_channels_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteChannelRequest.ProtoReflect.Descriptor instead.
func (*DeleteChannelRequest) Descriptor() ([]byte, []int) {
	return file_chat_v1_channels_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteChannelRequest) GetGuildId() uint64 {
	if x != nil {
		return x.GuildId
	}
	return 0
}

func (x *DeleteChannelRequest) GetChannelId() uint64 {
	if x != nil {
		return x.ChannelId
	}
	return 0
}

type GetGuildChannelsResponse_Channel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChannelId   uint64 `protobuf:"varint,1,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	ChannelName string `protobuf:"bytes,2,opt,name=channel_name,json=channelName,proto3" json:"channel_name,omitempty"`
	IsCategory  bool   `protobuf:"varint,3,opt,name=is_category,json=isCategory,proto3" json:"is_category,omitempty"`
	Kind        string `protobuf:"bytes,4,opt,name=kind,proto3" json:"kind,omitempty"`
}

func (x *GetGuildChannelsResponse_Channel) Reset() {
	*x = GetGuildChannelsResponse_Channel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_v1_channels_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGuildChannelsResponse_Channel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGuildChannelsResponse_Channel) ProtoMessage() {}

func (x *GetGuildChannelsResponse_Channel) ProtoReflect() protoreflect.Message {
	mi := &file_chat_v1_channels_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGuildChannelsResponse_Channel.ProtoReflect.Descriptor instead.
func (*GetGuildChannelsResponse_Channel) Descriptor() ([]byte, []int) {
	return file_chat_v1_channels_proto_rawDescGZIP(), []int{3, 0}
}

func (x *GetGuildChannelsResponse_Channel) GetChannelId() uint64 {
	if x != nil {
		return x.ChannelId
	}
	return 0
}

func (x *GetGuildChannelsResponse_Channel) GetChannelName() string {
	if x != nil {
		return x.ChannelName
	}
	return ""
}

func (x *GetGuildChannelsResponse_Channel) GetIsCategory() bool {
	if x != nil {
		return x.IsCategory
	}
	return false
}

func (x *GetGuildChannelsResponse_Channel) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

var File_chat_v1_channels_proto protoreflect.FileDescriptor

var file_chat_v1_channels_proto_rawDesc = []byte{
	0x0a, 0x16, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x76, 0x31, 0x22, 0xde, 0x01, 0x0a, 0x14, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x08, 0x67, 0x75, 0x69, 0x6c, 0x64, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x02, 0x30, 0x01, 0x52, 0x07, 0x67, 0x75, 0x69, 0x6c, 0x64,
	0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x73, 0x5f, 0x63, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x23, 0x0a, 0x0b, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f,
	0x75, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x42, 0x02, 0x30, 0x01, 0x52,
	0x0a, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x07, 0x6e,
	0x65, 0x78, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x42, 0x02, 0x30, 0x01,
	0x52, 0x06, 0x6e, 0x65, 0x78, 0x74, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x5f, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x4b, 0x69, 0x6e, 0x64, 0x22, 0x3a, 0x0a, 0x15, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x0a, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x02, 0x30, 0x01, 0x52, 0x09, 0x63, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x22, 0x38, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x47, 0x75,
	0x69, 0x6c, 0x64, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1d, 0x0a, 0x08, 0x67, 0x75, 0x69, 0x6c, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x42, 0x02, 0x30, 0x01, 0x52, 0x07, 0x67, 0x75, 0x69, 0x6c, 0x64, 0x49,
	0x64, 0x22, 0xf1, 0x01, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x47, 0x75, 0x69, 0x6c, 0x64, 0x43, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e,
	0x0a, 0x08, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x32, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x63, 0x68, 0x61, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x75, 0x69, 0x6c, 0x64, 0x43, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x43, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x08, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x73, 0x1a, 0x84,
	0x01, 0x0a, 0x07, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x21, 0x0a, 0x0a, 0x63, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x02,
	0x30, 0x01, 0x52, 0x09, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x21, 0x0a,
	0x0c, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x73, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6b, 0x69, 0x6e, 0x64, 0x22, 0x86, 0x01, 0x0a, 0x18, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1d, 0x0a, 0x08, 0x67, 0x75, 0x69, 0x6c, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x42, 0x02, 0x30, 0x01, 0x52, 0x07, 0x67, 0x75, 0x69, 0x6c, 0x64, 0x49,
	0x64, 0x12, 0x21, 0x0a, 0x0a, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x04, 0x42, 0x02, 0x30, 0x01, 0x52, 0x09, 0x63, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x10, 0x6e, 0x65, 0x77, 0x5f, 0x63, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e,
	0x6e, 0x65, 0x77, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x9f,
	0x01, 0x0a, 0x19, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x08,
	0x67, 0x75, 0x69, 0x6c, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x02,
	0x30, 0x01, 0x52, 0x07, 0x67, 0x75, 0x69, 0x6c, 0x64, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0a, 0x63,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x42,
	0x02, 0x30, 0x01, 0x52, 0x09, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x23,
	0x0a, 0x0b, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x04, 0x42, 0x02, 0x30, 0x01, 0x52, 0x0a, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75,
	0x73, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x07, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x04, 0x42, 0x02, 0x30, 0x01, 0x52, 0x06, 0x6e, 0x65, 0x78, 0x74, 0x49, 0x64,
	0x22, 0x58, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x08, 0x67, 0x75, 0x69, 0x6c,
	0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x02, 0x30, 0x01, 0x52, 0x07,
	0x67, 0x75, 0x69, 0x6c, 0x64, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0a, 0x63, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x42, 0x02, 0x30, 0x01, 0x52,
	0x09, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x61, 0x72, 0x6d, 0x6f, 0x6e, 0x79,
	0x2d, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x6c, 0x65, 0x67,
	0x61, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x76, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chat_v1_channels_proto_rawDescOnce sync.Once
	file_chat_v1_channels_proto_rawDescData = file_chat_v1_channels_proto_rawDesc
)

func file_chat_v1_channels_proto_rawDescGZIP() []byte {
	file_chat_v1_channels_proto_rawDescOnce.Do(func() {
		file_chat_v1_channels_proto_rawDescData = protoimpl.X.CompressGZIP(file_chat_v1_channels_proto_rawDescData)
	})
	return file_chat_v1_channels_proto_rawDescData
}

var file_chat_v1_channels_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_chat_v1_channels_proto_goTypes = []interface{}{
	(*CreateChannelRequest)(nil),             // 0: protocol.chat.v1.CreateChannelRequest
	(*CreateChannelResponse)(nil),            // 1: protocol.chat.v1.CreateChannelResponse
	(*GetGuildChannelsRequest)(nil),          // 2: protocol.chat.v1.GetGuildChannelsRequest
	(*GetGuildChannelsResponse)(nil),         // 3: protocol.chat.v1.GetGuildChannelsResponse
	(*UpdateChannelNameRequest)(nil),         // 4: protocol.chat.v1.UpdateChannelNameRequest
	(*UpdateChannelOrderRequest)(nil),        // 5: protocol.chat.v1.UpdateChannelOrderRequest
	(*DeleteChannelRequest)(nil),             // 6: protocol.chat.v1.DeleteChannelRequest
	(*GetGuildChannelsResponse_Channel)(nil), // 7: protocol.chat.v1.GetGuildChannelsResponse.Channel
}
var file_chat_v1_channels_proto_depIdxs = []int32{
	7, // 0: protocol.chat.v1.GetGuildChannelsResponse.channels:type_name -> protocol.chat.v1.GetGuildChannelsResponse.Channel
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_chat_v1_channels_proto_init() }
func file_chat_v1_channels_proto_init() {
	if File_chat_v1_channels_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_chat_v1_channels_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateChannelRequest); i {
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
		file_chat_v1_channels_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateChannelResponse); i {
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
		file_chat_v1_channels_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGuildChannelsRequest); i {
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
		file_chat_v1_channels_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGuildChannelsResponse); i {
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
		file_chat_v1_channels_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateChannelNameRequest); i {
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
		file_chat_v1_channels_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateChannelOrderRequest); i {
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
		file_chat_v1_channels_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteChannelRequest); i {
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
		file_chat_v1_channels_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGuildChannelsResponse_Channel); i {
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
			RawDescriptor: file_chat_v1_channels_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_chat_v1_channels_proto_goTypes,
		DependencyIndexes: file_chat_v1_channels_proto_depIdxs,
		MessageInfos:      file_chat_v1_channels_proto_msgTypes,
	}.Build()
	File_chat_v1_channels_proto = out.File
	file_chat_v1_channels_proto_rawDesc = nil
	file_chat_v1_channels_proto_goTypes = nil
	file_chat_v1_channels_proto_depIdxs = nil
}
