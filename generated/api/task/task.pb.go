// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: api/service/task.proto

package task

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type OrderField int32

const (
	OrderField_FieldUnknown       OrderField = 0
	OrderField_FieldTaskPriority  OrderField = 1
	OrderField_FieldTaskCreatedAt OrderField = 2
)

// Enum value maps for OrderField.
var (
	OrderField_name = map[int32]string{
		0: "FieldUnknown",
		1: "FieldTaskPriority",
		2: "FieldTaskCreatedAt",
	}
	OrderField_value = map[string]int32{
		"FieldUnknown":       0,
		"FieldTaskPriority":  1,
		"FieldTaskCreatedAt": 2,
	}
)

func (x OrderField) Enum() *OrderField {
	p := new(OrderField)
	*p = x
	return p
}

func (x OrderField) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OrderField) Descriptor() protoreflect.EnumDescriptor {
	return file_api_service_task_proto_enumTypes[0].Descriptor()
}

func (OrderField) Type() protoreflect.EnumType {
	return &file_api_service_task_proto_enumTypes[0]
}

func (x OrderField) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OrderField.Descriptor instead.
func (OrderField) EnumDescriptor() ([]byte, []int) {
	return file_api_service_task_proto_rawDescGZIP(), []int{0}
}

type OrderBy int32

const (
	OrderBy_Asc  OrderBy = 0
	OrderBy_Desc OrderBy = 1
)

// Enum value maps for OrderBy.
var (
	OrderBy_name = map[int32]string{
		0: "Asc",
		1: "Desc",
	}
	OrderBy_value = map[string]int32{
		"Asc":  0,
		"Desc": 1,
	}
)

func (x OrderBy) Enum() *OrderBy {
	p := new(OrderBy)
	*p = x
	return p
}

func (x OrderBy) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OrderBy) Descriptor() protoreflect.EnumDescriptor {
	return file_api_service_task_proto_enumTypes[1].Descriptor()
}

func (OrderBy) Type() protoreflect.EnumType {
	return &file_api_service_task_proto_enumTypes[1]
}

func (x OrderBy) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OrderBy.Descriptor instead.
func (OrderBy) EnumDescriptor() ([]byte, []int) {
	return file_api_service_task_proto_rawDescGZIP(), []int{1}
}

type NewTask struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	TaskName      string                 `protobuf:"bytes,2,opt,name=task_name,json=taskName,proto3" json:"task_name,omitempty"`
	TaskDesc      string                 `protobuf:"bytes,3,opt,name=task_desc,json=taskDesc,proto3" json:"task_desc,omitempty"`
	Priority      int64                  `protobuf:"varint,4,opt,name=priority,proto3" json:"priority,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NewTask) Reset() {
	*x = NewTask{}
	mi := &file_api_service_task_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NewTask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewTask) ProtoMessage() {}

func (x *NewTask) ProtoReflect() protoreflect.Message {
	mi := &file_api_service_task_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewTask.ProtoReflect.Descriptor instead.
func (*NewTask) Descriptor() ([]byte, []int) {
	return file_api_service_task_proto_rawDescGZIP(), []int{0}
}

func (x *NewTask) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *NewTask) GetTaskName() string {
	if x != nil {
		return x.TaskName
	}
	return ""
}

func (x *NewTask) GetTaskDesc() string {
	if x != nil {
		return x.TaskDesc
	}
	return ""
}

func (x *NewTask) GetPriority() int64 {
	if x != nil {
		return x.Priority
	}
	return 0
}

type CreateTaskRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Task          *NewTask               `protobuf:"bytes,1,opt,name=task,proto3" json:"task,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateTaskRequest) Reset() {
	*x = CreateTaskRequest{}
	mi := &file_api_service_task_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTaskRequest) ProtoMessage() {}

func (x *CreateTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_service_task_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTaskRequest.ProtoReflect.Descriptor instead.
func (*CreateTaskRequest) Descriptor() ([]byte, []int) {
	return file_api_service_task_proto_rawDescGZIP(), []int{1}
}

func (x *CreateTaskRequest) GetTask() *NewTask {
	if x != nil {
		return x.Task
	}
	return nil
}

type CreateTaskResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateTaskResponse) Reset() {
	*x = CreateTaskResponse{}
	mi := &file_api_service_task_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateTaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTaskResponse) ProtoMessage() {}

func (x *CreateTaskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_service_task_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTaskResponse.ProtoReflect.Descriptor instead.
func (*CreateTaskResponse) Descriptor() ([]byte, []int) {
	return file_api_service_task_proto_rawDescGZIP(), []int{2}
}

type SetTaskStatusRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TaskId        string                 `protobuf:"bytes,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	TaskStatus    string                 `protobuf:"bytes,2,opt,name=task_status,json=taskStatus,proto3" json:"task_status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SetTaskStatusRequest) Reset() {
	*x = SetTaskStatusRequest{}
	mi := &file_api_service_task_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetTaskStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetTaskStatusRequest) ProtoMessage() {}

func (x *SetTaskStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_service_task_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetTaskStatusRequest.ProtoReflect.Descriptor instead.
func (*SetTaskStatusRequest) Descriptor() ([]byte, []int) {
	return file_api_service_task_proto_rawDescGZIP(), []int{3}
}

func (x *SetTaskStatusRequest) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

func (x *SetTaskStatusRequest) GetTaskStatus() string {
	if x != nil {
		return x.TaskStatus
	}
	return ""
}

type SetTaskStatusResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SetTaskStatusResponse) Reset() {
	*x = SetTaskStatusResponse{}
	mi := &file_api_service_task_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetTaskStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetTaskStatusResponse) ProtoMessage() {}

func (x *SetTaskStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_service_task_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetTaskStatusResponse.ProtoReflect.Descriptor instead.
func (*SetTaskStatusResponse) Descriptor() ([]byte, []int) {
	return file_api_service_task_proto_rawDescGZIP(), []int{4}
}

type Task struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TaskId        string                 `protobuf:"bytes,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	UserId        string                 `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	TaskName      string                 `protobuf:"bytes,3,opt,name=task_name,json=taskName,proto3" json:"task_name,omitempty"`
	TaskDesk      string                 `protobuf:"bytes,4,opt,name=task_desk,json=taskDesk,proto3" json:"task_desk,omitempty"`
	TaskStatus    string                 `protobuf:"bytes,5,opt,name=task_status,json=taskStatus,proto3" json:"task_status,omitempty"`
	Priority      int64                  `protobuf:"varint,6,opt,name=priority,proto3" json:"priority,omitempty"`
	CreatedAt     string                 `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Task) Reset() {
	*x = Task{}
	mi := &file_api_service_task_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Task) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Task) ProtoMessage() {}

func (x *Task) ProtoReflect() protoreflect.Message {
	mi := &file_api_service_task_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Task.ProtoReflect.Descriptor instead.
func (*Task) Descriptor() ([]byte, []int) {
	return file_api_service_task_proto_rawDescGZIP(), []int{5}
}

func (x *Task) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

func (x *Task) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Task) GetTaskName() string {
	if x != nil {
		return x.TaskName
	}
	return ""
}

func (x *Task) GetTaskDesk() string {
	if x != nil {
		return x.TaskDesk
	}
	return ""
}

func (x *Task) GetTaskStatus() string {
	if x != nil {
		return x.TaskStatus
	}
	return ""
}

func (x *Task) GetPriority() int64 {
	if x != nil {
		return x.Priority
	}
	return 0
}

func (x *Task) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

type Filter struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	TaskStatus    string                 `protobuf:"bytes,2,opt,name=task_status,json=taskStatus,proto3" json:"task_status,omitempty"`
	Priority      int64                  `protobuf:"varint,3,opt,name=priority,proto3" json:"priority,omitempty"`
	OrderField    OrderField             `protobuf:"varint,4,opt,name=order_field,json=orderField,proto3,enum=OrderField" json:"order_field,omitempty"`
	OrderBy       OrderBy                `protobuf:"varint,5,opt,name=order_by,json=orderBy,proto3,enum=OrderBy" json:"order_by,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Filter) Reset() {
	*x = Filter{}
	mi := &file_api_service_task_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Filter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Filter) ProtoMessage() {}

func (x *Filter) ProtoReflect() protoreflect.Message {
	mi := &file_api_service_task_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Filter.ProtoReflect.Descriptor instead.
func (*Filter) Descriptor() ([]byte, []int) {
	return file_api_service_task_proto_rawDescGZIP(), []int{6}
}

func (x *Filter) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Filter) GetTaskStatus() string {
	if x != nil {
		return x.TaskStatus
	}
	return ""
}

func (x *Filter) GetPriority() int64 {
	if x != nil {
		return x.Priority
	}
	return 0
}

func (x *Filter) GetOrderField() OrderField {
	if x != nil {
		return x.OrderField
	}
	return OrderField_FieldUnknown
}

func (x *Filter) GetOrderBy() OrderBy {
	if x != nil {
		return x.OrderBy
	}
	return OrderBy_Asc
}

type GetTaskByFilterRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Filter        *Filter                `protobuf:"bytes,1,opt,name=filter,proto3" json:"filter,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetTaskByFilterRequest) Reset() {
	*x = GetTaskByFilterRequest{}
	mi := &file_api_service_task_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTaskByFilterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTaskByFilterRequest) ProtoMessage() {}

func (x *GetTaskByFilterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_service_task_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTaskByFilterRequest.ProtoReflect.Descriptor instead.
func (*GetTaskByFilterRequest) Descriptor() ([]byte, []int) {
	return file_api_service_task_proto_rawDescGZIP(), []int{7}
}

func (x *GetTaskByFilterRequest) GetFilter() *Filter {
	if x != nil {
		return x.Filter
	}
	return nil
}

type GetTaskByFilterResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Tasks         []*Task                `protobuf:"bytes,1,rep,name=tasks,proto3" json:"tasks,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetTaskByFilterResponse) Reset() {
	*x = GetTaskByFilterResponse{}
	mi := &file_api_service_task_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTaskByFilterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTaskByFilterResponse) ProtoMessage() {}

func (x *GetTaskByFilterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_service_task_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTaskByFilterResponse.ProtoReflect.Descriptor instead.
func (*GetTaskByFilterResponse) Descriptor() ([]byte, []int) {
	return file_api_service_task_proto_rawDescGZIP(), []int{8}
}

func (x *GetTaskByFilterResponse) GetTasks() []*Task {
	if x != nil {
		return x.Tasks
	}
	return nil
}

var File_api_service_task_proto protoreflect.FileDescriptor

var file_api_service_task_proto_rawDesc = string([]byte{
	0x0a, 0x16, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x74, 0x61,
	0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x78, 0x0a, 0x07, 0x4e, 0x65, 0x77, 0x54,
	0x61, 0x73, 0x6b, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09,
	0x74, 0x61, 0x73, 0x6b, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x74, 0x61, 0x73, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x61, 0x73,
	0x6b, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x61,
	0x73, 0x6b, 0x44, 0x65, 0x73, 0x63, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69,
	0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69,
	0x74, 0x79, 0x22, 0x31, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x4e, 0x65, 0x77, 0x54, 0x61, 0x73, 0x6b, 0x52,
	0x04, 0x74, 0x61, 0x73, 0x6b, 0x22, 0x14, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54,
	0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x50, 0x0a, 0x14, 0x53,
	0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b,
	0x74, 0x61, 0x73, 0x6b, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x74, 0x61, 0x73, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x17, 0x0a,
	0x15, 0x53, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xce, 0x01, 0x0a, 0x04, 0x54, 0x61, 0x73, 0x6b, 0x12,
	0x17, 0x0a, 0x07, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x61, 0x73, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b,
	0x0a, 0x09, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x64, 0x65, 0x73, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x74, 0x61, 0x73, 0x6b, 0x44, 0x65, 0x73, 0x6b, 0x12, 0x1f, 0x0a, 0x0b, 0x74,
	0x61, 0x73, 0x6b, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x74, 0x61, 0x73, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xb1, 0x01, 0x0a, 0x06, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x74,
	0x61, 0x73, 0x6b, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x74, 0x61, 0x73, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x2c, 0x0a, 0x0b, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x0a, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x23, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f,
	0x62, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x08, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x42, 0x79, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x22, 0x39, 0x0a, 0x16, 0x47,
	0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x42, 0x79, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x06,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x22, 0x36, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73,
	0x6b, 0x42, 0x79, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1b, 0x0a, 0x05, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x05, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x05, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x2a, 0x4d,
	0x0a, 0x0a, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x10, 0x0a, 0x0c,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12, 0x15,
	0x0a, 0x11, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x61, 0x73, 0x6b, 0x50, 0x72, 0x69, 0x6f, 0x72,
	0x69, 0x74, 0x79, 0x10, 0x01, 0x12, 0x16, 0x0a, 0x12, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x61,
	0x73, 0x6b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x10, 0x02, 0x2a, 0x1c, 0x0a,
	0x07, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x73, 0x63, 0x10,
	0x00, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x65, 0x73, 0x63, 0x10, 0x01, 0x32, 0xcc, 0x01, 0x0a, 0x07,
	0x54, 0x61, 0x73, 0x6b, 0x41, 0x70, 0x69, 0x12, 0x37, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x12, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61,
	0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x40, 0x0a, 0x0d, 0x53, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x15, 0x2e, 0x53, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x53, 0x65, 0x74, 0x54, 0x61,
	0x73, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x46, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x42, 0x79, 0x46,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x17, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x42,
	0x79, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18,
	0x2e, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x42, 0x79, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0a, 0x5a, 0x08, 0x61, 0x70,
	0x69, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_api_service_task_proto_rawDescOnce sync.Once
	file_api_service_task_proto_rawDescData []byte
)

func file_api_service_task_proto_rawDescGZIP() []byte {
	file_api_service_task_proto_rawDescOnce.Do(func() {
		file_api_service_task_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_service_task_proto_rawDesc), len(file_api_service_task_proto_rawDesc)))
	})
	return file_api_service_task_proto_rawDescData
}

var file_api_service_task_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_api_service_task_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_api_service_task_proto_goTypes = []any{
	(OrderField)(0),                 // 0: OrderField
	(OrderBy)(0),                    // 1: OrderBy
	(*NewTask)(nil),                 // 2: NewTask
	(*CreateTaskRequest)(nil),       // 3: CreateTaskRequest
	(*CreateTaskResponse)(nil),      // 4: CreateTaskResponse
	(*SetTaskStatusRequest)(nil),    // 5: SetTaskStatusRequest
	(*SetTaskStatusResponse)(nil),   // 6: SetTaskStatusResponse
	(*Task)(nil),                    // 7: Task
	(*Filter)(nil),                  // 8: Filter
	(*GetTaskByFilterRequest)(nil),  // 9: GetTaskByFilterRequest
	(*GetTaskByFilterResponse)(nil), // 10: GetTaskByFilterResponse
}
var file_api_service_task_proto_depIdxs = []int32{
	2,  // 0: CreateTaskRequest.task:type_name -> NewTask
	0,  // 1: Filter.order_field:type_name -> OrderField
	1,  // 2: Filter.order_by:type_name -> OrderBy
	8,  // 3: GetTaskByFilterRequest.filter:type_name -> Filter
	7,  // 4: GetTaskByFilterResponse.tasks:type_name -> Task
	3,  // 5: TaskApi.CreateTask:input_type -> CreateTaskRequest
	5,  // 6: TaskApi.SetTaskStatus:input_type -> SetTaskStatusRequest
	9,  // 7: TaskApi.GetTaskByFilter:input_type -> GetTaskByFilterRequest
	4,  // 8: TaskApi.CreateTask:output_type -> CreateTaskResponse
	6,  // 9: TaskApi.SetTaskStatus:output_type -> SetTaskStatusResponse
	10, // 10: TaskApi.GetTaskByFilter:output_type -> GetTaskByFilterResponse
	8,  // [8:11] is the sub-list for method output_type
	5,  // [5:8] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_api_service_task_proto_init() }
func file_api_service_task_proto_init() {
	if File_api_service_task_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_service_task_proto_rawDesc), len(file_api_service_task_proto_rawDesc)),
			NumEnums:      2,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_service_task_proto_goTypes,
		DependencyIndexes: file_api_service_task_proto_depIdxs,
		EnumInfos:         file_api_service_task_proto_enumTypes,
		MessageInfos:      file_api_service_task_proto_msgTypes,
	}.Build()
	File_api_service_task_proto = out.File
	file_api_service_task_proto_goTypes = nil
	file_api_service_task_proto_depIdxs = nil
}
