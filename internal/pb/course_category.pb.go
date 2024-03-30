// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: proto/course_category.proto

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

type Black struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Black) Reset() {
	*x = Black{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_course_category_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Black) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Black) ProtoMessage() {}

func (x *Black) ProtoReflect() protoreflect.Message {
	mi := &file_proto_course_category_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Black.ProtoReflect.Descriptor instead.
func (*Black) Descriptor() ([]byte, []int) {
	return file_proto_course_category_proto_rawDescGZIP(), []int{0}
}

type Category struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *Category) Reset() {
	*x = Category{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_course_category_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Category) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Category) ProtoMessage() {}

func (x *Category) ProtoReflect() protoreflect.Message {
	mi := &file_proto_course_category_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Category.ProtoReflect.Descriptor instead.
func (*Category) Descriptor() ([]byte, []int) {
	return file_proto_course_category_proto_rawDescGZIP(), []int{1}
}

func (x *Category) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Category) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Category) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type CreateCategoryReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *CreateCategoryReq) Reset() {
	*x = CreateCategoryReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_course_category_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCategoryReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCategoryReq) ProtoMessage() {}

func (x *CreateCategoryReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_course_category_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCategoryReq.ProtoReflect.Descriptor instead.
func (*CreateCategoryReq) Descriptor() ([]byte, []int) {
	return file_proto_course_category_proto_rawDescGZIP(), []int{2}
}

func (x *CreateCategoryReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateCategoryReq) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type CategoryGetByIdReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CategoryGetByIdReq) Reset() {
	*x = CategoryGetByIdReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_course_category_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CategoryGetByIdReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CategoryGetByIdReq) ProtoMessage() {}

func (x *CategoryGetByIdReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_course_category_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CategoryGetByIdReq.ProtoReflect.Descriptor instead.
func (*CategoryGetByIdReq) Descriptor() ([]byte, []int) {
	return file_proto_course_category_proto_rawDescGZIP(), []int{3}
}

func (x *CategoryGetByIdReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type CategoryList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Category []*Category `protobuf:"bytes,1,rep,name=category,proto3" json:"category,omitempty"`
}

func (x *CategoryList) Reset() {
	*x = CategoryList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_course_category_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CategoryList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CategoryList) ProtoMessage() {}

func (x *CategoryList) ProtoReflect() protoreflect.Message {
	mi := &file_proto_course_category_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CategoryList.ProtoReflect.Descriptor instead.
func (*CategoryList) Descriptor() ([]byte, []int) {
	return file_proto_course_category_proto_rawDescGZIP(), []int{4}
}

func (x *CategoryList) GetCategory() []*Category {
	if x != nil {
		return x.Category
	}
	return nil
}

var File_proto_course_category_proto protoreflect.FileDescriptor

var file_proto_course_category_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x5f, 0x63,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70,
	0x62, 0x22, 0x07, 0x0a, 0x05, 0x62, 0x6c, 0x61, 0x63, 0x6b, 0x22, 0x50, 0x0a, 0x08, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x49, 0x0a, 0x11,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65,
	0x71, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x24, 0x0a, 0x12, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x38, 0x0a,
	0x0c, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x28, 0x0a,
	0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x08, 0x63,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x32, 0xc8, 0x02, 0x0a, 0x0f, 0x43, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x37, 0x0a, 0x0e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x15, 0x2e,
	0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x52, 0x65, 0x71, 0x1a, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x22, 0x00, 0x12, 0x2f, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x12, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x62, 0x6c, 0x61, 0x63,
	0x6b, 0x1a, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x4c,
	0x69, 0x73, 0x74, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x43, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x49, 0x44, 0x12, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x0c,
	0x2e, 0x70, 0x62, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x22, 0x00, 0x12, 0x43,
	0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e,
	0x70, 0x62, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x22,
	0x00, 0x28, 0x01, 0x12, 0x4d, 0x0a, 0x20, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x42, 0x69, 0x64, 0x69, 0x72,
	0x65, 0x63, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x12, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x1a, 0x0c,
	0x2e, 0x70, 0x62, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x22, 0x00, 0x28, 0x01,
	0x30, 0x01, 0x42, 0x0d, 0x5a, 0x0b, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_course_category_proto_rawDescOnce sync.Once
	file_proto_course_category_proto_rawDescData = file_proto_course_category_proto_rawDesc
)

func file_proto_course_category_proto_rawDescGZIP() []byte {
	file_proto_course_category_proto_rawDescOnce.Do(func() {
		file_proto_course_category_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_course_category_proto_rawDescData)
	})
	return file_proto_course_category_proto_rawDescData
}

var file_proto_course_category_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_course_category_proto_goTypes = []interface{}{
	(*Black)(nil),              // 0: pb.black
	(*Category)(nil),           // 1: pb.Category
	(*CreateCategoryReq)(nil),  // 2: pb.CreateCategoryReq
	(*CategoryGetByIdReq)(nil), // 3: pb.CategoryGetByIdReq
	(*CategoryList)(nil),       // 4: pb.CategoryList
}
var file_proto_course_category_proto_depIdxs = []int32{
	1, // 0: pb.CategoryList.category:type_name -> pb.Category
	2, // 1: pb.CategoryService.CreateCategory:input_type -> pb.CreateCategoryReq
	0, // 2: pb.CategoryService.ListCategories:input_type -> pb.black
	3, // 3: pb.CategoryService.GetCategoryID:input_type -> pb.CategoryGetByIdReq
	2, // 4: pb.CategoryService.CreateCategoryStream:input_type -> pb.CreateCategoryReq
	2, // 5: pb.CategoryService.CreateCategoryStreamBidirecional:input_type -> pb.CreateCategoryReq
	1, // 6: pb.CategoryService.CreateCategory:output_type -> pb.Category
	4, // 7: pb.CategoryService.ListCategories:output_type -> pb.CategoryList
	1, // 8: pb.CategoryService.GetCategoryID:output_type -> pb.Category
	4, // 9: pb.CategoryService.CreateCategoryStream:output_type -> pb.CategoryList
	1, // 10: pb.CategoryService.CreateCategoryStreamBidirecional:output_type -> pb.Category
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_course_category_proto_init() }
func file_proto_course_category_proto_init() {
	if File_proto_course_category_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_course_category_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Black); i {
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
		file_proto_course_category_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Category); i {
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
		file_proto_course_category_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCategoryReq); i {
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
		file_proto_course_category_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CategoryGetByIdReq); i {
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
		file_proto_course_category_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CategoryList); i {
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
			RawDescriptor: file_proto_course_category_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_course_category_proto_goTypes,
		DependencyIndexes: file_proto_course_category_proto_depIdxs,
		MessageInfos:      file_proto_course_category_proto_msgTypes,
	}.Build()
	File_proto_course_category_proto = out.File
	file_proto_course_category_proto_rawDesc = nil
	file_proto_course_category_proto_goTypes = nil
	file_proto_course_category_proto_depIdxs = nil
}
