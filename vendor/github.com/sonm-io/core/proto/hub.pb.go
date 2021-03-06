// Code generated by protoc-gen-go. DO NOT EDIT.
// source: hub.proto

package sonm

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ListRequest struct {
}

func (m *ListRequest) Reset()                    { *m = ListRequest{} }
func (m *ListRequest) String() string            { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()               {}
func (*ListRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type ListReply struct {
	Info map[string]*ListReply_ListValue `protobuf:"bytes,1,rep,name=info" json:"info,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *ListReply) Reset()                    { *m = ListReply{} }
func (m *ListReply) String() string            { return proto.CompactTextString(m) }
func (*ListReply) ProtoMessage()               {}
func (*ListReply) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *ListReply) GetInfo() map[string]*ListReply_ListValue {
	if m != nil {
		return m.Info
	}
	return nil
}

type ListReply_ListValue struct {
	Values []string `protobuf:"bytes,1,rep,name=values" json:"values,omitempty"`
}

func (m *ListReply_ListValue) Reset()                    { *m = ListReply_ListValue{} }
func (m *ListReply_ListValue) String() string            { return proto.CompactTextString(m) }
func (*ListReply_ListValue) ProtoMessage()               {}
func (*ListReply_ListValue) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1, 0} }

func (m *ListReply_ListValue) GetValues() []string {
	if m != nil {
		return m.Values
	}
	return nil
}

type HubInfoRequest struct {
	Miner string `protobuf:"bytes,1,opt,name=miner" json:"miner,omitempty"`
}

func (m *HubInfoRequest) Reset()                    { *m = HubInfoRequest{} }
func (m *HubInfoRequest) String() string            { return proto.CompactTextString(m) }
func (*HubInfoRequest) ProtoMessage()               {}
func (*HubInfoRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *HubInfoRequest) GetMiner() string {
	if m != nil {
		return m.Miner
	}
	return ""
}

type TaskResourceRequirements struct {
	// Number of CPU cores.
	CPUCores uint64 `protobuf:"varint,1,opt,name=CPUCores,json=cPUCores" json:"CPUCores,omitempty"`
	// Memory in bytes required.
	MaxMemory uint64 `protobuf:"varint,2,opt,name=maxMemory" json:"maxMemory,omitempty"`
	// Describes whether a task requires GPU supoort.
	GPUSupport bool `protobuf:"varint,3,opt,name=GPUSupport,json=gPUSupport" json:"GPUSupport,omitempty"`
}

func (m *TaskResourceRequirements) Reset()                    { *m = TaskResourceRequirements{} }
func (m *TaskResourceRequirements) String() string            { return proto.CompactTextString(m) }
func (*TaskResourceRequirements) ProtoMessage()               {}
func (*TaskResourceRequirements) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *TaskResourceRequirements) GetCPUCores() uint64 {
	if m != nil {
		return m.CPUCores
	}
	return 0
}

func (m *TaskResourceRequirements) GetMaxMemory() uint64 {
	if m != nil {
		return m.MaxMemory
	}
	return 0
}

func (m *TaskResourceRequirements) GetGPUSupport() bool {
	if m != nil {
		return m.GPUSupport
	}
	return false
}

type TaskRequirements struct {
	// How much resources consumes a task. This is used both for scheduling and for cgroups configuration.
	Resources *TaskResourceRequirements `protobuf:"bytes,1,opt,name=resources" json:"resources,omitempty"`
	// Optional miner ID restrictions (currently IP:port), that are allowed to start a task.
	Miners []string `protobuf:"bytes,2,rep,name=miners" json:"miners,omitempty"`
}

func (m *TaskRequirements) Reset()                    { *m = TaskRequirements{} }
func (m *TaskRequirements) String() string            { return proto.CompactTextString(m) }
func (*TaskRequirements) ProtoMessage()               {}
func (*TaskRequirements) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *TaskRequirements) GetResources() *TaskResourceRequirements {
	if m != nil {
		return m.Resources
	}
	return nil
}

func (m *TaskRequirements) GetMiners() []string {
	if m != nil {
		return m.Miners
	}
	return nil
}

type HubStartTaskRequest struct {
	Requirements  *TaskRequirements `protobuf:"bytes,1,opt,name=requirements" json:"requirements,omitempty"`
	Registry      string            `protobuf:"bytes,2,opt,name=registry" json:"registry,omitempty"`
	Image         string            `protobuf:"bytes,3,opt,name=image" json:"image,omitempty"`
	Auth          string            `protobuf:"bytes,4,opt,name=auth" json:"auth,omitempty"`
	PublicKeyData string            `protobuf:"bytes,5,opt,name=PublicKeyData,json=publicKeyData" json:"PublicKeyData,omitempty"`
	CommitOnStop  bool              `protobuf:"varint,6,opt,name=commitOnStop" json:"commitOnStop,omitempty"`
	Env           map[string]string `protobuf:"bytes,7,rep,name=env" json:"env,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *HubStartTaskRequest) Reset()                    { *m = HubStartTaskRequest{} }
func (m *HubStartTaskRequest) String() string            { return proto.CompactTextString(m) }
func (*HubStartTaskRequest) ProtoMessage()               {}
func (*HubStartTaskRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

func (m *HubStartTaskRequest) GetRequirements() *TaskRequirements {
	if m != nil {
		return m.Requirements
	}
	return nil
}

func (m *HubStartTaskRequest) GetRegistry() string {
	if m != nil {
		return m.Registry
	}
	return ""
}

func (m *HubStartTaskRequest) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *HubStartTaskRequest) GetAuth() string {
	if m != nil {
		return m.Auth
	}
	return ""
}

func (m *HubStartTaskRequest) GetPublicKeyData() string {
	if m != nil {
		return m.PublicKeyData
	}
	return ""
}

func (m *HubStartTaskRequest) GetCommitOnStop() bool {
	if m != nil {
		return m.CommitOnStop
	}
	return false
}

func (m *HubStartTaskRequest) GetEnv() map[string]string {
	if m != nil {
		return m.Env
	}
	return nil
}

type HubStartTaskReply struct {
	Id       string   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Endpoint []string `protobuf:"bytes,2,rep,name=endpoint" json:"endpoint,omitempty"`
}

func (m *HubStartTaskReply) Reset()                    { *m = HubStartTaskReply{} }
func (m *HubStartTaskReply) String() string            { return proto.CompactTextString(m) }
func (*HubStartTaskReply) ProtoMessage()               {}
func (*HubStartTaskReply) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{6} }

func (m *HubStartTaskReply) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *HubStartTaskReply) GetEndpoint() []string {
	if m != nil {
		return m.Endpoint
	}
	return nil
}

type HubStatusMapRequest struct {
	Miner string `protobuf:"bytes,1,opt,name=miner" json:"miner,omitempty"`
}

func (m *HubStatusMapRequest) Reset()                    { *m = HubStatusMapRequest{} }
func (m *HubStatusMapRequest) String() string            { return proto.CompactTextString(m) }
func (*HubStatusMapRequest) ProtoMessage()               {}
func (*HubStatusMapRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{7} }

func (m *HubStatusMapRequest) GetMiner() string {
	if m != nil {
		return m.Miner
	}
	return ""
}

type HubStatusRequest struct {
}

func (m *HubStatusRequest) Reset()                    { *m = HubStatusRequest{} }
func (m *HubStatusRequest) String() string            { return proto.CompactTextString(m) }
func (*HubStatusRequest) ProtoMessage()               {}
func (*HubStatusRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{8} }

type HubStatusReply struct {
	MinerCount uint64 `protobuf:"varint,1,opt,name=minerCount" json:"minerCount,omitempty"`
	Uptime     uint64 `protobuf:"varint,2,opt,name=uptime" json:"uptime,omitempty"`
}

func (m *HubStatusReply) Reset()                    { *m = HubStatusReply{} }
func (m *HubStatusReply) String() string            { return proto.CompactTextString(m) }
func (*HubStatusReply) ProtoMessage()               {}
func (*HubStatusReply) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{9} }

func (m *HubStatusReply) GetMinerCount() uint64 {
	if m != nil {
		return m.MinerCount
	}
	return 0
}

func (m *HubStatusReply) GetUptime() uint64 {
	if m != nil {
		return m.Uptime
	}
	return 0
}

func init() {
	proto.RegisterType((*ListRequest)(nil), "sonm.ListRequest")
	proto.RegisterType((*ListReply)(nil), "sonm.ListReply")
	proto.RegisterType((*ListReply_ListValue)(nil), "sonm.ListReply.ListValue")
	proto.RegisterType((*HubInfoRequest)(nil), "sonm.HubInfoRequest")
	proto.RegisterType((*TaskResourceRequirements)(nil), "sonm.TaskResourceRequirements")
	proto.RegisterType((*TaskRequirements)(nil), "sonm.TaskRequirements")
	proto.RegisterType((*HubStartTaskRequest)(nil), "sonm.HubStartTaskRequest")
	proto.RegisterType((*HubStartTaskReply)(nil), "sonm.HubStartTaskReply")
	proto.RegisterType((*HubStatusMapRequest)(nil), "sonm.HubStatusMapRequest")
	proto.RegisterType((*HubStatusRequest)(nil), "sonm.HubStatusRequest")
	proto.RegisterType((*HubStatusReply)(nil), "sonm.HubStatusReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Hub service

type HubClient interface {
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingReply, error)
	Status(ctx context.Context, in *HubStatusRequest, opts ...grpc.CallOption) (*HubStatusReply, error)
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListReply, error)
	Info(ctx context.Context, in *HubInfoRequest, opts ...grpc.CallOption) (*InfoReply, error)
	StartTask(ctx context.Context, in *HubStartTaskRequest, opts ...grpc.CallOption) (*HubStartTaskReply, error)
	StopTask(ctx context.Context, in *StopTaskRequest, opts ...grpc.CallOption) (*StopTaskReply, error)
	TaskStatus(ctx context.Context, in *TaskStatusRequest, opts ...grpc.CallOption) (*TaskStatusReply, error)
	MinerStatus(ctx context.Context, in *HubStatusMapRequest, opts ...grpc.CallOption) (*StatusMapReply, error)
	TaskLogs(ctx context.Context, in *TaskLogsRequest, opts ...grpc.CallOption) (Hub_TaskLogsClient, error)
}

type hubClient struct {
	cc *grpc.ClientConn
}

func NewHubClient(cc *grpc.ClientConn) HubClient {
	return &hubClient{cc}
}

func (c *hubClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingReply, error) {
	out := new(PingReply)
	err := grpc.Invoke(ctx, "/sonm.Hub/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hubClient) Status(ctx context.Context, in *HubStatusRequest, opts ...grpc.CallOption) (*HubStatusReply, error) {
	out := new(HubStatusReply)
	err := grpc.Invoke(ctx, "/sonm.Hub/Status", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hubClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListReply, error) {
	out := new(ListReply)
	err := grpc.Invoke(ctx, "/sonm.Hub/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hubClient) Info(ctx context.Context, in *HubInfoRequest, opts ...grpc.CallOption) (*InfoReply, error) {
	out := new(InfoReply)
	err := grpc.Invoke(ctx, "/sonm.Hub/Info", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hubClient) StartTask(ctx context.Context, in *HubStartTaskRequest, opts ...grpc.CallOption) (*HubStartTaskReply, error) {
	out := new(HubStartTaskReply)
	err := grpc.Invoke(ctx, "/sonm.Hub/StartTask", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hubClient) StopTask(ctx context.Context, in *StopTaskRequest, opts ...grpc.CallOption) (*StopTaskReply, error) {
	out := new(StopTaskReply)
	err := grpc.Invoke(ctx, "/sonm.Hub/StopTask", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hubClient) TaskStatus(ctx context.Context, in *TaskStatusRequest, opts ...grpc.CallOption) (*TaskStatusReply, error) {
	out := new(TaskStatusReply)
	err := grpc.Invoke(ctx, "/sonm.Hub/TaskStatus", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hubClient) MinerStatus(ctx context.Context, in *HubStatusMapRequest, opts ...grpc.CallOption) (*StatusMapReply, error) {
	out := new(StatusMapReply)
	err := grpc.Invoke(ctx, "/sonm.Hub/MinerStatus", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hubClient) TaskLogs(ctx context.Context, in *TaskLogsRequest, opts ...grpc.CallOption) (Hub_TaskLogsClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Hub_serviceDesc.Streams[0], c.cc, "/sonm.Hub/TaskLogs", opts...)
	if err != nil {
		return nil, err
	}
	x := &hubTaskLogsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Hub_TaskLogsClient interface {
	Recv() (*TaskLogsChunk, error)
	grpc.ClientStream
}

type hubTaskLogsClient struct {
	grpc.ClientStream
}

func (x *hubTaskLogsClient) Recv() (*TaskLogsChunk, error) {
	m := new(TaskLogsChunk)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Hub service

type HubServer interface {
	Ping(context.Context, *PingRequest) (*PingReply, error)
	Status(context.Context, *HubStatusRequest) (*HubStatusReply, error)
	List(context.Context, *ListRequest) (*ListReply, error)
	Info(context.Context, *HubInfoRequest) (*InfoReply, error)
	StartTask(context.Context, *HubStartTaskRequest) (*HubStartTaskReply, error)
	StopTask(context.Context, *StopTaskRequest) (*StopTaskReply, error)
	TaskStatus(context.Context, *TaskStatusRequest) (*TaskStatusReply, error)
	MinerStatus(context.Context, *HubStatusMapRequest) (*StatusMapReply, error)
	TaskLogs(*TaskLogsRequest, Hub_TaskLogsServer) error
}

func RegisterHubServer(s *grpc.Server, srv HubServer) {
	s.RegisterService(&_Hub_serviceDesc, srv)
}

func _Hub_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HubServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.Hub/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HubServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Hub_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HubStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HubServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.Hub/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HubServer).Status(ctx, req.(*HubStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Hub_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HubServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.Hub/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HubServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Hub_Info_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HubInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HubServer).Info(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.Hub/Info",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HubServer).Info(ctx, req.(*HubInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Hub_StartTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HubStartTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HubServer).StartTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.Hub/StartTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HubServer).StartTask(ctx, req.(*HubStartTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Hub_StopTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HubServer).StopTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.Hub/StopTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HubServer).StopTask(ctx, req.(*StopTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Hub_TaskStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HubServer).TaskStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.Hub/TaskStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HubServer).TaskStatus(ctx, req.(*TaskStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Hub_MinerStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HubStatusMapRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HubServer).MinerStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.Hub/MinerStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HubServer).MinerStatus(ctx, req.(*HubStatusMapRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Hub_TaskLogs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(TaskLogsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(HubServer).TaskLogs(m, &hubTaskLogsServer{stream})
}

type Hub_TaskLogsServer interface {
	Send(*TaskLogsChunk) error
	grpc.ServerStream
}

type hubTaskLogsServer struct {
	grpc.ServerStream
}

func (x *hubTaskLogsServer) Send(m *TaskLogsChunk) error {
	return x.ServerStream.SendMsg(m)
}

var _Hub_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sonm.Hub",
	HandlerType: (*HubServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Hub_Ping_Handler,
		},
		{
			MethodName: "Status",
			Handler:    _Hub_Status_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Hub_List_Handler,
		},
		{
			MethodName: "Info",
			Handler:    _Hub_Info_Handler,
		},
		{
			MethodName: "StartTask",
			Handler:    _Hub_StartTask_Handler,
		},
		{
			MethodName: "StopTask",
			Handler:    _Hub_StopTask_Handler,
		},
		{
			MethodName: "TaskStatus",
			Handler:    _Hub_TaskStatus_Handler,
		},
		{
			MethodName: "MinerStatus",
			Handler:    _Hub_MinerStatus_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "TaskLogs",
			Handler:       _Hub_TaskLogs_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "hub.proto",
}

func init() { proto.RegisterFile("hub.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 698 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x54, 0xdb, 0x6e, 0xd3, 0x40,
	0x10, 0xad, 0x13, 0x37, 0xc4, 0x93, 0x5e, 0xb7, 0x37, 0xd7, 0x42, 0x55, 0x64, 0x10, 0x8a, 0x44,
	0x49, 0x51, 0x41, 0xa8, 0xaa, 0x10, 0x17, 0x85, 0x8a, 0x20, 0x5a, 0x11, 0xb9, 0x94, 0x77, 0x27,
	0xdd, 0x26, 0xab, 0xc6, 0xbb, 0xc6, 0xde, 0xad, 0xc8, 0x17, 0xf0, 0x41, 0x7c, 0x08, 0xbf, 0x84,
	0xf6, 0x62, 0x67, 0x93, 0xa6, 0x3c, 0xc5, 0x73, 0x76, 0x66, 0xce, 0xce, 0xd9, 0x33, 0x01, 0x6f,
	0x24, 0xfa, 0xed, 0x34, 0x63, 0x9c, 0x21, 0x37, 0x67, 0x34, 0x09, 0xd6, 0x09, 0x95, 0xbf, 0x94,
	0xc4, 0x1a, 0x0e, 0x57, 0xa1, 0x71, 0x4e, 0x72, 0x1e, 0xe1, 0x9f, 0x02, 0xe7, 0x3c, 0xfc, 0xe3,
	0x80, 0xa7, 0xe3, 0x74, 0x3c, 0x41, 0x2f, 0xc0, 0x25, 0xf4, 0x86, 0xf9, 0x4e, 0xb3, 0xda, 0x6a,
	0x1c, 0xef, 0xb7, 0x65, 0x69, 0xbb, 0x3c, 0x6e, 0x7f, 0xa1, 0x37, 0xec, 0x8c, 0xf2, 0x6c, 0x12,
	0xa9, 0xb4, 0xe0, 0x89, 0xae, 0xfd, 0x11, 0x8f, 0x05, 0x46, 0xbb, 0x50, 0xbb, 0x93, 0x1f, 0xb9,
	0xaa, 0xf6, 0x22, 0x13, 0x05, 0x11, 0x78, 0x65, 0x1d, 0xda, 0x80, 0xea, 0x2d, 0x9e, 0xf8, 0x4e,
	0xd3, 0x69, 0x79, 0x91, 0xfc, 0x44, 0x47, 0xb0, 0xac, 0x12, 0xfd, 0x4a, 0xd3, 0x59, 0xc4, 0x59,
	0x12, 0x44, 0x3a, 0xef, 0xb4, 0x72, 0xe2, 0x84, 0xcf, 0x60, 0xad, 0x2b, 0xfa, 0xb2, 0xad, 0x99,
	0x03, 0x6d, 0xc3, 0x72, 0x42, 0x28, 0xce, 0x4c, 0x6b, 0x1d, 0x84, 0x1c, 0xfc, 0xef, 0x71, 0x7e,
	0x1b, 0xe1, 0x9c, 0x89, 0x6c, 0x80, 0x65, 0x32, 0xc9, 0x70, 0x82, 0x29, 0xcf, 0x51, 0x00, 0xf5,
	0x4e, 0xef, 0xaa, 0xc3, 0x32, 0x75, 0x63, 0xa7, 0xe5, 0x46, 0xf5, 0x81, 0x89, 0xd1, 0x63, 0xf0,
	0x92, 0xf8, 0xd7, 0x05, 0x4e, 0x58, 0x36, 0x51, 0x17, 0x73, 0xa3, 0x29, 0x80, 0x0e, 0x00, 0x3e,
	0xf7, 0xae, 0x2e, 0x45, 0x9a, 0xb2, 0x8c, 0xfb, 0xd5, 0xa6, 0xd3, 0xaa, 0x47, 0x30, 0x2c, 0x91,
	0x70, 0x04, 0x1b, 0x9a, 0xd5, 0x62, 0x7b, 0x0b, 0x5e, 0x66, 0x6e, 0xa1, 0xe9, 0x1a, 0xc7, 0x07,
	0x7a, 0xd4, 0x87, 0x2e, 0x18, 0x4d, 0x0b, 0xa4, 0xb6, 0x6a, 0xa0, 0xdc, 0xaf, 0x68, 0x6d, 0x75,
	0x14, 0xfe, 0xad, 0xc0, 0x56, 0x57, 0xf4, 0x2f, 0x79, 0x9c, 0xf1, 0x82, 0x52, 0xaa, 0x71, 0x0a,
	0x2b, 0x99, 0xd5, 0xca, 0x10, 0xee, 0xda, 0x84, 0x16, 0xd1, 0x4c, 0xae, 0xd4, 0x25, 0xc3, 0x43,
	0x92, 0x73, 0x33, 0xba, 0x17, 0x95, 0xb1, 0x54, 0x99, 0x24, 0xf1, 0x10, 0xab, 0xa1, 0xbd, 0x48,
	0x07, 0x08, 0x81, 0x1b, 0x0b, 0x3e, 0xf2, 0x5d, 0x05, 0xaa, 0x6f, 0xf4, 0x14, 0x56, 0x7b, 0xa2,
	0x3f, 0x26, 0x83, 0xaf, 0x78, 0xf2, 0x29, 0xe6, 0xb1, 0xbf, 0xac, 0x0e, 0x57, 0x53, 0x1b, 0x44,
	0x21, 0xac, 0x0c, 0x58, 0x92, 0x10, 0xfe, 0x8d, 0x5e, 0x72, 0x96, 0xfa, 0x35, 0xa5, 0xe5, 0x0c,
	0x86, 0x5e, 0x43, 0x15, 0xd3, 0x3b, 0xff, 0x91, 0xb2, 0x64, 0xa8, 0x47, 0x58, 0x30, 0x73, 0xfb,
	0x8c, 0xde, 0x69, 0x6f, 0xca, 0xf4, 0xe0, 0x0d, 0xd4, 0x0b, 0x60, 0x81, 0xe9, 0xb6, 0x6d, 0xd3,
	0x79, 0xb6, 0xb3, 0xde, 0xc3, 0xe6, 0x6c, 0x73, 0xb9, 0x16, 0x6b, 0x50, 0x21, 0xd7, 0xa6, 0xbe,
	0x42, 0xae, 0xa5, 0x44, 0x98, 0x5e, 0xa7, 0x8c, 0x50, 0x6e, 0x1e, 0xa4, 0x8c, 0xc3, 0xe7, 0xc5,
	0x8b, 0x70, 0x91, 0x5f, 0xc4, 0xe9, 0xff, 0xfd, 0x89, 0x60, 0xa3, 0x4c, 0x2e, 0x36, 0xb2, 0xab,
	0xbc, 0x5d, 0x60, 0x92, 0xfe, 0x00, 0x40, 0xa5, 0x77, 0x98, 0xa0, 0xdc, 0x78, 0xd5, 0x42, 0xa4,
	0x3b, 0x44, 0xca, 0x49, 0x82, 0x8d, 0x55, 0x4d, 0x74, 0xfc, 0xdb, 0x85, 0x6a, 0x57, 0xf4, 0xd1,
	0x21, 0xb8, 0x3d, 0x42, 0x87, 0x68, 0x53, 0x8b, 0x27, 0xbf, 0x0d, 0x59, 0xb0, 0x6e, 0x43, 0xe9,
	0x78, 0x12, 0x2e, 0xa1, 0x13, 0xa8, 0x69, 0x72, 0xb4, 0x6b, 0x8b, 0x3d, 0xbd, 0x61, 0xb0, 0x7d,
	0x0f, 0xd7, 0x95, 0x87, 0xe0, 0xca, 0x6d, 0x2d, 0x78, 0xac, 0xbf, 0x99, 0x82, 0xa7, 0x5c, 0xeb,
	0x70, 0x09, 0x1d, 0x81, 0x2b, 0x17, 0x18, 0x4d, 0xbb, 0x59, 0xfb, 0x5c, 0x14, 0x68, 0x48, 0x17,
	0x7c, 0x04, 0xaf, 0x7c, 0x17, 0xb4, 0xff, 0xa0, 0x11, 0x82, 0xbd, 0x45, 0x47, 0xc5, 0x6c, 0x75,
	0xe9, 0x29, 0xd5, 0x61, 0x47, 0xa7, 0x15, 0x71, 0x51, 0xbd, 0x35, 0x0f, 0xeb, 0xca, 0x77, 0x00,
	0x32, 0x34, 0xca, 0xec, 0x4d, 0x37, 0x69, 0x56, 0x9a, 0x9d, 0xfb, 0x07, 0xba, 0xfe, 0x03, 0x34,
	0x2e, 0xe4, 0x8b, 0x99, 0x06, 0xfb, 0x73, 0x12, 0x4e, 0x9d, 0x52, 0xa8, 0x6b, 0xe1, 0xba, 0xc3,
	0x29, 0xd4, 0x65, 0xdb, 0x73, 0x36, 0xcc, 0x91, 0x45, 0x23, 0xe3, 0xb9, 0xbb, 0x17, 0x70, 0x67,
	0x24, 0xe8, 0x6d, 0xb8, 0xf4, 0xd2, 0xe9, 0xd7, 0xd4, 0x7f, 0xff, 0xab, 0x7f, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x3e, 0x36, 0xec, 0xb6, 0x1f, 0x06, 0x00, 0x00,
}
