// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.17.3
// source: shop/v1/shop.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RegisterReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Mobile    string `protobuf:"bytes,3,opt,name=mobile,proto3" json:"mobile,omitempty"`
	Username  string `protobuf:"bytes,4,opt,name=username,proto3" json:"username,omitempty"`
	Token     string `protobuf:"bytes,5,opt,name=token,proto3" json:"token,omitempty"`
	ExpiredAt int64  `protobuf:"varint,6,opt,name=expiredAt,proto3" json:"expiredAt,omitempty"`
}

func (x *RegisterReply) Reset() {
	*x = RegisterReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shop_v1_shop_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterReply) ProtoMessage() {}

func (x *RegisterReply) ProtoReflect() protoreflect.Message {
	mi := &file_shop_v1_shop_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterReply.ProtoReflect.Descriptor instead.
func (*RegisterReply) Descriptor() ([]byte, []int) {
	return file_shop_v1_shop_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterReply) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *RegisterReply) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *RegisterReply) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *RegisterReply) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *RegisterReply) GetExpiredAt() int64 {
	if x != nil {
		return x.ExpiredAt
	}
	return 0
}

type RegisterReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mobile   string `protobuf:"bytes,1,opt,name=mobile,proto3" json:"mobile,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *RegisterReq) Reset() {
	*x = RegisterReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shop_v1_shop_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterReq) ProtoMessage() {}

func (x *RegisterReq) ProtoReflect() protoreflect.Message {
	mi := &file_shop_v1_shop_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterReq.ProtoReflect.Descriptor instead.
func (*RegisterReq) Descriptor() ([]byte, []int) {
	return file_shop_v1_shop_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterReq) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *RegisterReq) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *RegisterReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type LoginReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mobile    string `protobuf:"bytes,1,opt,name=mobile,proto3" json:"mobile,omitempty"`
	Password  string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Captcha   string `protobuf:"bytes,3,opt,name=captcha,proto3" json:"captcha,omitempty"`
	CaptchaId string `protobuf:"bytes,4,opt,name=captchaId,proto3" json:"captchaId,omitempty"`
}

func (x *LoginReq) Reset() {
	*x = LoginReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shop_v1_shop_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginReq) ProtoMessage() {}

func (x *LoginReq) ProtoReflect() protoreflect.Message {
	mi := &file_shop_v1_shop_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginReq.ProtoReflect.Descriptor instead.
func (*LoginReq) Descriptor() ([]byte, []int) {
	return file_shop_v1_shop_proto_rawDescGZIP(), []int{2}
}

func (x *LoginReq) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *LoginReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *LoginReq) GetCaptcha() string {
	if x != nil {
		return x.Captcha
	}
	return ""
}

func (x *LoginReq) GetCaptchaId() string {
	if x != nil {
		return x.CaptchaId
	}
	return ""
}

type UserDetailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Mobile   string `protobuf:"bytes,2,opt,name=mobile,proto3" json:"mobile,omitempty"`
	NickName string `protobuf:"bytes,3,opt,name=nickName,proto3" json:"nickName,omitempty"`
	Birthday int64  `protobuf:"varint,4,opt,name=birthday,proto3" json:"birthday,omitempty"`
	Gender   string `protobuf:"bytes,5,opt,name=gender,proto3" json:"gender,omitempty"`
	Role     int32  `protobuf:"varint,6,opt,name=role,proto3" json:"role,omitempty"`
}

func (x *UserDetailResponse) Reset() {
	*x = UserDetailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shop_v1_shop_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserDetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserDetailResponse) ProtoMessage() {}

func (x *UserDetailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_shop_v1_shop_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserDetailResponse.ProtoReflect.Descriptor instead.
func (*UserDetailResponse) Descriptor() ([]byte, []int) {
	return file_shop_v1_shop_proto_rawDescGZIP(), []int{3}
}

func (x *UserDetailResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserDetailResponse) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *UserDetailResponse) GetNickName() string {
	if x != nil {
		return x.NickName
	}
	return ""
}

func (x *UserDetailResponse) GetBirthday() int64 {
	if x != nil {
		return x.Birthday
	}
	return 0
}

func (x *UserDetailResponse) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *UserDetailResponse) GetRole() int32 {
	if x != nil {
		return x.Role
	}
	return 0
}

type CaptchaReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CaptchaId string `protobuf:"bytes,1,opt,name=captchaId,proto3" json:"captchaId,omitempty"`
	PicPath   string `protobuf:"bytes,2,opt,name=picPath,proto3" json:"picPath,omitempty"`
}

func (x *CaptchaReply) Reset() {
	*x = CaptchaReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shop_v1_shop_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CaptchaReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CaptchaReply) ProtoMessage() {}

func (x *CaptchaReply) ProtoReflect() protoreflect.Message {
	mi := &file_shop_v1_shop_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CaptchaReply.ProtoReflect.Descriptor instead.
func (*CaptchaReply) Descriptor() ([]byte, []int) {
	return file_shop_v1_shop_proto_rawDescGZIP(), []int{4}
}

func (x *CaptchaReply) GetCaptchaId() string {
	if x != nil {
		return x.CaptchaId
	}
	return ""
}

func (x *CaptchaReply) GetPicPath() string {
	if x != nil {
		return x.PicPath
	}
	return ""
}

var File_shop_v1_shop_proto protoreflect.FileDescriptor

var file_shop_v1_shop_proto_rawDesc = []byte{
	0x0a, 0x12, 0x73, 0x68, 0x6f, 0x70, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x2e,
	0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x87, 0x01, 0x0a, 0x0d, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69,
	0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x41, 0x74, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x41, 0x74,
	0x22, 0x7b, 0x0a, 0x0b, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12,
	0x20, 0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0x98, 0x01, 0x0b, 0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c,
	0x65, 0x12, 0x25, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10, 0x03, 0x18, 0x0f, 0x52, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72,
	0x02, 0x10, 0x08, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x9d, 0x01,
	0x0a, 0x08, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x20, 0x0a, 0x06, 0x6d, 0x6f,
	0x62, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72,
	0x03, 0x98, 0x01, 0x0b, 0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x23, 0x0a, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07,
	0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x08, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x12, 0x23, 0x0a, 0x07, 0x63, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10, 0x05, 0x18, 0x05, 0x52, 0x07, 0x63,
	0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x12, 0x25, 0x0a, 0x09, 0x63, 0x61, 0x70, 0x74, 0x63, 0x68,
	0x61, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02,
	0x10, 0x01, 0x52, 0x09, 0x63, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x49, 0x64, 0x22, 0xa0, 0x01,
	0x0a, 0x12, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x6e, 0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6e, 0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x69, 0x72, 0x74,
	0x68, 0x64, 0x61, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x62, 0x69, 0x72, 0x74,
	0x68, 0x64, 0x61, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04,
	0x72, 0x6f, 0x6c, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65,
	0x22, 0x46, 0x0a, 0x0c, 0x43, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x1c, 0x0a, 0x09, 0x63, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x49, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x70, 0x69, 0x63, 0x50, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x70, 0x69, 0x63, 0x50, 0x61, 0x74, 0x68, 0x32, 0xff, 0x02, 0x0a, 0x04, 0x53, 0x68, 0x6f,
	0x70, 0x12, 0x62, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x19, 0x2e,
	0x73, 0x68, 0x6f, 0x70, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x1b, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x2e,
	0x73, 0x68, 0x6f, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x22, 0x13, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x3a, 0x01, 0x2a, 0x12, 0x59, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x16,
	0x2e, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x1b, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x73, 0x68,
	0x6f, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x22, 0x10, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x3a, 0x01, 0x2a,
	0x12, 0x59, 0x0a, 0x07, 0x43, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x12, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x1a, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x12, 0x12, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x73, 0x2f, 0x63, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x12, 0x5d, 0x0a, 0x06, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x20, 0x2e,
	0x73, 0x68, 0x6f, 0x70, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x12, 0x11, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x73, 0x2f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x42, 0x15, 0x5a, 0x13, 0x73, 0x68,
	0x6f, 0x70, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x76,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_shop_v1_shop_proto_rawDescOnce sync.Once
	file_shop_v1_shop_proto_rawDescData = file_shop_v1_shop_proto_rawDesc
)

func file_shop_v1_shop_proto_rawDescGZIP() []byte {
	file_shop_v1_shop_proto_rawDescOnce.Do(func() {
		file_shop_v1_shop_proto_rawDescData = protoimpl.X.CompressGZIP(file_shop_v1_shop_proto_rawDescData)
	})
	return file_shop_v1_shop_proto_rawDescData
}

var file_shop_v1_shop_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_shop_v1_shop_proto_goTypes = []interface{}{
	(*RegisterReply)(nil),      // 0: shop.shop.v1.RegisterReply
	(*RegisterReq)(nil),        // 1: shop.shop.v1.RegisterReq
	(*LoginReq)(nil),           // 2: shop.shop.v1.LoginReq
	(*UserDetailResponse)(nil), // 3: shop.shop.v1.UserDetailResponse
	(*CaptchaReply)(nil),       // 4: shop.shop.v1.CaptchaReply
	(*emptypb.Empty)(nil),      // 5: google.protobuf.Empty
}
var file_shop_v1_shop_proto_depIdxs = []int32{
	1, // 0: shop.shop.v1.Shop.Register:input_type -> shop.shop.v1.RegisterReq
	2, // 1: shop.shop.v1.Shop.Login:input_type -> shop.shop.v1.LoginReq
	5, // 2: shop.shop.v1.Shop.Captcha:input_type -> google.protobuf.Empty
	5, // 3: shop.shop.v1.Shop.Detail:input_type -> google.protobuf.Empty
	0, // 4: shop.shop.v1.Shop.Register:output_type -> shop.shop.v1.RegisterReply
	0, // 5: shop.shop.v1.Shop.Login:output_type -> shop.shop.v1.RegisterReply
	4, // 6: shop.shop.v1.Shop.Captcha:output_type -> shop.shop.v1.CaptchaReply
	3, // 7: shop.shop.v1.Shop.Detail:output_type -> shop.shop.v1.UserDetailResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_shop_v1_shop_proto_init() }
func file_shop_v1_shop_proto_init() {
	if File_shop_v1_shop_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_shop_v1_shop_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterReply); i {
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
		file_shop_v1_shop_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterReq); i {
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
		file_shop_v1_shop_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginReq); i {
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
		file_shop_v1_shop_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserDetailResponse); i {
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
		file_shop_v1_shop_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CaptchaReply); i {
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
			RawDescriptor: file_shop_v1_shop_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_shop_v1_shop_proto_goTypes,
		DependencyIndexes: file_shop_v1_shop_proto_depIdxs,
		MessageInfos:      file_shop_v1_shop_proto_msgTypes,
	}.Build()
	File_shop_v1_shop_proto = out.File
	file_shop_v1_shop_proto_rawDesc = nil
	file_shop_v1_shop_proto_goTypes = nil
	file_shop_v1_shop_proto_depIdxs = nil
}
