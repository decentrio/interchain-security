// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: interchain_security/ccv/child/v1/child.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/tendermint/tendermint/abci/types"
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

// Params defines the parameters for CCV child module
type Params struct {
	Enabled bool `protobuf:"varint,1,opt,name=Enabled,proto3" json:"Enabled,omitempty"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_d00ed7c78fea69b9, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

// CrossChainValidator defines the validators for CCV child module
type CrossChainValidator struct {
	Address         []byte                `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	ValidatorUpdate types.ValidatorUpdate `protobuf:"bytes,2,opt,name=validator_update,json=validatorUpdate,proto3" json:"validator_update" yaml:"last_update"`
}

func (m *CrossChainValidator) Reset()         { *m = CrossChainValidator{} }
func (m *CrossChainValidator) String() string { return proto.CompactTextString(m) }
func (*CrossChainValidator) ProtoMessage()    {}
func (*CrossChainValidator) Descriptor() ([]byte, []int) {
	return fileDescriptor_d00ed7c78fea69b9, []int{1}
}
func (m *CrossChainValidator) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CrossChainValidator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CrossChainValidator.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CrossChainValidator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CrossChainValidator.Merge(m, src)
}
func (m *CrossChainValidator) XXX_Size() int {
	return m.Size()
}
func (m *CrossChainValidator) XXX_DiscardUnknown() {
	xxx_messageInfo_CrossChainValidator.DiscardUnknown(m)
}

var xxx_messageInfo_CrossChainValidator proto.InternalMessageInfo

func (m *CrossChainValidator) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *CrossChainValidator) GetValidatorUpdate() types.ValidatorUpdate {
	if m != nil {
		return m.ValidatorUpdate
	}
	return types.ValidatorUpdate{}
}

func init() {
	proto.RegisterType((*Params)(nil), "interchain_security.ccv.child.v1.Params")
	proto.RegisterType((*CrossChainValidator)(nil), "interchain_security.ccv.child.v1.CrossChainValidator")
}

func init() {
	proto.RegisterFile("interchain_security/ccv/child/v1/child.proto", fileDescriptor_d00ed7c78fea69b9)
}

var fileDescriptor_d00ed7c78fea69b9 = []byte{
	// 304 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xb1, 0x4a, 0xc4, 0x30,
	0x18, 0xc7, 0x1b, 0x87, 0x53, 0xaa, 0xa0, 0x54, 0x87, 0xe3, 0x84, 0x5c, 0xe9, 0x74, 0x83, 0x26,
	0x9c, 0x0e, 0x82, 0xa3, 0x87, 0xbb, 0x16, 0x74, 0x70, 0x39, 0xd2, 0x24, 0xb4, 0x81, 0xb6, 0x29,
	0x49, 0x1a, 0xec, 0x5b, 0x38, 0xfa, 0x48, 0x37, 0xde, 0xe8, 0x74, 0x48, 0xfb, 0x06, 0x3e, 0x81,
	0xb4, 0xb5, 0x77, 0x22, 0x6e, 0xdf, 0x47, 0x7e, 0xff, 0x2f, 0x7f, 0x7e, 0xee, 0x85, 0xc8, 0x0d,
	0x57, 0x34, 0x21, 0x22, 0x5f, 0x6a, 0x4e, 0x4b, 0x25, 0x4c, 0x85, 0x29, 0xb5, 0x98, 0x26, 0x22,
	0x65, 0xd8, 0xce, 0xfb, 0x01, 0x15, 0x4a, 0x1a, 0xe9, 0xf9, 0xff, 0xd0, 0x88, 0x52, 0x8b, 0x7a,
	0xc8, 0xce, 0x27, 0x67, 0xb1, 0x8c, 0x65, 0x07, 0xe3, 0x76, 0xea, 0x73, 0x93, 0x73, 0xc3, 0x73,
	0xc6, 0x55, 0x26, 0x72, 0x83, 0x49, 0x44, 0x05, 0x36, 0x55, 0xc1, 0x75, 0xff, 0x18, 0x04, 0xee,
	0xe8, 0x81, 0x28, 0x92, 0x69, 0x6f, 0xec, 0xee, 0xdf, 0xe7, 0x24, 0x4a, 0x39, 0x1b, 0x03, 0x1f,
	0xcc, 0x0e, 0xc2, 0x61, 0x0d, 0xde, 0x81, 0x7b, 0xba, 0x50, 0x52, 0xeb, 0x45, 0xfb, 0xf7, 0x33,
	0x49, 0x05, 0x23, 0x46, 0xaa, 0x36, 0x41, 0x18, 0x53, 0x5c, 0xeb, 0x2e, 0x71, 0x14, 0x0e, 0xab,
	0x17, 0xbb, 0x27, 0x76, 0xc0, 0x96, 0x65, 0xc1, 0x88, 0xe1, 0xe3, 0x3d, 0x1f, 0xcc, 0x0e, 0xaf,
	0x7c, 0xb4, 0x6b, 0x83, 0xda, 0x36, 0x68, 0x7b, 0xef, 0xa9, 0xe3, 0xee, 0x26, 0xab, 0xcd, 0xd4,
	0xf9, 0xda, 0x4c, 0xbd, 0x8a, 0x64, 0xe9, 0x6d, 0x90, 0x12, 0x6d, 0x7e, 0x4e, 0x04, 0xe1, 0xb1,
	0xfd, 0x03, 0x3f, 0xae, 0x6a, 0x08, 0xd6, 0x35, 0x04, 0x9f, 0x35, 0x04, 0x6f, 0x0d, 0x74, 0xd6,
	0x0d, 0x74, 0x3e, 0x1a, 0xe8, 0xbc, 0xdc, 0xc4, 0xc2, 0x24, 0x65, 0x84, 0xa8, 0xcc, 0x30, 0x95,
	0x3a, 0x93, 0x1a, 0xef, 0xfc, 0x5d, 0x6e, 0x6d, 0xbf, 0xfe, 0xf2, 0xdd, 0x79, 0x89, 0x46, 0x9d,
	0x98, 0xeb, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0xae, 0xa0, 0x93, 0xbd, 0x9d, 0x01, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Enabled {
		i--
		if m.Enabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *CrossChainValidator) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CrossChainValidator) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CrossChainValidator) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.ValidatorUpdate.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintChild(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintChild(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintChild(dAtA []byte, offset int, v uint64) int {
	offset -= sovChild(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Enabled {
		n += 2
	}
	return n
}

func (m *CrossChainValidator) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovChild(uint64(l))
	}
	l = m.ValidatorUpdate.Size()
	n += 1 + l + sovChild(uint64(l))
	return n
}

func sovChild(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozChild(x uint64) (n int) {
	return sovChild(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowChild
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Enabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChild
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Enabled = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipChild(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthChild
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
func (m *CrossChainValidator) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowChild
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
			return fmt.Errorf("proto: CrossChainValidator: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CrossChainValidator: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChild
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthChild
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthChild
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = append(m.Address[:0], dAtA[iNdEx:postIndex]...)
			if m.Address == nil {
				m.Address = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorUpdate", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChild
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
				return ErrInvalidLengthChild
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthChild
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ValidatorUpdate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipChild(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthChild
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
func skipChild(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowChild
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
					return 0, ErrIntOverflowChild
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
					return 0, ErrIntOverflowChild
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
				return 0, ErrInvalidLengthChild
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupChild
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthChild
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthChild        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowChild          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupChild = fmt.Errorf("proto: unexpected end of group")
)
