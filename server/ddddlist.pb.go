// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.4
// source: ddddlist.proto

package server

import (
	reflect "reflect"
	sync "sync"

	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "trpc.group/trpc/trpc-protocol/pb/go/trpc/proto"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type HelloRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *HelloRequest) Reset() {
	*x = HelloRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ddddlist_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloRequest) ProtoMessage() {}

func (x *HelloRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ddddlist_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloRequest.ProtoReflect.Descriptor instead.
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return file_ddddlist_proto_rawDescGZIP(), []int{0}
}

func (x *HelloRequest) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type HelloResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *HelloResponse) Reset() {
	*x = HelloResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ddddlist_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloResponse) ProtoMessage() {}

func (x *HelloResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ddddlist_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloResponse.ProtoReflect.Descriptor instead.
func (*HelloResponse) Descriptor() ([]byte, []int) {
	return file_ddddlist_proto_rawDescGZIP(), []int{1}
}

func (x *HelloResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url     string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Country string `protobuf:"bytes,2,opt,name=country,proto3" json:"country,omitempty"`
	Length  int32  `protobuf:"varint,3,opt,name=length,proto3" json:"length,omitempty"`
	Offset  int32  `protobuf:"varint,4,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ddddlist_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_ddddlist_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_ddddlist_proto_rawDescGZIP(), []int{2}
}

func (x *Request) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Request) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *Request) GetLength() int32 {
	if x != nil {
		return x.Length
	}
	return 0
}

func (x *Request) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Video   *Data  `protobuf:"bytes,3,opt,name=video,proto3" json:"video,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ddddlist_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_ddddlist_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_ddddlist_proto_rawDescGZIP(), []int{3}
}

func (x *Response) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Response) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Response) GetVideo() *Data {
	if x != nil {
		return x.Video
	}
	return nil
}

type Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SeriesId     string            `protobuf:"bytes,1,opt,name=SeriesId,proto3" json:"SeriesId,omitempty"`
	SeriesTitle  string            `protobuf:"bytes,2,opt,name=SeriesTitle,proto3" json:"SeriesTitle,omitempty"`
	Descr        string   `protobuf:"bytes,3,opt,name=descr,proto3" json:"descr,omitempty"`
	VideoList    []*Video `protobuf:"bytes,4,rep,name=videoList,proto3" json:"videoList,omitempty"`
	IsVip        bool     `protobuf:"varint,5,opt,name=isVip,proto3" json:"isVip,omitempty"`
	CoverHzUrl   string            `protobuf:"bytes,6,opt,name=coverHzUrl,proto3" json:"coverHzUrl,omitempty"`
	CoverVtUrl   string            `protobuf:"bytes,7,opt,name=coverVtUrl,proto3" json:"coverVtUrl,omitempty"`
	TypeName     string            `protobuf:"bytes,8,opt,name=typeName,proto3" json:"typeName,omitempty"`
	CategoryName string            `protobuf:"bytes,9,opt,name=categoryName,proto3" json:"categoryName,omitempty"`
	VidNum       uint32            `protobuf:"varint,10,opt,name=vidNum,proto3" json:"vidNum,omitempty"`
	UpdateNum    uint32            `protobuf:"varint,11,opt,name=updateNum,proto3" json:"updateNum,omitempty"`
	Status       string            `protobuf:"bytes,12,opt,name=status,proto3" json:"status,omitempty"`
	CreateTime   string            `protobuf:"bytes,13,opt,name=createTime,proto3" json:"createTime,omitempty"`
	IsSeries     bool              `protobuf:"varint,14,opt,name=IsSeries,proto3" json:"IsSeries,omitempty"`
	Extra        map[string]string `protobuf:"bytes,15,rep,name=extra,proto3" json:"extra,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Data) Reset() {
	*x = Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ddddlist_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data) ProtoMessage() {}

func (x *Data) ProtoReflect() protoreflect.Message {
	mi := &file_ddddlist_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data.ProtoReflect.Descriptor instead.
func (*Data) Descriptor() ([]byte, []int) {
	return file_ddddlist_proto_rawDescGZIP(), []int{4}
}

func (x *Data) GetSeriesId() string {
	if x != nil {
		return x.SeriesId
	}
	return ""
}

func (x *Data) GetSeriesTitle() string {
	if x != nil {
		return x.SeriesTitle
	}
	return ""
}

func (x *Data) GetDescr() string {
	if x != nil {
		return x.Descr
	}
	return ""
}

func (x *Data) GetVideoList() []*Video {
	if x != nil {
		return x.VideoList
	}
	return nil
}

func (x *Data) GetIsVip() bool {
	if x != nil {
		return x.IsVip
	}
	return false
}

func (x *Data) GetCoverHzUrl() string {
	if x != nil {
		return x.CoverHzUrl
	}
	return ""
}

func (x *Data) GetCoverVtUrl() string {
	if x != nil {
		return x.CoverVtUrl
	}
	return ""
}

func (x *Data) GetTypeName() string {
	if x != nil {
		return x.TypeName
	}
	return ""
}

func (x *Data) GetCategoryName() string {
	if x != nil {
		return x.CategoryName
	}
	return ""
}

func (x *Data) GetVidNum() uint32 {
	if x != nil {
		return x.VidNum
	}
	return 0
}

func (x *Data) GetUpdateNum() uint32 {
	if x != nil {
		return x.UpdateNum
	}
	return 0
}

func (x *Data) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Data) GetCreateTime() string {
	if x != nil {
		return x.CreateTime
	}
	return ""
}

func (x *Data) GetIsSeries() bool {
	if x != nil {
		return x.IsSeries
	}
	return false
}

func (x *Data) GetExtra() map[string]string {
	if x != nil {
		return x.Extra
	}
	return nil
}

type Video struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EpisodeTitle string            `protobuf:"bytes,1,opt,name=EpisodeTitle,proto3" json:"EpisodeTitle,omitempty"`
	EpisodeId    string            `protobuf:"bytes,2,opt,name=EpisodeId,proto3" json:"EpisodeId,omitempty"`
	SubTitle     string            `protobuf:"bytes,3,opt,name=subTitle,proto3" json:"subTitle,omitempty"`
	CoverUrl     string            `protobuf:"bytes,4,opt,name=coverUrl,proto3" json:"coverUrl,omitempty"`
	IsTrailer    bool              `protobuf:"varint,5,opt,name=isTrailer,proto3" json:"isTrailer,omitempty"`
	IsVip        bool              `protobuf:"varint,8,opt,name=isVip,proto3" json:"isVip,omitempty"`
	Duration     uint32            `protobuf:"varint,9,opt,name=duration,proto3" json:"duration,omitempty"`
	Status       string            `protobuf:"bytes,11,opt,name=status,proto3" json:"status,omitempty"`
	Episode      uint32            `protobuf:"varint,12,opt,name=episode,proto3" json:"episode,omitempty"`
	Extra        map[string]string `protobuf:"bytes,13,rep,name=extra,proto3" json:"extra,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Video) Reset() {
	*x = Video{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ddddlist_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Video) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Video) ProtoMessage() {}

func (x *Video) ProtoReflect() protoreflect.Message {
	mi := &file_ddddlist_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Video.ProtoReflect.Descriptor instead.
func (*Video) Descriptor() ([]byte, []int) {
	return file_ddddlist_proto_rawDescGZIP(), []int{5}
}

func (x *Video) GetEpisodeTitle() string {
	if x != nil {
		return x.EpisodeTitle
	}
	return ""
}

func (x *Video) GetEpisodeId() string {
	if x != nil {
		return x.EpisodeId
	}
	return ""
}

func (x *Video) GetSubTitle() string {
	if x != nil {
		return x.SubTitle
	}
	return ""
}

func (x *Video) GetCoverUrl() string {
	if x != nil {
		return x.CoverUrl
	}
	return ""
}

func (x *Video) GetIsTrailer() bool {
	if x != nil {
		return x.IsTrailer
	}
	return false
}

func (x *Video) GetIsVip() bool {
	if x != nil {
		return x.IsVip
	}
	return false
}

func (x *Video) GetDuration() uint32 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *Video) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Video) GetEpisode() uint32 {
	if x != nil {
		return x.Episode
	}
	return 0
}

func (x *Video) GetExtra() map[string]string {
	if x != nil {
		return x.Extra
	}
	return nil
}

var File_ddddlist_proto protoreflect.FileDescriptor

var file_ddddlist_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x64, 0x64, 0x64, 0x64, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x1a, 0x1d, 0x74, 0x72, 0x70, 0x63, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x72, 0x70, 0x63, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x20, 0x0a, 0x0c, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d,
	0x73, 0x67, 0x22, 0x21, 0x0a, 0x0d, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x65, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75,
	0x72, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x16, 0x0a, 0x06,
	0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6c, 0x65,
	0x6e, 0x67, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x5c, 0x0a, 0x08,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x22, 0x0a, 0x05, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x44,
	0x61, 0x74, 0x61, 0x52, 0x05, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x22, 0x90, 0x04, 0x0a, 0x04, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x1a, 0x0a, 0x08, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x49, 0x64, 0x12,
	0x20, 0x0a, 0x0b, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x54, 0x69, 0x74, 0x6c,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x65, 0x73, 0x63, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x64, 0x65, 0x73, 0x63, 0x72, 0x12, 0x2b, 0x0a, 0x09, 0x76, 0x69, 0x64, 0x65, 0x6f,
	0x4c, 0x69, 0x73, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x09, 0x76, 0x69, 0x64, 0x65, 0x6f,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x73, 0x56, 0x69, 0x70, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x05, 0x69, 0x73, 0x56, 0x69, 0x70, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f,
	0x76, 0x65, 0x72, 0x48, 0x7a, 0x55, 0x72, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x48, 0x7a, 0x55, 0x72, 0x6c, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f,
	0x76, 0x65, 0x72, 0x56, 0x74, 0x55, 0x72, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x56, 0x74, 0x55, 0x72, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x79,
	0x70, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x79,
	0x70, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x69,
	0x64, 0x4e, 0x75, 0x6d, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x76, 0x69, 0x64, 0x4e,
	0x75, 0x6d, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x75, 0x6d, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x75, 0x6d,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x49, 0x73, 0x53, 0x65,
	0x72, 0x69, 0x65, 0x73, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x49, 0x73, 0x53, 0x65,
	0x72, 0x69, 0x65, 0x73, 0x12, 0x2d, 0x0a, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x18, 0x0f, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x44, 0x61, 0x74,
	0x61, 0x2e, 0x45, 0x78, 0x74, 0x72, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x65, 0x78,
	0x74, 0x72, 0x61, 0x1a, 0x38, 0x0a, 0x0a, 0x45, 0x78, 0x74, 0x72, 0x61, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xed, 0x02,
	0x0a, 0x05, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x12, 0x22, 0x0a, 0x0c, 0x45, 0x70, 0x69, 0x73, 0x6f,
	0x64, 0x65, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x45,
	0x70, 0x69, 0x73, 0x6f, 0x64, 0x65, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x45,
	0x70, 0x69, 0x73, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x45, 0x70, 0x69, 0x73, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x75, 0x62,
	0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x75, 0x62,
	0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x55, 0x72,
	0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x55, 0x72,
	0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x73, 0x54, 0x72, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x54, 0x72, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x12,
	0x14, 0x0a, 0x05, 0x69, 0x73, 0x56, 0x69, 0x70, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05,
	0x69, 0x73, 0x56, 0x69, 0x70, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x70, 0x69,
	0x73, 0x6f, 0x64, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x65, 0x70, 0x69, 0x73,
	0x6f, 0x64, 0x65, 0x12, 0x2e, 0x0a, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x18, 0x0d, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x56, 0x69, 0x64, 0x65,
	0x6f, 0x2e, 0x45, 0x78, 0x74, 0x72, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x65, 0x78,
	0x74, 0x72, 0x61, 0x1a, 0x38, 0x0a, 0x0a, 0x45, 0x78, 0x74, 0x72, 0x61, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x32, 0x71, 0x0a,
	0x06, 0x44, 0x44, 0x44, 0x44, 0x68, 0x6d, 0x12, 0x2f, 0x0a, 0x08, 0x44, 0x64, 0x64, 0x64, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x0f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x05, 0x48, 0x65, 0x6c, 0x6c,
	0x6f, 0x12, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x0e, 0x5a, 0x0c, 0x2e, 0x2e, 0x2f, 0x2e, 0x2e, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ddddlist_proto_rawDescOnce sync.Once
	file_ddddlist_proto_rawDescData = file_ddddlist_proto_rawDesc
)

func file_ddddlist_proto_rawDescGZIP() []byte {
	file_ddddlist_proto_rawDescOnce.Do(func() {
		file_ddddlist_proto_rawDescData = protoimpl.X.CompressGZIP(file_ddddlist_proto_rawDescData)
	})
	return file_ddddlist_proto_rawDescData
}

var file_ddddlist_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_ddddlist_proto_goTypes = []interface{}{
	(*HelloRequest)(nil),  // 0: server.HelloRequest
	(*HelloResponse)(nil), // 1: server.HelloResponse
	(*Request)(nil),       // 2: server.Request
	(*Response)(nil),      // 3: server.Response
	(*Data)(nil),          // 4: server.Data
	(*Video)(nil),         // 5: server.Video
	nil,                   // 6: server.Data.ExtraEntry
	nil,                   // 7: server.Video.ExtraEntry
}
var file_ddddlist_proto_depIdxs = []int32{
	4, // 0: server.Response.video:type_name -> server.Data
	5, // 1: server.Data.videoList:type_name -> server.Video
	6, // 2: server.Data.extra:type_name -> server.Data.ExtraEntry
	7, // 3: server.Video.extra:type_name -> server.Video.ExtraEntry
	2, // 4: server.DDDDhm.DdddList:input_type -> server.Request
	0, // 5: server.DDDDhm.Hello:input_type -> server.HelloRequest
	3, // 6: server.DDDDhm.DdddList:output_type -> server.Response
	1, // 7: server.DDDDhm.Hello:output_type -> server.HelloResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_ddddlist_proto_init() }
func file_ddddlist_proto_init() {
	if File_ddddlist_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ddddlist_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloRequest); i {
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
		file_ddddlist_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloResponse); i {
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
		file_ddddlist_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_ddddlist_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_ddddlist_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Data); i {
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
		file_ddddlist_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Video); i {
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
			RawDescriptor: file_ddddlist_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ddddlist_proto_goTypes,
		DependencyIndexes: file_ddddlist_proto_depIdxs,
		MessageInfos:      file_ddddlist_proto_msgTypes,
	}.Build()
	File_ddddlist_proto = out.File
	file_ddddlist_proto_rawDesc = nil
	file_ddddlist_proto_goTypes = nil
	file_ddddlist_proto_depIdxs = nil
}