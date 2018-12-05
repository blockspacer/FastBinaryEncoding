// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: test.fbe
// Version: 1.1.0.0

package test

import "fmt"
import "strconv"
import "strings"
import "../fbe"
import "../proto"

// Workaround for Go unused imports issue
var _ = fbe.Version
var _ = proto.Version

// Workaround for Go unused imports issue
var _ = fmt.Print
var _ = strconv.FormatInt

// StructSimple key
type StructSimpleKey struct {
    Uid int32
    value *StructSimple
}

// // Get StructSimple value
func (k *StructSimpleKey) Value() *StructSimple {
    return k.value
}

// Convert StructSimple flags key to string
func (k *StructSimpleKey) String() string {
    var sb strings.Builder
    sb.WriteString("StructSimpleKey(")
    sb.WriteString("uid=")
    sb.WriteString(strconv.FormatInt(int64(k.Uid), 10))
    sb.WriteString(")")
    return sb.String()
}

// StructSimple struct
type StructSimple struct {
    Uid int32 `json:"uid"`
    F1 bool `json:"f1"`
    F2 bool `json:"f2"`
    F3 byte `json:"f3"`
    F4 byte `json:"f4"`
    F5 rune `json:"f5"`
    F6 rune `json:"f6"`
    F7 rune `json:"f7"`
    F8 rune `json:"f8"`
    F9 int8 `json:"f9"`
    F10 int8 `json:"f10"`
    F11 uint8 `json:"f11"`
    F12 uint8 `json:"f12"`
    F13 int16 `json:"f13"`
    F14 int16 `json:"f14"`
    F15 uint16 `json:"f15"`
    F16 uint16 `json:"f16"`
    F17 int32 `json:"f17"`
    F18 int32 `json:"f18"`
    F19 uint32 `json:"f19"`
    F20 uint32 `json:"f20"`
    F21 int64 `json:"f21"`
    F22 int64 `json:"f22"`
    F23 uint64 `json:"f23"`
    F24 uint64 `json:"f24"`
    F25 float32 `json:"f25"`
    F26 float32 `json:"f26"`
    F27 float64 `json:"f27"`
    F28 float64 `json:"f28"`
    F29 fbe.Decimal `json:"f29"`
    F30 fbe.Decimal `json:"f30"`
    F31 string `json:"f31"`
    F32 string `json:"f32"`
    F33 fbe.Timestamp `json:"f33"`
    F34 fbe.Timestamp `json:"f34"`
    F35 fbe.Timestamp `json:"f35"`
    F36 fbe.UUID `json:"f36"`
    F37 fbe.UUID `json:"f37"`
    F38 fbe.UUID `json:"f38"`
    F39 proto.OrderSide `json:"f39"`
    F40 proto.OrderType `json:"f40"`
    F41 proto.Order `json:"f41"`
    F42 proto.Balance `json:"f42"`
    F43 proto.State `json:"f43"`
    F44 proto.Account `json:"f44"`
}

// Create a new StructSimple struct
func NewStructSimple() *StructSimple {
    return &StructSimple{
        Uid: 0,
        F1: false,
        F2: true,
        F3: 0,
        F4: byte(255),
        F5: '\000',
        F6: '!',
        F7: '\000',
        F8: rune(0x0444),
        F9: 0,
        F10: int8(127),
        F11: 0,
        F12: uint8(255),
        F13: 0,
        F14: int16(32767),
        F15: 0,
        F16: uint16(65535),
        F17: 0,
        F18: int32(2147483647),
        F19: 0,
        F20: uint32(0xFFFFFFFF),
        F21: 0,
        F22: int64(9223372036854775807),
        F23: 0,
        F24: uint64(0xFFFFFFFFFFFFFFFF),
        F25: 0.0,
        F26: float32(123.456),
        F27: 0.0,
        F28: float64(-123.456e+123),
        F29: fbe.DecimalZero(),
        F30: fbe.DecimalFromString("123456.123456"),
        F31: "",
        F32: "Initial string!",
        F33: fbe.TimestampEpoch(),
        F34: fbe.TimestampEpoch(),
        F35: fbe.TimestampUTC(),
        F36: fbe.UUIDNil(),
        F37: fbe.UUIDSequential(),
        F38: fbe.UUIDFromString("123e4567-e89b-12d3-a456-426655440000"),
        F39: *proto.NewOrderSide(),
        F40: *proto.NewOrderType(),
        F41: *proto.NewOrder(),
        F42: *proto.NewBalance(),
        F43: *proto.NewState(),
        F44: *proto.NewAccount(),
    }
}

// Create a new StructSimple struct from JSON
func NewStructSimpleFromJSON(buffer []byte) (*StructSimple, error) {
    result := *NewStructSimple()
    err := fbe.Json.Unmarshal(buffer, &result)
    if err != nil {
        return nil, err
    }
    return &result, nil
}

// Struct shallow copy
func (s *StructSimple) Copy() *StructSimple {
    var result = *s
    return &result
}

// Struct deep clone
func (s *StructSimple) Clone() *StructSimple {
    // Serialize the struct to the FBE stream
    writer := NewStructSimpleModel(fbe.NewEmptyBuffer())
    _, _ = writer.Serialize(s)

    // Deserialize the struct from the FBE stream
    reader := NewStructSimpleModel(writer.Buffer())
    result, _, _ := reader.Deserialize()
    return result
}

// Get the struct key
func (s *StructSimple) Key() StructSimpleKey {
    return StructSimpleKey{
        Uid: s.Uid,
    }
}

// Convert struct to optional
func (s *StructSimple) Optional() *StructSimple {
    return s
}

// Convert struct to string
func (s *StructSimple) String() string {
    var sb strings.Builder
    sb.WriteString("StructSimple(")
    sb.WriteString("uid=")
    sb.WriteString(strconv.FormatInt(int64(s.Uid), 10))
    sb.WriteString(",f1=")
    sb.WriteString(strconv.FormatBool(s.F1))
    sb.WriteString(",f2=")
    sb.WriteString(strconv.FormatBool(s.F2))
    sb.WriteString(",f3=")
    sb.WriteString(strconv.FormatUint(uint64(s.F3), 10))
    sb.WriteString(",f4=")
    sb.WriteString(strconv.FormatUint(uint64(s.F4), 10))
    sb.WriteString(",f5=")
    sb.WriteString("'" + string(s.F5) + "'")
    sb.WriteString(",f6=")
    sb.WriteString("'" + string(s.F6) + "'")
    sb.WriteString(",f7=")
    sb.WriteString("'" + string(s.F7) + "'")
    sb.WriteString(",f8=")
    sb.WriteString("'" + string(s.F8) + "'")
    sb.WriteString(",f9=")
    sb.WriteString(strconv.FormatInt(int64(s.F9), 10))
    sb.WriteString(",f10=")
    sb.WriteString(strconv.FormatInt(int64(s.F10), 10))
    sb.WriteString(",f11=")
    sb.WriteString(strconv.FormatUint(uint64(s.F11), 10))
    sb.WriteString(",f12=")
    sb.WriteString(strconv.FormatUint(uint64(s.F12), 10))
    sb.WriteString(",f13=")
    sb.WriteString(strconv.FormatInt(int64(s.F13), 10))
    sb.WriteString(",f14=")
    sb.WriteString(strconv.FormatInt(int64(s.F14), 10))
    sb.WriteString(",f15=")
    sb.WriteString(strconv.FormatUint(uint64(s.F15), 10))
    sb.WriteString(",f16=")
    sb.WriteString(strconv.FormatUint(uint64(s.F16), 10))
    sb.WriteString(",f17=")
    sb.WriteString(strconv.FormatInt(int64(s.F17), 10))
    sb.WriteString(",f18=")
    sb.WriteString(strconv.FormatInt(int64(s.F18), 10))
    sb.WriteString(",f19=")
    sb.WriteString(strconv.FormatUint(uint64(s.F19), 10))
    sb.WriteString(",f20=")
    sb.WriteString(strconv.FormatUint(uint64(s.F20), 10))
    sb.WriteString(",f21=")
    sb.WriteString(strconv.FormatInt(int64(s.F21), 10))
    sb.WriteString(",f22=")
    sb.WriteString(strconv.FormatInt(int64(s.F22), 10))
    sb.WriteString(",f23=")
    sb.WriteString(strconv.FormatUint(uint64(s.F23), 10))
    sb.WriteString(",f24=")
    sb.WriteString(strconv.FormatUint(uint64(s.F24), 10))
    sb.WriteString(",f25=")
    sb.WriteString(strconv.FormatFloat(float64(s.F25), 'g', -1, 32))
    sb.WriteString(",f26=")
    sb.WriteString(strconv.FormatFloat(float64(s.F26), 'g', -1, 32))
    sb.WriteString(",f27=")
    sb.WriteString(strconv.FormatFloat(float64(s.F27), 'g', -1, 64))
    sb.WriteString(",f28=")
    sb.WriteString(strconv.FormatFloat(float64(s.F28), 'g', -1, 64))
    sb.WriteString(",f29=")
    sb.WriteString(s.F29.String())
    sb.WriteString(",f30=")
    sb.WriteString(s.F30.String())
    sb.WriteString(",f31=")
    sb.WriteString("\"" + s.F31 + "\"")
    sb.WriteString(",f32=")
    sb.WriteString("\"" + s.F32 + "\"")
    sb.WriteString(",f33=")
    sb.WriteString(strconv.FormatInt(s.F33.UnixNano(), 10))
    sb.WriteString(",f34=")
    sb.WriteString(strconv.FormatInt(s.F34.UnixNano(), 10))
    sb.WriteString(",f35=")
    sb.WriteString(strconv.FormatInt(s.F35.UnixNano(), 10))
    sb.WriteString(",f36=")
    sb.WriteString("\"" + s.F36.String() + "\"")
    sb.WriteString(",f37=")
    sb.WriteString("\"" + s.F37.String() + "\"")
    sb.WriteString(",f38=")
    sb.WriteString("\"" + s.F38.String() + "\"")
    sb.WriteString(",f39=")
    sb.WriteString(s.F39.String())
    sb.WriteString(",f40=")
    sb.WriteString(s.F40.String())
    sb.WriteString(",f41=")
    sb.WriteString(s.F41.String())
    sb.WriteString(",f42=")
    sb.WriteString(s.F42.String())
    sb.WriteString(",f43=")
    sb.WriteString(s.F43.String())
    sb.WriteString(",f44=")
    sb.WriteString(s.F44.String())
    sb.WriteString(")")
    return sb.String()
}

// Convert struct to JSON
func (s *StructSimple) JSON() ([]byte, error) {
    return fbe.Json.Marshal(s)
}