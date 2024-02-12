// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.1
// source: consumerRPC.proto

package pb

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

type RequestDataStream struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Start bool `protobuf:"varint,1,opt,name=start,proto3" json:"start,omitempty"`
}

func (x *RequestDataStream) Reset() {
	*x = RequestDataStream{}
	if protoimpl.UnsafeEnabled {
		mi := &file_consumerRPC_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestDataStream) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestDataStream) ProtoMessage() {}

func (x *RequestDataStream) ProtoReflect() protoreflect.Message {
	mi := &file_consumerRPC_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestDataStream.ProtoReflect.Descriptor instead.
func (*RequestDataStream) Descriptor() ([]byte, []int) {
	return file_consumerRPC_proto_rawDescGZIP(), []int{0}
}

func (x *RequestDataStream) GetStart() bool {
	if x != nil {
		return x.Start
	}
	return false
}

type ResponseDataStream struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Date   *timestamppb.Timestamp     `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
	Bid    *ResponseDataStreamDecimal `protobuf:"bytes,2,opt,name=bid,proto3" json:"bid,omitempty"`
	Ask    *ResponseDataStreamDecimal `protobuf:"bytes,3,opt,name=ask,proto3" json:"ask,omitempty"`
	Symbol string                     `protobuf:"bytes,4,opt,name=symbol,proto3" json:"symbol,omitempty"`
}

func (x *ResponseDataStream) Reset() {
	*x = ResponseDataStream{}
	if protoimpl.UnsafeEnabled {
		mi := &file_consumerRPC_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseDataStream) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseDataStream) ProtoMessage() {}

func (x *ResponseDataStream) ProtoReflect() protoreflect.Message {
	mi := &file_consumerRPC_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseDataStream.ProtoReflect.Descriptor instead.
func (*ResponseDataStream) Descriptor() ([]byte, []int) {
	return file_consumerRPC_proto_rawDescGZIP(), []int{1}
}

func (x *ResponseDataStream) GetDate() *timestamppb.Timestamp {
	if x != nil {
		return x.Date
	}
	return nil
}

func (x *ResponseDataStream) GetBid() *ResponseDataStreamDecimal {
	if x != nil {
		return x.Bid
	}
	return nil
}

func (x *ResponseDataStream) GetAsk() *ResponseDataStreamDecimal {
	if x != nil {
		return x.Ask
	}
	return nil
}

func (x *ResponseDataStream) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

type ResponseDataStreamDecimal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value int64 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	Exp   int32 `protobuf:"varint,2,opt,name=exp,proto3" json:"exp,omitempty"`
}

func (x *ResponseDataStreamDecimal) Reset() {
	*x = ResponseDataStreamDecimal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_consumerRPC_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseDataStreamDecimal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseDataStreamDecimal) ProtoMessage() {}

func (x *ResponseDataStreamDecimal) ProtoReflect() protoreflect.Message {
	mi := &file_consumerRPC_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseDataStreamDecimal.ProtoReflect.Descriptor instead.
func (*ResponseDataStreamDecimal) Descriptor() ([]byte, []int) {
	return file_consumerRPC_proto_rawDescGZIP(), []int{1, 0}
}

func (x *ResponseDataStreamDecimal) GetValue() int64 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *ResponseDataStreamDecimal) GetExp() int32 {
	if x != nil {
		return x.Exp
	}
	return 0
}

var File_consumerRPC_proto protoreflect.FileDescriptor

var file_consumerRPC_proto_rawDesc = []byte{
	0x0a, 0x11, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x52, 0x50, 0x43, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x29, 0x0a, 0x11, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x14, 0x0a,
	0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x22, 0xf3, 0x01, 0x0a, 0x12, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x44, 0x61, 0x74, 0x61, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x2e, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x30, 0x0a, 0x03, 0x62, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x70, 0x62, 0x2e, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x61, 0x74, 0x61, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e,
	0x64, 0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x52, 0x03, 0x62, 0x69, 0x64, 0x12, 0x30, 0x0a, 0x03,
	0x61, 0x73, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x70, 0x62, 0x2e, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x61, 0x74, 0x61, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x2e, 0x64, 0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x52, 0x03, 0x61, 0x73, 0x6b, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x1a, 0x31, 0x0a, 0x07, 0x64, 0x65, 0x63, 0x69, 0x6d, 0x61,
	0x6c, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x78, 0x70, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x65, 0x78, 0x70, 0x32, 0x4b, 0x0a, 0x08, 0x43, 0x6f, 0x6e,
	0x73, 0x75, 0x6d, 0x65, 0x72, 0x12, 0x3f, 0x0a, 0x0a, 0x44, 0x61, 0x74, 0x61, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x12, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x44, 0x61, 0x74, 0x61, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x1a, 0x16, 0x2e, 0x70, 0x62, 0x2e,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x61, 0x74, 0x61, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x22, 0x00, 0x30, 0x01, 0x42, 0x05, 0x5a, 0x03, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_consumerRPC_proto_rawDescOnce sync.Once
	file_consumerRPC_proto_rawDescData = file_consumerRPC_proto_rawDesc
)

func file_consumerRPC_proto_rawDescGZIP() []byte {
	file_consumerRPC_proto_rawDescOnce.Do(func() {
		file_consumerRPC_proto_rawDescData = protoimpl.X.CompressGZIP(file_consumerRPC_proto_rawDescData)
	})
	return file_consumerRPC_proto_rawDescData
}

var file_consumerRPC_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_consumerRPC_proto_goTypes = []interface{}{
	(*RequestDataStream)(nil),         // 0: pb.requestDataStream
	(*ResponseDataStream)(nil),        // 1: pb.responseDataStream
	(*ResponseDataStreamDecimal)(nil), // 2: pb.responseDataStream.decimal
	(*timestamppb.Timestamp)(nil),     // 3: google.protobuf.Timestamp
}
var file_consumerRPC_proto_depIdxs = []int32{
	3, // 0: pb.responseDataStream.date:type_name -> google.protobuf.Timestamp
	2, // 1: pb.responseDataStream.bid:type_name -> pb.responseDataStream.decimal
	2, // 2: pb.responseDataStream.ask:type_name -> pb.responseDataStream.decimal
	0, // 3: pb.Consumer.DataStream:input_type -> pb.requestDataStream
	1, // 4: pb.Consumer.DataStream:output_type -> pb.responseDataStream
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_consumerRPC_proto_init() }
func file_consumerRPC_proto_init() {
	if File_consumerRPC_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_consumerRPC_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestDataStream); i {
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
		file_consumerRPC_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseDataStream); i {
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
		file_consumerRPC_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseDataStreamDecimal); i {
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
			RawDescriptor: file_consumerRPC_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_consumerRPC_proto_goTypes,
		DependencyIndexes: file_consumerRPC_proto_depIdxs,
		MessageInfos:      file_consumerRPC_proto_msgTypes,
	}.Build()
	File_consumerRPC_proto = out.File
	file_consumerRPC_proto_rawDesc = nil
	file_consumerRPC_proto_goTypes = nil
	file_consumerRPC_proto_depIdxs = nil
}
