// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/quizer.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Question struct {
	Id                   int64                      `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Text                 string                     `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	Options              map[int64]*Question_Option `protobuf:"bytes,4,rep,name=options,proto3" json:"options,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *Question) Reset()         { *m = Question{} }
func (m *Question) String() string { return proto.CompactTextString(m) }
func (*Question) ProtoMessage()    {}
func (*Question) Descriptor() ([]byte, []int) {
	return fileDescriptor_19b1cb92bbdf88b5, []int{0}
}

func (m *Question) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Question.Unmarshal(m, b)
}
func (m *Question) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Question.Marshal(b, m, deterministic)
}
func (m *Question) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Question.Merge(m, src)
}
func (m *Question) XXX_Size() int {
	return xxx_messageInfo_Question.Size(m)
}
func (m *Question) XXX_DiscardUnknown() {
	xxx_messageInfo_Question.DiscardUnknown(m)
}

var xxx_messageInfo_Question proto.InternalMessageInfo

func (m *Question) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Question) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *Question) GetOptions() map[int64]*Question_Option {
	if m != nil {
		return m.Options
	}
	return nil
}

type Question_Option struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Text                 string   `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Question_Option) Reset()         { *m = Question_Option{} }
func (m *Question_Option) String() string { return proto.CompactTextString(m) }
func (*Question_Option) ProtoMessage()    {}
func (*Question_Option) Descriptor() ([]byte, []int) {
	return fileDescriptor_19b1cb92bbdf88b5, []int{0, 0}
}

func (m *Question_Option) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Question_Option.Unmarshal(m, b)
}
func (m *Question_Option) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Question_Option.Marshal(b, m, deterministic)
}
func (m *Question_Option) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Question_Option.Merge(m, src)
}
func (m *Question_Option) XXX_Size() int {
	return xxx_messageInfo_Question_Option.Size(m)
}
func (m *Question_Option) XXX_DiscardUnknown() {
	xxx_messageInfo_Question_Option.DiscardUnknown(m)
}

var xxx_messageInfo_Question_Option proto.InternalMessageInfo

func (m *Question_Option) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Question_Option) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

type ServerQuestion struct {
	Question             *Question `protobuf:"bytes,1,opt,name=question,proto3" json:"question,omitempty"`
	CorrectAnswer        int64     `protobuf:"varint,2,opt,name=correct_answer,json=correctAnswer,proto3" json:"correct_answer,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ServerQuestion) Reset()         { *m = ServerQuestion{} }
func (m *ServerQuestion) String() string { return proto.CompactTextString(m) }
func (*ServerQuestion) ProtoMessage()    {}
func (*ServerQuestion) Descriptor() ([]byte, []int) {
	return fileDescriptor_19b1cb92bbdf88b5, []int{1}
}

func (m *ServerQuestion) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServerQuestion.Unmarshal(m, b)
}
func (m *ServerQuestion) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServerQuestion.Marshal(b, m, deterministic)
}
func (m *ServerQuestion) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServerQuestion.Merge(m, src)
}
func (m *ServerQuestion) XXX_Size() int {
	return xxx_messageInfo_ServerQuestion.Size(m)
}
func (m *ServerQuestion) XXX_DiscardUnknown() {
	xxx_messageInfo_ServerQuestion.DiscardUnknown(m)
}

var xxx_messageInfo_ServerQuestion proto.InternalMessageInfo

func (m *ServerQuestion) GetQuestion() *Question {
	if m != nil {
		return m.Question
	}
	return nil
}

func (m *ServerQuestion) GetCorrectAnswer() int64 {
	if m != nil {
		return m.CorrectAnswer
	}
	return 0
}

type Answer struct {
	QuestionId           int64    `protobuf:"varint,1,opt,name=question_id,json=questionId,proto3" json:"question_id,omitempty"`
	OptionId             int64    `protobuf:"varint,2,opt,name=option_id,json=optionId,proto3" json:"option_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Answer) Reset()         { *m = Answer{} }
func (m *Answer) String() string { return proto.CompactTextString(m) }
func (*Answer) ProtoMessage()    {}
func (*Answer) Descriptor() ([]byte, []int) {
	return fileDescriptor_19b1cb92bbdf88b5, []int{2}
}

func (m *Answer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Answer.Unmarshal(m, b)
}
func (m *Answer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Answer.Marshal(b, m, deterministic)
}
func (m *Answer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Answer.Merge(m, src)
}
func (m *Answer) XXX_Size() int {
	return xxx_messageInfo_Answer.Size(m)
}
func (m *Answer) XXX_DiscardUnknown() {
	xxx_messageInfo_Answer.DiscardUnknown(m)
}

var xxx_messageInfo_Answer proto.InternalMessageInfo

func (m *Answer) GetQuestionId() int64 {
	if m != nil {
		return m.QuestionId
	}
	return 0
}

func (m *Answer) GetOptionId() int64 {
	if m != nil {
		return m.OptionId
	}
	return 0
}

type Result struct {
	CorrectAnswers       int32    `protobuf:"varint,1,opt,name=correct_answers,json=correctAnswers,proto3" json:"correct_answers,omitempty"`
	BetterThen           int32    `protobuf:"varint,2,opt,name=better_then,json=betterThen,proto3" json:"better_then,omitempty"`
	ElapsedTime          float32  `protobuf:"fixed32,3,opt,name=elapsed_time,json=elapsedTime,proto3" json:"elapsed_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Result) Reset()         { *m = Result{} }
func (m *Result) String() string { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()    {}
func (*Result) Descriptor() ([]byte, []int) {
	return fileDescriptor_19b1cb92bbdf88b5, []int{3}
}

func (m *Result) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Result.Unmarshal(m, b)
}
func (m *Result) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Result.Marshal(b, m, deterministic)
}
func (m *Result) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Result.Merge(m, src)
}
func (m *Result) XXX_Size() int {
	return xxx_messageInfo_Result.Size(m)
}
func (m *Result) XXX_DiscardUnknown() {
	xxx_messageInfo_Result.DiscardUnknown(m)
}

var xxx_messageInfo_Result proto.InternalMessageInfo

func (m *Result) GetCorrectAnswers() int32 {
	if m != nil {
		return m.CorrectAnswers
	}
	return 0
}

func (m *Result) GetBetterThen() int32 {
	if m != nil {
		return m.BetterThen
	}
	return 0
}

func (m *Result) GetElapsedTime() float32 {
	if m != nil {
		return m.ElapsedTime
	}
	return 0
}

// quick fix and replacement for google.protobuf.Empty file not found error
type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_19b1cb92bbdf88b5, []int{4}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Question)(nil), "proto.Question")
	proto.RegisterMapType((map[int64]*Question_Option)(nil), "proto.Question.OptionsEntry")
	proto.RegisterType((*Question_Option)(nil), "proto.Question.Option")
	proto.RegisterType((*ServerQuestion)(nil), "proto.ServerQuestion")
	proto.RegisterType((*Answer)(nil), "proto.Answer")
	proto.RegisterType((*Result)(nil), "proto.Result")
	proto.RegisterType((*Empty)(nil), "proto.Empty")
}

func init() { proto.RegisterFile("proto/quizer.proto", fileDescriptor_19b1cb92bbdf88b5) }

var fileDescriptor_19b1cb92bbdf88b5 = []byte{
	// 387 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xd1, 0xaa, 0xd3, 0x40,
	0x10, 0xbd, 0x9b, 0xdc, 0xe4, 0xf6, 0x4e, 0x9a, 0x5c, 0x99, 0x07, 0x09, 0x51, 0x30, 0x06, 0xc4,
	0x80, 0xa5, 0x6a, 0x04, 0x11, 0xdf, 0x7c, 0xa8, 0x50, 0x10, 0xa4, 0x6b, 0xdf, 0x43, 0xdb, 0x0c,
	0x34, 0xd8, 0x26, 0xe9, 0xee, 0xa6, 0x5a, 0xbf, 0xd8, 0xcf, 0x90, 0xec, 0x26, 0xc5, 0x16, 0x1f,
	0x7c, 0xca, 0xcc, 0x99, 0x33, 0xe7, 0xec, 0x1c, 0x02, 0xd8, 0x88, 0x5a, 0xd5, 0xaf, 0x0f, 0x6d,
	0xf9, 0x8b, 0xc4, 0x54, 0x37, 0xe8, 0xe8, 0x4f, 0xf2, 0x9b, 0xc1, 0x68, 0xd1, 0x92, 0x54, 0x65,
	0x5d, 0x61, 0x00, 0x56, 0x59, 0x84, 0x2c, 0x66, 0xa9, 0xcd, 0xad, 0xb2, 0x40, 0x84, 0x5b, 0x45,
	0x3f, 0x55, 0x68, 0xc5, 0x2c, 0xbd, 0xe7, 0xba, 0xc6, 0xf7, 0x70, 0x57, 0x37, 0x1d, 0x5b, 0x86,
	0xb7, 0xb1, 0x9d, 0x7a, 0xd9, 0x53, 0x23, 0x38, 0x1d, 0x54, 0xa6, 0x5f, 0xcd, 0x78, 0x56, 0x29,
	0x71, 0xe2, 0x03, 0x39, 0x9a, 0x80, 0x6b, 0x06, 0xff, 0xe3, 0x12, 0x71, 0x18, 0xff, 0x2d, 0x83,
	0x8f, 0xc0, 0xfe, 0x4e, 0xa7, 0x7e, 0xa9, 0x2b, 0x71, 0x02, 0xce, 0x71, 0xb5, 0x6b, 0x49, 0xaf,
	0x79, 0xd9, 0xe3, 0x7f, 0xbf, 0x82, 0x1b, 0xd2, 0x47, 0xeb, 0x03, 0x4b, 0x0a, 0x08, 0xbe, 0x91,
	0x38, 0x92, 0x38, 0xdf, 0xfb, 0x0a, 0x46, 0x87, 0xbe, 0xd6, 0xd2, 0x5e, 0xf6, 0x70, 0x25, 0xc3,
	0xcf, 0x04, 0x7c, 0x01, 0xc1, 0xa6, 0x16, 0x82, 0x36, 0x2a, 0x5f, 0x55, 0xf2, 0x07, 0x09, 0xed,
	0x6c, 0x73, 0xbf, 0x47, 0x3f, 0x69, 0x30, 0xf9, 0x0c, 0xae, 0xa9, 0xf0, 0x19, 0x78, 0xc3, 0x72,
	0x7e, 0x3e, 0x18, 0x06, 0x68, 0x5e, 0xe0, 0x13, 0xb8, 0x37, 0xe9, 0x74, 0x63, 0x23, 0x36, 0x32,
	0xc0, 0xbc, 0x48, 0x5a, 0x70, 0x39, 0xc9, 0x76, 0xa7, 0xf0, 0x25, 0x3c, 0x5c, 0x1a, 0x4b, 0xad,
	0xe5, 0xf0, 0xe0, 0xc2, 0x59, 0x76, 0x86, 0x6b, 0x52, 0x8a, 0x44, 0xae, 0xb6, 0x54, 0x69, 0x45,
	0x87, 0x83, 0x81, 0x96, 0x5b, 0xaa, 0xf0, 0x39, 0x8c, 0x69, 0xb7, 0x6a, 0x24, 0x15, 0xb9, 0x2a,
	0xf7, 0x14, 0xda, 0x31, 0x4b, 0x2d, 0xee, 0xf5, 0xd8, 0xb2, 0xdc, 0x53, 0x72, 0x07, 0xce, 0x6c,
	0xdf, 0xa8, 0x53, 0x56, 0x83, 0xbb, 0xd0, 0xff, 0x0b, 0x66, 0xe0, 0x7f, 0x29, 0xa5, 0x1a, 0x22,
	0x91, 0x38, 0xee, 0x43, 0xd2, 0xc4, 0xe8, 0x3a, 0xb2, 0xe4, 0xe6, 0x0d, 0xc3, 0xb7, 0xe0, 0x73,
	0xda, 0xd4, 0xa2, 0x18, 0xde, 0xe6, 0xf7, 0x2c, 0xd3, 0x47, 0x43, 0x6b, 0x4e, 0x4c, 0x6e, 0x52,
	0xb6, 0x76, 0x35, 0xf2, 0xee, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x94, 0x5d, 0x39, 0xb2, 0xad,
	0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QuizerClient is the client API for Quizer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QuizerClient interface {
	// server-to-client streaming RPC
	// Questions are streamed rather than returned at once
	// in order to serve a bigger number of questions
	ListQuestions(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Quizer_ListQuestionsClient, error)
	// client-to-server streaming RPC
	// Accepts a stream of Answers and
	// returns a Result after all the answers are submitted
	RecordAnswers(ctx context.Context, opts ...grpc.CallOption) (Quizer_RecordAnswersClient, error)
}

type quizerClient struct {
	cc *grpc.ClientConn
}

func NewQuizerClient(cc *grpc.ClientConn) QuizerClient {
	return &quizerClient{cc}
}

func (c *quizerClient) ListQuestions(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Quizer_ListQuestionsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Quizer_serviceDesc.Streams[0], "/proto.Quizer/ListQuestions", opts...)
	if err != nil {
		return nil, err
	}
	x := &quizerListQuestionsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Quizer_ListQuestionsClient interface {
	Recv() (*Question, error)
	grpc.ClientStream
}

type quizerListQuestionsClient struct {
	grpc.ClientStream
}

func (x *quizerListQuestionsClient) Recv() (*Question, error) {
	m := new(Question)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *quizerClient) RecordAnswers(ctx context.Context, opts ...grpc.CallOption) (Quizer_RecordAnswersClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Quizer_serviceDesc.Streams[1], "/proto.Quizer/RecordAnswers", opts...)
	if err != nil {
		return nil, err
	}
	x := &quizerRecordAnswersClient{stream}
	return x, nil
}

type Quizer_RecordAnswersClient interface {
	Send(*Answer) error
	CloseAndRecv() (*Result, error)
	grpc.ClientStream
}

type quizerRecordAnswersClient struct {
	grpc.ClientStream
}

func (x *quizerRecordAnswersClient) Send(m *Answer) error {
	return x.ClientStream.SendMsg(m)
}

func (x *quizerRecordAnswersClient) CloseAndRecv() (*Result, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Result)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// QuizerServer is the server API for Quizer service.
type QuizerServer interface {
	// server-to-client streaming RPC
	// Questions are streamed rather than returned at once
	// in order to serve a bigger number of questions
	ListQuestions(*Empty, Quizer_ListQuestionsServer) error
	// client-to-server streaming RPC
	// Accepts a stream of Answers and
	// returns a Result after all the answers are submitted
	RecordAnswers(Quizer_RecordAnswersServer) error
}

// UnimplementedQuizerServer can be embedded to have forward compatible implementations.
type UnimplementedQuizerServer struct {
}

func (*UnimplementedQuizerServer) ListQuestions(req *Empty, srv Quizer_ListQuestionsServer) error {
	return status.Errorf(codes.Unimplemented, "method ListQuestions not implemented")
}
func (*UnimplementedQuizerServer) RecordAnswers(srv Quizer_RecordAnswersServer) error {
	return status.Errorf(codes.Unimplemented, "method RecordAnswers not implemented")
}

func RegisterQuizerServer(s *grpc.Server, srv QuizerServer) {
	s.RegisterService(&_Quizer_serviceDesc, srv)
}

func _Quizer_ListQuestions_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(QuizerServer).ListQuestions(m, &quizerListQuestionsServer{stream})
}

type Quizer_ListQuestionsServer interface {
	Send(*Question) error
	grpc.ServerStream
}

type quizerListQuestionsServer struct {
	grpc.ServerStream
}

func (x *quizerListQuestionsServer) Send(m *Question) error {
	return x.ServerStream.SendMsg(m)
}

func _Quizer_RecordAnswers_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(QuizerServer).RecordAnswers(&quizerRecordAnswersServer{stream})
}

type Quizer_RecordAnswersServer interface {
	SendAndClose(*Result) error
	Recv() (*Answer, error)
	grpc.ServerStream
}

type quizerRecordAnswersServer struct {
	grpc.ServerStream
}

func (x *quizerRecordAnswersServer) SendAndClose(m *Result) error {
	return x.ServerStream.SendMsg(m)
}

func (x *quizerRecordAnswersServer) Recv() (*Answer, error) {
	m := new(Answer)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Quizer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Quizer",
	HandlerType: (*QuizerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListQuestions",
			Handler:       _Quizer_ListQuestions_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "RecordAnswers",
			Handler:       _Quizer_RecordAnswers_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "proto/quizer.proto",
}