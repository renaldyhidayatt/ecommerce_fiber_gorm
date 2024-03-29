// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.12.4
// source: midtrans.proto

package pb

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

type CreateMidtransRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GrossAmount int32  `protobuf:"varint,1,opt,name=gross_amount,json=grossAmount,proto3" json:"gross_amount,omitempty"`
	FirstName   string `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName    string `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Email       string `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Phone       string `protobuf:"bytes,5,opt,name=phone,proto3" json:"phone,omitempty"`
}

func (x *CreateMidtransRequest) Reset() {
	*x = CreateMidtransRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_midtrans_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateMidtransRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMidtransRequest) ProtoMessage() {}

func (x *CreateMidtransRequest) ProtoReflect() protoreflect.Message {
	mi := &file_midtrans_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMidtransRequest.ProtoReflect.Descriptor instead.
func (*CreateMidtransRequest) Descriptor() ([]byte, []int) {
	return file_midtrans_proto_rawDescGZIP(), []int{0}
}

func (x *CreateMidtransRequest) GetGrossAmount() int32 {
	if x != nil {
		return x.GrossAmount
	}
	return 0
}

func (x *CreateMidtransRequest) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *CreateMidtransRequest) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *CreateMidtransRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CreateMidtransRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

type SnapResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token         string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	RedirectUrl   string   `protobuf:"bytes,2,opt,name=redirect_url,json=redirectUrl,proto3" json:"redirect_url,omitempty"`
	StatusCode    string   `protobuf:"bytes,3,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	ErrorMessages []string `protobuf:"bytes,4,rep,name=error_messages,json=errorMessages,proto3" json:"error_messages,omitempty"`
}

func (x *SnapResponse) Reset() {
	*x = SnapResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_midtrans_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SnapResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SnapResponse) ProtoMessage() {}

func (x *SnapResponse) ProtoReflect() protoreflect.Message {
	mi := &file_midtrans_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SnapResponse.ProtoReflect.Descriptor instead.
func (*SnapResponse) Descriptor() ([]byte, []int) {
	return file_midtrans_proto_rawDescGZIP(), []int{1}
}

func (x *SnapResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *SnapResponse) GetRedirectUrl() string {
	if x != nil {
		return x.RedirectUrl
	}
	return ""
}

func (x *SnapResponse) GetStatusCode() string {
	if x != nil {
		return x.StatusCode
	}
	return ""
}

func (x *SnapResponse) GetErrorMessages() []string {
	if x != nil {
		return x.ErrorMessages
	}
	return nil
}

var File_midtrans_proto protoreflect.FileDescriptor

var file_midtrans_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x6d, 0x69, 0x64, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x02, 0x70, 0x62, 0x22, 0xa2, 0x01, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d,
	0x69, 0x64, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21,
	0x0a, 0x0c, 0x67, 0x72, 0x6f, 0x73, 0x73, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x67, 0x72, 0x6f, 0x73, 0x73, 0x41, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x22, 0x8f, 0x01, 0x0a, 0x0c, 0x53, 0x6e,
	0x61, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x5f, 0x75, 0x72, 0x6c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74,
	0x55, 0x72, 0x6c, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x43, 0x6f, 0x64, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x32, 0x53, 0x0a, 0x0f, 0x4d,
	0x69, 0x64, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x40,
	0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d,
	0x69, 0x64, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10,
	0x2e, 0x70, 0x62, 0x2e, 0x53, 0x6e, 0x61, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x1d, 0x5a, 0x1b, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x5f, 0x66, 0x69,
	0x62, 0x65, 0x72, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_midtrans_proto_rawDescOnce sync.Once
	file_midtrans_proto_rawDescData = file_midtrans_proto_rawDesc
)

func file_midtrans_proto_rawDescGZIP() []byte {
	file_midtrans_proto_rawDescOnce.Do(func() {
		file_midtrans_proto_rawDescData = protoimpl.X.CompressGZIP(file_midtrans_proto_rawDescData)
	})
	return file_midtrans_proto_rawDescData
}

var file_midtrans_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_midtrans_proto_goTypes = []interface{}{
	(*CreateMidtransRequest)(nil), // 0: pb.CreateMidtransRequest
	(*SnapResponse)(nil),          // 1: pb.SnapResponse
}
var file_midtrans_proto_depIdxs = []int32{
	0, // 0: pb.MidtransService.CreateTransaction:input_type -> pb.CreateMidtransRequest
	1, // 1: pb.MidtransService.CreateTransaction:output_type -> pb.SnapResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_midtrans_proto_init() }
func file_midtrans_proto_init() {
	if File_midtrans_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_midtrans_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateMidtransRequest); i {
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
		file_midtrans_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SnapResponse); i {
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
			RawDescriptor: file_midtrans_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_midtrans_proto_goTypes,
		DependencyIndexes: file_midtrans_proto_depIdxs,
		MessageInfos:      file_midtrans_proto_msgTypes,
	}.Build()
	File_midtrans_proto = out.File
	file_midtrans_proto_rawDesc = nil
	file_midtrans_proto_goTypes = nil
	file_midtrans_proto_depIdxs = nil
}
