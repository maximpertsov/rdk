// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: proto/api/component/base/v1/base.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type MoveStraightRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of a base
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Desired travel distance in millimeters
	DistanceMm int64 `protobuf:"varint,2,opt,name=distance_mm,json=distanceMm,proto3" json:"distance_mm,omitempty"`
	// Desired travel velocity in millimeters/second
	MmPerSec float64 `protobuf:"fixed64,3,opt,name=mm_per_sec,json=mmPerSec,proto3" json:"mm_per_sec,omitempty"`
	// Whether the movement should block all other
	// movement commands to base until this movement is complete
	Block bool `protobuf:"varint,4,opt,name=block,proto3" json:"block,omitempty"`
}

func (x *MoveStraightRequest) Reset() {
	*x = MoveStraightRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_api_component_base_v1_base_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MoveStraightRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MoveStraightRequest) ProtoMessage() {}

func (x *MoveStraightRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_component_base_v1_base_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MoveStraightRequest.ProtoReflect.Descriptor instead.
func (*MoveStraightRequest) Descriptor() ([]byte, []int) {
	return file_proto_api_component_base_v1_base_proto_rawDescGZIP(), []int{0}
}

func (x *MoveStraightRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MoveStraightRequest) GetDistanceMm() int64 {
	if x != nil {
		return x.DistanceMm
	}
	return 0
}

func (x *MoveStraightRequest) GetMmPerSec() float64 {
	if x != nil {
		return x.MmPerSec
	}
	return 0
}

func (x *MoveStraightRequest) GetBlock() bool {
	if x != nil {
		return x.Block
	}
	return false
}

type MoveStraightResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MoveStraightResponse) Reset() {
	*x = MoveStraightResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_api_component_base_v1_base_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MoveStraightResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MoveStraightResponse) ProtoMessage() {}

func (x *MoveStraightResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_component_base_v1_base_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MoveStraightResponse.ProtoReflect.Descriptor instead.
func (*MoveStraightResponse) Descriptor() ([]byte, []int) {
	return file_proto_api_component_base_v1_base_proto_rawDescGZIP(), []int{1}
}

type MoveArcRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of a base
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Desired travel distance in millimeters
	DistanceMm int64 `protobuf:"varint,2,opt,name=distance_mm,json=distanceMm,proto3" json:"distance_mm,omitempty"`
	// Desired speed in millimeters per second
	MmPerSec float64 `protobuf:"fixed64,3,opt,name=mm_per_sec,json=mmPerSec,proto3" json:"mm_per_sec,omitempty"`
	// Desired angle in degrees
	AngleDeg float64 `protobuf:"fixed64,4,opt,name=angle_deg,json=angleDeg,proto3" json:"angle_deg,omitempty"`
	// Whether the movement should block all other
	// movement commands to base until this movement is complete
	Block bool `protobuf:"varint,5,opt,name=block,proto3" json:"block,omitempty"`
}

func (x *MoveArcRequest) Reset() {
	*x = MoveArcRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_api_component_base_v1_base_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MoveArcRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MoveArcRequest) ProtoMessage() {}

func (x *MoveArcRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_component_base_v1_base_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MoveArcRequest.ProtoReflect.Descriptor instead.
func (*MoveArcRequest) Descriptor() ([]byte, []int) {
	return file_proto_api_component_base_v1_base_proto_rawDescGZIP(), []int{2}
}

func (x *MoveArcRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MoveArcRequest) GetDistanceMm() int64 {
	if x != nil {
		return x.DistanceMm
	}
	return 0
}

func (x *MoveArcRequest) GetMmPerSec() float64 {
	if x != nil {
		return x.MmPerSec
	}
	return 0
}

func (x *MoveArcRequest) GetAngleDeg() float64 {
	if x != nil {
		return x.AngleDeg
	}
	return 0
}

func (x *MoveArcRequest) GetBlock() bool {
	if x != nil {
		return x.Block
	}
	return false
}

type MoveArcResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MoveArcResponse) Reset() {
	*x = MoveArcResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_api_component_base_v1_base_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MoveArcResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MoveArcResponse) ProtoMessage() {}

func (x *MoveArcResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_component_base_v1_base_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MoveArcResponse.ProtoReflect.Descriptor instead.
func (*MoveArcResponse) Descriptor() ([]byte, []int) {
	return file_proto_api_component_base_v1_base_proto_rawDescGZIP(), []int{3}
}

type SpinRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of a base
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Desired angle
	AngleDeg float64 `protobuf:"fixed64,2,opt,name=angle_deg,json=angleDeg,proto3" json:"angle_deg,omitempty"`
	// Desired angular velocity
	DegsPerSec float64 `protobuf:"fixed64,3,opt,name=degs_per_sec,json=degsPerSec,proto3" json:"degs_per_sec,omitempty"`
	// Whether the movement should block all other
	// movement commands to base until this movement is complete
	Block bool `protobuf:"varint,4,opt,name=block,proto3" json:"block,omitempty"`
}

func (x *SpinRequest) Reset() {
	*x = SpinRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_api_component_base_v1_base_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpinRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpinRequest) ProtoMessage() {}

func (x *SpinRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_component_base_v1_base_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpinRequest.ProtoReflect.Descriptor instead.
func (*SpinRequest) Descriptor() ([]byte, []int) {
	return file_proto_api_component_base_v1_base_proto_rawDescGZIP(), []int{4}
}

func (x *SpinRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SpinRequest) GetAngleDeg() float64 {
	if x != nil {
		return x.AngleDeg
	}
	return 0
}

func (x *SpinRequest) GetDegsPerSec() float64 {
	if x != nil {
		return x.DegsPerSec
	}
	return 0
}

func (x *SpinRequest) GetBlock() bool {
	if x != nil {
		return x.Block
	}
	return false
}

type SpinResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SpinResponse) Reset() {
	*x = SpinResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_api_component_base_v1_base_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpinResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpinResponse) ProtoMessage() {}

func (x *SpinResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_component_base_v1_base_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpinResponse.ProtoReflect.Descriptor instead.
func (*SpinResponse) Descriptor() ([]byte, []int) {
	return file_proto_api_component_base_v1_base_proto_rawDescGZIP(), []int{5}
}

type StopRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of a base
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *StopRequest) Reset() {
	*x = StopRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_api_component_base_v1_base_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StopRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StopRequest) ProtoMessage() {}

func (x *StopRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_component_base_v1_base_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StopRequest.ProtoReflect.Descriptor instead.
func (*StopRequest) Descriptor() ([]byte, []int) {
	return file_proto_api_component_base_v1_base_proto_rawDescGZIP(), []int{6}
}

func (x *StopRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type StopResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StopResponse) Reset() {
	*x = StopResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_api_component_base_v1_base_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StopResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StopResponse) ProtoMessage() {}

func (x *StopResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_component_base_v1_base_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StopResponse.ProtoReflect.Descriptor instead.
func (*StopResponse) Descriptor() ([]byte, []int) {
	return file_proto_api_component_base_v1_base_proto_rawDescGZIP(), []int{7}
}

var File_proto_api_component_base_v1_base_proto protoreflect.FileDescriptor

var file_proto_api_component_base_v1_base_proto_rawDesc = []byte{
	0x0a, 0x26, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x70,
	0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x61,
	0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x62, 0x61,
	0x73, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x7e, 0x0a, 0x13, 0x4d, 0x6f, 0x76, 0x65, 0x53, 0x74, 0x72, 0x61, 0x69,
	0x67, 0x68, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1f,
	0x0a, 0x0b, 0x64, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x6d, 0x6d, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0a, 0x64, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x4d, 0x6d, 0x12,
	0x1c, 0x0a, 0x0a, 0x6d, 0x6d, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x63, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x08, 0x6d, 0x6d, 0x50, 0x65, 0x72, 0x53, 0x65, 0x63, 0x12, 0x14, 0x0a,
	0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x22, 0x16, 0x0a, 0x14, 0x4d, 0x6f, 0x76, 0x65, 0x53, 0x74, 0x72, 0x61, 0x69,
	0x67, 0x68, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x96, 0x01, 0x0a, 0x0e,
	0x4d, 0x6f, 0x76, 0x65, 0x41, 0x72, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x6d,
	0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x64, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x4d, 0x6d, 0x12, 0x1c, 0x0a, 0x0a, 0x6d, 0x6d, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x73, 0x65,
	0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x6d, 0x6d, 0x50, 0x65, 0x72, 0x53, 0x65,
	0x63, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x6e, 0x67, 0x6c, 0x65, 0x5f, 0x64, 0x65, 0x67, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x61, 0x6e, 0x67, 0x6c, 0x65, 0x44, 0x65, 0x67, 0x12, 0x14,
	0x0a, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x22, 0x11, 0x0a, 0x0f, 0x4d, 0x6f, 0x76, 0x65, 0x41, 0x72, 0x63, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x76, 0x0a, 0x0b, 0x53, 0x70, 0x69, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x6e,
	0x67, 0x6c, 0x65, 0x5f, 0x64, 0x65, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x61,
	0x6e, 0x67, 0x6c, 0x65, 0x44, 0x65, 0x67, 0x12, 0x20, 0x0a, 0x0c, 0x64, 0x65, 0x67, 0x73, 0x5f,
	0x70, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x64,
	0x65, 0x67, 0x73, 0x50, 0x65, 0x72, 0x53, 0x65, 0x63, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x22,
	0x0e, 0x0a, 0x0c, 0x53, 0x70, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x21, 0x0a, 0x0b, 0x53, 0x74, 0x6f, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x0e, 0x0a, 0x0c, 0x53, 0x74, 0x6f, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x32, 0xf7, 0x04, 0x0a, 0x0b, 0x42, 0x61, 0x73, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0xad, 0x01, 0x0a, 0x0c, 0x4d, 0x6f, 0x76, 0x65, 0x53, 0x74, 0x72, 0x61, 0x69,
	0x67, 0x68, 0x74, 0x12, 0x30, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x4d, 0x6f, 0x76, 0x65, 0x53, 0x74, 0x72, 0x61, 0x69, 0x67, 0x68, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x62, 0x61, 0x73, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x6f, 0x76, 0x65, 0x53, 0x74, 0x72, 0x61, 0x69, 0x67, 0x68, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x38, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x32,
	0x22, 0x30, 0x2f, 0x76, 0x69, 0x61, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63,
	0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x7b, 0x6e,
	0x61, 0x6d, 0x65, 0x7d, 0x2f, 0x6d, 0x6f, 0x76, 0x65, 0x5f, 0x73, 0x74, 0x72, 0x61, 0x69, 0x67,
	0x68, 0x74, 0x12, 0x99, 0x01, 0x0a, 0x07, 0x4d, 0x6f, 0x76, 0x65, 0x41, 0x72, 0x63, 0x12, 0x2b,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f,
	0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x6f, 0x76,
	0x65, 0x41, 0x72, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e,
	0x74, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x6f, 0x76, 0x65, 0x41, 0x72,
	0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x33, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x2d, 0x22, 0x2b, 0x2f, 0x76, 0x69, 0x61, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f,
	0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x7b,
	0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x2f, 0x6d, 0x6f, 0x76, 0x65, 0x5f, 0x61, 0x72, 0x63, 0x12, 0x8c,
	0x01, 0x0a, 0x04, 0x53, 0x70, 0x69, 0x6e, 0x12, 0x28, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x62, 0x61,
	0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x70, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x29, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f,
	0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x70, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2f, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x29, 0x22, 0x27, 0x2f, 0x76, 0x69, 0x61, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2f, 0x62, 0x61, 0x73,
	0x65, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x2f, 0x73, 0x70, 0x69, 0x6e, 0x12, 0x8c, 0x01,
	0x0a, 0x04, 0x53, 0x74, 0x6f, 0x70, 0x12, 0x28, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x62, 0x61, 0x73,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x6f, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x29, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d,
	0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x74, 0x6f, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2f, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x29, 0x22, 0x27, 0x2f, 0x76, 0x69, 0x61, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2f, 0x62, 0x61, 0x73, 0x65,
	0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x2f, 0x73, 0x74, 0x6f, 0x70, 0x42, 0x57, 0x0a, 0x28,
	0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x69, 0x61, 0x6d, 0x2e, 0x72, 0x64, 0x6b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74,
	0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x5a, 0x2b, 0x67, 0x6f, 0x2e, 0x76, 0x69, 0x61,
	0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x64, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2f, 0x62, 0x61,
	0x73, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_api_component_base_v1_base_proto_rawDescOnce sync.Once
	file_proto_api_component_base_v1_base_proto_rawDescData = file_proto_api_component_base_v1_base_proto_rawDesc
)

func file_proto_api_component_base_v1_base_proto_rawDescGZIP() []byte {
	file_proto_api_component_base_v1_base_proto_rawDescOnce.Do(func() {
		file_proto_api_component_base_v1_base_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_api_component_base_v1_base_proto_rawDescData)
	})
	return file_proto_api_component_base_v1_base_proto_rawDescData
}

var file_proto_api_component_base_v1_base_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_proto_api_component_base_v1_base_proto_goTypes = []interface{}{
	(*MoveStraightRequest)(nil),  // 0: proto.api.component.base.v1.MoveStraightRequest
	(*MoveStraightResponse)(nil), // 1: proto.api.component.base.v1.MoveStraightResponse
	(*MoveArcRequest)(nil),       // 2: proto.api.component.base.v1.MoveArcRequest
	(*MoveArcResponse)(nil),      // 3: proto.api.component.base.v1.MoveArcResponse
	(*SpinRequest)(nil),          // 4: proto.api.component.base.v1.SpinRequest
	(*SpinResponse)(nil),         // 5: proto.api.component.base.v1.SpinResponse
	(*StopRequest)(nil),          // 6: proto.api.component.base.v1.StopRequest
	(*StopResponse)(nil),         // 7: proto.api.component.base.v1.StopResponse
}
var file_proto_api_component_base_v1_base_proto_depIdxs = []int32{
	0, // 0: proto.api.component.base.v1.BaseService.MoveStraight:input_type -> proto.api.component.base.v1.MoveStraightRequest
	2, // 1: proto.api.component.base.v1.BaseService.MoveArc:input_type -> proto.api.component.base.v1.MoveArcRequest
	4, // 2: proto.api.component.base.v1.BaseService.Spin:input_type -> proto.api.component.base.v1.SpinRequest
	6, // 3: proto.api.component.base.v1.BaseService.Stop:input_type -> proto.api.component.base.v1.StopRequest
	1, // 4: proto.api.component.base.v1.BaseService.MoveStraight:output_type -> proto.api.component.base.v1.MoveStraightResponse
	3, // 5: proto.api.component.base.v1.BaseService.MoveArc:output_type -> proto.api.component.base.v1.MoveArcResponse
	5, // 6: proto.api.component.base.v1.BaseService.Spin:output_type -> proto.api.component.base.v1.SpinResponse
	7, // 7: proto.api.component.base.v1.BaseService.Stop:output_type -> proto.api.component.base.v1.StopResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_api_component_base_v1_base_proto_init() }
func file_proto_api_component_base_v1_base_proto_init() {
	if File_proto_api_component_base_v1_base_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_api_component_base_v1_base_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MoveStraightRequest); i {
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
		file_proto_api_component_base_v1_base_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MoveStraightResponse); i {
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
		file_proto_api_component_base_v1_base_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MoveArcRequest); i {
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
		file_proto_api_component_base_v1_base_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MoveArcResponse); i {
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
		file_proto_api_component_base_v1_base_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SpinRequest); i {
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
		file_proto_api_component_base_v1_base_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SpinResponse); i {
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
		file_proto_api_component_base_v1_base_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StopRequest); i {
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
		file_proto_api_component_base_v1_base_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StopResponse); i {
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
			RawDescriptor: file_proto_api_component_base_v1_base_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_api_component_base_v1_base_proto_goTypes,
		DependencyIndexes: file_proto_api_component_base_v1_base_proto_depIdxs,
		MessageInfos:      file_proto_api_component_base_v1_base_proto_msgTypes,
	}.Build()
	File_proto_api_component_base_v1_base_proto = out.File
	file_proto_api_component_base_v1_base_proto_rawDesc = nil
	file_proto_api_component_base_v1_base_proto_goTypes = nil
	file_proto_api_component_base_v1_base_proto_depIdxs = nil
}
