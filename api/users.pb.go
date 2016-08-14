// Code generated by protoc-gen-go.
// source: users.proto
// DO NOT EDIT!

package api

import "github.com/golang/protobuf/proto"
import "fmt"
import "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type User struct {
	ID           string                     `protobuf:"bytes,1,opt,name=ID,json=iD" json:"ID,omitempty"`
	Login        string                     `protobuf:"bytes,2,opt,name=Login,json=login" json:"Login,omitempty"`
	Name         string                     `protobuf:"bytes,3,opt,name=Name,json=name" json:"Name,omitempty"`
	RegisteredAt *google_protobuf.Timestamp `protobuf:"bytes,4,opt,name=RegisteredAt,json=registeredAt" json:"RegisteredAt,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *User) GetRegisteredAt() *google_protobuf.Timestamp {
	if m != nil {
		return m.RegisteredAt
	}
	return nil
}

type UserSpec struct {
	Login string `protobuf:"bytes,1,opt,name=Login,json=login" json:"Login,omitempty"`
	ID    string `protobuf:"bytes,2,opt,name=ID,json=iD" json:"ID,omitempty"`
}

func (m *UserSpec) Reset()                    { *m = UserSpec{} }
func (m *UserSpec) String() string            { return proto.CompactTextString(m) }
func (*UserSpec) ProtoMessage()               {}
func (*UserSpec) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func init() {
	proto.RegisterType((*User)(nil), "api.User")
	proto.RegisterType((*UserSpec)(nil), "api.UserSpec")
}

func init() { proto.RegisterFile("users.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 189 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x2e, 0x2d, 0x4e, 0x2d,
	0x2a, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4e, 0x2c, 0xc8, 0x94, 0x92, 0x4f, 0xcf,
	0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x07, 0x0b, 0x25, 0x95, 0xa6, 0xe9, 0x97, 0x64, 0xe6, 0xa6, 0x16,
	0x97, 0x24, 0xe6, 0x16, 0x40, 0x54, 0x29, 0x35, 0x30, 0x72, 0xb1, 0x84, 0x02, 0x75, 0x09, 0xf1,
	0x71, 0x31, 0x79, 0xba, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x31, 0x65, 0xba, 0x08, 0x89,
	0x70, 0xb1, 0xfa, 0xe4, 0xa7, 0x67, 0xe6, 0x49, 0x30, 0x81, 0x85, 0x58, 0x73, 0x40, 0x1c, 0x21,
	0x21, 0x2e, 0x16, 0xbf, 0xc4, 0xdc, 0x54, 0x09, 0x66, 0xb0, 0x20, 0x4b, 0x1e, 0x90, 0x2d, 0x64,
	0xc7, 0xc5, 0x13, 0x94, 0x9a, 0x9e, 0x59, 0x5c, 0x92, 0x5a, 0x94, 0x9a, 0xe2, 0x58, 0x22, 0xc1,
	0x02, 0x94, 0xe3, 0x36, 0x92, 0xd2, 0x83, 0x58, 0xad, 0x07, 0xb3, 0x5a, 0x2f, 0x04, 0x66, 0x75,
	0x10, 0x4f, 0x11, 0x92, 0x7a, 0x25, 0x03, 0x2e, 0x0e, 0x90, 0x0b, 0x82, 0x0b, 0x52, 0x93, 0x11,
	0xb6, 0x32, 0x22, 0xdb, 0x0a, 0x71, 0x1b, 0x13, 0xcc, 0x6d, 0x49, 0x6c, 0x60, 0x33, 0x8d, 0x01,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x2b, 0x3e, 0xa3, 0x2a, 0xf0, 0x00, 0x00, 0x00,
}
