// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.12.3
// source: profile/v1/profile.proto

package profilev1

import (
	context "context"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type UserStatus int32

const (
	UserStatus_USER_STATUS_ONLINE_UNSPECIFIED UserStatus = 0
	UserStatus_USER_STATUS_STREAMING          UserStatus = 1
	UserStatus_USER_STATUS_DO_NOT_DISTURB     UserStatus = 2
	UserStatus_USER_STATUS_IDLE               UserStatus = 3
	UserStatus_USER_STATUS_OFFLINE            UserStatus = 4
)

// Enum value maps for UserStatus.
var (
	UserStatus_name = map[int32]string{
		0: "USER_STATUS_ONLINE_UNSPECIFIED",
		1: "USER_STATUS_STREAMING",
		2: "USER_STATUS_DO_NOT_DISTURB",
		3: "USER_STATUS_IDLE",
		4: "USER_STATUS_OFFLINE",
	}
	UserStatus_value = map[string]int32{
		"USER_STATUS_ONLINE_UNSPECIFIED": 0,
		"USER_STATUS_STREAMING":          1,
		"USER_STATUS_DO_NOT_DISTURB":     2,
		"USER_STATUS_IDLE":               3,
		"USER_STATUS_OFFLINE":            4,
	}
)

func (x UserStatus) Enum() *UserStatus {
	p := new(UserStatus)
	*p = x
	return p
}

func (x UserStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_profile_v1_profile_proto_enumTypes[0].Descriptor()
}

func (UserStatus) Type() protoreflect.EnumType {
	return &file_profile_v1_profile_proto_enumTypes[0]
}

func (x UserStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UserStatus.Descriptor instead.
func (UserStatus) EnumDescriptor() ([]byte, []int) {
	return file_profile_v1_profile_proto_rawDescGZIP(), []int{0}
}

type GetUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId uint64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetUserRequest) Reset() {
	*x = GetUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profile_v1_profile_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserRequest) ProtoMessage() {}

func (x *GetUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_profile_v1_profile_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserRequest.ProtoReflect.Descriptor instead.
func (*GetUserRequest) Descriptor() ([]byte, []int) {
	return file_profile_v1_profile_proto_rawDescGZIP(), []int{0}
}

func (x *GetUserRequest) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type GetUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserName   string     `protobuf:"bytes,1,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	UserAvatar string     `protobuf:"bytes,2,opt,name=user_avatar,json=userAvatar,proto3" json:"user_avatar,omitempty"`
	UserStatus UserStatus `protobuf:"varint,3,opt,name=user_status,json=userStatus,proto3,enum=protocol.profile.v1.UserStatus" json:"user_status,omitempty"`
}

func (x *GetUserResponse) Reset() {
	*x = GetUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profile_v1_profile_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserResponse) ProtoMessage() {}

func (x *GetUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_profile_v1_profile_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserResponse.ProtoReflect.Descriptor instead.
func (*GetUserResponse) Descriptor() ([]byte, []int) {
	return file_profile_v1_profile_proto_rawDescGZIP(), []int{1}
}

func (x *GetUserResponse) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *GetUserResponse) GetUserAvatar() string {
	if x != nil {
		return x.UserAvatar
	}
	return ""
}

func (x *GetUserResponse) GetUserStatus() UserStatus {
	if x != nil {
		return x.UserStatus
	}
	return UserStatus_USER_STATUS_ONLINE_UNSPECIFIED
}

type GetUserMetadataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppId string `protobuf:"bytes,1,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
}

func (x *GetUserMetadataRequest) Reset() {
	*x = GetUserMetadataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profile_v1_profile_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserMetadataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserMetadataRequest) ProtoMessage() {}

func (x *GetUserMetadataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_profile_v1_profile_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserMetadataRequest.ProtoReflect.Descriptor instead.
func (*GetUserMetadataRequest) Descriptor() ([]byte, []int) {
	return file_profile_v1_profile_proto_rawDescGZIP(), []int{2}
}

func (x *GetUserMetadataRequest) GetAppId() string {
	if x != nil {
		return x.AppId
	}
	return ""
}

type GetUserMetadataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metadata string `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *GetUserMetadataResponse) Reset() {
	*x = GetUserMetadataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profile_v1_profile_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserMetadataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserMetadataResponse) ProtoMessage() {}

func (x *GetUserMetadataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_profile_v1_profile_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserMetadataResponse.ProtoReflect.Descriptor instead.
func (*GetUserMetadataResponse) Descriptor() ([]byte, []int) {
	return file_profile_v1_profile_proto_rawDescGZIP(), []int{3}
}

func (x *GetUserMetadataResponse) GetMetadata() string {
	if x != nil {
		return x.Metadata
	}
	return ""
}

type UsernameUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserName string `protobuf:"bytes,1,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
}

func (x *UsernameUpdateRequest) Reset() {
	*x = UsernameUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profile_v1_profile_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UsernameUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UsernameUpdateRequest) ProtoMessage() {}

func (x *UsernameUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_profile_v1_profile_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UsernameUpdateRequest.ProtoReflect.Descriptor instead.
func (*UsernameUpdateRequest) Descriptor() ([]byte, []int) {
	return file_profile_v1_profile_proto_rawDescGZIP(), []int{4}
}

func (x *UsernameUpdateRequest) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

type UsernameUpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *UsernameUpdateResponse) Reset() {
	*x = UsernameUpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profile_v1_profile_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UsernameUpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UsernameUpdateResponse) ProtoMessage() {}

func (x *UsernameUpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_profile_v1_profile_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UsernameUpdateResponse.ProtoReflect.Descriptor instead.
func (*UsernameUpdateResponse) Descriptor() ([]byte, []int) {
	return file_profile_v1_profile_proto_rawDescGZIP(), []int{5}
}

func (x *UsernameUpdateResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type StatusUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NewStatus UserStatus `protobuf:"varint,1,opt,name=new_status,json=newStatus,proto3,enum=protocol.profile.v1.UserStatus" json:"new_status,omitempty"`
}

func (x *StatusUpdateRequest) Reset() {
	*x = StatusUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profile_v1_profile_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusUpdateRequest) ProtoMessage() {}

func (x *StatusUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_profile_v1_profile_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusUpdateRequest.ProtoReflect.Descriptor instead.
func (*StatusUpdateRequest) Descriptor() ([]byte, []int) {
	return file_profile_v1_profile_proto_rawDescGZIP(), []int{6}
}

func (x *StatusUpdateRequest) GetNewStatus() UserStatus {
	if x != nil {
		return x.NewStatus
	}
	return UserStatus_USER_STATUS_ONLINE_UNSPECIFIED
}

type StatusUpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *StatusUpdateResponse) Reset() {
	*x = StatusUpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profile_v1_profile_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusUpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusUpdateResponse) ProtoMessage() {}

func (x *StatusUpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_profile_v1_profile_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusUpdateResponse.ProtoReflect.Descriptor instead.
func (*StatusUpdateResponse) Descriptor() ([]byte, []int) {
	return file_profile_v1_profile_proto_rawDescGZIP(), []int{7}
}

func (x *StatusUpdateResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_profile_v1_profile_proto protoreflect.FileDescriptor

var file_profile_v1_profile_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x1a,
	0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x32, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x07, 0xfa, 0x42, 0x04,
	0x32, 0x02, 0x20, 0x00, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x91, 0x01, 0x0a,
	0x0f, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a,
	0x0b, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x40,
	0x0a, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x70,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x22, 0x38, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x06, 0x61, 0x70,
	0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72,
	0x02, 0x10, 0x01, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x22, 0x35, 0x0a, 0x17, 0x47, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x22, 0x3d, 0x0a, 0x15, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x09, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa,
	0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65,
	0x22, 0x32, 0x0a, 0x16, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x22, 0x5f, 0x0a, 0x13, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x48, 0x0a, 0x0a, 0x6e,
	0x65, 0x77, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x42, 0x08, 0xfa, 0x42, 0x05, 0x82, 0x01, 0x02, 0x10, 0x01, 0x52, 0x09, 0x6e, 0x65, 0x77, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x30, 0x0a, 0x14, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2a, 0x9a, 0x01, 0x0a, 0x0a, 0x55, 0x73, 0x65, 0x72,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x22, 0x0a, 0x1e, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x53,
	0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x4f, 0x4e, 0x4c, 0x49, 0x4e, 0x45, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x19, 0x0a, 0x15, 0x55, 0x53,
	0x45, 0x52, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x53, 0x54, 0x52, 0x45, 0x41, 0x4d,
	0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x1e, 0x0a, 0x1a, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x53, 0x54,
	0x41, 0x54, 0x55, 0x53, 0x5f, 0x44, 0x4f, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x44, 0x49, 0x53, 0x54,
	0x55, 0x52, 0x42, 0x10, 0x02, 0x12, 0x14, 0x0a, 0x10, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x53, 0x54,
	0x41, 0x54, 0x55, 0x53, 0x5f, 0x49, 0x44, 0x4c, 0x45, 0x10, 0x03, 0x12, 0x17, 0x0a, 0x13, 0x55,
	0x53, 0x45, 0x52, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x4f, 0x46, 0x46, 0x4c, 0x49,
	0x4e, 0x45, 0x10, 0x04, 0x32, 0xac, 0x03, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x56, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x12, 0x23, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x70, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x6e, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x2b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x70, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x2c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x6b, 0x0a, 0x0e, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x12, 0x2a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x65, 0x0a, 0x0c,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x28, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x13, 0x5a, 0x11, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x3b, 0x70,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_profile_v1_profile_proto_rawDescOnce sync.Once
	file_profile_v1_profile_proto_rawDescData = file_profile_v1_profile_proto_rawDesc
)

func file_profile_v1_profile_proto_rawDescGZIP() []byte {
	file_profile_v1_profile_proto_rawDescOnce.Do(func() {
		file_profile_v1_profile_proto_rawDescData = protoimpl.X.CompressGZIP(file_profile_v1_profile_proto_rawDescData)
	})
	return file_profile_v1_profile_proto_rawDescData
}

var file_profile_v1_profile_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_profile_v1_profile_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_profile_v1_profile_proto_goTypes = []interface{}{
	(UserStatus)(0),                 // 0: protocol.profile.v1.UserStatus
	(*GetUserRequest)(nil),          // 1: protocol.profile.v1.GetUserRequest
	(*GetUserResponse)(nil),         // 2: protocol.profile.v1.GetUserResponse
	(*GetUserMetadataRequest)(nil),  // 3: protocol.profile.v1.GetUserMetadataRequest
	(*GetUserMetadataResponse)(nil), // 4: protocol.profile.v1.GetUserMetadataResponse
	(*UsernameUpdateRequest)(nil),   // 5: protocol.profile.v1.UsernameUpdateRequest
	(*UsernameUpdateResponse)(nil),  // 6: protocol.profile.v1.UsernameUpdateResponse
	(*StatusUpdateRequest)(nil),     // 7: protocol.profile.v1.StatusUpdateRequest
	(*StatusUpdateResponse)(nil),    // 8: protocol.profile.v1.StatusUpdateResponse
}
var file_profile_v1_profile_proto_depIdxs = []int32{
	0, // 0: protocol.profile.v1.GetUserResponse.user_status:type_name -> protocol.profile.v1.UserStatus
	0, // 1: protocol.profile.v1.StatusUpdateRequest.new_status:type_name -> protocol.profile.v1.UserStatus
	1, // 2: protocol.profile.v1.ProfileService.GetUser:input_type -> protocol.profile.v1.GetUserRequest
	3, // 3: protocol.profile.v1.ProfileService.GetUserMetadata:input_type -> protocol.profile.v1.GetUserMetadataRequest
	5, // 4: protocol.profile.v1.ProfileService.UsernameUpdate:input_type -> protocol.profile.v1.UsernameUpdateRequest
	7, // 5: protocol.profile.v1.ProfileService.StatusUpdate:input_type -> protocol.profile.v1.StatusUpdateRequest
	2, // 6: protocol.profile.v1.ProfileService.GetUser:output_type -> protocol.profile.v1.GetUserResponse
	4, // 7: protocol.profile.v1.ProfileService.GetUserMetadata:output_type -> protocol.profile.v1.GetUserMetadataResponse
	6, // 8: protocol.profile.v1.ProfileService.UsernameUpdate:output_type -> protocol.profile.v1.UsernameUpdateResponse
	8, // 9: protocol.profile.v1.ProfileService.StatusUpdate:output_type -> protocol.profile.v1.StatusUpdateResponse
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_profile_v1_profile_proto_init() }
func file_profile_v1_profile_proto_init() {
	if File_profile_v1_profile_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_profile_v1_profile_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserRequest); i {
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
		file_profile_v1_profile_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserResponse); i {
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
		file_profile_v1_profile_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserMetadataRequest); i {
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
		file_profile_v1_profile_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserMetadataResponse); i {
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
		file_profile_v1_profile_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UsernameUpdateRequest); i {
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
		file_profile_v1_profile_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UsernameUpdateResponse); i {
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
		file_profile_v1_profile_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusUpdateRequest); i {
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
		file_profile_v1_profile_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusUpdateResponse); i {
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
			RawDescriptor: file_profile_v1_profile_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_profile_v1_profile_proto_goTypes,
		DependencyIndexes: file_profile_v1_profile_proto_depIdxs,
		EnumInfos:         file_profile_v1_profile_proto_enumTypes,
		MessageInfos:      file_profile_v1_profile_proto_msgTypes,
	}.Build()
	File_profile_v1_profile_proto = out.File
	file_profile_v1_profile_proto_rawDesc = nil
	file_profile_v1_profile_proto_goTypes = nil
	file_profile_v1_profile_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ProfileServiceClient is the client API for ProfileService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProfileServiceClient interface {
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error)
	GetUserMetadata(ctx context.Context, in *GetUserMetadataRequest, opts ...grpc.CallOption) (*GetUserMetadataResponse, error)
	UsernameUpdate(ctx context.Context, in *UsernameUpdateRequest, opts ...grpc.CallOption) (*UsernameUpdateResponse, error)
	StatusUpdate(ctx context.Context, in *StatusUpdateRequest, opts ...grpc.CallOption) (*StatusUpdateResponse, error)
}

type profileServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProfileServiceClient(cc grpc.ClientConnInterface) ProfileServiceClient {
	return &profileServiceClient{cc}
}

func (c *profileServiceClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error) {
	out := new(GetUserResponse)
	err := c.cc.Invoke(ctx, "/protocol.profile.v1.ProfileService/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileServiceClient) GetUserMetadata(ctx context.Context, in *GetUserMetadataRequest, opts ...grpc.CallOption) (*GetUserMetadataResponse, error) {
	out := new(GetUserMetadataResponse)
	err := c.cc.Invoke(ctx, "/protocol.profile.v1.ProfileService/GetUserMetadata", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileServiceClient) UsernameUpdate(ctx context.Context, in *UsernameUpdateRequest, opts ...grpc.CallOption) (*UsernameUpdateResponse, error) {
	out := new(UsernameUpdateResponse)
	err := c.cc.Invoke(ctx, "/protocol.profile.v1.ProfileService/UsernameUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileServiceClient) StatusUpdate(ctx context.Context, in *StatusUpdateRequest, opts ...grpc.CallOption) (*StatusUpdateResponse, error) {
	out := new(StatusUpdateResponse)
	err := c.cc.Invoke(ctx, "/protocol.profile.v1.ProfileService/StatusUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProfileServiceServer is the server API for ProfileService service.
type ProfileServiceServer interface {
	GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error)
	GetUserMetadata(context.Context, *GetUserMetadataRequest) (*GetUserMetadataResponse, error)
	UsernameUpdate(context.Context, *UsernameUpdateRequest) (*UsernameUpdateResponse, error)
	StatusUpdate(context.Context, *StatusUpdateRequest) (*StatusUpdateResponse, error)
}

// UnimplementedProfileServiceServer can be embedded to have forward compatible implementations.
type UnimplementedProfileServiceServer struct {
}

func (*UnimplementedProfileServiceServer) GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (*UnimplementedProfileServiceServer) GetUserMetadata(context.Context, *GetUserMetadataRequest) (*GetUserMetadataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserMetadata not implemented")
}
func (*UnimplementedProfileServiceServer) UsernameUpdate(context.Context, *UsernameUpdateRequest) (*UsernameUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UsernameUpdate not implemented")
}
func (*UnimplementedProfileServiceServer) StatusUpdate(context.Context, *StatusUpdateRequest) (*StatusUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StatusUpdate not implemented")
}

func RegisterProfileServiceServer(s *grpc.Server, srv ProfileServiceServer) {
	s.RegisterService(&_ProfileService_serviceDesc, srv)
}

func _ProfileService_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServiceServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.profile.v1.ProfileService/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServiceServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfileService_GetUserMetadata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserMetadataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServiceServer).GetUserMetadata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.profile.v1.ProfileService/GetUserMetadata",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServiceServer).GetUserMetadata(ctx, req.(*GetUserMetadataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfileService_UsernameUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UsernameUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServiceServer).UsernameUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.profile.v1.ProfileService/UsernameUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServiceServer).UsernameUpdate(ctx, req.(*UsernameUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfileService_StatusUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatusUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServiceServer).StatusUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.profile.v1.ProfileService/StatusUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServiceServer).StatusUpdate(ctx, req.(*StatusUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProfileService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protocol.profile.v1.ProfileService",
	HandlerType: (*ProfileServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUser",
			Handler:    _ProfileService_GetUser_Handler,
		},
		{
			MethodName: "GetUserMetadata",
			Handler:    _ProfileService_GetUserMetadata_Handler,
		},
		{
			MethodName: "UsernameUpdate",
			Handler:    _ProfileService_UsernameUpdate_Handler,
		},
		{
			MethodName: "StatusUpdate",
			Handler:    _ProfileService_StatusUpdate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "profile/v1/profile.proto",
}
