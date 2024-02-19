// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.2
// source: src/api/v1/inventory/inventory.proto

package inventory

import (
	character "github.com/VoidMesh/backend/src/api/v1/character"
	resource "github.com/VoidMesh/backend/src/api/v1/resource"
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

type UUID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *UUID) Reset() {
	*x = UUID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_api_v1_inventory_inventory_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UUID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UUID) ProtoMessage() {}

func (x *UUID) ProtoReflect() protoreflect.Message {
	mi := &file_src_api_v1_inventory_inventory_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UUID.ProtoReflect.Descriptor instead.
func (*UUID) Descriptor() ([]byte, []int) {
	return file_src_api_v1_inventory_inventory_proto_rawDescGZIP(), []int{0}
}

func (x *UUID) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type Inventory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          *UUID           `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CharacterId *character.UUID `protobuf:"bytes,2,opt,name=character_id,json=characterId,proto3" json:"character_id,omitempty"`
	Slots       []*Slot         `protobuf:"bytes,3,rep,name=slots,proto3" json:"slots,omitempty"`
}

func (x *Inventory) Reset() {
	*x = Inventory{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_api_v1_inventory_inventory_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Inventory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Inventory) ProtoMessage() {}

func (x *Inventory) ProtoReflect() protoreflect.Message {
	mi := &file_src_api_v1_inventory_inventory_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Inventory.ProtoReflect.Descriptor instead.
func (*Inventory) Descriptor() ([]byte, []int) {
	return file_src_api_v1_inventory_inventory_proto_rawDescGZIP(), []int{1}
}

func (x *Inventory) GetId() *UUID {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *Inventory) GetCharacterId() *character.UUID {
	if x != nil {
		return x.CharacterId
	}
	return nil
}

func (x *Inventory) GetSlots() []*Slot {
	if x != nil {
		return x.Slots
	}
	return nil
}

type Slot struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Resource *resource.Resource `protobuf:"bytes,1,opt,name=resource,proto3" json:"resource,omitempty"`
	Quantity int32              `protobuf:"varint,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
}

func (x *Slot) Reset() {
	*x = Slot{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_api_v1_inventory_inventory_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Slot) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Slot) ProtoMessage() {}

func (x *Slot) ProtoReflect() protoreflect.Message {
	mi := &file_src_api_v1_inventory_inventory_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Slot.ProtoReflect.Descriptor instead.
func (*Slot) Descriptor() ([]byte, []int) {
	return file_src_api_v1_inventory_inventory_proto_rawDescGZIP(), []int{2}
}

func (x *Slot) GetResource() *resource.Resource {
	if x != nil {
		return x.Resource
	}
	return nil
}

func (x *Slot) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type ReadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          *UUID           `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CharacterId *character.UUID `protobuf:"bytes,2,opt,name=character_id,json=characterId,proto3" json:"character_id,omitempty"`
}

func (x *ReadRequest) Reset() {
	*x = ReadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_api_v1_inventory_inventory_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadRequest) ProtoMessage() {}

func (x *ReadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_src_api_v1_inventory_inventory_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadRequest.ProtoReflect.Descriptor instead.
func (*ReadRequest) Descriptor() ([]byte, []int) {
	return file_src_api_v1_inventory_inventory_proto_rawDescGZIP(), []int{3}
}

func (x *ReadRequest) GetId() *UUID {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *ReadRequest) GetCharacterId() *character.UUID {
	if x != nil {
		return x.CharacterId
	}
	return nil
}

type ReadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Inventory *Inventory `protobuf:"bytes,1,opt,name=inventory,proto3" json:"inventory,omitempty"`
}

func (x *ReadResponse) Reset() {
	*x = ReadResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_api_v1_inventory_inventory_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadResponse) ProtoMessage() {}

func (x *ReadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_src_api_v1_inventory_inventory_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadResponse.ProtoReflect.Descriptor instead.
func (*ReadResponse) Descriptor() ([]byte, []int) {
	return file_src_api_v1_inventory_inventory_proto_rawDescGZIP(), []int{4}
}

func (x *ReadResponse) GetInventory() *Inventory {
	if x != nil {
		return x.Inventory
	}
	return nil
}

var File_src_api_v1_inventory_inventory_proto protoreflect.FileDescriptor

var file_src_api_v1_inventory_inventory_proto_rawDesc = []byte{
	0x0a, 0x24, 0x73, 0x72, 0x63, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e, 0x76,
	0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72,
	0x79, 0x1a, 0x24, 0x73, 0x72, 0x63, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x68,
	0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x2f, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x73, 0x72, 0x63, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1c, 0x0a, 0x04, 0x55,
	0x55, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x87, 0x01, 0x0a, 0x09, 0x49, 0x6e,
	0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x1f, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e,
	0x55, 0x55, 0x49, 0x44, 0x52, 0x02, 0x69, 0x64, 0x12, 0x32, 0x0a, 0x0c, 0x63, 0x68, 0x61, 0x72,
	0x61, 0x63, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f,
	0x2e, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x2e, 0x55, 0x55, 0x49, 0x44, 0x52,
	0x0b, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x05,
	0x73, 0x6c, 0x6f, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x69, 0x6e,
	0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x53, 0x6c, 0x6f, 0x74, 0x52, 0x05, 0x73, 0x6c,
	0x6f, 0x74, 0x73, 0x22, 0x52, 0x0a, 0x04, 0x53, 0x6c, 0x6f, 0x74, 0x12, 0x2e, 0x0a, 0x08, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x52, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x71,
	0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x71,
	0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x62, 0x0a, 0x0b, 0x52, 0x65, 0x61, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x55,
	0x55, 0x49, 0x44, 0x52, 0x02, 0x69, 0x64, 0x12, 0x32, 0x0a, 0x0c, 0x63, 0x68, 0x61, 0x72, 0x61,
	0x63, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e,
	0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x2e, 0x55, 0x55, 0x49, 0x44, 0x52, 0x0b,
	0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x49, 0x64, 0x22, 0x42, 0x0a, 0x0c, 0x52,
	0x65, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x09, 0x69,
	0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x49, 0x6e, 0x76, 0x65, 0x6e,
	0x74, 0x6f, 0x72, 0x79, 0x52, 0x09, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x32,
	0x49, 0x0a, 0x0c, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x53, 0x76, 0x63, 0x12,
	0x39, 0x0a, 0x04, 0x52, 0x65, 0x61, 0x64, 0x12, 0x16, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74,
	0x6f, 0x72, 0x79, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x17, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x52, 0x65, 0x61, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x32, 0x5a, 0x30, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x56, 0x6f, 0x69, 0x64, 0x4d, 0x65, 0x73,
	0x68, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_src_api_v1_inventory_inventory_proto_rawDescOnce sync.Once
	file_src_api_v1_inventory_inventory_proto_rawDescData = file_src_api_v1_inventory_inventory_proto_rawDesc
)

func file_src_api_v1_inventory_inventory_proto_rawDescGZIP() []byte {
	file_src_api_v1_inventory_inventory_proto_rawDescOnce.Do(func() {
		file_src_api_v1_inventory_inventory_proto_rawDescData = protoimpl.X.CompressGZIP(file_src_api_v1_inventory_inventory_proto_rawDescData)
	})
	return file_src_api_v1_inventory_inventory_proto_rawDescData
}

var file_src_api_v1_inventory_inventory_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_src_api_v1_inventory_inventory_proto_goTypes = []interface{}{
	(*UUID)(nil),              // 0: inventory.UUID
	(*Inventory)(nil),         // 1: inventory.Inventory
	(*Slot)(nil),              // 2: inventory.Slot
	(*ReadRequest)(nil),       // 3: inventory.ReadRequest
	(*ReadResponse)(nil),      // 4: inventory.ReadResponse
	(*character.UUID)(nil),    // 5: character.UUID
	(*resource.Resource)(nil), // 6: resource.Resource
}
var file_src_api_v1_inventory_inventory_proto_depIdxs = []int32{
	0, // 0: inventory.Inventory.id:type_name -> inventory.UUID
	5, // 1: inventory.Inventory.character_id:type_name -> character.UUID
	2, // 2: inventory.Inventory.slots:type_name -> inventory.Slot
	6, // 3: inventory.Slot.resource:type_name -> resource.Resource
	0, // 4: inventory.ReadRequest.id:type_name -> inventory.UUID
	5, // 5: inventory.ReadRequest.character_id:type_name -> character.UUID
	1, // 6: inventory.ReadResponse.inventory:type_name -> inventory.Inventory
	3, // 7: inventory.InventorySvc.Read:input_type -> inventory.ReadRequest
	4, // 8: inventory.InventorySvc.Read:output_type -> inventory.ReadResponse
	8, // [8:9] is the sub-list for method output_type
	7, // [7:8] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_src_api_v1_inventory_inventory_proto_init() }
func file_src_api_v1_inventory_inventory_proto_init() {
	if File_src_api_v1_inventory_inventory_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_src_api_v1_inventory_inventory_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UUID); i {
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
		file_src_api_v1_inventory_inventory_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Inventory); i {
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
		file_src_api_v1_inventory_inventory_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Slot); i {
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
		file_src_api_v1_inventory_inventory_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadRequest); i {
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
		file_src_api_v1_inventory_inventory_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadResponse); i {
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
			RawDescriptor: file_src_api_v1_inventory_inventory_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_src_api_v1_inventory_inventory_proto_goTypes,
		DependencyIndexes: file_src_api_v1_inventory_inventory_proto_depIdxs,
		MessageInfos:      file_src_api_v1_inventory_inventory_proto_msgTypes,
	}.Build()
	File_src_api_v1_inventory_inventory_proto = out.File
	file_src_api_v1_inventory_inventory_proto_rawDesc = nil
	file_src_api_v1_inventory_inventory_proto_goTypes = nil
	file_src_api_v1_inventory_inventory_proto_depIdxs = nil
}
