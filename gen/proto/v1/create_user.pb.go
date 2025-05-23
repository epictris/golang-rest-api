// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: proto/v1/create_user.proto

package apiv1

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
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

type CreateUserRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Email         string                 `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateUserRequest) Reset() {
	*x = CreateUserRequest{}
	mi := &file_proto_v1_create_user_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserRequest) ProtoMessage() {}

func (x *CreateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_create_user_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserRequest.ProtoReflect.Descriptor instead.
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return file_proto_v1_create_user_proto_rawDescGZIP(), []int{0}
}

func (x *CreateUserRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type CreateUserResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        int64                  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateUserResponse) Reset() {
	*x = CreateUserResponse{}
	mi := &file_proto_v1_create_user_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserResponse) ProtoMessage() {}

func (x *CreateUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_create_user_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserResponse.ProtoReflect.Descriptor instead.
func (*CreateUserResponse) Descriptor() ([]byte, []int) {
	return file_proto_v1_create_user_proto_rawDescGZIP(), []int{1}
}

func (x *CreateUserResponse) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

var File_proto_v1_create_user_proto protoreflect.FileDescriptor

const file_proto_v1_create_user_proto_rawDesc = "" +
	"\n" +
	"\x1aproto/v1/create_user.proto\x12\bproto.v1\x1a\x1bbuf/validate/validate.proto\"2\n" +
	"\x11CreateUserRequest\x12\x1d\n" +
	"\x05email\x18\x01 \x01(\tB\a\xbaH\x04r\x02`\x01R\x05email\"5\n" +
	"\x12CreateUserResponse\x12\x1f\n" +
	"\auser_id\x18\x01 \x01(\x03B\x06\xbaH\x03\xc8\x01\x01R\x06userIdB+Z)github.com/epictris/go/gen/proto/v1;apiv1b\x06proto3"

var (
	file_proto_v1_create_user_proto_rawDescOnce sync.Once
	file_proto_v1_create_user_proto_rawDescData []byte
)

func file_proto_v1_create_user_proto_rawDescGZIP() []byte {
	file_proto_v1_create_user_proto_rawDescOnce.Do(func() {
		file_proto_v1_create_user_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_v1_create_user_proto_rawDesc), len(file_proto_v1_create_user_proto_rawDesc)))
	})
	return file_proto_v1_create_user_proto_rawDescData
}

var file_proto_v1_create_user_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_v1_create_user_proto_goTypes = []any{
	(*CreateUserRequest)(nil),  // 0: proto.v1.CreateUserRequest
	(*CreateUserResponse)(nil), // 1: proto.v1.CreateUserResponse
}
var file_proto_v1_create_user_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_v1_create_user_proto_init() }
func file_proto_v1_create_user_proto_init() {
	if File_proto_v1_create_user_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_v1_create_user_proto_rawDesc), len(file_proto_v1_create_user_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_v1_create_user_proto_goTypes,
		DependencyIndexes: file_proto_v1_create_user_proto_depIdxs,
		MessageInfos:      file_proto_v1_create_user_proto_msgTypes,
	}.Build()
	File_proto_v1_create_user_proto = out.File
	file_proto_v1_create_user_proto_goTypes = nil
	file_proto_v1_create_user_proto_depIdxs = nil
}
