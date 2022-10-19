// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: teleport/attestation/v1/attestation.proto

package v1

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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

// AttestationStatement is an attestation statement for a hardware private key.
type AttestationStatement struct {
	// Types that are valid to be assigned to AttestationStatement:
	//
	//	*AttestationStatement_YubikeyAttestationStatement
	AttestationStatement isAttestationStatement_AttestationStatement `protobuf_oneof:"attestation_statement"`
	XXX_NoUnkeyedLiteral struct{}                                    `json:"-"`
	XXX_unrecognized     []byte                                      `json:"-"`
	XXX_sizecache        int32                                       `json:"-"`
}

func (m *AttestationStatement) Reset()         { *m = AttestationStatement{} }
func (m *AttestationStatement) String() string { return proto.CompactTextString(m) }
func (*AttestationStatement) ProtoMessage()    {}
func (*AttestationStatement) Descriptor() ([]byte, []int) {
	return fileDescriptor_99516272d9ee5937, []int{0}
}
func (m *AttestationStatement) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AttestationStatement) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AttestationStatement.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AttestationStatement) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AttestationStatement.Merge(m, src)
}
func (m *AttestationStatement) XXX_Size() int {
	return m.Size()
}
func (m *AttestationStatement) XXX_DiscardUnknown() {
	xxx_messageInfo_AttestationStatement.DiscardUnknown(m)
}

var xxx_messageInfo_AttestationStatement proto.InternalMessageInfo

type isAttestationStatement_AttestationStatement interface {
	isAttestationStatement_AttestationStatement()
	MarshalTo([]byte) (int, error)
	Size() int
}

type AttestationStatement_YubikeyAttestationStatement struct {
	YubikeyAttestationStatement *YubiKeyAttestationStatement `protobuf:"bytes,1,opt,name=yubikey_attestation_statement,json=yubikeyAttestationStatement,proto3,oneof" json:"yubikey_attestation_statement,omitempty"`
}

func (*AttestationStatement_YubikeyAttestationStatement) isAttestationStatement_AttestationStatement() {
}

func (m *AttestationStatement) GetAttestationStatement() isAttestationStatement_AttestationStatement {
	if m != nil {
		return m.AttestationStatement
	}
	return nil
}

func (m *AttestationStatement) GetYubikeyAttestationStatement() *YubiKeyAttestationStatement {
	if x, ok := m.GetAttestationStatement().(*AttestationStatement_YubikeyAttestationStatement); ok {
		return x.YubikeyAttestationStatement
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*AttestationStatement) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*AttestationStatement_YubikeyAttestationStatement)(nil),
	}
}

// YubiKeyAttestationStatement is an attestation statement for a specific YubiKey PIV slot.
type YubiKeyAttestationStatement struct {
	// slot_cert is an attestation certificate generated from a YubiKey PIV
	// slot's public key and signed by the YubiKey's attestation certificate.
	SlotCert []byte `protobuf:"bytes,1,opt,name=slot_cert,json=slotCert,proto3" json:"slot_cert,omitempty"`
	// attestation_cert is the YubiKey's unique attestation certificate, signed by a Yubico CA.
	AttestationCert      []byte   `protobuf:"bytes,2,opt,name=attestation_cert,json=attestationCert,proto3" json:"attestation_cert,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *YubiKeyAttestationStatement) Reset()         { *m = YubiKeyAttestationStatement{} }
func (m *YubiKeyAttestationStatement) String() string { return proto.CompactTextString(m) }
func (*YubiKeyAttestationStatement) ProtoMessage()    {}
func (*YubiKeyAttestationStatement) Descriptor() ([]byte, []int) {
	return fileDescriptor_99516272d9ee5937, []int{1}
}
func (m *YubiKeyAttestationStatement) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *YubiKeyAttestationStatement) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_YubiKeyAttestationStatement.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *YubiKeyAttestationStatement) XXX_Merge(src proto.Message) {
	xxx_messageInfo_YubiKeyAttestationStatement.Merge(m, src)
}
func (m *YubiKeyAttestationStatement) XXX_Size() int {
	return m.Size()
}
func (m *YubiKeyAttestationStatement) XXX_DiscardUnknown() {
	xxx_messageInfo_YubiKeyAttestationStatement.DiscardUnknown(m)
}

var xxx_messageInfo_YubiKeyAttestationStatement proto.InternalMessageInfo

func (m *YubiKeyAttestationStatement) GetSlotCert() []byte {
	if m != nil {
		return m.SlotCert
	}
	return nil
}

func (m *YubiKeyAttestationStatement) GetAttestationCert() []byte {
	if m != nil {
		return m.AttestationCert
	}
	return nil
}

func init() {
	proto.RegisterType((*AttestationStatement)(nil), "teleport.attestation.v1.AttestationStatement")
	proto.RegisterType((*YubiKeyAttestationStatement)(nil), "teleport.attestation.v1.YubiKeyAttestationStatement")
}

func init() {
	proto.RegisterFile("teleport/attestation/v1/attestation.proto", fileDescriptor_99516272d9ee5937)
}

var fileDescriptor_99516272d9ee5937 = []byte{
	// 243 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x2c, 0x49, 0xcd, 0x49,
	0x2d, 0xc8, 0x2f, 0x2a, 0xd1, 0x4f, 0x2c, 0x29, 0x49, 0x2d, 0x2e, 0x49, 0x2c, 0xc9, 0xcc, 0xcf,
	0xd3, 0x2f, 0x33, 0x44, 0xe6, 0xea, 0x15, 0x14, 0xe5, 0x97, 0xe4, 0x0b, 0x89, 0xc3, 0x94, 0xea,
	0x21, 0xcb, 0x95, 0x19, 0x2a, 0xad, 0x66, 0xe4, 0x12, 0x71, 0x44, 0x08, 0x05, 0x97, 0x24, 0x96,
	0xa4, 0xe6, 0xa6, 0xe6, 0x95, 0x08, 0x55, 0x71, 0xc9, 0x56, 0x96, 0x26, 0x65, 0x66, 0xa7, 0x56,
	0xc6, 0x23, 0x69, 0x89, 0x2f, 0x86, 0x29, 0x90, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x36, 0x32, 0xd1,
	0xc3, 0x61, 0xb2, 0x5e, 0x64, 0x69, 0x52, 0xa6, 0x77, 0x6a, 0x25, 0x36, 0xc3, 0x3d, 0x18, 0x82,
	0xa4, 0xa1, 0x86, 0x63, 0x93, 0x76, 0x12, 0xe7, 0x12, 0xc5, 0x6a, 0xa7, 0x52, 0x2a, 0x97, 0x34,
	0x1e, 0x63, 0x85, 0xa4, 0xb9, 0x38, 0x8b, 0x73, 0xf2, 0x4b, 0xe2, 0x93, 0x53, 0x8b, 0x20, 0xee,
	0xe3, 0x09, 0xe2, 0x00, 0x09, 0x38, 0xa7, 0x16, 0x95, 0x08, 0x69, 0x72, 0x09, 0x20, 0x1b, 0x0a,
	0x56, 0xc3, 0x04, 0x56, 0xc3, 0x8f, 0x24, 0x0e, 0x52, 0xea, 0xe4, 0x7f, 0xe2, 0x91, 0x1c, 0xe3,
	0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0x46, 0x39, 0xa6, 0x67, 0x96, 0x64, 0x94, 0x26,
	0xe9, 0x25, 0xe7, 0xe7, 0xea, 0xa7, 0x17, 0x25, 0x96, 0x65, 0x42, 0x54, 0x26, 0xe6, 0xe8, 0x23,
	0xc2, 0xbf, 0x20, 0x53, 0x3f, 0x3d, 0x35, 0x4f, 0x1f, 0x1c, 0xd8, 0xfa, 0xe9, 0xf9, 0x68, 0x11,
	0x92, 0xc4, 0x06, 0x96, 0x30, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xb1, 0x5e, 0x52, 0x76, 0xb2,
	0x01, 0x00, 0x00,
}

func (m *AttestationStatement) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AttestationStatement) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AttestationStatement) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.AttestationStatement != nil {
		{
			size := m.AttestationStatement.Size()
			i -= size
			if _, err := m.AttestationStatement.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *AttestationStatement_YubikeyAttestationStatement) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AttestationStatement_YubikeyAttestationStatement) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.YubikeyAttestationStatement != nil {
		{
			size, err := m.YubikeyAttestationStatement.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintAttestation(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}
func (m *YubiKeyAttestationStatement) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *YubiKeyAttestationStatement) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *YubiKeyAttestationStatement) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.AttestationCert) > 0 {
		i -= len(m.AttestationCert)
		copy(dAtA[i:], m.AttestationCert)
		i = encodeVarintAttestation(dAtA, i, uint64(len(m.AttestationCert)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.SlotCert) > 0 {
		i -= len(m.SlotCert)
		copy(dAtA[i:], m.SlotCert)
		i = encodeVarintAttestation(dAtA, i, uint64(len(m.SlotCert)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintAttestation(dAtA []byte, offset int, v uint64) int {
	offset -= sovAttestation(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *AttestationStatement) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.AttestationStatement != nil {
		n += m.AttestationStatement.Size()
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *AttestationStatement_YubikeyAttestationStatement) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.YubikeyAttestationStatement != nil {
		l = m.YubikeyAttestationStatement.Size()
		n += 1 + l + sovAttestation(uint64(l))
	}
	return n
}
func (m *YubiKeyAttestationStatement) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.SlotCert)
	if l > 0 {
		n += 1 + l + sovAttestation(uint64(l))
	}
	l = len(m.AttestationCert)
	if l > 0 {
		n += 1 + l + sovAttestation(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovAttestation(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAttestation(x uint64) (n int) {
	return sovAttestation(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AttestationStatement) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAttestation
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
			return fmt.Errorf("proto: AttestationStatement: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AttestationStatement: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field YubikeyAttestationStatement", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAttestation
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
				return ErrInvalidLengthAttestation
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAttestation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &YubiKeyAttestationStatement{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.AttestationStatement = &AttestationStatement_YubikeyAttestationStatement{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAttestation(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAttestation
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *YubiKeyAttestationStatement) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAttestation
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
			return fmt.Errorf("proto: YubiKeyAttestationStatement: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: YubiKeyAttestationStatement: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SlotCert", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAttestation
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
				return ErrInvalidLengthAttestation
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthAttestation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SlotCert = append(m.SlotCert[:0], dAtA[iNdEx:postIndex]...)
			if m.SlotCert == nil {
				m.SlotCert = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AttestationCert", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAttestation
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
				return ErrInvalidLengthAttestation
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthAttestation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AttestationCert = append(m.AttestationCert[:0], dAtA[iNdEx:postIndex]...)
			if m.AttestationCert == nil {
				m.AttestationCert = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAttestation(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAttestation
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipAttestation(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAttestation
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
					return 0, ErrIntOverflowAttestation
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
					return 0, ErrIntOverflowAttestation
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
				return 0, ErrInvalidLengthAttestation
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAttestation
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAttestation
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAttestation        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAttestation          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAttestation = fmt.Errorf("proto: unexpected end of group")
)
