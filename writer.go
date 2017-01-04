package d2protocol

// Writer is the minimal interface required to serialiaze any
// protocol message/type
type Writer interface {
	WriteBoolean(bool) error
	WriteInt8(int8) error
	WriteUInt8(uint8) error
	WriteInt16(int16) error
	WriteUInt16(uint16) error
	WriteInt32(int32) error
	WriteUInt32(uint32) error
	WriteFloat(float32) error
	WriteDouble(float64) error
	WriteString(string) error
	WriteVarInt16(int16) error
	WriteVarUInt16(uint16) error
	WriteVarInt32(int32) error
	WriteVarUInt32(uint32) error
	WriteVarInt64(int64) error
	WriteVarUInt64(uint64) error
}
