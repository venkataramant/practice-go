// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.1
// source: pb/backoffice/invoce.proto

package backoffice

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

type LineItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SKU    string `protobuf:"bytes,1,opt,name=SKU,proto3" json:"SKU,omitempty"`
	Amount int64  `protobuf:"varint,2,opt,name=Amount,proto3" json:"Amount,omitempty"`
	Price  int64  `protobuf:"varint,3,opt,name=Price,proto3" json:"Price,omitempty"`
}

func (x *LineItem) Reset() {
	*x = LineItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_backoffice_invoce_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LineItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LineItem) ProtoMessage() {}

func (x *LineItem) ProtoReflect() protoreflect.Message {
	mi := &file_pb_backoffice_invoce_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LineItem.ProtoReflect.Descriptor instead.
func (*LineItem) Descriptor() ([]byte, []int) {
	return file_pb_backoffice_invoce_proto_rawDescGZIP(), []int{0}
}

func (x *LineItem) GetSKU() string {
	if x != nil {
		return x.SKU
	}
	return ""
}

func (x *LineItem) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *LineItem) GetPrice() int64 {
	if x != nil {
		return x.Price
	}
	return 0
}

type Invoice struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID       string                 `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Time     *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=Time,proto3" json:"Time,omitempty"`
	Customer string                 `protobuf:"bytes,3,opt,name=Customer,proto3" json:"Customer,omitempty"`
	Items    []*LineItem            `protobuf:"bytes,4,rep,name=Items,proto3" json:"Items,omitempty"`
}

func (x *Invoice) Reset() {
	*x = Invoice{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_backoffice_invoce_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Invoice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Invoice) ProtoMessage() {}

func (x *Invoice) ProtoReflect() protoreflect.Message {
	mi := &file_pb_backoffice_invoce_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Invoice.ProtoReflect.Descriptor instead.
func (*Invoice) Descriptor() ([]byte, []int) {
	return file_pb_backoffice_invoce_proto_rawDescGZIP(), []int{1}
}

func (x *Invoice) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Invoice) GetTime() *timestamppb.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *Invoice) GetCustomer() string {
	if x != nil {
		return x.Customer
	}
	return ""
}

func (x *Invoice) GetItems() []*LineItem {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_pb_backoffice_invoce_proto protoreflect.FileDescriptor

var file_pb_backoffice_invoce_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x62, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x6f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x2f,
	0x69, 0x6e, 0x76, 0x6f, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4a, 0x0a,
	0x08, 0x4c, 0x69, 0x6e, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x10, 0x0a, 0x03, 0x53, 0x4b, 0x55,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x53, 0x4b, 0x55, 0x12, 0x16, 0x0a, 0x06, 0x41,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x41, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x22, 0x86, 0x01, 0x0a, 0x07, 0x49, 0x6e,
	0x76, 0x6f, 0x69, 0x63, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x2e, 0x0a, 0x04, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x04, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x12, 0x1f, 0x0a, 0x05, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x09, 0x2e, 0x4c, 0x69, 0x6e, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x49, 0x74, 0x65,
	0x6d, 0x73, 0x42, 0x0f, 0x5a, 0x0d, 0x70, 0x62, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x6f, 0x66, 0x66,
	0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_backoffice_invoce_proto_rawDescOnce sync.Once
	file_pb_backoffice_invoce_proto_rawDescData = file_pb_backoffice_invoce_proto_rawDesc
)

func file_pb_backoffice_invoce_proto_rawDescGZIP() []byte {
	file_pb_backoffice_invoce_proto_rawDescOnce.Do(func() {
		file_pb_backoffice_invoce_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_backoffice_invoce_proto_rawDescData)
	})
	return file_pb_backoffice_invoce_proto_rawDescData
}

var file_pb_backoffice_invoce_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pb_backoffice_invoce_proto_goTypes = []interface{}{
	(*LineItem)(nil),              // 0: LineItem
	(*Invoice)(nil),               // 1: Invoice
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_pb_backoffice_invoce_proto_depIdxs = []int32{
	2, // 0: Invoice.Time:type_name -> google.protobuf.Timestamp
	0, // 1: Invoice.Items:type_name -> LineItem
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_pb_backoffice_invoce_proto_init() }
func file_pb_backoffice_invoce_proto_init() {
	if File_pb_backoffice_invoce_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_backoffice_invoce_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LineItem); i {
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
		file_pb_backoffice_invoce_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Invoice); i {
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
			RawDescriptor: file_pb_backoffice_invoce_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pb_backoffice_invoce_proto_goTypes,
		DependencyIndexes: file_pb_backoffice_invoce_proto_depIdxs,
		MessageInfos:      file_pb_backoffice_invoce_proto_msgTypes,
	}.Build()
	File_pb_backoffice_invoce_proto = out.File
	file_pb_backoffice_invoce_proto_rawDesc = nil
	file_pb_backoffice_invoce_proto_goTypes = nil
	file_pb_backoffice_invoce_proto_depIdxs = nil
}
