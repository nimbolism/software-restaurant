// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.26.1
// source: food_service.proto

package proto

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

type Food struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Category    string                 `protobuf:"bytes,4,opt,name=category,proto3" json:"category,omitempty"`
	Meal        string                 `protobuf:"bytes,5,opt,name=meal,proto3" json:"meal,omitempty"`
	CreatedAt   *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt   *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Food) Reset() {
	*x = Food{}
	if protoimpl.UnsafeEnabled {
		mi := &file_food_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Food) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Food) ProtoMessage() {}

func (x *Food) ProtoReflect() protoreflect.Message {
	mi := &file_food_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Food.ProtoReflect.Descriptor instead.
func (*Food) Descriptor() ([]byte, []int) {
	return file_food_service_proto_rawDescGZIP(), []int{0}
}

func (x *Food) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Food) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Food) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Food) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *Food) GetMeal() string {
	if x != nil {
		return x.Meal
	}
	return ""
}

func (x *Food) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Food) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type SideDish struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	CreatedAt   *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt   *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *SideDish) Reset() {
	*x = SideDish{}
	if protoimpl.UnsafeEnabled {
		mi := &file_food_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SideDish) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SideDish) ProtoMessage() {}

func (x *SideDish) ProtoReflect() protoreflect.Message {
	mi := &file_food_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SideDish.ProtoReflect.Descriptor instead.
func (*SideDish) Descriptor() ([]byte, []int) {
	return file_food_service_proto_rawDescGZIP(), []int{1}
}

func (x *SideDish) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SideDish) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SideDish) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *SideDish) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *SideDish) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type FoodIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FoodName string `protobuf:"bytes,1,opt,name=food_name,json=foodName,proto3" json:"food_name,omitempty"`
}

func (x *FoodIdRequest) Reset() {
	*x = FoodIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_food_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FoodIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FoodIdRequest) ProtoMessage() {}

func (x *FoodIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_food_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FoodIdRequest.ProtoReflect.Descriptor instead.
func (*FoodIdRequest) Descriptor() ([]byte, []int) {
	return file_food_service_proto_rawDescGZIP(), []int{2}
}

func (x *FoodIdRequest) GetFoodName() string {
	if x != nil {
		return x.FoodName
	}
	return ""
}

type SideDishIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SideDishName string `protobuf:"bytes,1,opt,name=side_dish_name,json=sideDishName,proto3" json:"side_dish_name,omitempty"`
}

func (x *SideDishIdRequest) Reset() {
	*x = SideDishIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_food_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SideDishIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SideDishIdRequest) ProtoMessage() {}

func (x *SideDishIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_food_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SideDishIdRequest.ProtoReflect.Descriptor instead.
func (*SideDishIdRequest) Descriptor() ([]byte, []int) {
	return file_food_service_proto_rawDescGZIP(), []int{3}
}

func (x *SideDishIdRequest) GetSideDishName() string {
	if x != nil {
		return x.SideDishName
	}
	return ""
}

var File_food_service_proto protoreflect.FileDescriptor

var file_food_service_proto_rawDesc = []byte{
	0x0a, 0x12, 0x66, 0x6f, 0x6f, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf2, 0x01, 0x0a, 0x04, 0x46, 0x6f, 0x6f, 0x64, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x12, 0x12, 0x0a, 0x04, 0x6d, 0x65, 0x61, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6d, 0x65, 0x61, 0x6c, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xc6, 0x01, 0x0a, 0x08, 0x53,
	0x69, 0x64, 0x65, 0x44, 0x69, 0x73, 0x68, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x39, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x22, 0x2c, 0x0a, 0x0d, 0x46, 0x6f, 0x6f, 0x64, 0x49, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x6f, 0x6f, 0x64, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x6f, 0x6f, 0x64, 0x4e, 0x61, 0x6d,
	0x65, 0x22, 0x39, 0x0a, 0x11, 0x53, 0x69, 0x64, 0x65, 0x44, 0x69, 0x73, 0x68, 0x49, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0e, 0x73, 0x69, 0x64, 0x65, 0x5f, 0x64,
	0x69, 0x73, 0x68, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x73, 0x69, 0x64, 0x65, 0x44, 0x69, 0x73, 0x68, 0x4e, 0x61, 0x6d, 0x65, 0x32, 0x77, 0x0a, 0x0b,
	0x46, 0x6f, 0x6f, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2d, 0x0a, 0x14, 0x47,
	0x65, 0x74, 0x46, 0x6f, 0x6f, 0x64, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x42, 0x79, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x0e, 0x2e, 0x46, 0x6f, 0x6f, 0x64, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x05, 0x2e, 0x46, 0x6f, 0x6f, 0x64, 0x12, 0x39, 0x0a, 0x18, 0x47, 0x65,
	0x74, 0x53, 0x69, 0x64, 0x65, 0x44, 0x69, 0x73, 0x68, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73,
	0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x2e, 0x53, 0x69, 0x64, 0x65, 0x44, 0x69, 0x73,
	0x68, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x09, 0x2e, 0x53, 0x69, 0x64,
	0x65, 0x44, 0x69, 0x73, 0x68, 0x42, 0x46, 0x5a, 0x44, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x69, 0x6d, 0x62, 0x6f, 0x6c, 0x69, 0x73, 0x6d, 0x2f, 0x73, 0x6f,
	0x66, 0x74, 0x77, 0x61, 0x72, 0x65, 0x2d, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e,
	0x74, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x2d, 0x65, 0x6e, 0x64, 0x2f, 0x66, 0x6f, 0x6f, 0x64, 0x2d,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_food_service_proto_rawDescOnce sync.Once
	file_food_service_proto_rawDescData = file_food_service_proto_rawDesc
)

func file_food_service_proto_rawDescGZIP() []byte {
	file_food_service_proto_rawDescOnce.Do(func() {
		file_food_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_food_service_proto_rawDescData)
	})
	return file_food_service_proto_rawDescData
}

var file_food_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_food_service_proto_goTypes = []interface{}{
	(*Food)(nil),                  // 0: Food
	(*SideDish)(nil),              // 1: SideDish
	(*FoodIdRequest)(nil),         // 2: FoodIdRequest
	(*SideDishIdRequest)(nil),     // 3: SideDishIdRequest
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_food_service_proto_depIdxs = []int32{
	4, // 0: Food.created_at:type_name -> google.protobuf.Timestamp
	4, // 1: Food.updated_at:type_name -> google.protobuf.Timestamp
	4, // 2: SideDish.created_at:type_name -> google.protobuf.Timestamp
	4, // 3: SideDish.updated_at:type_name -> google.protobuf.Timestamp
	2, // 4: FoodService.GetFoodDetailsByName:input_type -> FoodIdRequest
	3, // 5: FoodService.GetSideDishDetailsByName:input_type -> SideDishIdRequest
	0, // 6: FoodService.GetFoodDetailsByName:output_type -> Food
	1, // 7: FoodService.GetSideDishDetailsByName:output_type -> SideDish
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_food_service_proto_init() }
func file_food_service_proto_init() {
	if File_food_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_food_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Food); i {
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
		file_food_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SideDish); i {
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
		file_food_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FoodIdRequest); i {
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
		file_food_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SideDishIdRequest); i {
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
			RawDescriptor: file_food_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_food_service_proto_goTypes,
		DependencyIndexes: file_food_service_proto_depIdxs,
		MessageInfos:      file_food_service_proto_msgTypes,
	}.Build()
	File_food_service_proto = out.File
	file_food_service_proto_rawDesc = nil
	file_food_service_proto_goTypes = nil
	file_food_service_proto_depIdxs = nil
}
