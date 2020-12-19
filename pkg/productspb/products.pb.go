// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: api/products.proto

package productspb

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

// downloads csv of form product_name;price by given url
// writes downloaded products to mongo updating price as necessary with update count and time
type FetchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *FetchRequest) Reset() {
	*x = FetchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_products_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchRequest) ProtoMessage() {}

func (x *FetchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_products_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchRequest.ProtoReflect.Descriptor instead.
func (*FetchRequest) Descriptor() ([]byte, []int) {
	return file_api_products_proto_rawDescGZIP(), []int{0}
}

func (x *FetchRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

// empty
type FetchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FetchResponse) Reset() {
	*x = FetchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_products_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchResponse) ProtoMessage() {}

func (x *FetchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_products_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchResponse.ProtoReflect.Descriptor instead.
func (*FetchResponse) Descriptor() ([]byte, []int) {
	return file_api_products_proto_rawDescGZIP(), []int{1}
}

// returns a requested page of products
// able to sort by any product's field
// what if I change sorting method for arbitrary page?
type ListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paging  *ListRequest_Paging  `protobuf:"bytes,1,opt,name=paging,proto3" json:"paging,omitempty"`
	Sorting *ListRequest_Sorting `protobuf:"bytes,2,opt,name=sorting,proto3" json:"sorting,omitempty"`
}

func (x *ListRequest) Reset() {
	*x = ListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_products_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRequest) ProtoMessage() {}

func (x *ListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_products_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRequest.ProtoReflect.Descriptor instead.
func (*ListRequest) Descriptor() ([]byte, []int) {
	return file_api_products_proto_rawDescGZIP(), []int{2}
}

func (x *ListRequest) GetPaging() *ListRequest_Paging {
	if x != nil {
		return x.Paging
	}
	return nil
}

func (x *ListRequest) GetSorting() *ListRequest_Sorting {
	if x != nil {
		return x.Sorting
	}
	return nil
}

type ListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Products []*ListResponse_Product `protobuf:"bytes,1,rep,name=products,proto3" json:"products,omitempty"`
}

func (x *ListResponse) Reset() {
	*x = ListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_products_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResponse) ProtoMessage() {}

func (x *ListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_products_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListResponse.ProtoReflect.Descriptor instead.
func (*ListResponse) Descriptor() ([]byte, []int) {
	return file_api_products_proto_rawDescGZIP(), []int{3}
}

func (x *ListResponse) GetProducts() []*ListResponse_Product {
	if x != nil {
		return x.Products
	}
	return nil
}

type ListRequest_Paging struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset uint32 `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit  uint32 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *ListRequest_Paging) Reset() {
	*x = ListRequest_Paging{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_products_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRequest_Paging) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRequest_Paging) ProtoMessage() {}

func (x *ListRequest_Paging) ProtoReflect() protoreflect.Message {
	mi := &file_api_products_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRequest_Paging.ProtoReflect.Descriptor instead.
func (*ListRequest_Paging) Descriptor() ([]byte, []int) {
	return file_api_products_proto_rawDescGZIP(), []int{2, 0}
}

func (x *ListRequest_Paging) GetOffset() uint32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ListRequest_Paging) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type ListRequest_Sorting struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ascending bool   `protobuf:"varint,1,opt,name=ascending,proto3" json:"ascending,omitempty"`
	SortBy    string `protobuf:"bytes,2,opt,name=sortBy,proto3" json:"sortBy,omitempty"`
}

func (x *ListRequest_Sorting) Reset() {
	*x = ListRequest_Sorting{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_products_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRequest_Sorting) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRequest_Sorting) ProtoMessage() {}

func (x *ListRequest_Sorting) ProtoReflect() protoreflect.Message {
	mi := &file_api_products_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRequest_Sorting.ProtoReflect.Descriptor instead.
func (*ListRequest_Sorting) Descriptor() ([]byte, []int) {
	return file_api_products_proto_rawDescGZIP(), []int{2, 1}
}

func (x *ListRequest_Sorting) GetAscending() bool {
	if x != nil {
		return x.Ascending
	}
	return false
}

func (x *ListRequest_Sorting) GetSortBy() string {
	if x != nil {
		return x.SortBy
	}
	return ""
}

type ListResponse_Product struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name             string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Price            string                 `protobuf:"bytes,3,opt,name=price,proto3" json:"price,omitempty"`
	PriceUpdateCount uint32                 `protobuf:"varint,4,opt,name=priceUpdateCount,proto3" json:"priceUpdateCount,omitempty"`
	LastModified     *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=lastModified,proto3" json:"lastModified,omitempty"`
}

func (x *ListResponse_Product) Reset() {
	*x = ListResponse_Product{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_products_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListResponse_Product) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResponse_Product) ProtoMessage() {}

func (x *ListResponse_Product) ProtoReflect() protoreflect.Message {
	mi := &file_api_products_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListResponse_Product.ProtoReflect.Descriptor instead.
func (*ListResponse_Product) Descriptor() ([]byte, []int) {
	return file_api_products_proto_rawDescGZIP(), []int{3, 0}
}

func (x *ListResponse_Product) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ListResponse_Product) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ListResponse_Product) GetPrice() string {
	if x != nil {
		return x.Price
	}
	return ""
}

func (x *ListResponse_Product) GetPriceUpdateCount() uint32 {
	if x != nil {
		return x.PriceUpdateCount
	}
	return 0
}

func (x *ListResponse_Product) GetLastModified() *timestamppb.Timestamp {
	if x != nil {
		return x.LastModified
	}
	return nil
}

var File_api_products_proto protoreflect.FileDescriptor

var file_api_products_proto_rawDesc = []byte{
	0x0a, 0x12, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x20, 0x0a, 0x0c, 0x46, 0x65, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72,
	0x6c, 0x22, 0x0f, 0x0a, 0x0d, 0x46, 0x65, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0xf5, 0x01, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x34, 0x0a, 0x06, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x67,
	0x52, 0x06, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x12, 0x37, 0x0a, 0x07, 0x73, 0x6f, 0x72, 0x74,
	0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x2e, 0x53, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x07, 0x73, 0x6f, 0x72, 0x74, 0x69, 0x6e,
	0x67, 0x1a, 0x36, 0x0a, 0x06, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x6f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x6f, 0x66, 0x66,
	0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x1a, 0x3f, 0x0a, 0x07, 0x53, 0x6f, 0x72,
	0x74, 0x69, 0x6e, 0x67, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x73, 0x63, 0x65, 0x6e, 0x64, 0x69, 0x6e,
	0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x61, 0x73, 0x63, 0x65, 0x6e, 0x64, 0x69,
	0x6e, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x72, 0x74, 0x42, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x72, 0x74, 0x42, 0x79, 0x22, 0xfc, 0x01, 0x0a, 0x0c, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x08, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x08, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x1a, 0xaf, 0x01, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x2a, 0x0a,
	0x10, 0x70, 0x72, 0x69, 0x63, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x10, 0x70, 0x72, 0x69, 0x63, 0x65, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x3e, 0x0a, 0x0c, 0x6c, 0x61, 0x73,
	0x74, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0c, 0x6c, 0x61, 0x73,
	0x74, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x32, 0x7f, 0x0a, 0x08, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x73, 0x12, 0x3a, 0x0a, 0x05, 0x46, 0x65, 0x74, 0x63, 0x68, 0x12, 0x16,
	0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x73, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x37, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x17, 0x5a, 0x15, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x70, 0x62, 0x3b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x73, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_products_proto_rawDescOnce sync.Once
	file_api_products_proto_rawDescData = file_api_products_proto_rawDesc
)

func file_api_products_proto_rawDescGZIP() []byte {
	file_api_products_proto_rawDescOnce.Do(func() {
		file_api_products_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_products_proto_rawDescData)
	})
	return file_api_products_proto_rawDescData
}

var file_api_products_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_api_products_proto_goTypes = []interface{}{
	(*FetchRequest)(nil),          // 0: products.FetchRequest
	(*FetchResponse)(nil),         // 1: products.FetchResponse
	(*ListRequest)(nil),           // 2: products.ListRequest
	(*ListResponse)(nil),          // 3: products.ListResponse
	(*ListRequest_Paging)(nil),    // 4: products.ListRequest.Paging
	(*ListRequest_Sorting)(nil),   // 5: products.ListRequest.Sorting
	(*ListResponse_Product)(nil),  // 6: products.ListResponse.Product
	(*timestamppb.Timestamp)(nil), // 7: google.protobuf.Timestamp
}
var file_api_products_proto_depIdxs = []int32{
	4, // 0: products.ListRequest.paging:type_name -> products.ListRequest.Paging
	5, // 1: products.ListRequest.sorting:type_name -> products.ListRequest.Sorting
	6, // 2: products.ListResponse.products:type_name -> products.ListResponse.Product
	7, // 3: products.ListResponse.Product.lastModified:type_name -> google.protobuf.Timestamp
	0, // 4: products.Products.Fetch:input_type -> products.FetchRequest
	2, // 5: products.Products.List:input_type -> products.ListRequest
	1, // 6: products.Products.Fetch:output_type -> products.FetchResponse
	3, // 7: products.Products.List:output_type -> products.ListResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_api_products_proto_init() }
func file_api_products_proto_init() {
	if File_api_products_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_products_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchRequest); i {
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
		file_api_products_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchResponse); i {
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
		file_api_products_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRequest); i {
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
		file_api_products_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListResponse); i {
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
		file_api_products_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRequest_Paging); i {
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
		file_api_products_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRequest_Sorting); i {
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
		file_api_products_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListResponse_Product); i {
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
			RawDescriptor: file_api_products_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_products_proto_goTypes,
		DependencyIndexes: file_api_products_proto_depIdxs,
		MessageInfos:      file_api_products_proto_msgTypes,
	}.Build()
	File_api_products_proto = out.File
	file_api_products_proto_rawDesc = nil
	file_api_products_proto_goTypes = nil
	file_api_products_proto_depIdxs = nil
}
