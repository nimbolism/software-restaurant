// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.26.1
// source: voucher_service.proto

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

// Define the message types
type OrderHelper struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Username    string                 `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Email       string                 `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	PhoneNumber string                 `protobuf:"bytes,4,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	Foods       []*FoodHelper          `protobuf:"bytes,5,rep,name=foods,proto3" json:"foods,omitempty"`
	SideDishes  []*SideDishHelper      `protobuf:"bytes,6,rep,name=side_dishes,json=sideDishes,proto3" json:"side_dishes,omitempty"`
	Paid        bool                   `protobuf:"varint,7,opt,name=paid,proto3" json:"paid,omitempty"`
	CreatedAt   *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt   *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *OrderHelper) Reset() {
	*x = OrderHelper{}
	if protoimpl.UnsafeEnabled {
		mi := &file_voucher_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderHelper) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderHelper) ProtoMessage() {}

func (x *OrderHelper) ProtoReflect() protoreflect.Message {
	mi := &file_voucher_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderHelper.ProtoReflect.Descriptor instead.
func (*OrderHelper) Descriptor() ([]byte, []int) {
	return file_voucher_service_proto_rawDescGZIP(), []int{0}
}

func (x *OrderHelper) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *OrderHelper) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *OrderHelper) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *OrderHelper) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *OrderHelper) GetFoods() []*FoodHelper {
	if x != nil {
		return x.Foods
	}
	return nil
}

func (x *OrderHelper) GetSideDishes() []*SideDishHelper {
	if x != nil {
		return x.SideDishes
	}
	return nil
}

func (x *OrderHelper) GetPaid() bool {
	if x != nil {
		return x.Paid
	}
	return false
}

func (x *OrderHelper) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *OrderHelper) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type FoodHelper struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name         string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description  string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	CategoryName string `protobuf:"bytes,3,opt,name=category_name,json=categoryName,proto3" json:"category_name,omitempty"`
	MealName     string `protobuf:"bytes,4,opt,name=meal_name,json=mealName,proto3" json:"meal_name,omitempty"`
}

func (x *FoodHelper) Reset() {
	*x = FoodHelper{}
	if protoimpl.UnsafeEnabled {
		mi := &file_voucher_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FoodHelper) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FoodHelper) ProtoMessage() {}

func (x *FoodHelper) ProtoReflect() protoreflect.Message {
	mi := &file_voucher_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FoodHelper.ProtoReflect.Descriptor instead.
func (*FoodHelper) Descriptor() ([]byte, []int) {
	return file_voucher_service_proto_rawDescGZIP(), []int{1}
}

func (x *FoodHelper) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FoodHelper) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *FoodHelper) GetCategoryName() string {
	if x != nil {
		return x.CategoryName
	}
	return ""
}

func (x *FoodHelper) GetMealName() string {
	if x != nil {
		return x.MealName
	}
	return ""
}

type SideDishHelper struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *SideDishHelper) Reset() {
	*x = SideDishHelper{}
	if protoimpl.UnsafeEnabled {
		mi := &file_voucher_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SideDishHelper) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SideDishHelper) ProtoMessage() {}

func (x *SideDishHelper) ProtoReflect() protoreflect.Message {
	mi := &file_voucher_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SideDishHelper.ProtoReflect.Descriptor instead.
func (*SideDishHelper) Descriptor() ([]byte, []int) {
	return file_voucher_service_proto_rawDescGZIP(), []int{2}
}

func (x *SideDishHelper) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SideDishHelper) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

// Define request and response messages for storing order details
type StoreOrderDetailsRequestHelper struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Order *OrderHelper `protobuf:"bytes,1,opt,name=order,proto3" json:"order,omitempty"`
}

func (x *StoreOrderDetailsRequestHelper) Reset() {
	*x = StoreOrderDetailsRequestHelper{}
	if protoimpl.UnsafeEnabled {
		mi := &file_voucher_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoreOrderDetailsRequestHelper) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoreOrderDetailsRequestHelper) ProtoMessage() {}

func (x *StoreOrderDetailsRequestHelper) ProtoReflect() protoreflect.Message {
	mi := &file_voucher_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoreOrderDetailsRequestHelper.ProtoReflect.Descriptor instead.
func (*StoreOrderDetailsRequestHelper) Descriptor() ([]byte, []int) {
	return file_voucher_service_proto_rawDescGZIP(), []int{3}
}

func (x *StoreOrderDetailsRequestHelper) GetOrder() *OrderHelper {
	if x != nil {
		return x.Order
	}
	return nil
}

type StoreOrderDetailsResponseHelper struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *StoreOrderDetailsResponseHelper) Reset() {
	*x = StoreOrderDetailsResponseHelper{}
	if protoimpl.UnsafeEnabled {
		mi := &file_voucher_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoreOrderDetailsResponseHelper) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoreOrderDetailsResponseHelper) ProtoMessage() {}

func (x *StoreOrderDetailsResponseHelper) ProtoReflect() protoreflect.Message {
	mi := &file_voucher_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoreOrderDetailsResponseHelper.ProtoReflect.Descriptor instead.
func (*StoreOrderDetailsResponseHelper) Descriptor() ([]byte, []int) {
	return file_voucher_service_proto_rawDescGZIP(), []int{4}
}

func (x *StoreOrderDetailsResponseHelper) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

// Define request and response messages for retrieving all orders
type GetAllOrdersRequestHelper struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JwtToken string `protobuf:"bytes,1,opt,name=jwt_token,json=jwtToken,proto3" json:"jwt_token,omitempty"`
}

func (x *GetAllOrdersRequestHelper) Reset() {
	*x = GetAllOrdersRequestHelper{}
	if protoimpl.UnsafeEnabled {
		mi := &file_voucher_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllOrdersRequestHelper) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllOrdersRequestHelper) ProtoMessage() {}

func (x *GetAllOrdersRequestHelper) ProtoReflect() protoreflect.Message {
	mi := &file_voucher_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllOrdersRequestHelper.ProtoReflect.Descriptor instead.
func (*GetAllOrdersRequestHelper) Descriptor() ([]byte, []int) {
	return file_voucher_service_proto_rawDescGZIP(), []int{5}
}

func (x *GetAllOrdersRequestHelper) GetJwtToken() string {
	if x != nil {
		return x.JwtToken
	}
	return ""
}

type GetAllOrdersResponseHelper struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Orders []*OrderHelper `protobuf:"bytes,1,rep,name=orders,proto3" json:"orders,omitempty"`
}

func (x *GetAllOrdersResponseHelper) Reset() {
	*x = GetAllOrdersResponseHelper{}
	if protoimpl.UnsafeEnabled {
		mi := &file_voucher_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllOrdersResponseHelper) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllOrdersResponseHelper) ProtoMessage() {}

func (x *GetAllOrdersResponseHelper) ProtoReflect() protoreflect.Message {
	mi := &file_voucher_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllOrdersResponseHelper.ProtoReflect.Descriptor instead.
func (*GetAllOrdersResponseHelper) Descriptor() ([]byte, []int) {
	return file_voucher_service_proto_rawDescGZIP(), []int{6}
}

func (x *GetAllOrdersResponseHelper) GetOrders() []*OrderHelper {
	if x != nil {
		return x.Orders
	}
	return nil
}

var File_voucher_service_proto protoreflect.FileDescriptor

var file_voucher_service_proto_rawDesc = []byte{
	0x0a, 0x15, 0x76, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd1, 0x02, 0x0a, 0x0b, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x48, 0x65, 0x6c, 0x70, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x21, 0x0a,
	0x05, 0x66, 0x6f, 0x6f, 0x64, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x46,
	0x6f, 0x6f, 0x64, 0x48, 0x65, 0x6c, 0x70, 0x65, 0x72, 0x52, 0x05, 0x66, 0x6f, 0x6f, 0x64, 0x73,
	0x12, 0x30, 0x0a, 0x0b, 0x73, 0x69, 0x64, 0x65, 0x5f, 0x64, 0x69, 0x73, 0x68, 0x65, 0x73, 0x18,
	0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x53, 0x69, 0x64, 0x65, 0x44, 0x69, 0x73, 0x68,
	0x48, 0x65, 0x6c, 0x70, 0x65, 0x72, 0x52, 0x0a, 0x73, 0x69, 0x64, 0x65, 0x44, 0x69, 0x73, 0x68,
	0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x04, 0x70, 0x61, 0x69, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x84, 0x01, 0x0a,
	0x0a, 0x46, 0x6f, 0x6f, 0x64, 0x48, 0x65, 0x6c, 0x70, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65, 0x61, 0x6c, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x65, 0x61, 0x6c, 0x4e,
	0x61, 0x6d, 0x65, 0x22, 0x46, 0x0a, 0x0e, 0x53, 0x69, 0x64, 0x65, 0x44, 0x69, 0x73, 0x68, 0x48,
	0x65, 0x6c, 0x70, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x44, 0x0a, 0x1e, 0x53,
	0x74, 0x6f, 0x72, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x65, 0x6c, 0x70, 0x65, 0x72, 0x12, 0x22, 0x0a,
	0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x48, 0x65, 0x6c, 0x70, 0x65, 0x72, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x22, 0x3b, 0x0a, 0x1f, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x65,
	0x6c, 0x70, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x38,
	0x0a, 0x19, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x65, 0x6c, 0x70, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x6a,
	0x77, 0x74, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6a, 0x77, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x42, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x48, 0x65, 0x6c, 0x70, 0x65, 0x72, 0x12, 0x24, 0x0a, 0x06, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x48, 0x65,
	0x6c, 0x70, 0x65, 0x72, 0x52, 0x06, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x32, 0xb1, 0x01, 0x0a,
	0x0e, 0x56, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x56, 0x0a, 0x11, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x73, 0x12, 0x1f, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48,
	0x65, 0x6c, 0x70, 0x65, 0x72, 0x1a, 0x20, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x48, 0x65, 0x6c, 0x70, 0x65, 0x72, 0x12, 0x47, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x12, 0x1a, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x65, 0x6c,
	0x70, 0x65, 0x72, 0x1a, 0x1b, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x65, 0x6c, 0x70, 0x65, 0x72,
	0x42, 0x49, 0x5a, 0x47, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e,
	0x69, 0x6d, 0x62, 0x6f, 0x6c, 0x69, 0x73, 0x6d, 0x2f, 0x73, 0x6f, 0x66, 0x74, 0x77, 0x61, 0x72,
	0x65, 0x2d, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x2f, 0x62, 0x61, 0x63,
	0x6b, 0x2d, 0x65, 0x6e, 0x64, 0x2f, 0x76, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x72, 0x2d, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_voucher_service_proto_rawDescOnce sync.Once
	file_voucher_service_proto_rawDescData = file_voucher_service_proto_rawDesc
)

func file_voucher_service_proto_rawDescGZIP() []byte {
	file_voucher_service_proto_rawDescOnce.Do(func() {
		file_voucher_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_voucher_service_proto_rawDescData)
	})
	return file_voucher_service_proto_rawDescData
}

var file_voucher_service_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_voucher_service_proto_goTypes = []interface{}{
	(*OrderHelper)(nil),                     // 0: OrderHelper
	(*FoodHelper)(nil),                      // 1: FoodHelper
	(*SideDishHelper)(nil),                  // 2: SideDishHelper
	(*StoreOrderDetailsRequestHelper)(nil),  // 3: StoreOrderDetailsRequestHelper
	(*StoreOrderDetailsResponseHelper)(nil), // 4: StoreOrderDetailsResponseHelper
	(*GetAllOrdersRequestHelper)(nil),       // 5: GetAllOrdersRequestHelper
	(*GetAllOrdersResponseHelper)(nil),      // 6: GetAllOrdersResponseHelper
	(*timestamppb.Timestamp)(nil),           // 7: google.protobuf.Timestamp
}
var file_voucher_service_proto_depIdxs = []int32{
	1, // 0: OrderHelper.foods:type_name -> FoodHelper
	2, // 1: OrderHelper.side_dishes:type_name -> SideDishHelper
	7, // 2: OrderHelper.created_at:type_name -> google.protobuf.Timestamp
	7, // 3: OrderHelper.updated_at:type_name -> google.protobuf.Timestamp
	0, // 4: StoreOrderDetailsRequestHelper.order:type_name -> OrderHelper
	0, // 5: GetAllOrdersResponseHelper.orders:type_name -> OrderHelper
	3, // 6: VoucherService.StoreOrderDetails:input_type -> StoreOrderDetailsRequestHelper
	5, // 7: VoucherService.GetAllOrders:input_type -> GetAllOrdersRequestHelper
	4, // 8: VoucherService.StoreOrderDetails:output_type -> StoreOrderDetailsResponseHelper
	6, // 9: VoucherService.GetAllOrders:output_type -> GetAllOrdersResponseHelper
	8, // [8:10] is the sub-list for method output_type
	6, // [6:8] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_voucher_service_proto_init() }
func file_voucher_service_proto_init() {
	if File_voucher_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_voucher_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderHelper); i {
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
		file_voucher_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FoodHelper); i {
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
		file_voucher_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SideDishHelper); i {
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
		file_voucher_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoreOrderDetailsRequestHelper); i {
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
		file_voucher_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoreOrderDetailsResponseHelper); i {
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
		file_voucher_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllOrdersRequestHelper); i {
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
		file_voucher_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllOrdersResponseHelper); i {
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
			RawDescriptor: file_voucher_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_voucher_service_proto_goTypes,
		DependencyIndexes: file_voucher_service_proto_depIdxs,
		MessageInfos:      file_voucher_service_proto_msgTypes,
	}.Build()
	File_voucher_service_proto = out.File
	file_voucher_service_proto_rawDesc = nil
	file_voucher_service_proto_goTypes = nil
	file_voucher_service_proto_depIdxs = nil
}
