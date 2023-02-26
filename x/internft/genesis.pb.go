// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/internft/v1alpha1/genesis.proto

package internft

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// GenesisState defines the module's genesis state.
type GenesisState struct {
	// all the paramaters of the module
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	// classes
	Classes []GenesisClass `protobuf:"bytes,2,rep,name=classes,proto3" json:"classes"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_3c25a780c407fed1, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetClasses() []GenesisClass {
	if m != nil {
		return m.Classes
	}
	return nil
}

// GenesisClass defines a class and its relevant data.
type GenesisClass struct {
	// identifier of the class
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// traits of the class
	Traits []Trait `protobuf:"bytes,2,rep,name=traits,proto3" json:"traits"`
	// last minted nft of the class
	LastMintedNftId github_com_cosmos_cosmos_sdk_types.Uint `protobuf:"bytes,3,opt,name=last_minted_nft_id,json=lastMintedNftId,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Uint" json:"last_minted_nft_id"`
	// nfts of the class
	Nfts []GenesisNFT `protobuf:"bytes,4,rep,name=nfts,proto3" json:"nfts"`
}

func (m *GenesisClass) Reset()         { *m = GenesisClass{} }
func (m *GenesisClass) String() string { return proto.CompactTextString(m) }
func (*GenesisClass) ProtoMessage()    {}
func (*GenesisClass) Descriptor() ([]byte, []int) {
	return fileDescriptor_3c25a780c407fed1, []int{1}
}
func (m *GenesisClass) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisClass) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisClass.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisClass) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisClass.Merge(m, src)
}
func (m *GenesisClass) XXX_Size() int {
	return m.Size()
}
func (m *GenesisClass) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisClass.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisClass proto.InternalMessageInfo

func (m *GenesisClass) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GenesisClass) GetTraits() []Trait {
	if m != nil {
		return m.Traits
	}
	return nil
}

func (m *GenesisClass) GetNfts() []GenesisNFT {
	if m != nil {
		return m.Nfts
	}
	return nil
}

// GenesisNFT defines an nft and its relevant data.
type GenesisNFT struct {
	// identifier of the nft
	Id github_com_cosmos_cosmos_sdk_types.Uint `protobuf:"bytes,1,opt,name=id,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Uint" json:"id"`
	// properties of the nft
	Properties []Property `protobuf:"bytes,2,rep,name=properties,proto3" json:"properties"`
	// owner of the nft
	Owner string `protobuf:"bytes,3,opt,name=owner,proto3" json:"owner,omitempty"`
}

func (m *GenesisNFT) Reset()         { *m = GenesisNFT{} }
func (m *GenesisNFT) String() string { return proto.CompactTextString(m) }
func (*GenesisNFT) ProtoMessage()    {}
func (*GenesisNFT) Descriptor() ([]byte, []int) {
	return fileDescriptor_3c25a780c407fed1, []int{2}
}
func (m *GenesisNFT) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisNFT) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisNFT.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisNFT) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisNFT.Merge(m, src)
}
func (m *GenesisNFT) XXX_Size() int {
	return m.Size()
}
func (m *GenesisNFT) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisNFT.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisNFT proto.InternalMessageInfo

func (m *GenesisNFT) GetProperties() []Property {
	if m != nil {
		return m.Properties
	}
	return nil
}

func (m *GenesisNFT) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "cosmos.internft.v1alpha1.GenesisState")
	proto.RegisterType((*GenesisClass)(nil), "cosmos.internft.v1alpha1.GenesisClass")
	proto.RegisterType((*GenesisNFT)(nil), "cosmos.internft.v1alpha1.GenesisNFT")
}

func init() {
	proto.RegisterFile("cosmos/internft/v1alpha1/genesis.proto", fileDescriptor_3c25a780c407fed1)
}

var fileDescriptor_3c25a780c407fed1 = []byte{
	// 403 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x4f, 0x8b, 0xda, 0x40,
	0x18, 0xc6, 0x13, 0xb5, 0x96, 0x8e, 0xa5, 0x85, 0xc1, 0x43, 0xf0, 0x10, 0x25, 0x88, 0xf5, 0xd2,
	0x04, 0xed, 0xb9, 0x96, 0x5a, 0xb0, 0xed, 0xa1, 0x52, 0x52, 0x7b, 0x29, 0x05, 0x19, 0xcd, 0x24,
	0x0e, 0x35, 0x99, 0x90, 0x99, 0xee, 0xae, 0xdf, 0x62, 0x4f, 0xfb, 0x2d, 0xf6, 0x7b, 0x78, 0xf4,
	0xb8, 0xec, 0x41, 0x16, 0xfd, 0x18, 0x7b, 0x59, 0xe6, 0x4f, 0x74, 0x2f, 0x71, 0xd9, 0x53, 0x12,
	0xe6, 0xf7, 0x3e, 0xcf, 0x93, 0x67, 0x5e, 0xd0, 0x99, 0x53, 0x16, 0x53, 0xe6, 0x91, 0x84, 0xe3,
	0x2c, 0x09, 0xb9, 0x77, 0xd6, 0x43, 0xcb, 0x74, 0x81, 0x7a, 0x5e, 0x84, 0x13, 0xcc, 0x08, 0x73,
	0xd3, 0x8c, 0x72, 0x0a, 0x2d, 0xc5, 0xb9, 0x39, 0xe7, 0xe6, 0x5c, 0xa3, 0x1e, 0xd1, 0x88, 0x4a,
	0xc8, 0x13, 0x6f, 0x8a, 0x6f, 0xb4, 0x0b, 0x75, 0xf9, 0x2a, 0xc5, 0x5a, 0xd5, 0xb9, 0x32, 0xc1,
	0xeb, 0xaf, 0xca, 0xe7, 0x17, 0x47, 0x1c, 0xc3, 0x01, 0xa8, 0xa6, 0x28, 0x43, 0x31, 0xb3, 0xcc,
	0x96, 0xd9, 0xad, 0xf5, 0x5b, 0x6e, 0x91, 0xaf, 0xfb, 0x53, 0x72, 0xc3, 0xca, 0x7a, 0xdb, 0x34,
	0x7c, 0x3d, 0x05, 0x47, 0xe0, 0xe5, 0x7c, 0x89, 0x18, 0xc3, 0xcc, 0x2a, 0xb5, 0xca, 0xdd, 0x5a,
	0xbf, 0x53, 0x2c, 0xa0, 0x8d, 0xbf, 0x08, 0x5e, 0xcb, 0xe4, 0xc3, 0xce, 0xfd, 0x31, 0x98, 0x3c,
	0x87, 0x6f, 0x40, 0x89, 0x04, 0x32, 0xd4, 0x2b, 0xbf, 0x44, 0x02, 0xf8, 0x11, 0x54, 0x79, 0x86,
	0x08, 0xcf, 0x7d, 0x9a, 0xc5, 0x3e, 0x13, 0xc1, 0xe5, 0x39, 0xd5, 0x10, 0xfc, 0x0b, 0xe0, 0x12,
	0x31, 0x3e, 0x8d, 0x05, 0x1e, 0x4c, 0x93, 0x90, 0x4f, 0x49, 0x60, 0x95, 0x85, 0xfc, 0xd0, 0x13,
	0xe4, 0xed, 0xb6, 0xf9, 0x2e, 0x22, 0x7c, 0xf1, 0x7f, 0xe6, 0xce, 0x69, 0xec, 0xe9, 0x36, 0xd5,
	0xe3, 0x3d, 0x0b, 0xfe, 0xe9, 0x1a, 0x7f, 0x93, 0x84, 0xfb, 0x6f, 0x85, 0xd4, 0x0f, 0xa9, 0x34,
	0x0e, 0xf9, 0xf7, 0x00, 0x0e, 0x40, 0x25, 0x09, 0x39, 0xb3, 0x2a, 0x32, 0x5a, 0xfb, 0xc9, 0x0a,
	0xc6, 0xa3, 0x89, 0xce, 0x27, 0xe7, 0x9c, 0x6b, 0x13, 0x80, 0xe3, 0x11, 0xfc, 0x74, 0xfc, 0xf7,
	0xe7, 0x87, 0x13, 0x65, 0x7d, 0x03, 0x20, 0xcd, 0x68, 0x8a, 0x33, 0x4e, 0x0e, 0x17, 0xe3, 0x9c,
	0xb8, 0x59, 0xc5, 0xae, 0x74, 0xa6, 0x47, 0xb3, 0xb0, 0x0e, 0x5e, 0xd0, 0xf3, 0x04, 0x67, 0xaa,
	0x2a, 0x5f, 0x7d, 0x0c, 0x3f, 0xaf, 0x77, 0xb6, 0xb9, 0xd9, 0xd9, 0xe6, 0xdd, 0xce, 0x36, 0x2f,
	0xf7, 0xb6, 0xb1, 0xd9, 0xdb, 0xc6, 0xcd, 0xde, 0x36, 0xfe, 0x9c, 0x8e, 0x79, 0x71, 0x58, 0xcf,
	0x59, 0x55, 0x2e, 0xe4, 0x87, 0x87, 0x00, 0x00, 0x00, 0xff, 0xff, 0x09, 0x5e, 0x06, 0xe9, 0x10,
	0x03, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Classes) > 0 {
		for iNdEx := len(m.Classes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Classes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *GenesisClass) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisClass) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisClass) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Nfts) > 0 {
		for iNdEx := len(m.Nfts) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Nfts[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	{
		size := m.LastMintedNftId.Size()
		i -= size
		if _, err := m.LastMintedNftId.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.Traits) > 0 {
		for iNdEx := len(m.Traits) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Traits[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GenesisNFT) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisNFT) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisNFT) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Properties) > 0 {
		for iNdEx := len(m.Properties) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Properties[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size := m.Id.Size()
		i -= size
		if _, err := m.Id.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.Classes) > 0 {
		for _, e := range m.Classes {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *GenesisClass) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.Traits) > 0 {
		for _, e := range m.Traits {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	l = m.LastMintedNftId.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.Nfts) > 0 {
		for _, e := range m.Nfts {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *GenesisNFT) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Id.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.Properties) > 0 {
		for _, e := range m.Properties {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Classes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Classes = append(m.Classes, GenesisClass{})
			if err := m.Classes[len(m.Classes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GenesisClass) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GenesisClass: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisClass: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Traits", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Traits = append(m.Traits, Trait{})
			if err := m.Traits[len(m.Traits)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastMintedNftId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.LastMintedNftId.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nfts", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Nfts = append(m.Nfts, GenesisNFT{})
			if err := m.Nfts[len(m.Nfts)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GenesisNFT) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GenesisNFT: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisNFT: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Id.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Properties", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Properties = append(m.Properties, Property{})
			if err := m.Properties[len(m.Properties)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
