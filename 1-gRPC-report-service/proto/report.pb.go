// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.12.4
// source: proto/report.proto

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

type ReportRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReportType string `protobuf:"bytes,1,opt,name=report_type,json=reportType,proto3" json:"report_type,omitempty"`
}

func (x *ReportRequest) Reset() {
	*x = ReportRequest{}
	mi := &file_proto_report_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReportRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportRequest) ProtoMessage() {}

func (x *ReportRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_report_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportRequest.ProtoReflect.Descriptor instead.
func (*ReportRequest) Descriptor() ([]byte, []int) {
	return file_proto_report_proto_rawDescGZIP(), []int{0}
}

func (x *ReportRequest) GetReportType() string {
	if x != nil {
		return x.ReportType
	}
	return ""
}

type ReportResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReportId string `protobuf:"bytes,1,opt,name=report_id,json=reportId,proto3" json:"report_id,omitempty"`
}

func (x *ReportResponse) Reset() {
	*x = ReportResponse{}
	mi := &file_proto_report_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReportResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportResponse) ProtoMessage() {}

func (x *ReportResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_report_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportResponse.ProtoReflect.Descriptor instead.
func (*ReportResponse) Descriptor() ([]byte, []int) {
	return file_proto_report_proto_rawDescGZIP(), []int{1}
}

func (x *ReportResponse) GetReportId() string {
	if x != nil {
		return x.ReportId
	}
	return ""
}

type StatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReportId string `protobuf:"bytes,1,opt,name=report_id,json=reportId,proto3" json:"report_id,omitempty"`
}

func (x *StatusRequest) Reset() {
	*x = StatusRequest{}
	mi := &file_proto_report_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusRequest) ProtoMessage() {}

func (x *StatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_report_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusRequest.ProtoReflect.Descriptor instead.
func (*StatusRequest) Descriptor() ([]byte, []int) {
	return file_proto_report_proto_rawDescGZIP(), []int{2}
}

func (x *StatusRequest) GetReportId() string {
	if x != nil {
		return x.ReportId
	}
	return ""
}

type StatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReportId  string `protobuf:"bytes,1,opt,name=report_id,json=reportId,proto3" json:"report_id,omitempty"`
	Status    string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	ReportUrl string `protobuf:"bytes,3,opt,name=report_url,json=reportUrl,proto3" json:"report_url,omitempty"`
}

func (x *StatusResponse) Reset() {
	*x = StatusResponse{}
	mi := &file_proto_report_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusResponse) ProtoMessage() {}

func (x *StatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_report_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusResponse.ProtoReflect.Descriptor instead.
func (*StatusResponse) Descriptor() ([]byte, []int) {
	return file_proto_report_proto_rawDescGZIP(), []int{3}
}

func (x *StatusResponse) GetReportId() string {
	if x != nil {
		return x.ReportId
	}
	return ""
}

func (x *StatusResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *StatusResponse) GetReportUrl() string {
	if x != nil {
		return x.ReportUrl
	}
	return ""
}

var File_proto_report_proto protoreflect.FileDescriptor

var file_proto_report_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x22, 0x30, 0x0a, 0x0d,
	0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a,
	0x0b, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x22, 0x2d,
	0x0a, 0x0e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x49, 0x64, 0x22, 0x2c, 0x0a,
	0x0d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b,
	0x0a, 0x09, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x49, 0x64, 0x22, 0x64, 0x0a, 0x0e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a,
	0x09, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x75, 0x72, 0x6c,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x55, 0x72,
	0x6c, 0x32, 0x97, 0x01, 0x0a, 0x0d, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x3f, 0x0a, 0x0e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x15, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x52,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x72,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x12, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x15, 0x2e, 0x72, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x42, 0x3d, 0x5a, 0x3b, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x34, 0x30, 0x34, 0x37, 0x31, 0x2f,
	0x67, 0x6f, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73,
	0x2f, 0x67, 0x52, 0x50, 0x43, 0x2d, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2d, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_proto_report_proto_rawDescOnce sync.Once
	file_proto_report_proto_rawDescData = file_proto_report_proto_rawDesc
)

func file_proto_report_proto_rawDescGZIP() []byte {
	file_proto_report_proto_rawDescOnce.Do(func() {
		file_proto_report_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_report_proto_rawDescData)
	})
	return file_proto_report_proto_rawDescData
}

var file_proto_report_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_report_proto_goTypes = []any{
	(*ReportRequest)(nil),  // 0: report.ReportRequest
	(*ReportResponse)(nil), // 1: report.ReportResponse
	(*StatusRequest)(nil),  // 2: report.StatusRequest
	(*StatusResponse)(nil), // 3: report.StatusResponse
}
var file_proto_report_proto_depIdxs = []int32{
	0, // 0: report.ReportService.GenerateReport:input_type -> report.ReportRequest
	2, // 1: report.ReportService.StreamReportStatus:input_type -> report.StatusRequest
	1, // 2: report.ReportService.GenerateReport:output_type -> report.ReportResponse
	3, // 3: report.ReportService.StreamReportStatus:output_type -> report.StatusResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_report_proto_init() }
func file_proto_report_proto_init() {
	if File_proto_report_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_report_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_report_proto_goTypes,
		DependencyIndexes: file_proto_report_proto_depIdxs,
		MessageInfos:      file_proto_report_proto_msgTypes,
	}.Build()
	File_proto_report_proto = out.File
	file_proto_report_proto_rawDesc = nil
	file_proto_report_proto_goTypes = nil
	file_proto_report_proto_depIdxs = nil
}
