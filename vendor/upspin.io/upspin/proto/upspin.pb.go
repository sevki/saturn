// Code generated by protoc-gen-go. DO NOT EDIT.
// source: upspin.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	upspin.proto

It has these top-level messages:
	Endpoint
	Location
	Refdata
	EndpointRequest
	EndpointResponse
	StoreGetRequest
	StoreGetResponse
	StorePutRequest
	StorePutResponse
	StoreDeleteRequest
	StoreDeleteResponse
	User
	KeyLookupRequest
	KeyLookupResponse
	KeyPutRequest
	KeyPutResponse
	EntryError
	EntriesError
	DirLookupRequest
	DirPutRequest
	DirGlobRequest
	DirDeleteRequest
	DirWhichAccessRequest
	DirWatchRequest
	Event
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

// Endpoint mirrors upspin.Endpoint.
type Endpoint struct {
	Transport int32  `protobuf:"varint,1,opt,name=transport" json:"transport,omitempty"`
	NetAddr   string `protobuf:"bytes,2,opt,name=net_addr,json=netAddr" json:"net_addr,omitempty"`
}

func (m *Endpoint) Reset()                    { *m = Endpoint{} }
func (m *Endpoint) String() string            { return proto1.CompactTextString(m) }
func (*Endpoint) ProtoMessage()               {}
func (*Endpoint) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Endpoint) GetTransport() int32 {
	if m != nil {
		return m.Transport
	}
	return 0
}

func (m *Endpoint) GetNetAddr() string {
	if m != nil {
		return m.NetAddr
	}
	return ""
}

// Location mirrors upspin.Location.
type Location struct {
	Endpoint  *Endpoint `protobuf:"bytes,1,opt,name=endpoint" json:"endpoint,omitempty"`
	Reference string    `protobuf:"bytes,2,opt,name=reference" json:"reference,omitempty"`
}

func (m *Location) Reset()                    { *m = Location{} }
func (m *Location) String() string            { return proto1.CompactTextString(m) }
func (*Location) ProtoMessage()               {}
func (*Location) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Location) GetEndpoint() *Endpoint {
	if m != nil {
		return m.Endpoint
	}
	return nil
}

func (m *Location) GetReference() string {
	if m != nil {
		return m.Reference
	}
	return ""
}

// Refdata mirrors upspin.Refdata.
type Refdata struct {
	Reference string `protobuf:"bytes,1,opt,name=reference" json:"reference,omitempty"`
	Volatile  bool   `protobuf:"varint,2,opt,name=volatile" json:"volatile,omitempty"`
	Duration  int64  `protobuf:"varint,3,opt,name=duration" json:"duration,omitempty"`
}

func (m *Refdata) Reset()                    { *m = Refdata{} }
func (m *Refdata) String() string            { return proto1.CompactTextString(m) }
func (*Refdata) ProtoMessage()               {}
func (*Refdata) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Refdata) GetReference() string {
	if m != nil {
		return m.Reference
	}
	return ""
}

func (m *Refdata) GetVolatile() bool {
	if m != nil {
		return m.Volatile
	}
	return false
}

func (m *Refdata) GetDuration() int64 {
	if m != nil {
		return m.Duration
	}
	return 0
}

type EndpointRequest struct {
}

func (m *EndpointRequest) Reset()                    { *m = EndpointRequest{} }
func (m *EndpointRequest) String() string            { return proto1.CompactTextString(m) }
func (*EndpointRequest) ProtoMessage()               {}
func (*EndpointRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type EndpointResponse struct {
	Endpoint *Endpoint `protobuf:"bytes,1,opt,name=endpoint" json:"endpoint,omitempty"`
}

func (m *EndpointResponse) Reset()                    { *m = EndpointResponse{} }
func (m *EndpointResponse) String() string            { return proto1.CompactTextString(m) }
func (*EndpointResponse) ProtoMessage()               {}
func (*EndpointResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *EndpointResponse) GetEndpoint() *Endpoint {
	if m != nil {
		return m.Endpoint
	}
	return nil
}

type StoreGetRequest struct {
	Reference string `protobuf:"bytes,1,opt,name=reference" json:"reference,omitempty"`
}

func (m *StoreGetRequest) Reset()                    { *m = StoreGetRequest{} }
func (m *StoreGetRequest) String() string            { return proto1.CompactTextString(m) }
func (*StoreGetRequest) ProtoMessage()               {}
func (*StoreGetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *StoreGetRequest) GetReference() string {
	if m != nil {
		return m.Reference
	}
	return ""
}

type StoreGetResponse struct {
	Data      []byte      `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Refdata   *Refdata    `protobuf:"bytes,2,opt,name=refdata" json:"refdata,omitempty"`
	Locations []*Location `protobuf:"bytes,3,rep,name=locations" json:"locations,omitempty"`
	Error     []byte      `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
}

func (m *StoreGetResponse) Reset()                    { *m = StoreGetResponse{} }
func (m *StoreGetResponse) String() string            { return proto1.CompactTextString(m) }
func (*StoreGetResponse) ProtoMessage()               {}
func (*StoreGetResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *StoreGetResponse) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *StoreGetResponse) GetRefdata() *Refdata {
	if m != nil {
		return m.Refdata
	}
	return nil
}

func (m *StoreGetResponse) GetLocations() []*Location {
	if m != nil {
		return m.Locations
	}
	return nil
}

func (m *StoreGetResponse) GetError() []byte {
	if m != nil {
		return m.Error
	}
	return nil
}

type StorePutRequest struct {
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *StorePutRequest) Reset()                    { *m = StorePutRequest{} }
func (m *StorePutRequest) String() string            { return proto1.CompactTextString(m) }
func (*StorePutRequest) ProtoMessage()               {}
func (*StorePutRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *StorePutRequest) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type StorePutResponse struct {
	Refdata *Refdata `protobuf:"bytes,1,opt,name=refdata" json:"refdata,omitempty"`
	Error   []byte   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (m *StorePutResponse) Reset()                    { *m = StorePutResponse{} }
func (m *StorePutResponse) String() string            { return proto1.CompactTextString(m) }
func (*StorePutResponse) ProtoMessage()               {}
func (*StorePutResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *StorePutResponse) GetRefdata() *Refdata {
	if m != nil {
		return m.Refdata
	}
	return nil
}

func (m *StorePutResponse) GetError() []byte {
	if m != nil {
		return m.Error
	}
	return nil
}

type StoreDeleteRequest struct {
	Reference string `protobuf:"bytes,1,opt,name=reference" json:"reference,omitempty"`
}

func (m *StoreDeleteRequest) Reset()                    { *m = StoreDeleteRequest{} }
func (m *StoreDeleteRequest) String() string            { return proto1.CompactTextString(m) }
func (*StoreDeleteRequest) ProtoMessage()               {}
func (*StoreDeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *StoreDeleteRequest) GetReference() string {
	if m != nil {
		return m.Reference
	}
	return ""
}

type StoreDeleteResponse struct {
	Error []byte `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
}

func (m *StoreDeleteResponse) Reset()                    { *m = StoreDeleteResponse{} }
func (m *StoreDeleteResponse) String() string            { return proto1.CompactTextString(m) }
func (*StoreDeleteResponse) ProtoMessage()               {}
func (*StoreDeleteResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *StoreDeleteResponse) GetError() []byte {
	if m != nil {
		return m.Error
	}
	return nil
}

type User struct {
	Name      string      `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Dirs      []*Endpoint `protobuf:"bytes,2,rep,name=dirs" json:"dirs,omitempty"`
	Stores    []*Endpoint `protobuf:"bytes,3,rep,name=stores" json:"stores,omitempty"`
	PublicKey string      `protobuf:"bytes,4,opt,name=public_key,json=publicKey" json:"public_key,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto1.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetDirs() []*Endpoint {
	if m != nil {
		return m.Dirs
	}
	return nil
}

func (m *User) GetStores() []*Endpoint {
	if m != nil {
		return m.Stores
	}
	return nil
}

func (m *User) GetPublicKey() string {
	if m != nil {
		return m.PublicKey
	}
	return ""
}

type KeyLookupRequest struct {
	UserName string `protobuf:"bytes,1,opt,name=user_name,json=userName" json:"user_name,omitempty"`
}

func (m *KeyLookupRequest) Reset()                    { *m = KeyLookupRequest{} }
func (m *KeyLookupRequest) String() string            { return proto1.CompactTextString(m) }
func (*KeyLookupRequest) ProtoMessage()               {}
func (*KeyLookupRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *KeyLookupRequest) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

type KeyLookupResponse struct {
	User  *User  `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
	Error []byte `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (m *KeyLookupResponse) Reset()                    { *m = KeyLookupResponse{} }
func (m *KeyLookupResponse) String() string            { return proto1.CompactTextString(m) }
func (*KeyLookupResponse) ProtoMessage()               {}
func (*KeyLookupResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *KeyLookupResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *KeyLookupResponse) GetError() []byte {
	if m != nil {
		return m.Error
	}
	return nil
}

type KeyPutRequest struct {
	User *User `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
}

func (m *KeyPutRequest) Reset()                    { *m = KeyPutRequest{} }
func (m *KeyPutRequest) String() string            { return proto1.CompactTextString(m) }
func (*KeyPutRequest) ProtoMessage()               {}
func (*KeyPutRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *KeyPutRequest) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type KeyPutResponse struct {
	Error []byte `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
}

func (m *KeyPutResponse) Reset()                    { *m = KeyPutResponse{} }
func (m *KeyPutResponse) String() string            { return proto1.CompactTextString(m) }
func (*KeyPutResponse) ProtoMessage()               {}
func (*KeyPutResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func (m *KeyPutResponse) GetError() []byte {
	if m != nil {
		return m.Error
	}
	return nil
}

type EntryError struct {
	Entry []byte `protobuf:"bytes,1,opt,name=entry,proto3" json:"entry,omitempty"`
	Error []byte `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (m *EntryError) Reset()                    { *m = EntryError{} }
func (m *EntryError) String() string            { return proto1.CompactTextString(m) }
func (*EntryError) ProtoMessage()               {}
func (*EntryError) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{16} }

func (m *EntryError) GetEntry() []byte {
	if m != nil {
		return m.Entry
	}
	return nil
}

func (m *EntryError) GetError() []byte {
	if m != nil {
		return m.Error
	}
	return nil
}

type EntriesError struct {
	Entries [][]byte `protobuf:"bytes,1,rep,name=entries,proto3" json:"entries,omitempty"`
	Error   []byte   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (m *EntriesError) Reset()                    { *m = EntriesError{} }
func (m *EntriesError) String() string            { return proto1.CompactTextString(m) }
func (*EntriesError) ProtoMessage()               {}
func (*EntriesError) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{17} }

func (m *EntriesError) GetEntries() [][]byte {
	if m != nil {
		return m.Entries
	}
	return nil
}

func (m *EntriesError) GetError() []byte {
	if m != nil {
		return m.Error
	}
	return nil
}

type DirLookupRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *DirLookupRequest) Reset()                    { *m = DirLookupRequest{} }
func (m *DirLookupRequest) String() string            { return proto1.CompactTextString(m) }
func (*DirLookupRequest) ProtoMessage()               {}
func (*DirLookupRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{18} }

func (m *DirLookupRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type DirPutRequest struct {
	Entry []byte `protobuf:"bytes,1,opt,name=entry,proto3" json:"entry,omitempty"`
}

func (m *DirPutRequest) Reset()                    { *m = DirPutRequest{} }
func (m *DirPutRequest) String() string            { return proto1.CompactTextString(m) }
func (*DirPutRequest) ProtoMessage()               {}
func (*DirPutRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{19} }

func (m *DirPutRequest) GetEntry() []byte {
	if m != nil {
		return m.Entry
	}
	return nil
}

type DirGlobRequest struct {
	Pattern string `protobuf:"bytes,1,opt,name=pattern" json:"pattern,omitempty"`
}

func (m *DirGlobRequest) Reset()                    { *m = DirGlobRequest{} }
func (m *DirGlobRequest) String() string            { return proto1.CompactTextString(m) }
func (*DirGlobRequest) ProtoMessage()               {}
func (*DirGlobRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{20} }

func (m *DirGlobRequest) GetPattern() string {
	if m != nil {
		return m.Pattern
	}
	return ""
}

type DirDeleteRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *DirDeleteRequest) Reset()                    { *m = DirDeleteRequest{} }
func (m *DirDeleteRequest) String() string            { return proto1.CompactTextString(m) }
func (*DirDeleteRequest) ProtoMessage()               {}
func (*DirDeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{21} }

func (m *DirDeleteRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type DirWhichAccessRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *DirWhichAccessRequest) Reset()                    { *m = DirWhichAccessRequest{} }
func (m *DirWhichAccessRequest) String() string            { return proto1.CompactTextString(m) }
func (*DirWhichAccessRequest) ProtoMessage()               {}
func (*DirWhichAccessRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{22} }

func (m *DirWhichAccessRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type DirWatchRequest struct {
	Name     string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Sequence int64  `protobuf:"varint,2,opt,name=sequence" json:"sequence,omitempty"`
}

func (m *DirWatchRequest) Reset()                    { *m = DirWatchRequest{} }
func (m *DirWatchRequest) String() string            { return proto1.CompactTextString(m) }
func (*DirWatchRequest) ProtoMessage()               {}
func (*DirWatchRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{23} }

func (m *DirWatchRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DirWatchRequest) GetSequence() int64 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

// The first response in the stream is whether dir.Watch succeeded. If it
// didn't, the error field contains the error and no streaming happens. If it
// did succeed the error is nil and subsequent streams are from the Events
// channel.
type Event struct {
	Entry    []byte `protobuf:"bytes,1,opt,name=entry,proto3" json:"entry,omitempty"`
	Sequence int64  `protobuf:"varint,2,opt,name=sequence" json:"sequence,omitempty"`
	Delete   bool   `protobuf:"varint,3,opt,name=delete" json:"delete,omitempty"`
	Error    []byte `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
}

func (m *Event) Reset()                    { *m = Event{} }
func (m *Event) String() string            { return proto1.CompactTextString(m) }
func (*Event) ProtoMessage()               {}
func (*Event) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{24} }

func (m *Event) GetEntry() []byte {
	if m != nil {
		return m.Entry
	}
	return nil
}

func (m *Event) GetSequence() int64 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

func (m *Event) GetDelete() bool {
	if m != nil {
		return m.Delete
	}
	return false
}

func (m *Event) GetError() []byte {
	if m != nil {
		return m.Error
	}
	return nil
}

func init() {
	proto1.RegisterType((*Endpoint)(nil), "proto.Endpoint")
	proto1.RegisterType((*Location)(nil), "proto.Location")
	proto1.RegisterType((*Refdata)(nil), "proto.Refdata")
	proto1.RegisterType((*EndpointRequest)(nil), "proto.EndpointRequest")
	proto1.RegisterType((*EndpointResponse)(nil), "proto.EndpointResponse")
	proto1.RegisterType((*StoreGetRequest)(nil), "proto.StoreGetRequest")
	proto1.RegisterType((*StoreGetResponse)(nil), "proto.StoreGetResponse")
	proto1.RegisterType((*StorePutRequest)(nil), "proto.StorePutRequest")
	proto1.RegisterType((*StorePutResponse)(nil), "proto.StorePutResponse")
	proto1.RegisterType((*StoreDeleteRequest)(nil), "proto.StoreDeleteRequest")
	proto1.RegisterType((*StoreDeleteResponse)(nil), "proto.StoreDeleteResponse")
	proto1.RegisterType((*User)(nil), "proto.User")
	proto1.RegisterType((*KeyLookupRequest)(nil), "proto.KeyLookupRequest")
	proto1.RegisterType((*KeyLookupResponse)(nil), "proto.KeyLookupResponse")
	proto1.RegisterType((*KeyPutRequest)(nil), "proto.KeyPutRequest")
	proto1.RegisterType((*KeyPutResponse)(nil), "proto.KeyPutResponse")
	proto1.RegisterType((*EntryError)(nil), "proto.EntryError")
	proto1.RegisterType((*EntriesError)(nil), "proto.EntriesError")
	proto1.RegisterType((*DirLookupRequest)(nil), "proto.DirLookupRequest")
	proto1.RegisterType((*DirPutRequest)(nil), "proto.DirPutRequest")
	proto1.RegisterType((*DirGlobRequest)(nil), "proto.DirGlobRequest")
	proto1.RegisterType((*DirDeleteRequest)(nil), "proto.DirDeleteRequest")
	proto1.RegisterType((*DirWhichAccessRequest)(nil), "proto.DirWhichAccessRequest")
	proto1.RegisterType((*DirWatchRequest)(nil), "proto.DirWatchRequest")
	proto1.RegisterType((*Event)(nil), "proto.Event")
}

func init() { proto1.RegisterFile("upspin.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 839 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0x5f, 0x8f, 0xdb, 0x44,
	0x10, 0x8f, 0xcf, 0xf9, 0xe3, 0x4c, 0xd2, 0x4b, 0x6e, 0xdb, 0x2b, 0xae, 0x29, 0x22, 0x5a, 0xd4,
	0x12, 0x71, 0xa2, 0x3d, 0x42, 0x85, 0xfa, 0x52, 0x20, 0x22, 0xd1, 0x49, 0xa4, 0x42, 0x95, 0x51,
	0xc5, 0x63, 0xe4, 0x8b, 0xb7, 0x9c, 0xd5, 0xd4, 0x36, 0xeb, 0xf5, 0x49, 0xf9, 0x04, 0x3c, 0xf3,
	0xc0, 0x87, 0xe1, 0xc3, 0x21, 0xa1, 0x5d, 0xef, 0xae, 0xd7, 0x8e, 0x13, 0x40, 0xf7, 0x74, 0x99,
	0x9d, 0xdf, 0x6f, 0xe6, 0x37, 0x33, 0x9e, 0x39, 0x18, 0xe6, 0x69, 0x96, 0x46, 0xf1, 0xb3, 0x94,
	0x26, 0x2c, 0x41, 0x1d, 0xf1, 0x07, 0xff, 0x00, 0xce, 0x32, 0x0e, 0xd3, 0x24, 0x8a, 0x19, 0x7a,
	0x0c, 0x7d, 0x46, 0x83, 0x38, 0x4b, 0x13, 0xca, 0x5c, 0x6b, 0x62, 0x4d, 0x3b, 0x7e, 0xf9, 0x80,
	0x1e, 0x81, 0x13, 0x13, 0xb6, 0x0e, 0xc2, 0x90, 0xba, 0x27, 0x13, 0x6b, 0xda, 0xf7, 0x7b, 0x31,
	0x61, 0xf3, 0x30, 0xa4, 0xf8, 0x2d, 0x38, 0xaf, 0x93, 0x4d, 0xc0, 0xa2, 0x24, 0x46, 0x17, 0xe0,
	0x10, 0x19, 0x50, 0xc4, 0x18, 0xcc, 0x46, 0x45, 0xc6, 0x67, 0x2a, 0x8f, 0xaf, 0x01, 0x3c, 0x23,
	0x25, 0xef, 0x08, 0x25, 0xf1, 0x86, 0xc8, 0xa0, 0xe5, 0x03, 0x5e, 0x43, 0xcf, 0x27, 0xef, 0xc2,
	0x80, 0x05, 0x55, 0xa0, 0x55, 0x03, 0x22, 0x0f, 0x9c, 0xdb, 0x64, 0x1b, 0xb0, 0x68, 0x5b, 0x44,
	0x71, 0x7c, 0x6d, 0x73, 0x5f, 0x98, 0x53, 0xa1, 0xcd, 0xb5, 0x27, 0xd6, 0xd4, 0xf6, 0xb5, 0x8d,
	0xcf, 0x60, 0xa4, 0x45, 0x91, 0xdf, 0x72, 0x92, 0x31, 0xfc, 0x1d, 0x8c, 0xcb, 0xa7, 0x2c, 0x4d,
	0xe2, 0x8c, 0xfc, 0xaf, 0x92, 0xf0, 0x73, 0x18, 0xfd, 0xcc, 0x12, 0x4a, 0xae, 0x88, 0x8a, 0x79,
	0x5c, 0x3c, 0xfe, 0xd3, 0x82, 0x71, 0xc9, 0x90, 0x29, 0x11, 0xb4, 0x79, 0xdd, 0x02, 0x3d, 0xf4,
	0xc5, 0x6f, 0x34, 0x85, 0x1e, 0x2d, 0xda, 0x21, 0x8a, 0x1c, 0xcc, 0x4e, 0xa5, 0x0a, 0xd9, 0x24,
	0x5f, 0xb9, 0xd1, 0x97, 0xd0, 0xdf, 0xca, 0x79, 0x64, 0xae, 0x3d, 0xb1, 0x0d, 0xc5, 0x6a, 0x4e,
	0x7e, 0x89, 0x40, 0x0f, 0xa0, 0x43, 0x28, 0x4d, 0xa8, 0xdb, 0x16, 0xd9, 0x0a, 0x03, 0x3f, 0x91,
	0x85, 0xbc, 0xc9, 0x75, 0x21, 0x0d, 0xaa, 0xb0, 0x2f, 0xd5, 0x0b, 0x98, 0x54, 0x6f, 0x28, 0xb5,
	0x8e, 0x2b, 0xd5, 0xa9, 0x4f, 0xcc, 0xd4, 0x33, 0x40, 0x22, 0xe6, 0x82, 0x6c, 0x09, 0x23, 0xff,
	0xad, 0x8d, 0x17, 0x70, 0xbf, 0xc2, 0x91, 0x52, 0x74, 0x02, 0xcb, 0x4c, 0xf0, 0xbb, 0x05, 0xed,
	0xb7, 0x19, 0xa1, 0xbc, 0xa2, 0x38, 0xf8, 0xa0, 0xc2, 0x89, 0xdf, 0xe8, 0x33, 0x68, 0x87, 0x11,
	0xcd, 0xdc, 0x93, 0x4a, 0xe3, 0xf4, 0xa8, 0x85, 0x13, 0x7d, 0x0e, 0xdd, 0x8c, 0xa7, 0xab, 0xf7,
	0x57, 0xc3, 0xa4, 0x1b, 0x7d, 0x02, 0x90, 0xe6, 0xd7, 0xdb, 0x68, 0xb3, 0x7e, 0x4f, 0x76, 0xa2,
	0xc3, 0x7d, 0xbf, 0x5f, 0xbc, 0xac, 0xc8, 0x0e, 0x3f, 0x87, 0xf1, 0x8a, 0xec, 0x5e, 0x27, 0xc9,
	0xfb, 0x3c, 0x55, 0x85, 0x7e, 0x0c, 0xfd, 0x3c, 0x23, 0x74, 0x6d, 0x28, 0x73, 0xf8, 0xc3, 0x4f,
	0xc1, 0x07, 0x82, 0x7f, 0x84, 0x33, 0x83, 0x20, 0xab, 0xfc, 0x14, 0xda, 0x1c, 0x20, 0xbb, 0x3d,
	0x90, 0x5a, 0x78, 0x85, 0xbe, 0x70, 0x1c, 0xe8, 0xf3, 0x25, 0xdc, 0x5b, 0x91, 0x9d, 0x31, 0xe0,
	0x7f, 0x8b, 0x83, 0x9f, 0xc2, 0xa9, 0x62, 0x1c, 0x6d, 0xf0, 0x4b, 0x80, 0x65, 0xcc, 0xe8, 0x6e,
	0xc9, 0x2d, 0x81, 0xe1, 0x96, 0xc6, 0x70, 0xe3, 0x80, 0xa6, 0x6f, 0x61, 0xc8, 0x99, 0x11, 0xc9,
	0x0a, 0xae, 0x0b, 0x3d, 0x52, 0xd8, 0xae, 0x35, 0xb1, 0xa7, 0x43, 0x5f, 0x99, 0x07, 0xf8, 0x4f,
	0x61, 0xbc, 0x88, 0x68, 0xb5, 0xa1, 0x0d, 0x53, 0xc6, 0x4f, 0xe0, 0xde, 0x22, 0xa2, 0x46, 0xed,
	0x8d, 0x22, 0xf1, 0x17, 0x70, 0xba, 0x88, 0xe8, 0xd5, 0x36, 0xb9, 0x56, 0x38, 0x17, 0x7a, 0x69,
	0xc0, 0x18, 0xa1, 0xb1, 0x8c, 0xa7, 0x4c, 0x99, 0xba, 0xfa, 0xd1, 0x36, 0xa5, 0xbe, 0x80, 0xf3,
	0x45, 0x44, 0x7f, 0xb9, 0x89, 0x36, 0x37, 0xf3, 0xcd, 0x86, 0x64, 0xd9, 0x31, 0xf0, 0x1c, 0x46,
	0x1c, 0x1c, 0xb0, 0xcd, 0xcd, 0x11, 0x18, 0x3f, 0x73, 0x19, 0x77, 0xab, 0x43, 0x6a, 0xfb, 0xda,
	0xc6, 0xbf, 0x42, 0x67, 0x79, 0x4b, 0xe2, 0x03, 0x25, 0x1e, 0xa3, 0xa2, 0x87, 0xd0, 0x0d, 0x45,
	0x3d, 0xe2, 0x76, 0x3a, 0xbe, 0xb4, 0x9a, 0x4f, 0xc6, 0xec, 0x6f, 0x0b, 0x3a, 0x62, 0x09, 0xd1,
	0x2b, 0xe3, 0xdf, 0xca, 0xc3, 0xfa, 0x6a, 0x14, 0x65, 0x78, 0x1f, 0xed, 0xbd, 0x17, 0x9f, 0x14,
	0x6e, 0xa1, 0x97, 0x60, 0x5f, 0x91, 0x92, 0x59, 0x3b, 0xa8, 0x9a, 0x59, 0x3f, 0x9b, 0x05, 0xf3,
	0x4d, 0x5e, 0x63, 0x96, 0x43, 0xae, 0x32, 0x8d, 0xcf, 0x18, 0xb7, 0xd0, 0x1c, 0xba, 0xc5, 0xe8,
	0xd0, 0x23, 0x13, 0x54, 0x19, 0xa7, 0xe7, 0x35, 0xb9, 0x54, 0x88, 0xd9, 0x5f, 0x16, 0xd8, 0x2b,
	0xb2, 0xbb, 0x6b, 0xf5, 0xaf, 0xa0, 0x5b, 0x7c, 0xbf, 0x48, 0x81, 0xea, 0x27, 0xc2, 0x73, 0xf7,
	0x1d, 0x9a, 0xfe, 0xa2, 0x68, 0xc1, 0x83, 0x12, 0x62, 0x34, 0xe0, 0xbc, 0xf6, 0xaa, 0xb5, 0xff,
	0x61, 0x83, 0xbd, 0x88, 0xe8, 0x5d, 0xb5, 0x7f, 0xb3, 0xa7, 0xbd, 0xbe, 0x8d, 0xde, 0x99, 0x66,
	0xab, 0x03, 0x81, 0x5b, 0xe8, 0xb2, 0x2a, 0xba, 0xb2, 0x9a, 0xcd, 0x8c, 0x17, 0xd0, 0xe6, 0x6b,
	0x89, 0xce, 0x4b, 0x8a, 0xb1, 0xa6, 0xde, 0x7d, 0x83, 0xa3, 0x8e, 0x49, 0xa1, 0x4f, 0x4e, 0xd9,
	0xd0, 0x57, 0x9d, 0x71, 0x63, 0xb6, 0xef, 0x61, 0x60, 0x2c, 0x2c, 0x7a, 0x5c, 0x92, 0xf7, 0xf7,
	0xb8, 0x39, 0xc2, 0x57, 0xd0, 0x11, 0x5b, 0xac, 0xbb, 0x5a, 0x5b, 0x6b, 0x6f, 0xa8, 0x58, 0x7c,
	0x57, 0x71, 0xeb, 0xd2, 0xba, 0xee, 0x8a, 0x87, 0xaf, 0xff, 0x09, 0x00, 0x00, 0xff, 0xff, 0xe7,
	0x95, 0x7f, 0xb4, 0xba, 0x09, 0x00, 0x00,
}