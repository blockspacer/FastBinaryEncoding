// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: test.fbe
// Version: 1.1.0.0

package test

import "strconv"
import "strings"
import "encoding/json"
import "../fbe"
import "../proto"

// Workaround for Go unused imports issue
var _ = fbe.Version
var _ = proto.Version

// StructOptional key
type StructOptionalKey struct {
    StructSimpleKey
}

// Convert StructOptional flags key to string
func (k StructOptionalKey) String() string {
    var sb strings.Builder
    return sb.String()
}

// StructOptional struct
type StructOptional struct {
    StructSimple
    F100 *bool `json:"f100"`
    F101 *bool `json:"f101"`
    F102 *bool `json:"f102"`
    F103 *byte `json:"f103"`
    F104 *byte `json:"f104"`
    F105 *byte `json:"f105"`
    F106 *rune `json:"f106"`
    F107 *rune `json:"f107"`
    F108 *rune `json:"f108"`
    F109 *rune `json:"f109"`
    F110 *rune `json:"f110"`
    F111 *rune `json:"f111"`
    F112 *int8 `json:"f112"`
    F113 *int8 `json:"f113"`
    F114 *int8 `json:"f114"`
    F115 *uint8 `json:"f115"`
    F116 *uint8 `json:"f116"`
    F117 *uint8 `json:"f117"`
    F118 *int16 `json:"f118"`
    F119 *int16 `json:"f119"`
    F120 *int16 `json:"f120"`
    F121 *uint16 `json:"f121"`
    F122 *uint16 `json:"f122"`
    F123 *uint16 `json:"f123"`
    F124 *int32 `json:"f124"`
    F125 *int32 `json:"f125"`
    F126 *int32 `json:"f126"`
    F127 *uint32 `json:"f127"`
    F128 *uint32 `json:"f128"`
    F129 *uint32 `json:"f129"`
    F130 *int64 `json:"f130"`
    F131 *int64 `json:"f131"`
    F132 *int64 `json:"f132"`
    F133 *uint64 `json:"f133"`
    F134 *uint64 `json:"f134"`
    F135 *uint64 `json:"f135"`
    F136 *float32 `json:"f136"`
    F137 *float32 `json:"f137"`
    F138 *float32 `json:"f138"`
    F139 *float64 `json:"f139"`
    F140 *float64 `json:"f140"`
    F141 *float64 `json:"f141"`
    F142 *fbe.Decimal `json:"f142"`
    F143 *fbe.Decimal `json:"f143"`
    F144 *fbe.Decimal `json:"f144"`
    F145 *string `json:"f145"`
    F146 *string `json:"f146"`
    F147 *string `json:"f147"`
    F148 *fbe.Timestamp `json:"f148"`
    F149 *fbe.Timestamp `json:"f149"`
    F150 *fbe.Timestamp `json:"f150"`
    F151 *fbe.UUID `json:"f151"`
    F152 *fbe.UUID `json:"f152"`
    F153 *fbe.UUID `json:"f153"`
    F154 *proto.OrderSide `json:"f154"`
    F155 *proto.OrderSide `json:"f155"`
    F156 *proto.OrderType `json:"f156"`
    F157 *proto.OrderType `json:"f157"`
    F158 *proto.Order `json:"f158"`
    F159 *proto.Order `json:"f159"`
    F160 *proto.Balance `json:"f160"`
    F161 *proto.Balance `json:"f161"`
    F162 *proto.State `json:"f162"`
    F163 *proto.State `json:"f163"`
    F164 *proto.Account `json:"f164"`
    F165 *proto.Account `json:"f165"`
}

// Create a new StructOptional struct
func NewStructOptional() *StructOptional {
    return &StructOptional{
        StructSimple: *NewStructSimple(),
        F100: nil,
        F101: fbe.OptionalBool(true),
        F102: nil,
        F103: nil,
        F104: fbe.OptionalByte(byte(255)),
        F105: nil,
        F106: nil,
        F107: fbe.OptionalRune('!'),
        F108: nil,
        F109: nil,
        F110: fbe.OptionalRune(rune(0x0444)),
        F111: nil,
        F112: nil,
        F113: fbe.OptionalInt8(int8(127)),
        F114: nil,
        F115: nil,
        F116: fbe.OptionalUInt8(uint8(255)),
        F117: nil,
        F118: nil,
        F119: fbe.OptionalInt16(int16(32767)),
        F120: nil,
        F121: nil,
        F122: fbe.OptionalUInt16(uint16(65535)),
        F123: nil,
        F124: nil,
        F125: fbe.OptionalInt32(int32(2147483647)),
        F126: nil,
        F127: nil,
        F128: fbe.OptionalUInt32(uint32(0xFFFFFFFF)),
        F129: nil,
        F130: nil,
        F131: fbe.OptionalInt64(int64(9223372036854775807)),
        F132: nil,
        F133: nil,
        F134: fbe.OptionalUInt64(uint64(0xFFFFFFFFFFFFFFFF)),
        F135: nil,
        F136: nil,
        F137: fbe.OptionalFloat32(float32(123.456)),
        F138: nil,
        F139: nil,
        F140: fbe.OptionalFloat64(float64(-123.456e+123)),
        F141: nil,
        F142: nil,
        F143: fbe.OptionalDecimal(fbe.DecimalFromString("123456.123456")),
        F144: nil,
        F145: nil,
        F146: fbe.OptionalString("Initial string!"),
        F147: nil,
        F148: nil,
        F149: fbe.OptionalTimestamp(fbe.TimestampUTC()),
        F150: nil,
        F151: nil,
        F152: fbe.OptionalUUID(fbe.UUIDFromString("123e4567-e89b-12d3-a456-426655440000")),
        F153: nil,
        F154: nil,
        F155: nil,
        F156: nil,
        F157: nil,
        F158: nil,
        F159: nil,
        F160: nil,
        F161: nil,
        F162: nil,
        F163: nil,
        F164: nil,
        F165: nil,
    }
}

// Create a new StructOptional struct from JSON
func NewStructOptionalFromJSON(buffer []byte) (*StructOptional, error) {
    var result StructOptional
    err := json.Unmarshal(buffer, &result)
    if err != nil {
        return nil, err
    }
    return &result, nil
}

// Struct shallow copy
func (s StructOptional) Copy() *StructOptional {
    var result = s
    return &result
}

// Struct deep clone
func (s StructOptional) Clone() *StructOptional {
    var result = s
    return &result
}

// Get the struct key
func (s StructOptional) Key() StructOptionalKey {
    return StructOptionalKey{
        s.StructSimple.Key(),
    }
}

// Convert struct to optional
func (s StructOptional) Optional() *StructOptional {
    return &s
}

// Convert struct to string
func (s StructOptional) String() string {
    var sb strings.Builder
    return sb.String()
}

// Convert struct to JSON
func (s StructOptional) JSON() ([]byte, error) {
    return json.Marshal(s)
}
